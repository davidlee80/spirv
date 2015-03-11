// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpDecorationGroup represents a collector of decorations from OpDecorate
// instructions.
//
// All such instructions must precede this instruction. Subsequent OpGroupDecorate
// and OpGroupMemberDecorate instructions can consume the Result <id> to apply
// multiple decorations to multiple target <id>s. Those are the only
// instructions allowed to consume the Result <id>.
type OpDecorationGroup uint32

func (c OpDecorationGroup) Opcode() uint32 { return 49 }

func bindOpDecorationGroup(set *InstructionSet) {
	set.Set(
		OpDecorationGroup(0).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 1 {
					return nil, ErrMissingInstructionArgs
				}

				return OpDecorationGroup(argv[0]), nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(OpDecorationGroup)
				out[0] = EncodeOpcode(2, v.Opcode())
				out[1] = uint32(v)
				return nil
			},
		},
	)
}
