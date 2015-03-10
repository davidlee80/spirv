// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package lib

import (
	"fmt"

	"github.com/jteeuwen/spirv"
)

// InstructionSet maps opcodes to an instruction encoder/decoder.
type InstructionSet struct {
	data map[uint32]Codec // List of registered instruction codecs.
	buf  [128]uint32      // Encoding scratch buffer.
}

// Decode decodes the given sequence of words in an Instruction.
// Returns an error if there is no matching instruction or the
// loading failed.
func (set *InstructionSet) Decode(words []uint32) (Instruction, error) {
	if len(words) == 0 {
		return nil, spirv.ErrUnexpectedEOF
	}

	wordCount := words[0] >> 16
	opcode := words[0] & 0xffff

	if wordCount == 0 {
		return nil, spirv.ErrInvalidInstructionSize
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
		return nil, spirv.ErrInvalidInstructionSize
	}

	return set.buf[:words], nil
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

	codec, ok := set.data[opcode]
	return codec, ok
}

// BindDefaults loads all default instruction codecs into the set.
func (set *InstructionSet) BindDefaults() {
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
