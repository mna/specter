package vm

const (
	labelsCap = 10 // Labels map capacity increments by this number
)

type program struct {
	start int // Instruction index of the (optional) start label
	// TODO : Why not labels["start"]?

	instrs *oSlice
	labels map[string]int
}
