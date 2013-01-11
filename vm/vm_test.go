package vm

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestExamples(t *testing.T) {
	const dir string = "../cmd/examples"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	for _, fi := range files {
		if filepath.Ext(fi.Name()) == ".vm" {
			if f, err := os.Open(filepath.Join(dir, fi.Name())); err != nil {
				t.Error(err)
			} else {
				runExample(t, f)
			}
		}
	}
}

func runExample(t *testing.T, f *os.File) {
	defer func() {
		e := recover()
		if e != nil && !strings.HasPrefix(filepath.Base(f.Name()), "err_") {
			t.Error(e)
		}
		if e == nil && strings.HasPrefix(filepath.Base(f.Name()), "err_") {
			t.Error(e)
		}
	}()

	vm := New()
	vm.Run(f)
}
