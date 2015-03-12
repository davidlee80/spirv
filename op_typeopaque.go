// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpTypeOpaque declares a named structure type with no body specified.
type OpTypeOpaque struct {
	// The <id> of the new array type.
	ResultId uint32

	// The name of the opaque type.
	Name String
}

func (c *OpTypeOpaque) Opcode() uint32 { return 19 }

func bindOpTypeOpaque(set *InstructionSet) {
	set.Set(
		(&OpTypeOpaque{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 1 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpTypeOpaque{
					ResultId: argv[0],
					Name:     DecodeString(argv[1:]),
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpTypeOpaque)
				size := v.Name.EncodedLen()

				out[0] = EncodeOpcode(2+size, v.Opcode())
				out[1] = v.ResultId
				v.Name.Encode(out[2:])
				return nil
			},
		},
	)
}
