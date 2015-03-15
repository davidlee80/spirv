// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import (
	"bytes"
	"reflect"
	"testing"
)

var mod = NewModule()

func init() {
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
}

func TestModuleVerify(t *testing.T) {
	err := mod.Verify()
	if err != nil {
		t.Fatal(err)
	}
}

func TestModuleStrip(t *testing.T) {
	mod.Strip()
	if len(mod.Code) != 5 {
		t.Fatalf("Strip error: Expected 5 remaining instructions; have: %d", len(mod.Code))
	}
}

func TestModuleRoundtrip(t *testing.T) {
	var out bytes.Buffer
	err := mod.Save(&out)
	if err != nil {
		t.Fatalf("save: %v", err)
	}

	in := bytes.NewBuffer(out.Bytes())
	modb, err := Load(in)
	if err != nil {
		t.Fatalf("load: %v", err)
	}

	if !reflect.DeepEqual(mod, modb) {
		t.Fatalf("rountrip failure:\nHave: %v\nWant: %v", mod, modb)
	}
}
