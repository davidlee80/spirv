// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpConstantSampler declares a new null sampler constant.
type OpConstantSampler struct {
	ResultType uint32
	ResultId   uint32

	// Addressing is the addressing Mode.
	Addressing uint32

	// Param is one of:
	//
	//   0: Nonparametric
	//   1: Parametric
	//
	Param uint32

	// Filter is the filter mode.
	Filter uint32
}

func (c *OpConstantSampler) Opcode() uint32 { return 31 }

func init() {
	Bind(
		(&OpConstantSampler{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 5 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpConstantSampler{
					ResultType: argv[0],
					ResultId:   argv[1],
					Addressing: argv[2],
					Param:      argv[3],
					Filter:     argv[4],
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpConstantSampler)
				out[0] = EncodeOpcode(6, v.Opcode())
				out[1] = v.ResultType
				out[2] = v.ResultId
				out[3] = v.Addressing
				out[4] = v.Param
				out[5] = v.Filter
				return nil
			},
		},
	)
}
