// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpNop represents the OpNop instruction.
//
// Its use is not allowed, but it is explicitely defined in the specification.
// We will therefor have the decoder and encoder return an appropriate error
// when it is being used.
type OpNop struct{}

func (c OpNop) Opcode() uint32 { return 0 }

func init() {
	Bind(
		(&OpNop{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				return nil, ErrUnacceptable
			},
			Encode: func(i Instruction, out []uint32) (uint32, error) {
				return 0, ErrUnacceptable
			},
		},
	)
}

// OpUndef makes an intermediate object with no initialization.
type OpUndef struct {
	ResultType uint32
	ResultId   uint32
}

func (c *OpUndef) Opcode() uint32 { return 45 }

func init() {
	Bind(
		(&OpUndef{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) != 2 {
					return nil, ErrInvalidInstructionSize
				}

				return &OpUndef{
					ResultType: argv[0],
					ResultId:   argv[1],
				}, nil
			},
			Encode: func(i Instruction, out []uint32) (uint32, error) {
				v := i.(*OpUndef)
				out[0] = v.ResultType
				out[1] = v.ResultId
				return 2, nil
			},
		},
	)
}
