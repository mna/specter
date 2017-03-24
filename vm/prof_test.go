package vm

import (
	"os"
	"testing"
)

const file = "/Users/martin/go/src/github.com/mna/specter/cmd/examples/loop.vm"

// Benchmark function used for profiling (see ../bench/Makefile : run-prof)
func BenchmarkForProfiling(b *testing.B) {
	for i := 0; i < b.N; i++ {
		vm := New()
		if f, err := os.Open(file); err != nil {
			b.Fatal(err)
		} else {
			vm.Run(f)
		}
	}
}
