package vm

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	// Initial lines slice capacity
	_LINES_CAP = 1000

	// Maximum number of arguments for an opcode
	_MAX_ARGS = 2
)

// Parse the content from the reader, load it into the program so that it can
// be executed. The parser makes minimal checks, because it is assumed that a
// compiler/code generator produced its input, and it is up to this program
// to give relevant errors. The VM is not there to give helpful errors to correct
// its code.
func (vm *VM) parse(r io.Reader) {
	bio := bufio.NewReader(r)
	lines := make([][]string, 0, _LINES_CAP)
	lIdx := 0

	// First pass is to load label definitions and instructions only, since it has to be there
	// before parsing the arguments (jump-to-label).
	for l, err := bio.ReadString('\n'); true; l, err = bio.ReadString('\n') {
		// Split the line in tokens (ReadString returns the delimiter)
		lines = append(lines, strings.FieldsFunc(l, func(r rune) bool {
			// Split on either a space, a comma or a tab (or newline)
			switch r {
			case ' ', ',', '\t', '\n':
				return true
			}
			return false
		}))

		hasInstr := false
		// Loop through the tokens, store labels and instructions
		for _, tok := range lines[lIdx] {
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
		}

		// If EOF or error, return
		if err != nil {
			if err != io.EOF {
				panic(err)
			} else {
				break
			}
		}
		// Increment line index
		lIdx++
	}

	// Here we know exactly the number of instructions, so allocate the right size
	// for the arguments slice
	vm.p.args = make([]*int32, len(vm.p.instrs)*2)

	// Next, parse instruction arguments one line at a time, a single line can contain at most one instruction,
	// possibly zero if it is only a label (a line may also contain a label AND an
	// instruction).
	instrIdx := -1
	for _, toks := range lines {
		// Loop through the tokens, store arguments
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
			if strings.HasSuffix(tok, ":") {
				continue
			}

			// Is it an instruction (opcode)?
			if _, ok := opsMap[tok]; ok {
				instrIdx++
				hasInstr = true
				continue
			}

			// It is not a comment, nor a label definition, nor an instruction, so this is
			// an argument. Make sure an instruction has been found.
			if !hasInstr {
				panic(fmt.Sprintf("found argument token '%s' without an instruction", tok))
			} else if argIdx >= _MAX_ARGS {
				panic(fmt.Sprintf("found excessive argument token '%s' after %d arguments", tok, _MAX_ARGS))
			}
			if vm.parseRegister(tok, instrIdx, argIdx) {
				argIdx++
				continue
			}
			if vm.parseLabelVal(tok, instrIdx, argIdx) {
				argIdx++
				continue
			}
			if vm.parseAddress(tok, instrIdx, argIdx) {
				argIdx++
				continue
			}
			// Parse value panics if the value is invalid, so must be last, and no need 
			// to add a panic after the call (or a continue)
			if vm.parseValue(tok, instrIdx, argIdx) {
				argIdx++
			}
		}
	}
	// Insert a program-ending instruction, useful in execution loop
	vm.p.instrs = append(vm.p.instrs, _OP_END)
}

// Parse a literal value (with an optional base code)
func (vm *VM) parseValue(tok string, instrIdx int, argIdx int) bool {
	// In Go, it is totally legal to grab the address of a stack variable, so
	// we can avoid the p.values slice altogether.
	i32 := toValue(tok)
	vm.p.args[(instrIdx*2)+argIdx] = &i32
	return true
}

func toValue(tok string) int32 {
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
	// ParseInt natively supports decimals and hexadecimals (if value starts with 0x).
	// Other bases must use the | notation.
	if i, err := strconv.ParseInt(val, base, 32); err != nil {
		panic(err)
	} else {
		return int32(i)
	}
	panic("unreachable")
}

// Parse an address (heap) pointer, format: [123]
func (vm *VM) parseAddress(tok string, instrIdx int, argIdx int) bool {
	if strings.HasPrefix(tok, "[") {
		i := toValue(tok[1 : len(tok)-1])
		vm.p.args[(instrIdx*2)+argIdx] = &vm.m.heap[i]
		return true
	}

	return false
}

// Parse a register name.
func (vm *VM) parseRegister(tok string, instrIdx int, argIdx int) bool {
	if reg, ok := rgsMap[tok]; ok {
		vm.p.args[(instrIdx*2)+argIdx] = &vm.m.registers[reg]
		return true
	}

	return false
}

// Parse an instruction code (opcode).
func (vm *VM) parseInstr(tok string) bool {
	if op, ok := opsMap[tok]; ok {
		// This is an instruction token
		vm.p.instrs = append(vm.p.instrs, op)
		return true
	}

	return false
}

// Parse a label value (label used as argument, i.e. to a jump).
func (vm *VM) parseLabelVal(tok string, instrIdx int, argIdx int) bool {
	if instr, ok := vm.p.labels[tok]; ok {
		// In Go, it is totally legal to grab the address of a stack variable, so
		// we can avoid the p.values slice altogether.
		var i32 int32 = int32(instr)
		vm.p.args[(instrIdx*2)+argIdx] = &i32
		return true
	}

	return false
}

// Parse a label definition.
func (vm *VM) parseLabelDef(tok string) bool {
	if strings.HasSuffix(tok, ":") {
		// This is a label
		lbl := tok[:len(tok)-1]

		// Check if this is a register name (invalid label)
		if _, ok := rgsMap[lbl]; ok {
			// This label uses a register name
			panic(fmt.Sprintf("the register name '%s' cannot be used as label", lbl))
		}
		// Check if this is a duplicate
		if _, ok := vm.p.labels[lbl]; ok {
			// This label already exists
			panic(fmt.Sprintf("a label '%s' already exists", lbl))
		}
		// Store it with a pointer to the next instruction
		vm.p.labels[lbl] = int32(len(vm.p.instrs))
		// If this is the special-case "start" label, store the start instruction
		if lbl == "start" {
			vm.p.start = int32(len(vm.p.instrs))
		}

		return true
	}

	return false
}
