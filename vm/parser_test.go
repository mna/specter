package vm

import (
	"os"
	"testing"
)

// Reading from file through a bufio is slightly faster than doing a ioutil.ReadAll,
// no optimization there.
func BenchmarkParseBufFile(b *testing.B) {
	b.StopTimer()
	f, err := os.Open("../cmd/examples/fact.vm")
	if err != nil {
		b.Fatal(err)
	}
	vm := New()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		vm.parse(f)
	}
}

/*
func BenchmarkParseReadAll(b *testing.B) {
	b.StopTimer()
	f, err := os.Open("../cmd/examples/fact.vm")
	if err != nil {
		b.Fatal(err)
	}
	vm := New()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		vm.parse2(f)
	}
}
*/
