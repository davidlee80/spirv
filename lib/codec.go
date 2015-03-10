// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package lib

// Codec defines handlers used to encode or decode
// the a specific type of instruction
type Codec struct {
	// Decoder for an instruction's arguments.
	//
	// The provided set of arguments is guaranteed to have the size
	// defined in the instruction's word count. However, this does not
	// mean it is the amount actually expected by the instruction. So
	// a size check on this slice is warrented. ErrMissingInstructionArgs
	// should returned if this check fails.
	Decode func(argv []uint32) (Instruction, error)

	// Encoder for an instruction
	//
	// The word set must define the FULL instruction. This includes the first
	// word with the word count and opcode.
	//
	// The provided word buffer is large enough to hold a full instruction.
	Encode func(Instruction, []uint32) error
}
