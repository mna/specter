package vm

const (
	_STACK_CAP = 10 // Allocate the stack in increments of a capacity of 10 elements
)

type regcode int32

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

// TODO :  Start with this straightforward implementation (using 12 bytes), and later
// try with an interface{} field, and with a [4]byte array.
type register struct {
	i32     int32
	i32_ptr *int32
	i16_h   int16
	i16_l   int16
}

type memory struct {
	// Special-use "registers"
	FLAGS     int
	remainder int

	// A fixed array for the regular registers
	registers [rg_count]register

	// Different approach than TinyVM for the stack, since it can only hold
	// integers, use a slice that grows by `stackCap` increments.
	stack *oSlice
}

func newMemory() *memory {
	var m memory

	// Create the stack with initial capacity
	m.stack = newOSlice(_STACK_CAP)
	return &m
}

func (m *memory) pushStack(i int32) {
	m.stack.addIncr(i)
}

func (m *memory) popStack(i *int32) {
	m.stack.decr()
	*i = m.stack.sl[m.stack.size]
}
