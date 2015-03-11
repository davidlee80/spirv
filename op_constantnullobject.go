// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpConstantNullObject declares a new null object constant.
// The objerct can be a queue, event or reservation id.
type OpConstantNullObject struct {
	// Result Type must be the scalar Boolean type.
	ResultType uint32

	// The <id> of the new constant type.
	ResultId uint32
}

func (c *OpConstantNullObject) Opcode() uint32 { return 33 }

func bindOpConstantNullObject(set *InstructionSet) {
	set.Set(
		(&OpConstantNullObject{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 2 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpConstantNullObject{
					ResultType: argv[0],
					ResultId:   argv[1],
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpConstantNullObject)
				out[0] = EncodeOpcode(3, v.Opcode())
				out[1] = v.ResultType
				out[2] = v.ResultId
				return nil
			},
		},
	)
}
