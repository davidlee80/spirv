// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpEntryPoint represents the OpEntryPoint instruction.
// It declares an entry point and its execution model.
type OpEntryPoint struct {
	ExecutionModel uint32 // Execution model for the entry point and its static call tree.
	Id             uint32 // Must the Result <id> of an OpFunction instruction.
}

func (c *OpEntryPoint) Opcode() uint32 { return 6 }

func bindOpEntryPoint(set *InstructionSet) {
	set.Set(
		(&OpEntryPoint{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 2 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpEntryPoint{
					ExecutionModel: argv[0],
					Id:             argv[1],
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpEntryPoint)
				out[0] = EncodeOpcode(3, v.Opcode())
				out[1] = v.ExecutionModel
				out[2] = v.Id
				return nil
			},
		},
	)
}
