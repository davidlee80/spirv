// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "sync"

// InstructionSet maps opcodes to an instruction encoder/decoder.
type InstructionSet struct {
	sync.RWMutex
	data map[uint32]Codec // List of registered instruction codecs.
}

// Global, internal instruction set.
// This has instructions registered atomically during init.
var instructions = InstructionSet{
	data: make(map[uint32]Codec),
}

// Bind registers the given codec for the specified opcode.
func Bind(opcode uint32, codec Codec) {
	instructions.Lock()
	instructions.data[opcode] = codec
	instructions.Unlock()
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
