package clgo

import "testing"

// Should load string as environment input.
func Test_Environment_LoadString(t *testing.T) {
	str := "(+ 1 2 3)"

	var env Environment
	env.LoadString(str)

	if env.input != str {
		t.Error("Input not loaded.")
	}

	if env.inputType != INPUT_TYPE_STRING {
		t.Error("Input type must be INPUT_TYPE_STRING")
	}
}

// Should set file path as environment input.
func Test_Environment_LoadFile(t *testing.T) {
	str := "(+ 1 2 3)"

	var env Environment
	env.LoadFile(str)

	if env.input != str {
		t.Error("Input not loaded.")
	}

	if env.inputType != INPUT_TYPE_FILE {
		t.Error("Input type must be INPUT_TYPE_FILE")
	}
}
