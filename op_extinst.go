// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

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
