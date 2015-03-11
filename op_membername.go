// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpMemberName defines the OpMemberName instruction.
//
// It names a member of a structure type.
// This has no semantic impact and can safely be removed from a module.
type OpMemberName struct {
	Type   uint32
	Member uint32
	Name   String
}

func (c *OpMemberName) Opcode() uint32 { return 55 }

// NewOpMemberName creates a new codec for the OpMemberName instruction.
func NewOpMemberName() Codec {
	return Codec{
		Decode: func(argv []uint32) (Instruction, error) {
			if len(argv) < 2 {
				return nil, ErrMissingInstructionArgs
			}

			return &OpMemberName{
				Type:   argv[0],
				Member: argv[1],
				Name:   DecodeString(argv[2:]),
			}, nil
		},
		Encode: func(i Instruction, out []uint32) error {
			v := i.(*OpMemberName)
			nameSize := v.Name.EncodedLen()

			out[0] = EncodeOpcode(3+nameSize, v.Opcode())
			out[1] = v.Type
			out[2] = v.Member
			v.Name.Encode(out[3:])
			return nil
		},
	}
}
