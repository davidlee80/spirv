// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package lib

import "errors"

var ErrMissingInstructionArgs = errors.New("insufficient instruction arguments")

// Instruction defines a generic instruction.
type Instruction interface {
	// Opcode returns the opcode for this instruction.
	// It is used by the encoder to find the correct codec in the
	// instruction set library.
	Opcode() uint32
}
