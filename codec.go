// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// Codec defines handlers used to encode or decode
// a specific type of instruction
type Codec struct {
	// Decoder for an instruction's arguments.
	//
	// The provided set of arguments is guaranteed to have the size
	// defined in the instruction's word count. However, this does not
	// mean it is the amount actually expected as per the specification.
	// So a size check on this slice is warrented.
	//
	// ErrMissingInstructionArgs should be returned if this check fails.
	Decode func(argv []uint32) (Instruction, error)

	// Encoder for instruction arguments.
	//
	// The provided word set must encode all arguments. It can be assumed
	// to be large enough to hold a full instruction.
	//
	// The function must return the number of words it has written.
	Encode func(Instruction, []uint32) (uint32, error)
}
