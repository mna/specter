package vm

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// The VM, with its program and memory abstractions.
type VM struct {
	p *program
	m *memory
	b *bufio.Writer
}

var opCalls map[opcode]func(*VM, *int32, *int32, *int32)

func init() {
	opCalls = make(map[opcode]func(*VM, *int32, *int32, *int32), opcode_count)
	opCalls[_OP_MOV] = func(_ *VM, i, a0, a1 *int32) {
		*a0 = *a1
	}
	opCalls[_OP_PUSH] = func(vm *VM, i, a0, a1 *int32) {
		vm.m.pushStack(*a0)
	}
	opCalls[_OP_POP] = func(vm *VM, i, a0, a1 *int32) {
		vm.m.popStack(a0)
	}
	opCalls[_OP_PUSHF] = func(vm *VM, i, a0, a1 *int32) {
		vm.m.pushStack(vm.m.FLAGS)
	}
	opCalls[_OP_POPF] = func(vm *VM, i, a0, a1 *int32) {
		vm.m.popStack(a0)
	}
	opCalls[_OP_INC] = func(_ *VM, i, a0, a1 *int32) {
		(*a0)++
	}
	opCalls[_OP_DEC] = func(_ *VM, i, a0, a1 *int32) {
		(*a0)--
	}
	opCalls[_OP_ADD] = func(_ *VM, i, a0, a1 *int32) {
		*a0 += *a1
	}
	opCalls[_OP_SUB] = func(_ *VM, i, a0, a1 *int32) {
		*a0 -= *a1
	}
	opCalls[_OP_MUL] = func(_ *VM, i, a0, a1 *int32) {
		*a0 *= *a1
	}
	opCalls[_OP_DIV] = func(_ *VM, i, a0, a1 *int32) {
		*a0 /= *a1
	}
	opCalls[_OP_CMP] = func(vm *VM, i, a0, a1 *int32) {
		if *a0 == *a1 {
			vm.m.FLAGS = 0x1
		} else if *a0 > *a1 {
			vm.m.FLAGS = 0x2
		} else {
			vm.m.FLAGS = 0x0
		}
	}
	opCalls[_OP_CALL] = func(vm *VM, i, a0, a1 *int32) {
		vm.m.pushStack(*i)
		*i = *a0 - 1
	}
	opCalls[_OP_JMP] = func(_ *VM, i, a0, a1 *int32) {
		*i = *a0 - 1
	}
	opCalls[_OP_JL] = func(vm *VM, i, a0, a1 *int32) {
		if vm.m.FLAGS&0x3 == 0 {
			*i = *a0 - 1
		}
	}
	opCalls[_OP_JLE] = func(vm *VM, i, a0, a1 *int32) {
		if vm.m.FLAGS&0x2 == 0 {
			*i = *a0 - 1
		}
	}
	opCalls[_OP_PRN] = func(vm *VM, i, a0, a1 *int32) {
		vm.b.WriteString(strconv.FormatInt(int64(*a0), 10))
		vm.b.WriteRune('\n')
	}
	opCalls[_OP_RET] = func(vm *VM, i, a0, a1 *int32) {
		vm.m.popStack(i)
	}
}

// Create a new VM.
func New() *VM {
	return &VM{newProgram(), newMemory(), bufio.NewWriter(os.Stdout)}
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
	a0, a1 := vm.p.args[*instrIndex][0], vm.p.args[*instrIndex][1]

	//printInstr("before", *instrIndex, opcode(vm.p.instrs[*instrIndex]), a0, a1)

	f, ok := opCalls[vm.p.instrs[*instrIndex]]
	if ok {
		f(vm, instrIndex, a0, a1)
	} else {
		panic(fmt.Sprintf("missing function for opcode %s", vm.p.instrs[*instrIndex]))
	}

	/*
		if *instrIndex >= 0 {
			printInstr("after", *instrIndex, opcode(vm.p.instrs[*instrIndex]), a0, a1)
		} else {
			printInstr("after", *instrIndex, opcode(vm.p.instrs[*instrIndex+1]), a0, a1)
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
