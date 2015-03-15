// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "reflect"

// InstructionFunc defines a constructor for an instruction.
type instructionFunc func() Instruction

// InstructionSet maps opcodes to an instruction  constructor.
type instructionSet struct {
	data map[uint32]instructionFunc
}

// Global, internal instruction set.
// This has instructions registered atomically during init.
var instructions = instructionSet{
	data: make(map[uint32]instructionFunc),
}

// Bind registers the given instruction.
//
// This call panics if the instruction type defined by the constructor
// is not a pointer type. Additionally, this panics if there is already an
// entry for the instruction's opcode.
//
// All instructions are meant to be registered during package initialisation
func bind(fun instructionFunc) {
	obj := fun()
	rv := reflect.ValueOf(obj)

	if rv.Kind() != reflect.Ptr {
		panic(ErrInstructionNotPointer)
	}

	opcode := obj.Opcode()

	_, ok := instructions.data[opcode]
	if ok {
		panic(ErrDuplicateInstruction)
	}

	instructions.data[opcode] = fun
}

// Len returns the number of registered instructions.
func (set *instructionSet) Len() int { return len(set.data) }

// Get returns the codec for the given opcode.
// Returns false if it is not in the set.
func (set *instructionSet) Get(opcode uint32) (instructionFunc, bool) {
	constructor, ok := set.data[opcode]
	return constructor, ok
}
