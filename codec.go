// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// Codec defines handlers used to encode or decode
// a specific type of instruction
type Codec struct {
	// New should return a new, default version of the instruction.
	// The value returned by this must be a pointer to a struct.
	New func() Instruction
}
