// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package v99

import "github.com/jteeuwen/spirv"

// OpTypeBool represents the OpTypeBool instruction.
type OpTypeBool uint32

func (c OpTypeBool) Opcode() uint32 { return 9 }

// NewOpTypeBool creates a new codec for the OpTypeBool instruction.
func NewOpTypeBool() spirv.Codec {
	return spirv.Codec{
		Decode: func(argv []uint32) (spirv.Instruction, error) {
			if len(argv) < 1 {
				return nil, spirv.ErrMissingInstructionArgs
			}

			return OpTypeBool(argv[0]), nil
		},
		Encode: func(i spirv.Instruction, out []uint32) error {
			v := i.(OpTypeBool)
			out[0] = spirv.EncodeOpcode(2, 9)
			out[1] = uint32(v)
			return nil
		},
	}
}
