// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package v99

import "github.com/jteeuwen/spirv"

// OpTypeFloat represents the OpTypeFloat instruction.
// It declares a new floating point type.
type OpTypeFloat struct {
	// The <id> of the new floating-point type.
	Result uint32

	// Specifies how many bits wide the type is. The bit pattern of a
	// floating-point value is as described by the IEEE 754 standard.
	Width uint32
}

func (c *OpTypeFloat) Opcode() uint32 { return 11 }

// NewOpTypeFloat creates a new codec for the OpTypeFloat instruction.
func NewOpTypeFloat() spirv.Codec {
	return spirv.Codec{
		Decode: func(argv []uint32) (spirv.Instruction, error) {
			if len(argv) < 2 {
				return nil, spirv.ErrMissingInstructionArgs
			}

			return &OpTypeFloat{
				Result: argv[0],
				Width:  argv[1],
			}, nil
		},
		Encode: func(i spirv.Instruction, out []uint32) error {
			v := i.(*OpTypeFloat)
			out[0] = spirv.EncodeOpcode(3, 11)
			out[1] = v.Result
			out[2] = v.Width
			return nil
		},
	}
}
