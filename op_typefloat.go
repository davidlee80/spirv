// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpTypeFloat represents the OpTypeFloat instruction.
// It declares a new floating point type.
type OpTypeFloat struct {
	// The <id> of the new floating-point type.
	Result uint32

	// Specifies how many bits wide the type is. The bit pattern of a
	// floating-point value is as described by the IEEE 754 standard.
	Width uint32
}

func init() {
	DefaultInstructionSet[opTypeFloat] = Codec{
		Decode: func(argv []uint32) (interface{}, error) {
			if len(argv) < 2 {
				return nil, ErrMissingInstructionArgs
			}

			return &OpTypeFloat{
				Result: argv[0],
				Width:  argv[1],
			}, nil
		},
		Encode: func(instr interface{}) ([]uint32, error) {
			f := instr.(*OpTypeFloat)
			return []uint32{
				f.Result,
				f.Width,
			}, nil
		},
	}
}
