package vm

const (
	_LABELS_CAP = 10 // Labels map capacity increments by this number
	_INSTRS_CAP = 10
	_ARGS_CAP   = _INSTRS_CAP * 2 // Maximum of 2 args per instruction in this VM
)

type program struct {
	start int // Instruction index of the (optional) start label (or 0 - start at beginning)

	instrs *oSlice
	// TODO: Args has to be pointers to int? cannot work with this oSlice
	args   *oSlice
	labels map[string]int
}

// TODO : Both args have to be allocated even if there are no (or only one) arg for the instr
