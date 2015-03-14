// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "errors"

var (
	ErrUnexpectedEOF          = errors.New("unexpected EOF")
	ErrInvalidInstructionSize = errors.New("instruction has invalid size")
	ErrMissingInstructionArgs = errors.New("insufficient instruction arguments")
	ErrUnacceptable           = errors.New("use of this instruction is not allowed")
	ErrInstructionNotPointer  = errors.New("value from Codec.New is not a pointer type")
	ErrDuplicateInstruction   = errors.New("duplicate opcode being registered")
)
