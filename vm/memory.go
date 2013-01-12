package vm

const (
	// Since the stack can grow, start smaller than TinyVM (256 bytes)
	_STACK_CAP = 256 / 4 // Allocate the stack in increments of a capacity of n elements
)

// i32_ptr is not needed, TinyVM uses it for registers that hold stack limits.
// i16 (high and low) is unused in TinyVM, dropped from this implementation.
// So register is basically an int32!
// Don't even use a type, causes problems to store in args which is *int32 :
//
// type register int32 

type memory struct {
	// Special-use "registers"
	// FLAGS is similar to x86 register: 
	// 0x1 = equal
	// 0x2 = greater
	FLAGS     int32
	remainder int32

	// A fixed array for the regular registers
	registers [rg_count]int32

	// Different approach than TinyVM for the stack, since it can only hold
	// integers, use a slice that grows by `_STACK_CAP` increments.
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
