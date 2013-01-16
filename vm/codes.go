package vm

// The operation code type.
type opcode int32

// List of available opcodes.
const (
	_OP_END opcode = iota - 1
	_OP_NOP
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

var (
	// Reverse lookup of opcodes (opcode index = opcode string name)
	opsRev = []string{
		_OP_NOP:   "nop",
		_OP_INT:   "int",
		_OP_MOV:   "mov",
		_OP_PUSH:  "push",
		_OP_POP:   "pop",
		_OP_PUSHF: "pushf",
		_OP_POPF:  "popf",
		_OP_INC:   "inc",
		_OP_DEC:   "dec",
		_OP_ADD:   "add",
		_OP_SUB:   "sub",
		_OP_MUL:   "mul",
		_OP_DIV:   "div",
		_OP_MOD:   "mod",
		_OP_REM:   "rem",
		_OP_NOT:   "not",
		_OP_XOR:   "xor",
		_OP_OR:    "or",
		_OP_AND:   "and",
		_OP_SHL:   "shl",
		_OP_SHR:   "shr",
		_OP_CMP:   "cmp",
		_OP_CALL:  "call",
		_OP_JMP:   "jmp",
		_OP_RET:   "ret",
		_OP_JE:    "je",
		_OP_JNE:   "jne",
		_OP_JG:    "jg",
		_OP_JGE:   "jge",
		_OP_JL:    "jl",
		_OP_JLE:   "jle",
		_OP_PRN:   "prn",
	}

	// Lookup of opcodes (opcode string name key = opcode integer value)
	opsMap = map[string]opcode{
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
)

// Stringer implementation for debugging purpose.
func (o opcode) String() string {
	return opsRev[o]
}

// Register code type.
type regcode int32

// List of available register codes.
const (
	_RG_EAX regcode = iota
	_RG_EBX
	_RG_ECX
	_RG_EDX
	_RG_ESI
	_RG_EDI
	_RG_ESP
	_RG_EBP
	_RG_EIP
	_RG_R08
	_RG_R09
	_RG_R10
	_RG_R11
	_RG_R12
	_RG_R13
	_RG_R14
	_RG_R15
	rg_count
)

// Lookup map of registers (register string name key = register integer code)
var rgsMap = map[string]regcode{
	"eax": _RG_EAX,
	"ebx": _RG_EBX,
	"ecx": _RG_ECX,
	"edx": _RG_EDX,
	"esi": _RG_ESI,
	"edi": _RG_EDI,
	"esp": _RG_ESP,
	"ebp": _RG_EBP,
	"eip": _RG_EIP,
	"r08": _RG_R08,
	"r09": _RG_R09,
	"r10": _RG_R10,
	"r11": _RG_R11,
	"r12": _RG_R12,
	"r13": _RG_R13,
	"r14": _RG_R14,
	"r15": _RG_R15,
}
