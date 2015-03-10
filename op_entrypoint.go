// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpEntryPoint represents the OpEntryPoint instruction.
// It declares an entry point and its execution model.
type OpEntryPoint struct {
	Execution ExecutionModel // Execution model for the entry point and its static call tree.
	Id        uint32         // Must the Result <id> of an OpFunction instruction.
}

func init() {
	DefaultInstructionSet[opEntryPoint] = Codec{
		Decode: func(argv []uint32) (interface{}, error) {
			if len(argv) < 2 {
				return nil, ErrMissingInstructionArgs
			}

			return &OpEntryPoint{
				Execution: ExecutionModel(argv[0]),
				Id:        argv[1],
			}, nil
		},
		Encode: func(instr interface{}) ([]uint32, error) {
			ep := instr.(*OpEntryPoint)
			return []uint32{
				uint32(ep.Execution),
				ep.Id,
			}, nil
		},
	}
}
