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

	/*
		switch {
		case a0 == nil && a1 == nil:
			fmt.Printf("instr=%d: %x (%s)\n", *instrIndex, vm.p.instrs.sl[*instrIndex], opcode(vm.p.instrs.sl[*instrIndex]))
		case a1 == nil:
			fmt.Printf("instr=%d: %x (%s) a0=%d\n", *instrIndex, vm.p.instrs.sl[*instrIndex], opcode(vm.p.instrs.sl[*instrIndex]), *a0)
		case a0 == nil:
			fmt.Printf("instr=%d: %x (%s) a0=nil, a1=%d\n", *instrIndex, vm.p.instrs.sl[*instrIndex], opcode(vm.p.instrs.sl[*instrIndex]), *a1)
		default:
			fmt.Printf("instr=%d: %x (%s) a0=%d, a1=%d\n", *instrIndex, vm.p.instrs.sl[*instrIndex], opcode(vm.p.instrs.sl[*instrIndex]), *a0, *a1)
		}
	*/
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
		fmt.Printf("before increment, content of *a0=%d\n", *a0)
		(*a0)++
		fmt.Printf("after increment, content of *a0=%d\n", *a0)
	case _OP_CMP:
		if *a0 == *a1 {
			vm.m.FLAGS = 0x1
		} else if *a0 > *a1 {
			vm.m.FLAGS = 0x2
		} else {
			vm.m.FLAGS = 0x0
		}
	case _OP_JL:
		if vm.m.FLAGS&0x3 == 0 {
			*instrIndex = *a0 - 1
			fmt.Printf("jump to instr %d+1\n", *instrIndex)
		}
	}
}
