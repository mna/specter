package vm

import (
	"bufio"
	"io"
	"os"
	"runtime"
	"strconv"
	"sync"
)

// The VM, with its program and memory abstractions.
type VM struct {
	p *program
	m *memory
	b *bufio.Writer
}

// Create a new VM.
func New() *VM {
	return NewWithWriter(os.Stdout)
}

// Create a new VM with the specified output stream.
func NewWithWriter(w io.Writer) *VM {
	runtime.GOMAXPROCS(4)
	return &VM{newProgram(), newMemory(), bufio.NewWriter(w)}
}

func printOutput(vm *VM, wg *sync.WaitGroup, c <-chan int32) {
	for v := range c {
		vm.b.WriteString(strconv.FormatInt(int64(v), 10))
		vm.b.WriteByte('\n')
	}
	vm.b.Flush()
	wg.Done()
}

// Run executes the vm bytecode read by the reader.
func (vm *VM) Run(r io.Reader) {
	var i int32
	var c chan int32
	var wg sync.WaitGroup

	// Parse the content to execute.
	vm.parse(r)

	// Create the output channel and start the goroutine
	// Tried with buffer of 1000, 10000 and 100000, not much difference in performance.
	c = make(chan int32, 100000)
	wg.Add(1)
	go printOutput(vm, &wg, c)

	// Execution loop.
	for i = vm.p.start; vm.p.instrs[i] != _OP_END; i++ {
		vm.runInstruction(&i, c)
	}
	close(c)
	wg.Wait()
}

// Run a single instruction.
func (vm *VM) runInstruction(instrIndex *int32, c chan<- int32) {
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
		c <- *a0
	}
}
