// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpExtension defines the OpExtension instruction.
//
// It declares use of an extension to SPIR-V. This allows validation of
// additional instructions, tokens, semantics, etc
type OpExtension String

func (c OpExtension) Opcode() uint32 { return 3 }

func init() {
	Bind(
		OpExtension("").Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) == 0 {
					return nil, ErrInvalidInstructionSize
				}

				return OpExtension(
					DecodeString(argv),
				), nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(OpExtension)
				size := String(v).EncodedLen()
				out[0] = EncodeOpcode(size+1, v.Opcode())
				String(v).Encode(out[1:])
				return nil
			},
		},
	)
}

// OpExtInstImport defines the OpExtInstImport instruction.
//
// It defines an import of an extended set of instructions.
// It can later be referenced by the Result <id>
type OpExtInstImport struct {
	Name     String
	ResultId uint32
}

func (c *OpExtInstImport) Opcode() uint32 { return 4 }

func init() {
	Bind(
		(&OpExtInstImport{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 2 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpExtInstImport{
					ResultId: argv[0],
					Name:     DecodeString(argv[1:]),
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpExtInstImport)
				size := v.Name.EncodedLen()

				out[0] = EncodeOpcode(2+size, v.Opcode())
				out[1] = v.ResultId
				v.Name.Encode(out[2:])
				return nil
			},
		},
	)
}

// OpExtInst defines an instruction in an imported set of extended instructions.
type OpExtInst struct {
	ResultType  uint32
	ResultId    uint32
	Set         uint32   // Result of an OpExtInstImport instruction.
	Instruction uint32   // Enumerant of the instruction to execute within the extended instruction Set.
	Operands    []uint32 // Operands to the extended instruction.
}

func (c *OpExtInst) Opcode() uint32 { return 44 }

func init() {
	Bind(
		(&OpExtInst{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 4 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpExtInst{
					ResultType:  argv[0],
					ResultId:    argv[1],
					Set:         argv[2],
					Instruction: argv[3],
					Operands:    Copy(argv[4:]),
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpExtInst)
				out[0] = EncodeOpcode(5+uint32(len(v.Operands)), v.Opcode())
				out[1] = v.ResultType
				out[2] = v.ResultId
				out[3] = v.Set
				out[4] = v.Instruction
				copy(out[5:], v.Operands)
				return nil
			},
		},
	)
}
