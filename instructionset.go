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

// Bind registers the given instruction.
//
// This call panics if the instruction type defined by the constructor
// is not a pointer type. Additionally, this panics if there is already an
// entry for the instruction's opcode.
//
// All instructions are meant to be registered during package initialisation
func Bind(fun InstructionFunc) {
	obj := fun()
	rv := reflect.ValueOf(obj)

	if rv.Kind() != reflect.Ptr {
		panic(ErrInstructionNotPointer)
	}

	opcode := obj.Opcode()

	instructions.Lock()
	_, ok := instructions.data[opcode]
	if ok {
		panic(ErrDuplicateInstruction)
	}

	instructions.data[opcode] = fun
	instructions.Unlock()
}

// Len returns the number of registered instructions.
func (set *InstructionSet) Len() int {
	set.RLock()
	v := len(set.data)
	set.RUnlock()
	return v
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
	set.Lock()
	set.data = make(map[uint32]InstructionFunc)
	set.Unlock()
}
