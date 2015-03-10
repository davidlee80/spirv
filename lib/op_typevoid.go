// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package lib

import "github.com/jteeuwen/spirv"

// OpTypeVoid represents the OpTypeVoid instruction.
type OpTypeVoid uint32

func (c OpTypeVoid) Opcode() uint32 { return 8 }

// NewOpTypeVoid creates a new codec for the OpTypeVoid instruction.
func NewOpTypeVoid() Codec {
	return Codec{
		Decode: func(argv []uint32) (Instruction, error) {
			if len(argv) < 1 {
				return nil, ErrMissingInstructionArgs
			}

			return OpTypeVoid(argv[0]), nil
		},
		Encode: func(i Instruction, out []uint32) error {
			v := i.(OpTypeVoid)
			out[0] = spirv.EncodeOpcode(2, 8)
			out[1] = uint32(v)
			return nil
		},
	}
}
