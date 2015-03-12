// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpMemoryModel represents the OpMemoryModel instruction.
//
// It sets addressing model and memory model for the entire module.
type OpMemoryModel struct {
	AddressingMode uint32 // Selects the module’s addressing model
	MemoryMode     uint32 // Selects the module’s memory mode
}

func (c *OpMemoryModel) Opcode() uint32 { return 5 }

func bindOpMemoryModel(set *InstructionSet) {
	set.Set(
		(&OpMemoryModel{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 2 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpMemoryModel{
					AddressingMode: argv[0],
					MemoryMode:     argv[1],
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpMemoryModel)
				out[0] = EncodeOpcode(3, v.Opcode())
				out[1] = v.AddressingMode
				out[2] = v.MemoryMode
				return nil
			},
		},
	)
}
