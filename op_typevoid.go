// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpTypeVoid represents the OpTypeVoid instruction.
type OpTypeVoid uint32

func (c OpTypeVoid) Opcode() uint32 { return 8 }

// NewOpTypeVoid creates a new codec for the OpTypeVoid instruction.
func NewOpTypeVoid() Codec {
	return Codec{
		Decode: func(argv []uint32) (Instruction, error) {
			if len(argv) < 1 {
				return nil, ErrMissingInstructionArgs
			}

			return OpTypeVoid(argv[0]), nil
		},
		Encode: func(i Instruction) ([]uint32, error) {
			id := i.(OpTypeVoid)
			return []uint32{uint32(id)}, nil
		},
	}
}
