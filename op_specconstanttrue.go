// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpSpecConstantTrue declares a Boolean-type scalar specialization
// constant with a default value of true.
//
// This instruction can be specialized to become either an OpConstantTrue
// or OpConstantFalse instruction.
type OpSpecConstantTrue struct {
	// Result Type must be the scalar Boolean type.
	ResultType uint32

	// The <id> of the new constant type.
	ResultId uint32
}

func (c *OpSpecConstantTrue) Opcode() uint32 { return 34 }

func init() {
	Bind(
		(&OpSpecConstantTrue{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 2 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpSpecConstantTrue{
					ResultType: argv[0],
					ResultId:   argv[1],
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpSpecConstantTrue)
				out[0] = EncodeOpcode(3, v.Opcode())
				out[1] = v.ResultType
				out[2] = v.ResultId
				return nil
			},
		},
	)
}
