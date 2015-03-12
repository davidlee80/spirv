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
		(&OpSource{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 2 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpSource{
					SourceLanguage: argv[0],
					Version:        argv[1],
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpSource)
				out[0] = EncodeOpcode(3, v.Opcode())
				out[1] = v.SourceLanguage
				out[2] = v.Version
				return nil
			},
		},
	)
}

// OpSourceExtension defines optional extensions used within the source language.
//
// It documents an extension to the source language. This has no semantic
// impact and can safely be removed from a module.
type OpSourceExtension String

func (c OpSourceExtension) Opcode() uint32 { return 2 }

func init() {
	Bind(
		OpSourceExtension("").Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) == 0 {
					return nil, ErrInvalidInstructionSize
				}

				return OpSourceExtension(
					DecodeString(argv),
				), nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(OpSourceExtension)
				size := String(v).EncodedLen()
				out[0] = EncodeOpcode(size+1, v.Opcode())
				String(v).Encode(out[1:])
				return nil
			},
		},
	)
}

// OpName defines the OpName instruction.
//
// It names a result ID.
// This has no semantic impact and can safely be removed from a module.
type OpName struct {
	Name   String
	Target uint32
}

func (c *OpName) Opcode() uint32 { return 54 }

func init() {
	Bind(
		(&OpName{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 2 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpName{
					Target: argv[0],
					Name:   DecodeString(argv[1:]),
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpName)
				nameSize := v.Name.EncodedLen()

				out[0] = EncodeOpcode(2+nameSize, v.Opcode())
				out[1] = v.Target
				v.Name.Encode(out[2:])
				return nil
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
		(&OpMemberName{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 2 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpMemberName{
					Type:   argv[0],
					Member: argv[1],
					Name:   DecodeString(argv[2:]),
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpMemberName)
				nameSize := v.Name.EncodedLen()

				out[0] = EncodeOpcode(3+nameSize, v.Opcode())
				out[1] = v.Type
				out[2] = v.Member
				v.Name.Encode(out[3:])
				return nil
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
		(&OpLine{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) != 4 {
					return nil, ErrInvalidInstructionSize
				}

				return &OpLine{
					Target: argv[0],
					File:   argv[1],
					Line:   argv[2],
					Column: argv[3],
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpLine)
				out[0] = EncodeOpcode(5, v.Opcode())
				out[1] = v.Target
				out[2] = v.File
				out[3] = v.Line
				out[4] = v.Column
				return nil
			},
		},
	)
}
