// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpConstantNullPointer declares a new null pointer constant.
type OpConstantNullPointer struct {
	// Result Type must be the scalar Boolean type.
	ResultType uint32

	// The <id> of the new constant type.
	ResultId uint32
}

func (c *OpConstantNullPointer) Opcode() uint32 { return 32 }

func init() {
	Bind(
		(&OpConstantNullPointer{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 2 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpConstantNullPointer{
					ResultType: argv[0],
					ResultId:   argv[1],
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpConstantNullPointer)
				out[0] = EncodeOpcode(3, v.Opcode())
				out[1] = v.ResultType
				out[2] = v.ResultId
				return nil
			},
		},
	)
}
