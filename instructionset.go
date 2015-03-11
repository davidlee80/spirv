// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "fmt"

// InstructionSet maps opcodes to an instruction encoder/decoder.
type InstructionSet struct {
	data map[uint32]Codec // List of registered instruction codecs.
	buf  [64]uint32       // Encoding scratch buffer.
}

// NewInstructionSet creates a new instruction set with all default
// instructions bound to it.
func NewInstructionSet() *InstructionSet {
	set := &InstructionSet{
		data: make(map[uint32]Codec),
	}

	bindOpNop(set)
	bindOpSource(set)
	bindOpSourceExtension(set)
	bindOpExtension(set)
	bindOpExtInstImport(set)
	bindOpMemoryModel(set)
	bindOpEntryPoint(set)
	bindOpExecutionMode(set)
	bindOpTypeVoid(set)
	bindOpTypeBool(set)
	bindOpTypeInt(set)
	bindOpTypeFloat(set)
	bindOpTypeVector(set)
	bindOpExtInst(set)
	bindOpUndef(set)
	bindOpName(set)
	bindOpMemberName(set)
	bindOpString(set)
	bindOpLine(set)
	bindOpCompileFlag(set)
	return set
}

// Decode decodes the given sequence of words in an Instruction.
// Returns an error if there is no matching instruction or the
// loading failed.
func (set *InstructionSet) Decode(words []uint32) (Instruction, error) {
	if len(words) == 0 {
		return nil, ErrUnexpectedEOF
	}

	wordCount := words[0] >> 16
	opcode := words[0] & 0xffff

	if wordCount == 0 {
		return nil, ErrInvalidInstructionSize
	}

	codec, ok := set.data[opcode]
	if !ok {
		return nil, fmt.Errorf("unknown instruction: %08x", opcode)
	}

	return codec.Decode(words[1:])
}

// Encode encodes the given instruction into a list of words.
// Returns an error if there is no matching encoder for the
// instruction.
//
// The returned slice is valid until the next call to Encode.
func (set *InstructionSet) Encode(i Instruction) ([]uint32, error) {
	opcode := i.Opcode()
	codec, ok := set.data[opcode]
	if !ok {
		return nil, fmt.Errorf("unknown instruction: %08x", opcode)
	}

	// If the encoder fails to write the first word,
	// we want to know about it.
	set.buf[0] = 0

	err := codec.Encode(i, set.buf[:])
	if err != nil {
		return nil, err
	}

	words := (set.buf[0] >> 16) & 0xffff
	if words <= 0 {
		return nil, ErrInvalidInstructionSize
	}

	return set.buf[:words], nil
}

// Add adds a new codec to the instruction set.
func (set *InstructionSet) Set(opcode uint32, codec Codec) {
	set.data[opcode] = codec
}

// Get returns the codec for the given opcode.
// Returns false if it is not in the set.
func (set *InstructionSet) Get(opcode uint32) (Codec, bool) {
	codec, ok := set.data[opcode]
	return codec, ok
}

// Clear unbinds all instructions.
func (set *InstructionSet) Clear() {
	set.data = make(map[uint32]Codec)
}