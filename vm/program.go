package vm

const (
	_LABELS_CAP = 10 // Labels map capacity increments by this number
	_INSTRS_CAP = 10
)

type program struct {
	start  int32 // Instruction index of the (optional) start label (or 0 - start at beginning)
	instrs *oSlice

	// TODO : extract interface from oSlice, use same pattern for args and labels?
	args   [][2]*int32
	labels map[string]int32
}

// TODO : Both args have to be allocated even if there are no (or only one) arg for the instr

func newProgram() *program {
	return &program{
		0,
		newOSlice(_INSTRS_CAP),
		nil,
		make(map[string]int32, _LABELS_CAP),
	}
}
