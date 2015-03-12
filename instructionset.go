// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import (
	"fmt"
	"sync"
)

// Copy copies the given slice as-is.
func Copy(v []uint32) []uint32 {
	if len(v) == 0 {
		// `nil` is not the same as `make([]T, 0)`. So be explicit, lest
		// our unit tests complain when comparing the return value of
		// this function with a nil slice in the comparison target.
		//
		// Ref: http://play.golang.org/p/5vO9v00gl7
		return nil
	}

	new := make([]uint32, len(v))
	copy(new, v)
	return v
}

// InstructionSet maps opcodes to an instruction encoder/decoder.
type InstructionSet struct {
	sync.RWMutex
	data map[uint32]Codec // List of registered instruction codecs.
	buf  []uint32
}

// Global, internal instruction set.
// This has instructions registered atomically during init.
var instructions = InstructionSet{
	data: make(map[uint32]Codec),
	buf:  make([]uint32, 16),
}

// Bind registers the given codec for the specified opcode.
func Bind(opcode uint32, codec Codec) {
	instructions.Lock()
	instructions.data[opcode] = codec
	instructions.Unlock()
}

// Decode decodes the given sequence of words into an Instruction.
// This requires the given set of words to have the exact length needed
// to hold the entire instruction, but no more.
//
// Returns an error if there is no matching instruction or the
// decoding failed.
func Decode(words []uint32) (Instruction, error) {
	if len(words) == 0 {
		return nil, ErrUnexpectedEOF
	}

	wordCount := words[0] >> 16
	opcode := words[0] & 0xffff

	if wordCount == 0 {
		return nil, ErrInvalidInstructionSize
	}

	instructions.RLock()
	defer instructions.RUnlock()

	codec, ok := instructions.data[opcode]
	if !ok {
		return nil, fmt.Errorf("unknown instruction: %08x", opcode)
	}

	return codec.Decode(words[1:])
}

// Encode encodes the given instruction into a list of words.
// Returns an error if there is no matching encoder for the
// instruction.
func Encode(i Instruction) ([]uint32, error) {
	instructions.Lock()
	defer instructions.Unlock()

	opcode := i.Opcode()
	codec, ok := instructions.data[opcode]
	if !ok {
		return nil, fmt.Errorf("unknown instruction: %08x", opcode)
	}

	// Make sure the word buffer has the correct size.
	size := EncodedLen(i)
	if size > len(instructions.buf) {
		instructions.buf = make([]uint32, size)
	}

	// Encode the instruction arguments.
	argc, err := codec.Encode(i, instructions.buf[1:])
	if err != nil {
		return nil, err
	}

	argc++

	// Set the first instruction word.
	instructions.buf[0] = EncodeOpcode(argc, i.Opcode())
	return instructions.buf[:argc], nil
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
