package vm

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

const _PBUF_CAP = 10

// The VM, with its program and memory abstractions.
type VM struct {
	p    *program
	m    *memory
	b    *bufio.Writer
	pbuf []byte
}

// Create a new VM.
func New() *VM {
	return NewWithWriter(os.Stdout)
}

// Create a new VM with the specified output stream.
func NewWithWriter(w io.Writer) *VM {
	return &VM{newProgram(), newMemory(), bufio.NewWriter(w), make([]byte, 0, _PBUF_CAP)}
}

// Run executes the vm bytecode read by the reader.
func (vm *VM) Run(r io.Reader) {
	var i int32

	// Parse the content to execute.
	vm.parse(r)

	// Execution loop.
	defer vm.b.Flush()
	for i = vm.p.start; vm.p.instrs[i] != _OP_END; i++ {
		vm.runInstruction(&i)
	}
}

// Run a single instruction.
func (vm *VM) runInstruction(instrIndex *int32) {
	a0, a1 := vm.p.args[((*instrIndex)*2)+0], vm.p.args[((*instrIndex)*2)+1]

	switch vm.p.instrs[*instrIndex] {
	case _OP_NOP:
		// Nothing
	case _OP_INT:
		// Not implemented
	case _OP_MOV:
		*a0 = *a1
	case _OP_PUSH:
		vm.m.pushStack(*a0)
	case _OP_POP:
		vm.m.popStack(a0)
	case _OP_PUSHF:
		vm.m.pushStack(vm.m.FLAGS)
	case _OP_POPF:
		vm.m.popStack(a0)
	case _OP_INC:
		(*a0)++
	case _OP_DEC:
		(*a0)--
	case _OP_ADD:
		*a0 += *a1
	case _OP_SUB:
		*a0 -= *a1
	case _OP_MUL:
		*a0 *= *a1
	case _OP_DIV:
		*a0 /= *a1
	case _OP_MOD:
		vm.m.remainder = *a0 % *a1
	case _OP_REM:
		*a0 = vm.m.remainder
	case _OP_NOT:
		*a0 = ^(*a0)
	case _OP_XOR:
		*a0 ^= *a1
	case _OP_OR:
		*a0 |= *a1
	case _OP_AND:
		*a0 &= *a1
	case _OP_SHL:
		// cannot shift on signed int32
		if *a1 > 0 {
			*a0 <<= uint(*a1)
		}
	case _OP_SHR:
		// cannot shift on signed int32
		if *a1 > 0 {
			*a0 >>= uint(*a1)
		}
	case _OP_CMP:
		if *a0 == *a1 {
			vm.m.FLAGS = 0x1
		} else if *a0 > *a1 {
			vm.m.FLAGS = 0x2
		} else {
			vm.m.FLAGS = 0x0
		}
	case _OP_CALL:
		vm.m.pushStack(*instrIndex)
		fallthrough
	case _OP_JMP:
		*instrIndex = *a0 - 1
	case _OP_RET:
		vm.m.popStack(instrIndex)
	case _OP_JE:
		if vm.m.FLAGS&0x1 != 0 {
			*instrIndex = *a0 - 1
		}
	case _OP_JNE:
		if vm.m.FLAGS&0x1 == 0 {
			*instrIndex = *a0 - 1
		}
	case _OP_JG:
		if vm.m.FLAGS&0x2 != 0 {
			*instrIndex = *a0 - 1
		}
	case _OP_JGE:
		if vm.m.FLAGS&0x3 != 0 {
			*instrIndex = *a0 - 1
		}
	case _OP_JL:
		if vm.m.FLAGS&0x3 == 0 {
			*instrIndex = *a0 - 1
		}
	case _OP_JLE:
		if vm.m.FLAGS&0x2 == 0 {
			*instrIndex = *a0 - 1
		}
	case _OP_PRN:
		vm.pbuf = vm.pbuf[:0] // Clear buffer
		vm.pbuf = strconv.AppendInt(vm.pbuf, int64(*a0), 10)
		vm.pbuf = append(vm.pbuf, '\n')
		vm.b.Write(vm.pbuf)
	}
}
