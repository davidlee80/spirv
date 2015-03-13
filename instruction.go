// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// Instruction defines a generic instruction.
type Instruction interface {
	// Opcode returns the opcode for this instruction.
	// It is used by the encoder to find the correct codec in the
	// instruction set library.
	Opcode() uint32

	// Verify checks the instruction contents and validates its values.
	// Some instructions have constraints on the values which are allowed in
	// a given field.
	Verify() error
}
