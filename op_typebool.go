// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpTypeBool represents the OpTypeBool instruction.
type OpTypeBool uint32

func (c OpTypeBool) Opcode() uint32 { return 9 }

// NewOpTypeBool creates a new codec for the OpTypeBool instruction.
func NewOpTypeBool() Codec {
	return Codec{
		Decode: func(argv []uint32) (Instruction, error) {
			if len(argv) < 1 {
				return nil, ErrMissingInstructionArgs
			}

			return OpTypeBool(argv[0]), nil
		},
		Encode: func(i Instruction, out []uint32) error {
			v := i.(OpTypeBool)
			out[0] = EncodeOpcode(2, v.Opcode())
			out[1] = uint32(v)
			return nil
		},
	}
}
