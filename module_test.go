// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "testing"

func TestModule(t *testing.T) {
	mod := NewModule()
	mod.Code = []Instruction{
		&OpSource{SourceLanguageGLSL, 450},
		&OpExtInst{
			ResultType:  1,
			ResultId:    2,
			Set:         3,
			Instruction: 4,
			Operands:    []Id{5, 4, 5},
		},
		&OpFunction{
			ResultType:   0,
			ResultId:     1,
			ControlMask:  FunctionControlMaskInLine,
			FunctionType: 2,
		},
		&OpFunctionParameter{
			ResultType: 0,
			ResultId:   1,
		},
		&OpFunctionParameter{
			ResultType: 0,
			ResultId:   1,
		},
		&OpFunctionEnd{},
	}

	err := mod.Verify()
	if err != nil {
		t.Fatal(err)
	}
}
