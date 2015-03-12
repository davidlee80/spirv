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

func init() {
	Bind(
		OpDecorationGroup(0).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 1 {
					return nil, ErrMissingInstructionArgs
				}

				return OpDecorationGroup(argv[0]), nil
			},
			Encode: func(i Instruction, out []uint32) (uint32, error) {
				v := i.(OpDecorationGroup)
				out[0] = uint32(v)
				return 1, nil
			},
		},
	)
}

// OpDecorate represents the OpDecorate instruction.
// It adds a decoration to another <id>.
type OpDecorate struct {
	// Target is the <id> to decorate. It can potentially be any <id> that
	// is a forward reference. A set of decorations can be grouped together
	// by having multiple OpDecorate instructions target the same
	// OpDecorationGroup instruction.
	Target uint32

	// The decoration type to apply.
	Decoration uint32

	// Optional list of decoration arguments.
	//
	// Refer to the Decoration constant documentation for more details
	// on which values require which arguments.
	Argv []uint32
}

func (c *OpDecorate) Opcode() uint32 { return 50 }

func init() {
	Bind(
		(&OpDecorate{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 2 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpDecorate{
					Target:     argv[0],
					Decoration: argv[1],
					Argv:       Copy(argv[2:]),
				}, nil
			},
			Encode: func(i Instruction, out []uint32) (uint32, error) {
				v := i.(*OpDecorate)
				size := uint32(len(v.Argv))

				out[0] = v.Target
				out[1] = v.Decoration
				copy(out[2:], v.Argv)
				return 2 + size, nil
			},
		},
	)
}

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

func init() {
	Bind(
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
			Encode: func(i Instruction, out []uint32) (uint32, error) {
				v := i.(*OpMemberDecorate)
				size := uint32(len(v.Argv))

				out[0] = v.StructType
				out[1] = v.Member
				out[2] = v.Decoration
				copy(out[3:], v.Argv)
				return 3 + size, nil
			},
		},
	)
}

// OpGroupDecorate represents the OpGroupDecorate instruction.
// It adds a group of decorations to another <id>.
type OpGroupDecorate struct {
	// Targets are the target <id>s to decorate with the groups of decorations.
	Targets []uint32

	// Decoration group is the <id> of an OpDecorationGroup instruction.
	Group uint32
}

func (c *OpGroupDecorate) Opcode() uint32 { return 52 }

func init() {
	Bind(
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
			Encode: func(i Instruction, out []uint32) (uint32, error) {
				v := i.(*OpGroupDecorate)
				size := uint32(len(v.Targets))

				out[0] = v.Group
				copy(out[1:], v.Targets)
				return 1 + size, nil
			},
		},
	)
}

// OpGroupMemberDecorate represents the OpGroupMemberDecorate instruction.
// It adds a decoration to a member of a structure type
type OpGroupMemberDecorate struct {
	// Targets are the target <id>s to decorate with the groups of decorations.
	Targets []uint32

	// The <id> of a OpDecorationGroup instruction.
	Group uint32
}

func (c *OpGroupMemberDecorate) Opcode() uint32 { return 53 }

func init() {
	Bind(
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
			Encode: func(i Instruction, out []uint32) (uint32, error) {
				v := i.(*OpGroupMemberDecorate)
				size := uint32(len(v.Targets))

				out[0] = v.Group
				copy(out[1:], v.Targets)
				return 1 + size, nil
			},
		},
	)
}
