// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// InstructionFunc defines a constructor for an instruction codec.
type InstructionFunc func() Codec

// Instruction defines a generic instruction.
type Instruction interface {
	// Opcode returns the opcode for this instruction.
	// It is used by the encoder to find the correct codec in the
	// instruction set library.
	Opcode() uint32
}

// Codec defines handlers used to encode or decode
// a the operand list for a specific type of instruction
type Codec struct {
	// Decoder for an instruction's arguments.
	//
	// The provided set of arguments is guaranteed to have the size
	// defined in the instruction's word count. However, this does not
	// mean it is the amount actually expected by the instruction. So
	// a size check on this slice is warrented. ErrMissingInstructionArgs
	// should returned if this check fails.
	Decode func(argv []uint32) (Instruction, error)

	// Encoder for an instruction's arguments.
	//
	// The word set being returned must not include the instruction's
	// first word with opcode or word count. This will be generated
	// by the module encoder.
	Encode func(Instruction) ([]uint32, error)
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

// BindDefault adds all default instruction codecs to the given set.
func BindDefault(set *InstructionSet) {
	set.Set(1, NewOpSource())
	set.Set(2, NewOpSourceExtension())
	set.Set(3, NewOpExtension())
	set.Set(4, NewOpExtInstImport())
	set.Set(5, NewOpMemoryModel())
	set.Set(6, NewOpEntryPoint())
	set.Set(7, NewOpExecutionMode())
	set.Set(8, NewOpTypeVoid())
	set.Set(9, NewOpTypeBool())
	set.Set(10, NewOpTypeInt())
	set.Set(11, NewOpTypeFloat())
	set.Set(12, NewOpTypeVector())
	set.Set(44, NewOpExtInst())
	set.Set(218, NewOpCompileFlag())
}
