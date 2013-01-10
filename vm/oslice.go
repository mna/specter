package vm

import (
	"fmt"
)

// "Optimized" slice, for the needs of the VM. Basically, the length is available
// by reading a field (no len() call), and when it needs to be expanded, it expands
// by increments of its original capacity.
type oSlice struct {
	sl     []int32
	size   int
	expand int
}

func newOSlice(c int) *oSlice {
	if c <= 0 {
		panic("the capacity must be greater than zero")
	}
	return &oSlice{
		make([]int32, c, c),
		0,
		c,
	}
}

func (o *oSlice) addIncr(val int32) {
	fmt.Printf("o.size=%d, o.expand=%d\n", o.size, o.expand)
	fmt.Printf("o.len=%d, o.cap=%d\n", len(o.sl), cap(o.sl))
	if o.size > 0 && (o.size%o.expand) == 0 {
		// Need to allocate more memory
		fmt.Println("call alloc")
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
	newSz := o.size + o.expand
	sl := make([]int32, newSz, newSz)
	copy(sl, o.sl)
	o.sl = sl
	fmt.Printf("After alloc, o.len=%d, o.cap=%d\n", len(o.sl), cap(o.sl))
}
