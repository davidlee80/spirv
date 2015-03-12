// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpSource defines the OpSource instruction.
//
// It documents what source language a module was translated from.
// This has no semantic impact and can safely be removed from a module.
type OpSource struct {
	SourceLanguage uint32
	Version        uint32
}

func (c *OpSource) Opcode() uint32 { return 1 }

func init() {
	Bind(
		Codec{
			New: func() Instruction {
				return &OpSource{}
			},
		},
	)
}

// OpSourceExtension defines optional extensions used within the source language.
//
// It documents an extension to the source language. This has no semantic
// impact and can safely be removed from a module.
type OpSourceExtension struct {
	Extension String
}

func (c OpSourceExtension) Opcode() uint32 { return 2 }

func init() {
	Bind(
		Codec{
			New: func() Instruction {
				return &OpSourceExtension{}
			},
		},
	)
}

// OpName defines the OpName instruction.
//
// It names a result ID.
// This has no semantic impact and can safely be removed from a module.
type OpName struct {
	Target uint32
	Name   String
}

func (c *OpName) Opcode() uint32 { return 54 }

func init() {
	Bind(
		Codec{
			New: func() Instruction {
				return &OpName{}
			},
		},
	)
}

// OpMemberName defines the OpMemberName instruction.
//
// It names a member of a structure type.
// This has no semantic impact and can safely be removed from a module.
type OpMemberName struct {
	Type   uint32
	Member uint32
	Name   String
}

func (c *OpMemberName) Opcode() uint32 { return 55 }

func init() {
	Bind(
		Codec{
			New: func() Instruction {
				return &OpMemberName{}
			},
		},
	)
}

// OpString defines the OpString instruction.
//
// It names a string for use with other debug instructions.
// This has no semantic impact and can safely be removed from a module.
type OpString struct {
	ResultId uint32
	String   String
}

func (c *OpString) Opcode() uint32 { return 56 }

func init() {
	Bind(
		Codec{
			New: func() Instruction {
				return &OpString{}
			},
		},
	)
}

// OpLine defines the OpLine instruction.
//
// It adds source-level location information.
// This has no semantic impact and can safely be removed from a module.
type OpLine struct {
	Target uint32
	File   uint32
	Line   uint32
	Column uint32
}

func (c *OpLine) Opcode() uint32 { return 57 }

func init() {
	Bind(
		Codec{
			New: func() Instruction {
				return &OpLine{}
			},
		},
	)
}
