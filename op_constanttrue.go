// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpConstantTrue declares a true Boolean-type scalar constant.
type OpConstantTrue struct {
	// Result Type must be the scalar Boolean type.
	ResultType uint32

	// The <id> of the new constant type.
	ResultId uint32
}

func (c *OpConstantTrue) Opcode() uint32 { return 27 }

func bindOpConstantTrue(set *InstructionSet) {
	set.Set(
		(&OpConstantTrue{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 2 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpConstantTrue{
					ResultType: argv[0],
					ResultId:   argv[1],
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpConstantTrue)
				out[0] = EncodeOpcode(3, v.Opcode())
				out[1] = v.ResultType
				out[2] = v.ResultId
				return nil
			},
		},
	)
}
