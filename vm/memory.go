package vm

const (
	_STACK_CAP = 10 // Allocate the stack in increments of a capacity of 10 elements
)

// TODO :  Start with this straightforward implementation (using 12 bytes), and later
// try with an interface{} field, and with a [4]byte array.

// i32_ptr is not needed, TinyVM uses it for registers that hold stack limits.
// i16 (high and low) is unused in TinyVM, dropped from this implementation.
// So register is basically an int32!
// Don't even use a type, causes problems to store in args which is *int32
// type register int32 

type memory struct {
	// Special-use "registers"
	FLAGS     int32
	remainder int32

	// A fixed array for the regular registers
	registers [rg_count]int32

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
