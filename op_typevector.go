// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpTypeVector represents the OpTypeVector instruction.
// It declares a new vector type.
type OpTypeVector struct {
	// The <id> of the new vector type.
	ResultId uint32

	// Specifies the type of each component in the resulting type.
	ComponentType uint32

	// Specifies the number of compononents in the resulting type.
	// It must be at least 2.
	ComponentCount uint32
}

func (c *OpTypeVector) Opcode() uint32 { return 12 }

func bindOpTypeVector(set *InstructionSet) {
	set.Set(
		(&OpTypeVector{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 3 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpTypeVector{
					ResultId:       argv[0],
					ComponentType:  argv[1],
					ComponentCount: argv[2],
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpTypeVector)
				out[0] = EncodeOpcode(4, v.Opcode())
				out[1] = v.ResultId
				out[2] = v.ComponentType
				out[3] = v.ComponentCount
				return nil
			},
		},
	)
}
