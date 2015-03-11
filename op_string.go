// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpString defines the OpString instruction.
//
// It names a string for use with other debug instructions.
// This has no semantic impact and can safely be removed from a module.
type OpString struct {
	ResultId uint32
	String   string
}

func (c *OpString) Opcode() uint32 { return 56 }

// NewOpString creates a new codec for the OpString instruction.
func NewOpString() Codec {
	return Codec{
		Decode: func(argv []uint32) (Instruction, error) {
			if len(argv) < 2 {
				return nil, ErrMissingInstructionArgs
			}

			return &OpString{
				ResultId: argv[0],
				String:   DecodeString(argv[1:]),
			}, nil
		},
		Encode: func(i Instruction, out []uint32) error {
			v := i.(*OpString)
			name_size := EncodedStringLen(v.String)

			out[0] = EncodeOpcode(2+uint32(name_size), v.Opcode())
			out[1] = v.ResultId
			EncodeString(v.String, out[2:])
			return nil
		},
	}
}
