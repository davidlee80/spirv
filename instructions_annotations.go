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
type OpDecorationGroup struct {
	ResultId Id
}

func (c *OpDecorationGroup) Opcode() uint32 { return 49 }
func (c *OpDecorationGroup) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpDecorationGroup{}
	})
}

// OpDecorate represents the OpDecorate instruction.
// It adds a decoration to another <id>.
type OpDecorate struct {
	// Target is the <id> to decorate. It can potentially be any <id> that
	// is a forward reference. A set of decorations can be grouped together
	// by having multiple OpDecorate instructions target the same
	// OpDecorationGroup instruction.
	Target Id

	// The decoration type to apply.
	Decoration Decoration

	// Optional list of decoration arguments.
	//
	// Refer to the Decoration constant documentation for more details
	// on which values require which arguments.
	Argv []uint32
}

func (c *OpDecorate) Opcode() uint32 { return 50 }
func (c *OpDecorate) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpDecorate{}
	})
}

// OpMemberDecorate represents the OpMemberDecorate instruction.
// It adds a decoration to a member of a structure type.
type OpMemberDecorate struct {
	// The <id> of a type from OpTypeStruct.
	StructType Id

	// Member is the number of the member to decorate in the structure.
	// The first member is member 0, the next is member 1, . . .
	Member uint32

	// The decoration type to apply.
	Decoration Decoration

	// Optional list of decoration arguments.
	//
	// Refer to the Decoration constant documentation for more details
	// on which values require which arguments.
	Argv []uint32
}

func (c *OpMemberDecorate) Opcode() uint32 { return 51 }
func (c *OpMemberDecorate) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpMemberDecorate{}
	})
}

// OpGroupDecorate represents the OpGroupDecorate instruction.
// It adds a group of decorations to another <id>.
type OpGroupDecorate struct {
	// Decoration group is the <id> of an OpDecorationGroup instruction.
	Group Id

	// Targets are the target <id>s to decorate with the groups of decorations.
	Targets []Id
}

func (c *OpGroupDecorate) Opcode() uint32 { return 52 }
func (c *OpGroupDecorate) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpGroupDecorate{}
	})
}

// OpGroupMemberDecorate represents the OpGroupMemberDecorate instruction.
// It adds a decoration to a member of a structure type
type OpGroupMemberDecorate struct {
	// The <id> of a OpDecorationGroup instruction.
	Group Id

	// Targets are the target <id>s to decorate with the groups of decorations.
	Targets []Id
}

func (c *OpGroupMemberDecorate) Opcode() uint32 { return 53 }
func (c *OpGroupMemberDecorate) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpGroupMemberDecorate{}
	})
}
