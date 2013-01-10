package vm

import (
	"bufio"
	"fmt"
	"io"
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
			if vm.parseLabel(tok) {
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
			}
			if vm.parseRegister(tok, argIdx) {
				argIdx++
				continue
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

func (vm *VM) parseLabel(tok string) bool {
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
