// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

/*
Package lib defines the instruction set for version 99 of the SPIR-V
specification. It converts sequences of words from a decoder into typed
structures and vice-versa.

	// Create an instruction set and load up all default instructions.
	var lib InstructionSet
	lib.LoadDefaults()

	// Decode a slice of words.
	instr, err := lib.Decode(words)

Similarly, to encode data:

	// Create an instruction set and load up all default instructions.
	var lib InstructionSet
	lib.LoadDefaults()

	// Encode an instruction into a set of words.
	words, err := lib.Encode(instr)


*/
package lib
