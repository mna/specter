package vm

import (
	"os"
	"path/filepath"
	"testing"
)

var (
	assertOps = map[string][]opcode{
		"nop.vm": []opcode{
			_OP_MOV,
			_OP_NOP,
			_OP_INC,
			_OP_CMP,
			_OP_JL,
			_OP_END,
		},
	}

	assertLabels = map[string]map[string]int32{
		"nop.vm": map[string]int32{
			"loop": 1,
		},
	}
)

func TestParse(t *testing.T) {
	fi := getAllExampleFiles()
	for _, path := range fi {
		base := filepath.Base(path)
		// Ignore err_ files (they would panic anyway)
		if expCodes, ok := assertOps[base]; ok {
			testFile(path, expCodes, assertLabels[base], t)
		}
	}
}

func testFile(path string, expCodes []opcode, expLabels map[string]int32, t *testing.T) {
	base := filepath.Base(path)
	f, err := os.Open(path)
	if err != nil {
		t.Error(err)
	} else {
		vm := New()
		vm.parse(f)

		// Instructions
		if len(vm.p.instrs) != len(expCodes) {
			t.Errorf("file %s: expected %d instructions, got %d", base, len(expCodes), len(vm.p.instrs))
		} else {
			for i, code := range vm.p.instrs {
				if code != expCodes[i] {
					t.Errorf("file %s: expected opcode %s at index %d, got %s", base, expCodes[i], i, code)
				}
			}
		}

		// Labels
		if len(vm.p.labels) != len(expLabels) {
			t.Errorf("file %s: expected %d labels, got %d", base, len(expLabels), len(vm.p.labels))
		} else {
			for k, v := range vm.p.labels {
				exp, ok := expLabels[k]
				if !ok {
					t.Errorf("file %s: unexpected label %s", base, k)
				} else if v != exp {
					t.Errorf("file %s: expected label %s to jump to %d, got %d", base, k, exp, v)
				}
			}
		}
	}
}
