// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpTypeFloat represents the OpTypeFloat instruction.
// It declares a new floating point type.
type OpTypeFloat struct {
	// The <id> of the new floating-point type.
	ResultId uint32

	// Specifies how many bits wide the type is. The bit pattern of a
	// floating-point value is as described by the IEEE 754 standard.
	Width uint32
}

func (c *OpTypeFloat) Opcode() uint32 { return 11 }

func init() {
	Bind(
		(&OpTypeFloat{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 2 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpTypeFloat{
					ResultId: argv[0],
					Width:    argv[1],
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpTypeFloat)
				out[0] = EncodeOpcode(3, v.Opcode())
				out[1] = v.ResultId
				out[2] = v.Width
				return nil
			},
		},
	)
}
