// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpMemoryModel represents the OpMemoryModel instruction.
//
// It sets addressing model and memory model for the entire module.
type OpMemoryModel struct {
	Addressing AddressingMode // Selects the module’s addressing model
	Memory     MemoryMode     // Selects the module’s memory model
}

func init() {
	DefaultInstructionSet[opMemoryModel] = Codec{
		Decode: func(argv []uint32) (interface{}, error) {
			if len(argv) < 2 {
				return nil, ErrMissingInstructionArgs
			}

			return &OpMemoryModel{
				Addressing: AddressingMode(argv[0]),
				Memory:     MemoryMode(argv[1]),
			}, nil
		},
		Encode: func(instr interface{}) ([]uint32, error) {
			mm := instr.(*OpMemoryModel)
			return []uint32{
				uint32(mm.Addressing),
				uint32(mm.Memory),
			}, nil
		},
	}
}
