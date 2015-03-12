// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpConstant declares a new Integer-type or Floating-point-type
// scalar constant.
type OpConstant struct {
	// Result Type must be a scalar Integer type or Floating-point type.
	ResultType uint32

	// The <id> of the new constant type.
	ResultId uint32

	// Value is the bit pattern for the constant.
	//
	// Types 32 bits wide or smaller take one word. Larger types take
	// multiple words, with low-order words appearing first.
	Value []uint32
}

func (c *OpConstant) Opcode() uint32 { return 29 }

func init() {
	Bind(
		(&OpConstant{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 2 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpConstant{
					ResultType: argv[0],
					ResultId:   argv[1],
					Value:      Copy(argv[2:]),
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpConstant)
				size := uint32(len(v.Value))
				out[0] = EncodeOpcode(3+size, v.Opcode())
				out[1] = v.ResultType
				out[2] = v.ResultId
				copy(out[3:], v.Value)
				return nil
			},
		},
	)
}
