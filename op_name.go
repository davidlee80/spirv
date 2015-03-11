// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpName defines the OpName instruction.
//
// It names a result ID.
// This has no semantic impact and can safely be removed from a module.
type OpName struct {
	Name   String
	Target uint32
}

func (c *OpName) Opcode() uint32 { return 54 }

// NewOpName creates a new codec for the OpName instruction.
func NewOpName() Codec {
	return Codec{
		Decode: func(argv []uint32) (Instruction, error) {
			if len(argv) < 2 {
				return nil, ErrMissingInstructionArgs
			}

			return &OpName{
				Target: argv[0],
				Name:   DecodeString(argv[1:]),
			}, nil
		},
		Encode: func(i Instruction, out []uint32) error {
			v := i.(*OpName)
			nameSize := v.Name.EncodedLen()

			out[0] = EncodeOpcode(2+nameSize, v.Opcode())
			out[1] = v.Target
			v.Name.Encode(out[2:])
			return nil
		},
	}
}
