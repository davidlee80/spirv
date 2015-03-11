// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpGroupDecorate represents the OpGroupDecorate instruction.
// It adds a group of decorations to another <id>.
type OpGroupDecorate struct {
	// Targets are the target <id>s to decorate with the groups of decorations.
	Targets []uint32

	// Decoration group is the <id> of an OpDecorationGroup instruction.
	Group uint32
}

func (c *OpGroupDecorate) Opcode() uint32 { return 52 }

func bindOpGroupDecorate(set *InstructionSet) {
	set.Set(
		(&OpGroupDecorate{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 1 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpGroupDecorate{
					Group:   argv[0],
					Targets: Copy(argv[1:]),
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpGroupDecorate)
				size := uint32(len(v.Targets))

				out[0] = EncodeOpcode(2+size, v.Opcode())
				out[1] = v.Group
				copy(out[2:], v.Targets)
				return nil
			},
		},
	)
}
