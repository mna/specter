package vm

const (
	// Labels map initial capacity, it grows as needed
	_LABELS_CAP = 100

	// n times more instrs than labels 
	_INSTRS_CAP = _LABELS_CAP * 5
)

// The program struct holds the instructions to execute, the arguments, the labels,
// and the start instruction index.
type program struct {
	start  int32 // Instruction index of the (optional) start label (or 0 - start at beginning)
	instrs []opcode
	args   [][_MAX_ARGS]*int32
	labels map[string]int32
}

// Create a program.
func newProgram() *program {
	return &program{
		0,
		make([]opcode, 0, _INSTRS_CAP),
		nil,
		make(map[string]int32, _LABELS_CAP),
	}
}
