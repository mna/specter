package vm

// "Optimized" slice, for the needs of the VM. Basically, the length is available
// by reading a field (no len() call), and when it needs to be expanded, it expands
// by increments of its original capacity.
type oSlice struct {
	sl     []int32
	size   int
	expand int
}

func newOSlice(cap int) *oSlice {
	if cap <= 0 {
		panic("the capacity must be greater than zero")
	}
	return &oSlice{
		make([]int32, 0, cap),
		0,
		cap,
	}
}

func (o *oSlice) addIncr(val int32) {
	if o.size > 0 && o.size%o.expand == 0 {
		// Need to allocate more memory
		o.alloc()
	}

	o.sl[o.size] = val
	o.size++
}

func (o *oSlice) decr() {
	o.size--
}

func (o *oSlice) alloc() {
	// Allocate by increments of o.expand, instead of for each element once the initial
	// capacity is reached
	sl := make([]int32, 0, o.size+o.expand)
	copy(sl, o.sl)
	o.sl = sl
}
