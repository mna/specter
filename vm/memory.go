package vm

const (
	// Since the stack can grow, start smaller than TinyVM (256 bytes)
	_STACK_CAP = 256 / 4

	// The heap is 64Mb in TinyVM, but the stack uses 2Mb of it. Since it is implemented
	// here as an array of int32, divide by 4 bytes, so that this is the number of elements
	// in the heap array.
	_HEAP_CAP = (62 * 1024 * 1024) / 4
)

// Register: changes vs TinyVM
//
// i32_ptr is not needed, TinyVM uses it for registers that hold stack limits, not
// needed given how the stack is implemented in specter (not a pointer into the
// memory used for the heap).
// i16 (high and low) is unused in TinyVM, dropped from this implementation.
// So register is basically an int32!

// The memory struct holds the memory representation of the VM (the registers,
// the stack and the heap).
type memory struct {
	// Special-use "registers"
	// FLAGS is similar to x86 register: 
	// 0x1 = equal
	// 0x2 = greater
	FLAGS     int32
	remainder int32

	// A fixed array for the regular registers
	registers []int32

	// Different approach than TinyVM for the stack, since it can only hold
	// integers, use a slice.
	stack    []int32
	stackPos int32
	heap     []int32
}

// Create a memory struct
func newMemory() *memory {
	var m memory

	// Create the registers
	m.registers = make([]int32, rg_count)
	// Create the stack with initial capacity
	m.stack = make([]int32, 0, _STACK_CAP)
	// The heap is a fixed amount of memory, set the length so that it is completely indexable
	m.heap = make([]int32, _HEAP_CAP)
	return &m
}

// Push value on the stack.
func (m *memory) pushStack(i int32) {
	m.stack = append(m.stack, i)
	m.stackPos++
}

// Pop value from the stack.
func (m *memory) popStack(i *int32) {
	m.stackPos--
	*i = m.stack[m.stackPos]
	m.stack = m.stack[:m.stackPos]
}
