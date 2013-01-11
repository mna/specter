package vm

const (
	_STACK_CAP = 10 // Allocate the stack in increments of a capacity of 10 elements
)

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
	FLAGS     int32
	remainder int32

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
