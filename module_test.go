// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import (
	"bytes"
	"reflect"
	"testing"
)

var mod = NewModule()

func TestModuleVerifyLogicalLayout(t *testing.T) {
	mod.Code = []Instruction{
		&OpCompileFlag{},
		&OpSource{},
		&OpMemoryModel{},
		&OpEntryPoint{},
		&OpExecutionMode{},

		&OpFunction{},
		&OpFunctionEnd{},
	}

	err := mod.verifyLogicalLayout()
	if err == nil {
		t.Fatalf("expected failure")
	}

	mod.Code = []Instruction{
		&OpCompileFlag{},
		&OpMemoryModel{},
		&OpEntryPoint{},
		&OpExecutionMode{},

		&OpFunction{},
		&OpFunctionEnd{},
	}

	err = mod.verifyLogicalLayout()
	if err != nil {
		t.Fatal(err)
	}

	// Faulty module with 2 functions.. latter with parameters and a body,
	// but missing the OpFunctionEnd.
	mod.Code = []Instruction{
		&OpMemoryModel{},
		&OpEntryPoint{},
		&OpExecutionMode{},

		&OpFunction{},
		&OpFunctionEnd{},

		&OpFunction{},
		&OpFunctionParameter{},
		&OpFunctionParameter{},
		&OpFunctionParameter{},
		&OpIAdd{},
		&OpIAdd{},
		&OpIAdd{},
		&OpIAdd{},
	}

	err = mod.verifyLogicalLayout()
	if err == nil {
		t.Fatalf("expected failure")
	}

	// Complete module with 2 functions.. latter with parameters and a body.
	mod.Code = []Instruction{
		&OpMemoryModel{},
		&OpEntryPoint{},
		&OpExecutionMode{},

		&OpFunction{},
		&OpFunctionEnd{},

		&OpFunction{},
		&OpFunctionParameter{},
		&OpFunctionParameter{},
		&OpFunctionParameter{},
		&OpIAdd{},
		&OpIAdd{},
		&OpIAdd{},
		&OpIAdd{},
		&OpFunctionEnd{},
	}

	err = mod.verifyLogicalLayout()
	if err != nil {
		t.Fatal(err)
	}
}

func TestModuleStrip(t *testing.T) {
	mod.Code = []Instruction{
		&OpSource{SourceLanguageGLSL, 450},
		&OpCompileFlag{
			Flag: "test",
		},
		&OpMemoryModel{
			AddressingModel: AddressingModeLogical,
			MemoryModel:     MemoryModelGLSL450,
		},
	}

	mod.Strip()

	want := 2
	have := len(mod.Code)
	if have != want {
		t.Fatalf("Strip error: Expected %d remaining instructions; have: %d", want, have)
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
