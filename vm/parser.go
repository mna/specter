package vm

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func (vm *VM) parse(r io.Reader) {
	bio := bufio.NewReader(r)

	// Read one line at a time, a single line can contain at most one instruction,
	// possibly zero if it is only a label (a line may also contain a label AND an
	// instruction).
	for l, err := bio.ReadString('\n'); len(l) > 0; l, err = bio.ReadString('\n') {
		// Split the line in tokens
		toks := strings.FieldsFunc(l, func(r rune) bool {
			// Split on either a space, a comma or a tab
			switch r {
			case ' ', ',', '\t':
				return true
			}
			return false
		})

		// Loop through the tokens, store labels, instructions and arguments
		hasInstr := false
		argIdx := 0
		for _, tok := range toks {
			if strings.HasPrefix(tok, "#") {
				// This is a comment, ignore all other tokens on this line
				break
			}

			// Ignore empty tokens
			if tok == "" {
				continue
			}

			// Is it a label definition?
			if vm.parseLabelDef(tok) {
				if hasInstr {
					panic(fmt.Sprintf("cannot define label '%s' after an instruction on the same line", tok))
				}
				continue
			}

			// Is it an instruction (opcode)?
			if vm.parseInstr(tok) {
				hasInstr = true
				continue
			}

			// It is not a comment, nor a label definition, nor an instruction, so this is
			// an argument. Make sure an instruction has been found.
			if !hasInstr {
				panic(fmt.Sprintf("found argument token '%s' without an instruction", tok))
			} else if argIdx > 1 {
				panic(fmt.Sprintf("found excessive argument token '%s' after two arguments", tok))
			}
			if vm.parseRegister(tok, argIdx) {
				argIdx++
				continue
			}
			if vm.parseLabelVal(tok, argIdx) {
				argIdx++
				continue
			}

			// Parse value panics if the value is invalid, so must be last, and no need 
			// to add a panic after the call (or a continue)
			if vm.parseValue(tok, argIdx) {
				argIdx++
			}
		}

		// If EOF or error, return
		if err != nil {
			if err != io.EOF {
				panic(err)
			} else {
				break
			}
		}
	}
}

func (vm *VM) parseValue(tok string, argIdx int) bool {
	sepIdx := strings.IndexRune(tok, '|')
	base := 0
	val := tok

	if sepIdx > 0 && sepIdx < (len(tok)-1) {
		val = tok[:sepIdx]
		switch tok[sepIdx+1:] {
		case "h":
			base = 16
		case "d":
			base = 10
		case "o":
			base = 8
		case "b":
			base = 2
		default:
			panic(fmt.Sprintf("invalid base notation for value token '%s'", tok))
		}
	}
	if i, err := strconv.ParseInt(val, base, 32); err != nil {
		panic(err)
	} else {
		// The instructions pointer points on the next instruction slot at this point,
		// so use minus one.
		// In Go, it is totally legal to grab the address of a stack variable, so
		// we can avoid the p.values slice altogether.
		var i32 int32 = int32(i)
		vm.p.args[vm.p.instrs.size-1][argIdx] = &i32
		return true
	}
	panic("unreachable")
}

func (vm *VM) parseRegister(tok string, argIdx int) bool {
	if reg, ok := rgsMap[tok]; ok {
		// The instructions pointer points on the next instruction slot at this point,
		// so use minus one.
		vm.p.args[vm.p.instrs.size-1][argIdx] = &vm.m.registers[reg].i32
		return true
	}

	return false
}

func (vm *VM) parseInstr(tok string) bool {
	if op, ok := opsMap[tok]; ok {
		// This is an instruction token
		vm.p.instrs.addIncr(int32(op))
		return true
	}

	return false
}

func (vm *VM) parseLabelVal(tok string, argIdx int) bool {
	if instr, ok := vm.p.labels[tok]; ok {
		// The instructions pointer points on the next instruction slot at this point,
		// so use minus one.
		// In Go, it is totally legal to grab the address of a stack variable, so
		// we can avoid the p.values slice altogether.
		var i32 int32 = int32(instr)
		vm.p.args[vm.p.instrs.size-1][argIdx] = &i32
		return true
	}

	return false
}

func (vm *VM) parseLabelDef(tok string) bool {
	if strings.HasSuffix(tok, ":") {
		// This is a label
		lbl := tok[:len(tok)-1]

		// Check if this is a duplicate TODO : Return error instead?
		if _, ok := vm.p.labels[lbl]; ok {
			// This label already exists
			panic(fmt.Sprintf("a label '%s' already exists", lbl))
		}
		// Store it with a pointer to the next instruction
		vm.p.labels[lbl] = vm.p.instrs.size
		// If this is the special-case "start" label, store the start instruction
		if lbl == "start" {
			vm.p.start = vm.p.instrs.size
		}

		return true
	}

	return false
}
