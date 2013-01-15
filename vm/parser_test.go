package vm

import (
	"os"
	"path/filepath"
	"testing"
)

var (
	assertOps = map[string][]opcode{
		"euler1.vm": []opcode{
			_OP_MOV,
			_OP_MOV,
			_OP_MOV,
			_OP_MOD,
			_OP_REM,
			_OP_CMP,
			_OP_JNE,
			_OP_ADD,
			_OP_JE,
			_OP_MOV,
			_OP_MOD,
			_OP_REM,
			_OP_CMP,
			_OP_JNE,
			_OP_ADD,
			_OP_INC,
			_OP_CMP,
			_OP_JL,
			_OP_PRN,
			_OP_END,
		},
		"euler1_nodiv.vm": []opcode{
			_OP_MOV,
			_OP_XOR,
			_OP_MOV,
			_OP_ADD,
			_OP_ADD,
			_OP_CMP,
			_OP_JL,
			_OP_MOV,
			_OP_ADD,
			_OP_ADD,
			_OP_CMP,
			_OP_JL,
			_OP_MOV,
			_OP_ADD,
			_OP_ADD,
			_OP_CMP,
			_OP_JL,
			_OP_PRN,
			_OP_END,
		},
		"euler2.vm": []opcode{
			_OP_MOV,
			_OP_MOV,
			_OP_MOV,
			_OP_ADD,
			_OP_ADD,
			_OP_MOV,
			_OP_AND,
			_OP_CMP,
			_OP_JNE,
			_OP_ADD,
			_OP_MOV,
			_OP_AND,
			_OP_CMP,
			_OP_JNE,
			_OP_ADD,
			_OP_CMP,
			_OP_JG,
			_OP_CMP,
			_OP_JL,
			_OP_PRN,
			_OP_END,
		},
		"euler7": []opcode{
			_OP_MOV,
			_OP_MOV,
			_OP_CMP,
			_OP_JE,
			_OP_MOD,
			_OP_REM,
			_OP_CMP,
			_OP_JE,
			_OP_INC,
			_OP_JMP,
			_OP_INC,
			_OP_CMP,
			_OP_JE,
			_OP_INC,
			_OP_JMP,
			_OP_PRN,
			_OP_END,
		},
		"fact.vm": []opcode{
			_OP_PUSH,
			_OP_PUSH,
			_OP_PUSH,
			_OP_MOV,
			_OP_MOV,
			_OP_CMP,
			_OP_JLE,
			_OP_MOV,
			_OP_DEC,
			_OP_CALL,
			_OP_MUL,
			_OP_POP,
			_OP_POP,
			_OP_POP,
			_OP_RET,
			_OP_MOV,
			_OP_INC,
			_OP_CALL,
			_OP_PRN,
			_OP_CMP,
			_OP_JL,
			_OP_END,
		},
		"fib.vm": []opcode{
			_OP_MOV,
			_OP_MOV,
			_OP_ADD,
			_OP_ADD,
			_OP_PRN,
			_OP_PRN,
			_OP_CMP,
			_OP_JL,
			_OP_CMP,
			_OP_JG,
			_OP_END,
		},
		"jsr.vm": []opcode{
			_OP_PUSH,
			_OP_MOV,
			_OP_PRN,
			_OP_POP,
			_OP_RET,
			_OP_MOV,
			_OP_CALL,
			_OP_MOV,
			_OP_CALL,
			_OP_END,
		},
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
		"euler1.vm": map[string]int32{
			"start": 0,
			"L0":    2,
			"L1":    9,
			"check": 15,
		},
		"euler1_nodiv.vm": map[string]int32{
			"start": 0,
			"loop0": 3,
			"loop1": 8,
			"loop2": 13,
		},
		"euler2.vm": map[string]int32{
			"start": 0,
			"loop":  3,
			"L0":    10,
			"L1":    15,
			"end":   19,
		},
		"euler7.vm": map[string]int32{
			"start":       0,
			"checkPrime":  1,
			"checkFactor": 2,
			"primeFound":  10,
			"nextPrime":   13,
			"printResult": 15,
		},
		"fact.vm": map[string]int32{
			"fact":     0,
			"end_fact": 11,
			"start":    15,
			"loop":     16,
		},
		"fib.vm": map[string]int32{
			"start": 0,
			"loop":  2,
			"end":   10,
		},
		"jsr.vm": map[string]int32{
			"print_eax": 0,
			"start":     5,
		},
		"nop.vm": map[string]int32{
			"loop": 1,
		},
	}
)

// Test all the .vm files that don't fail.
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

// Validate the start position, the instructions, and the labels (the arguments
// will be validated by the behaviour tests, not the parser).
func testFile(path string, expCodes []opcode, expLabels map[string]int32, t *testing.T) {
	base := filepath.Base(path)
	f, err := os.Open(path)
	if err != nil {
		t.Error(err)
	} else {
		vm := New()
		vm.parse(f)

		// Assert the start position
		if strt, ok := expLabels["start"]; ok {
			if vm.p.start != strt {
				t.Errorf("file %s: expected start instruction to be %d, got %d", base, strt, vm.p.start)
			}
		}

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
