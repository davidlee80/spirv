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
		Encode: func(i Instruction) ([]uint32, error) {
			id := i.(OpTypeBool)
			return []uint32{uint32(id)}, nil
		},
	}
}
