package vm

import (
	"os"
	"testing"
)

const file = "/Users/martin/go/src/github.com/PuerkitoBio/specter/cmd/examples/loop.vm"

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
