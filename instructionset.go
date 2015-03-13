// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import (
	"reflect"
	"sync"
)

// InstructionFunc defines a constructor for an instruction.
type InstructionFunc func() Instruction

// InstructionSet maps opcodes to an instruction  constructor.
type InstructionSet struct {
	sync.RWMutex
	data map[uint32]InstructionFunc
}

// Global, internal instruction set.
// This has instructions registered atomically during init.
var instructions = InstructionSet{
	data: make(map[uint32]InstructionFunc),
}

// Bind registers the given codec.
//
// This call panics if the instruction type defined by the codec
// is not a pointer type.
func Bind(fun InstructionFunc) {
	obj := fun()
	rv := reflect.ValueOf(obj)

	if rv.Kind() != reflect.Ptr {
		panic(ErrInstructionNotPointer)
	}

	instructions.Lock()
	instructions.data[obj.Opcode()] = fun
	instructions.Unlock()
}

// Get returns the codec for the given opcode.
// Returns false if it is not in the set.
func (set *InstructionSet) Get(opcode uint32) (InstructionFunc, bool) {
	instructions.RLock()
	constructor, ok := set.data[opcode]
	instructions.RUnlock()
	return constructor, ok
}

// Clear unbinds all instructions.
func (set *InstructionSet) Clear() {
	set.data = make(map[uint32]InstructionFunc)
}
