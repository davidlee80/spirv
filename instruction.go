// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// Instruction defines a generic instruction.
type Instruction interface {
	// Opcode returns the opcode for this instruction.
	// It is used by the encoder to find the correct codec in the
	// instruction set library.
	Opcode() uint32
}

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

// InstructionSet maps opcodes to an instruction encoder/decoder.
type InstructionSet struct {
	data map[uint32]Codec
}

// Add adds a new codec to the instruction set.
func (set *InstructionSet) Set(opcode uint32, codec Codec) {
	if set.data == nil {
		set.data = make(map[uint32]Codec)
	}
	set.data[opcode] = codec
}

// Get returns the codec for the given opcode.
// Returns false if it is not in the set.
func (set *InstructionSet) Get(opcode uint32) (Codec, bool) {
	if set.data == nil {
		return Codec{}, false
	}

	c, ok := set.data[opcode]
	return c, ok
}
