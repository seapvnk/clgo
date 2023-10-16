package clgo

type InputType uint8

const (
	INPUT_TYPE_FILE = iota
	INPUT_TYPE_STRING
)

type Environment struct {
	input     string
	inputType InputType
}

func (env *Environment) LoadString(str string) {
	env.input = str
	env.inputType = INPUT_TYPE_STRING
}

func (env *Environment) LoadFile(str string) {
	env.input = str
	env.inputType = INPUT_TYPE_FILE
}
