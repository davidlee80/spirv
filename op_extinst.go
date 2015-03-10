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

// NewOpExtInst creates a new codec for the OpExtInst instruction.
func NewOpExtInst() Codec {
	return Codec{
		Decode: func(argv []uint32) (Instruction, error) {
			if len(argv) < 4 {
				return nil, ErrMissingInstructionArgs
			}

			operands := make([]uint32, len(argv)-4)
			copy(operands, argv[4:])

			return &OpExtInst{
				ResultType:  argv[0],
				ResultId:    argv[1],
				Set:         argv[2],
				Instruction: argv[3],
				Operands:    operands,
			}, nil
		},
		Encode: func(i Instruction) ([]uint32, error) {
			ext := i.(*OpExtInst)
			out := make([]uint32, 4+len(ext.Operands))
			out[0] = ext.ResultType
			out[1] = ext.ResultId
			out[2] = ext.Set
			out[3] = ext.Instruction
			copy(out[4:], ext.Operands)
			return out, nil
		},
	}
}
