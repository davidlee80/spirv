// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpUndef defines the OpUndef instruction.
//
// It makes an intermediate object with no initialization.
type OpUndef struct {
	ResultType uint32
	ResultId   uint32
}

func (c *OpUndef) Opcode() uint32 { return 45 }

func bindOpUndef(set *InstructionSet) {
	set.Set(
		(&OpUndef{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) != 2 {
					return nil, ErrInvalidInstructionSize
				}

				return &OpUndef{
					ResultType: argv[0],
					ResultId:   argv[1],
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpUndef)
				out[0] = EncodeOpcode(3, v.Opcode())
				out[1] = v.ResultType
				out[2] = v.ResultId
				return nil
			},
		},
	)
}
