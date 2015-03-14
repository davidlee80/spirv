// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

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
