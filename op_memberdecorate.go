// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpMemberDecorate represents the OpMemberDecorate instruction.
// It adds a decoration to a member of a structure type.
type OpMemberDecorate struct {
	// The <id> of a type from OpTypeStruct.
	StructType uint32

	// Member is the number of the member to decorate in the structure.
	// The first member is member 0, the next is member 1, . . .
	Member uint32

	// The decoration type to apply.
	Decoration uint32

	// Optional list of decoration arguments.
	//
	// Refer to the Decoration constant documentation for more details
	// on which values require which arguments.
	Argv []uint32
}

func (c *OpMemberDecorate) Opcode() uint32 { return 51 }

func bindOpMemberDecorate(set *InstructionSet) {
	set.Set(
		(&OpMemberDecorate{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 3 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpMemberDecorate{
					StructType: argv[0],
					Member:     argv[1],
					Decoration: argv[2],
					Argv:       Copy(argv[3:]),
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpMemberDecorate)
				size := uint32(len(v.Argv))

				out[0] = EncodeOpcode(4+size, v.Opcode())
				out[1] = v.StructType
				out[2] = v.Member
				out[3] = v.Decoration
				copy(out[4:], v.Argv)
				return nil
			},
		},
	)
}
