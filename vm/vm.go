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

}
