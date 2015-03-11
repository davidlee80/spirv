// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpString defines the OpString instruction.
//
// It names a string for use with other debug instructions.
// This has no semantic impact and can safely be removed from a module.
type OpString struct {
	ResultId uint32
	String   String
}

func (c *OpString) Opcode() uint32 { return 56 }

func bindOpString(set *InstructionSet) {
	set.Set(
		(&OpString{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 2 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpString{
					ResultId: argv[0],
					String:   DecodeString(argv[1:]),
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpString)
				nameSize := v.String.EncodedLen()

				out[0] = EncodeOpcode(2+nameSize, v.Opcode())
				out[1] = v.ResultId
				v.String.Encode(out[2:])
				return nil
			},
		},
	)
}
