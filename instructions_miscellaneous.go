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
		Codec{
			New: func() Instruction {
				return &OpNop{}
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
		Codec{
			New: func() Instruction {
				return &OpUndef{}
			},
		},
	)
}
