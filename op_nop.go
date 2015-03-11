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

// NewOpNop creates a new codec for the OpNop instruction.
func NewOpNop() Codec {
	return Codec{
		Decode: func(argv []uint32) (Instruction, error) {
			return nil, ErrUnacceptable
		},
		Encode: func(i Instruction, out []uint32) error {
			return ErrUnacceptable
		},
	}
}
