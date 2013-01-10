package vm

import (
	"io"
)

type opcode int32

const (
	_OP_NOP opcode = iota
	_OP_INT
	_OP_MOV
	_OP_PUSH
	_OP_POP
	_OP_PUSHF
	_OP_POPF
	_OP_INC
	_OP_DEC
	_OP_ADD
	_OP_SUB
	_OP_MUL
	_OP_DIV
	_OP_MOD
	_OP_REM
	_OP_NOT
	_OP_XOR
	_OP_OR
	_OP_AND
	_OP_SHL
	_OP_SHR
	_OP_CMP
	_OP_CALL
	_OP_JMP
	_OP_RET
	_OP_JE
	_OP_JNE
	_OP_JG
	_OP_JGE
	_OP_JL
	_OP_JLE
	_OP_PRN
)

var opsMap = map[string]opcode{
	"nop":   _OP_NOP,
	"int":   _OP_INT,
	"mov":   _OP_MOV,
	"push":  _OP_PUSH,
	"pop":   _OP_POP,
	"pushf": _OP_PUSHF,
	"popf":  _OP_POPF,
	"inc":   _OP_INC,
	"dec":   _OP_DEC,
	"add":   _OP_ADD,
	"sub":   _OP_SUB,
	"mul":   _OP_MUL,
	"div":   _OP_DIV,
	"mod":   _OP_MOD,
	"rem":   _OP_REM,
	"not":   _OP_NOT,
	"xor":   _OP_XOR,
	"or":    _OP_OR,
	"and":   _OP_AND,
	"shl":   _OP_SHL,
	"shr":   _OP_SHR,
	"cmp":   _OP_CMP,
	"call":  _OP_CALL,
	"jmp":   _OP_JMP,
	"ret":   _OP_RET,
	"je":    _OP_JE,
	"jne":   _OP_JNE,
	"jg":    _OP_JG,
	"jge":   _OP_JGE,
	"jl":    _OP_JL,
	"jle":   _OP_JLE,
	"prn":   _OP_PRN,
}

type VM struct {
	p *program
	m *memory
}

func New() *VM {
	return &VM{&program{}, newMemory()}
}

// Run executes the vm bytecode read by the reader.
func (vm *VM) Run(r io.Reader) {
	vm.parse(r)
	for i := vm.p.start; vm.p.instrs.sl[i] != -1; i++ {
		vm.runInstruction(&i)
	}
}

func (vm *VM) runInstruction(instrIndex *int) {
	//a0, a1 := vm.p.args.sl[(*instrIndex)*2], vm.p.args.sl[(*instrIndex)*2+1]
	switch opcode(vm.p.instrs.sl[*instrIndex]) {
	case _OP_NOP:
	case _OP_INT:
	case _OP_MOV:
		//*a0 = *a1
	}
}
