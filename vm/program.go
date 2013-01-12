package vm

const (
	_LABELS_CAP = 100             // Labels map initial capacity, it grows as needed
	_INSTRS_CAP = _LABELS_CAP * 3 // n times more instrs than labels
)

type program struct {
	start  int32 // Instruction index of the (optional) start label (or 0 - start at beginning)
	instrs *oSlice

	// TODO : extract interface from oSlice, use same pattern for args and labels?
	args   [][_MAX_ARGS]*int32
	labels map[string]int32
}

func newProgram() *program {
	return &program{
		0,
		newOSlice(_INSTRS_CAP),
		nil,
		make(map[string]int32, _LABELS_CAP),
	}
}
