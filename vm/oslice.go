package vm

// "Optimized" slice, for the needs of the VM. Basically, the length is available
// by reading a field (no len() call), and when it needs to be expanded, it expands
// by increments of its original capacity.
type oSlice struct {
	sl     []int32
	size   int32
	expand int32
}

// Create an optimized slice.
func newOSlice(c int32) *oSlice {
	if c <= 0 {
		panic("the capacity must be greater than zero")
	}
	return &oSlice{
		make([]int32, c, c),
		0,
		c,
	}
}

// Add the value to the slice, and increment the length.
func (o *oSlice) addIncr(val int32) {
	if o.size > 0 && (o.size%o.expand) == 0 {
		// Need to allocate more memory
		o.alloc()
	}

	o.sl[o.size] = val
	o.size++
}

// Decrement the length (used by stacks)
func (o *oSlice) decr() {
	o.size--
}

// Expand the slice (allocate memory).
func (o *oSlice) alloc() {
	// Allocate by increments of o.expand, instead of for each element once the initial
	// capacity is reached
	newSz := o.size + o.expand
	sl := make([]int32, newSz, newSz)
	copy(sl, o.sl)
	o.sl = sl
}
