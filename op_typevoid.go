// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpTypeVoid represents the OpTypeVoid instruction.
type OpTypeVoid uint32

func init() {
	DefaultInstructionSet[opTypeVoid] = Codec{
		Decode: func(argv []uint32) (interface{}, error) {
			if len(argv) < 1 {
				return nil, ErrMissingInstructionArgs
			}

			return OpTypeVoid(argv[0]), nil
		},
		Encode: func(instr interface{}) ([]uint32, error) {
			id := instr.(OpTypeVoid)
			return []uint32{uint32(id)}, nil
		},
	}
}
