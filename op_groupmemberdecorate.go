// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpGroupMemberDecorate represents the OpGroupMemberDecorate instruction.
// It adds a decoration to a member of a structure type
type OpGroupMemberDecorate struct {
	// Targets are the target <id>s to decorate with the groups of decorations.
	Targets []uint32

	// The <id> of a OpDecorationGroup instruction.
	Group uint32
}

func (c *OpGroupMemberDecorate) Opcode() uint32 { return 53 }

func bindOpGroupMemberDecorate(set *InstructionSet) {
	set.Set(
		(&OpGroupMemberDecorate{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 1 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpGroupMemberDecorate{
					Group:   argv[0],
					Targets: Copy(argv[1:]),
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpGroupMemberDecorate)
				size := uint32(len(v.Targets))

				out[0] = EncodeOpcode(2+size, v.Opcode())
				out[1] = v.Group
				copy(out[2:], v.Targets)
				return nil
			},
		},
	)
}
