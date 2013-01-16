package vm

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func getAllExampleFiles() []string {
	const dir string = "../cmd/examples"

	var ret []string

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for _, fi := range files {
		if filepath.Ext(fi.Name()) == ".vm" {
			ret = append(ret, filepath.Join(dir, fi.Name()))
		}
	}
	return ret
}

// Run all examples (*.vm files in ../cmd/examples/)
func TestExamples(t *testing.T) {
	files := getAllExampleFiles()
	for _, fi := range files {
		if f, err := os.Open(fi); err != nil {
			t.Error(err)
		} else {
			runExample(t, f)
		}
	}
}

func runExample(t *testing.T, f *os.File) {
	defer func() {
		e := recover()
		// Error is NOT expected for files not beginning with "err_"
		if e != nil && !strings.HasPrefix(filepath.Base(f.Name()), "err_") {
			t.Error(e)
		}
		// Error is expected for files beginning with "err_"
		if e == nil && strings.HasPrefix(filepath.Base(f.Name()), "err_") {
			t.Error(e)
		}
	}()

	var b bytes.Buffer
	fmt.Printf("Running example %s\n", filepath.Base(f.Name()))
	vm := NewWithWriter(&b)
	// The file execution panics if there is an error
	vm.Run(f)
}
