package vm

import (
	"io"
)

type VM struct {
	p *program
	m *memory
}

func New() *VM {
	return &VM{&program{}, newMemory()}
}

// Run executes the vm bytecode read by the reader.
func (vm *VM) Run(r io.Reader) error {
	// TODO : Lex and Parse content (use something similar to Go's scanner and parser)

	for i := vm.p.start; vm.p.instrs.sl[i] != -1; i++ {
		vm.runInstruction(&i)
	}
	return nil
}

func (vm *VM) runInstruction(instrIndex *int) {
	//a0, a1 := vm.p.args.sl[(*instrIndex)*2], vm.p.args.sl[(*instrIndex)*2+1]
	switch vm.p.instrs.sl[*instrIndex] {
	case 0x0:
		// no-op
	case 0x1:
	// int (unimplemented)
	case 0x2:
		// mov
		//*a0 = *a1
	}
}
