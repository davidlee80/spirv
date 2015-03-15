// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "reflect"

type layout struct {
	target     uint32
	preceeding []uint32
}

// logicalLayout holds descriptors defining the order in which instructions
// must appear in a module.
var logicalLayout = []layout{
	{
		target: opcodeSource,
	},
	{
		target:     opcodeSourceExtension,
		preceeding: []uint32{opcodeSource},
	},
}

// verifyInstruction iterates over the instruction fields and calls
// Verify() on all of those which implement it. We then call Verify
// on the instruction itself.
func verifyInstruction(i Instruction) error {
	rv := reflect.ValueOf(i)
	rv = reflect.Indirect(rv)

	err := verifyValue(rv)
	if err != nil {
		return err
	}

	return i.Verify()
}

func verifyValue(rv reflect.Value) error {
	switch rv.Kind() {
	case reflect.Struct:
		return verifyStruct(rv)
	case reflect.Slice:
		return verifySlice(rv)
	}

	data := rv.Interface()
	verifiable, ok := data.(Verifiable)
	if ok {
		return verifiable.Verify()
	}

	return nil
}

func verifyStruct(rv reflect.Value) error {
	for i := 0; i < rv.NumField(); i++ {
		err := verifyValue(rv.Field(i))
		if err != nil {
			return err
		}
	}

	return nil
}

func verifySlice(rv reflect.Value) error {
	for i := 0; i < rv.Len(); i++ {
		err := verifyValue(rv.Index(i))
		if err != nil {
			return err
		}
	}

	return nil
}
