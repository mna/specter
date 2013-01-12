package vm

import (
	"fmt"
	"io"
)

type VM struct {
	p *program
	m *memory
}

func New() *VM {
	return &VM{newProgram(), newMemory()}
}

// Run executes the vm bytecode read by the reader.
func (vm *VM) Run(r io.Reader) {
	var i int32

	vm.parse(r)
	for i = vm.p.start; vm.p.instrs.sl[i] != int32(_OP_END); i++ {
		vm.runInstruction(&i)
	}
}

func (vm *VM) runInstruction(instrIndex *int32) {
	a0, a1 := vm.p.args[*instrIndex][0], vm.p.args[*instrIndex][1]

	//printInstr("before", *instrIndex, opcode(vm.p.instrs.sl[*instrIndex]), a0, a1)

	switch opcode(vm.p.instrs.sl[*instrIndex]) {
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
		// TODO : Unimplemented, cannot shift on signed int32
		//*a0 <<= *a1
	case _OP_SHR:
		// TODO : Unimplemented, cannot shift on signed int32
		//*a0 >>= *a1
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
		fmt.Println(*a0)
	}
	/*
		if *instrIndex >= 0 {
			printInstr("after", *instrIndex, opcode(vm.p.instrs.sl[*instrIndex]), a0, a1)
		} else {
			printInstr("after", *instrIndex, opcode(vm.p.instrs.sl[*instrIndex+1]), a0, a1)
		}
	*/
}

func printInstr(prefix string, idx int32, op opcode, a0, a1 *int32) {
	switch {
	case a0 == nil && a1 == nil:
		fmt.Printf("[%s] instr=%d: %d (%s) a0=nil, a1=nil\n", prefix, idx, op, op)
	case a1 == nil:
		fmt.Printf("[%s] instr=%d: %d (%s) a0=%d, a1=nil\n", prefix, idx, op, op, *a0)
	case a0 == nil:
		fmt.Printf("[%s] instr=%d: %d (%s) a0=nil, a1=%d\n", prefix, idx, op, op, *a1)
	default:
		fmt.Printf("[%s] instr=%d: %d (%s) a0=%d, a1=%d\n", prefix, idx, op, op, *a0, *a1)
	}
}
