// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package lib

import "github.com/jteeuwen/spirv"

// OpCompileFlag represents the OpCompileFlag instruction.
type OpCompileFlag string

func (c OpCompileFlag) Opcode() uint32 { return 218 }

// NewOpCompileFlag creates a new codec for the OpCompileFlag instruction.
func NewOpCompileFlag() Codec {
	return Codec{
		Decode: func(argv []uint32) (Instruction, error) {
			return OpCompileFlag(
				spirv.DecodeString(argv),
			), nil
		},
		Encode: func(i Instruction, out []uint32) error {
			cf := i.(OpCompileFlag)
			size := spirv.EncodedStringLen(string(cf))
			out[0] = spirv.EncodeOpcode(uint32(size+1), 218)
			spirv.EncodeString(string(cf), out[1:])
			return nil
		},
	}
}
