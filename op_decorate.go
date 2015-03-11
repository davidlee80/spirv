// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpDecorate represents the OpDecorate instruction.
// It adds a decoration to another <id>.
type OpDecorate struct {
	// Target is the <id> to decorate. It can potentially be any <id> that
	// is a forward reference. A set of decorations can be grouped together
	// by having multiple OpDecorate instructions target the same
	// OpDecorationGroup instruction.
	Target uint32

	// The decoration type to apply.
	Decoration Decoration

	// Optional list of decoration arguments.
	//
	// Refer to the Decoration constant documentation for more details
	// on which values require which arguments.
	Argv []uint32
}

func (c *OpDecorate) Opcode() uint32 { return 50 }

func bindOpDecorate(set *InstructionSet) {
	set.Set(
		(&OpDecorate{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 2 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpDecorate{
					Target:     argv[0],
					Decoration: Decoration(argv[1]),
					Argv:       Copy(argv[2:]),
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpDecorate)
				size := uint32(len(v.Argv))

				out[0] = EncodeOpcode(3+size, v.Opcode())
				out[1] = v.Target
				out[2] = uint32(v.Decoration)
				copy(out[3:], v.Argv)
				return nil
			},
		},
	)
}
