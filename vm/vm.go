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
	// TODO : Init program and memory here instead of in new vm
	vm.parse(r)
	for i := vm.p.start; vm.p.instrs.sl[i] != int32(_OP_END); i++ {
		vm.runInstruction(&i)
	}
}

func (vm *VM) runInstruction(instrIndex *int) {
	a0, a1 := vm.p.args[*instrIndex][0], vm.p.args[*instrIndex][1]

	switch {
	case a0 == nil && a1 == nil:
		fmt.Printf("instr=%d: %x (%s)\n", *instrIndex, vm.p.instrs.sl[*instrIndex], opsRev[vm.p.instrs.sl[*instrIndex]])
	case a1 == nil:
		fmt.Printf("instr=%d: %x (%s) a0=%d\n", *instrIndex, vm.p.instrs.sl[*instrIndex], opsRev[vm.p.instrs.sl[*instrIndex]], *a0)
	default:
		fmt.Printf("instr=%d: %x (%s) a0=%d, a1=%d\n", *instrIndex, vm.p.instrs.sl[*instrIndex], opsRev[vm.p.instrs.sl[*instrIndex]], *a0, *a1)
	}
	/*
		//a0, a1 := vm.p.args.sl[(*instrIndex)*2], vm.p.args.sl[(*instrIndex)*2+1]
		switch opcode(vm.p.instrs.sl[*instrIndex]) {
		case _OP_NOP:
		case _OP_INT:
		case _OP_MOV:
			//*a0 = *a1
		}
	*/
}
