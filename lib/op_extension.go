// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package lib

import "github.com/jteeuwen/spirv"

// OpExtension defines the OpExtension instruction.
//
// It declares use of an extension to SPIR-V. This allows validation of
// additional instructions, tokens, semantics, etc
type OpExtension string

func (c OpExtension) Opcode() uint32 { return 3 }

// NewOpExtension creates a new codec for the OpExtension instruction.
func NewOpExtension() Codec {
	return Codec{
		Decode: func(argv []uint32) (Instruction, error) {
			if len(argv) < 1 {
				return nil, ErrMissingInstructionArgs
			}

			return OpExtension(
				spirv.DecodeString(argv),
			), nil
		},
		Encode: func(i Instruction, out []uint32) error {
			cf := i.(OpExtension)
			size := spirv.EncodedStringLen(string(cf))
			out[0] = spirv.EncodeOpcode(uint32(size)+1, 3)
			spirv.EncodeString(string(cf), out[1:])
			return nil
		},
	}
}