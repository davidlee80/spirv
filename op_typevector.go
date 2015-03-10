// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpTypeVector represents the OpTypeVector instruction.
// It declares a new vector type.
type OpTypeVector struct {
	// The <id> of the new vector type.
	Result uint32

	// Specifies the type of each component in the resulting type.
	ComponentType uint32

	// Specifies the number of compononents in the resulting type.
	// It must be at least 2.
	ComponentCount uint32
}

func init() {
	DefaultInstructionSet[opTypeVector] = Codec{
		Decode: func(argv []uint32) (interface{}, error) {
			if len(argv) < 3 {
				return nil, ErrMissingInstructionArgs
			}

			return &OpTypeVector{
				Result:         argv[0],
				ComponentType:  argv[1],
				ComponentCount: argv[2],
			}, nil
		},
		Encode: func(instr interface{}) ([]uint32, error) {
			v := instr.(*OpTypeVector)
			return []uint32{
				v.Result,
				v.ComponentType,
				v.ComponentCount,
			}, nil
		},
	}
}