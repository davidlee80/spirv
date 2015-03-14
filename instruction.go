// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "reflect"

// Verifiable defines any type which implements verification semantics.
// This may entail simple range checks on numeric fields and constants, or
// as complex as semantic rule validation in a whole module.
type Verifiable interface {
	Verify() error
}

// Instruction defines a generic instruction.
type Instruction interface {
	Verifiable

	// Opcode returns the opcode for this instruction.
	// It is used by the encoder to find the correct codec in the
	// instruction set library.
	Opcode() uint32
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
