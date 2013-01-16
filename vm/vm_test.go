package vm

import (
	"bytes"
	"testing"
)

type testInfo struct {
	op   opcode
	name string
	code string
	exp  string
}

var tests = [...]testInfo{
	testInfo{
		_OP_MOV,
		"",
		`
    mov eax, 17
    prn eax
    `,
		"17\n",
	},
	testInfo{
		_OP_PUSH,
		"",
		`
    push 1
    push 2
    push 3
    pop eax
    prn eax
    pop eax
    prn eax
    pop eax
    prn eax
    `,
		"3\n2\n1\n",
	},
	testInfo{
		_OP_PUSH,
		"",
		`
    push 10
    push 20
    push 30
    pop eax
    pop ebx
    pop ecx
    prn eax
    prn ebx
    prn ecx
    `,
		"30\n20\n10\n",
	},
	testInfo{
		_OP_PUSHF,
		"",
		`
    mov eax, 10
    cmp eax, 3
    pushf
    pop ecx
    prn ecx
    `,
		"2\n",
	},
	testInfo{
		_OP_POPF,
		"",
		`
    mov eax, 10
    cmp eax, 30
    pushf
    popf eax
    prn eax
    `,
		"0\n",
	},
	testInfo{
		_OP_INC,
		"",
		`
    mov eax, 10
    inc eax
    prn eax
    `,
		"11\n",
	},
	testInfo{
		_OP_DEC,
		"",
		`
    mov edx, 10
    dec edx
    prn edx
    `,
		"9\n",
	},
	testInfo{
		_OP_ADD,
		"",
		`
    mov eax, 10
    mov ebx, 7
    add eax, ebx
    prn eax
    `,
		"17\n",
	},
	testInfo{
		_OP_SUB,
		"",
		`
    mov eax, 10
    mov ebx, 7
    sub eax, ebx
    prn eax
    `,
		"3\n",
	},
	testInfo{
		_OP_MUL,
		"",
		`
    mov eax, 10
    mov ebx, 7
    mul eax, ebx
    prn eax
    `,
		"70\n",
	},
	testInfo{
		_OP_DIV,
		"",
		`
    mov eax, 30
    mov ebx, 7
    div eax, ebx
    prn eax
    `,
		"4\n",
	},
	testInfo{
		_OP_MOD,
		"",
		`
    mov eax, 30
    mov ebx, 7
    mod eax, ebx
    rem edx
    prn edx
    `,
		"2\n",
	},
	testInfo{
		_OP_NOT,
		"",
		`
    mov eax, 1
    not eax
    prn eax
    `,
		"-2\n",
	},
	testInfo{
		_OP_XOR,
		"",
		`
    mov eax, 137
    mov ebx, 145
    xor eax, ebx
    prn eax
    `,
		"24\n",
	},
	testInfo{
		_OP_OR,
		"",
		`
    mov eax, 195
    mov ebx, 224
    or eax, ebx
    prn eax
    `,
		"227\n",
	},
	testInfo{
		_OP_AND,
		"",
		`
    mov eax, 232
    mov ebx, 42
    and eax, ebx
    prn eax
    `,
		"40\n",
	},
	testInfo{
		_OP_SHL,
		"",
		`
    mov eax, 1045
    mov ebx, 3
    shl eax, ebx
    prn eax
    `,
		"8360\n",
	},
	testInfo{
		_OP_SHR,
		"",
		`
    mov eax, 1045
    mov ebx, 3
    shr eax, ebx
    prn eax
    `,
		"130\n",
	},
	testInfo{
		_OP_CMP,
		"",
		`
    mov eax, 12
    mov ebx, 9
    cmp eax, ebx
    pushf
    popf ecx
    prn ecx
    `,
		"2\n",
	},
	testInfo{
		_OP_CALL, // Tests RET too
		"",
		`
    call 3
    prn 1
    jmp 5
    prn 3
    ret
    nop
    `,
		"3\n1\n",
	},
	testInfo{
		_OP_JMP,
		"",
		`
    jmp 2
    prn 1
    prn 3
    `,
		"3\n",
	},
	testInfo{
		_OP_JE,
		"",
		`
    mov eax 4
    mov ebx 4
    cmp eax, ebx
    je 5
    prn 1
    prn 2
    `,
		"2\n",
	},
	testInfo{
		_OP_JE,
		"false",
		`
    mov eax 5
    mov ebx 4
    cmp eax, ebx
    je 5
    prn 1
    prn 2
    `,
		"1\n2\n",
	},
	testInfo{
		_OP_JNE,
		"",
		`
    mov eax 12
    mov ebx 4
    cmp eax, ebx
    jne 5
    prn 1
    prn 2
    `,
		"2\n",
	},
	testInfo{
		_OP_JNE,
		"false",
		`
    mov eax 4
    mov ebx 4
    cmp eax, ebx
    jne 5
    prn 1
    prn 2
    `,
		"1\n2\n",
	},
	testInfo{
		_OP_JG,
		"",
		`
    mov eax 6
    mov ebx 4
    cmp eax, ebx
    jg 5
    prn 1
    prn 2
    `,
		"2\n",
	},
	testInfo{
		_OP_JG,
		"false, equal",
		`
    mov eax 4
    mov ebx 4
    cmp eax, ebx
    jg 5
    prn 1
    prn 2
    `,
		"1\n2\n",
	},
	testInfo{
		_OP_JG,
		"false, lower",
		`
    mov eax 2
    mov ebx 4
    cmp eax, ebx
    jg 5
    prn 1
    prn 2
    `,
		"1\n2\n",
	},
	testInfo{
		_OP_JGE,
		"",
		`
    mov eax 6
    mov ebx 4
    cmp eax, ebx
    jge 5
    prn 1
    prn 2
    `,
		"2\n",
	},
	testInfo{
		_OP_JGE,
		"true, equal",
		`
    mov eax 4
    mov ebx 4
    cmp eax, ebx
    jge 5
    prn 1
    prn 2
    `,
		"2\n",
	},
	testInfo{
		_OP_JGE,
		"false",
		`
    mov eax 2
    mov ebx 4
    cmp eax, ebx
    jge 5
    prn 1
    prn 2
    `,
		"1\n2\n",
	},
	testInfo{
		_OP_JL,
		"",
		`
    mov eax 2
    mov ebx 4
    cmp eax, ebx
    jl 5
    prn 1
    prn 2
    `,
		"2\n",
	},
	testInfo{
		_OP_JL,
		"false, equal",
		`
    mov eax 4
    mov ebx 4
    cmp eax, ebx
    jl 5
    prn 1
    prn 2
    `,
		"1\n2\n",
	},
	testInfo{
		_OP_JL,
		"false, greater",
		`
    mov eax 6
    mov ebx 4
    cmp eax, ebx
    jl 5
    prn 1
    prn 2
    `,
		"1\n2\n",
	},
	testInfo{
		_OP_JLE,
		"",
		`
    mov eax 2
    mov ebx 4
    cmp eax, ebx
    jle 5
    prn 1
    prn 2
    `,
		"2\n",
	},
	testInfo{
		_OP_JLE,
		"true, equal",
		`
    mov eax 4
    mov ebx 4
    cmp eax, ebx
    jle 5
    prn 1
    prn 2
    `,
		"2\n",
	},
	testInfo{
		_OP_JLE,
		"false",
		`
    mov eax 6
    mov ebx 4
    cmp eax, ebx
    jle 5
    prn 1
    prn 2
    `,
		"1\n2\n",
	},
}

func TestCodes(t *testing.T) {
	for _, ti := range tests {
		testCode(ti, t)
	}
}

func testCode(ti testInfo, t *testing.T) {
	// Arrange
	b := new(bytes.Buffer)
	r := bytes.NewBufferString(ti.code)
	v := NewWithWriter(b)

	// Act
	v.Run(r)

	// Assert
	res := b.String()
	if res != ti.exp {
		if len(ti.name) > 0 {
			t.Errorf("test %s (%s): expected %s, got %s", ti.op, ti.name, ti.exp, res)
		} else {
			t.Errorf("test %s: expected %s, got %s", ti.op, ti.exp, res)
		}
	}
}
