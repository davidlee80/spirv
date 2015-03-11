// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpTypeQueue declares an OpenCL queue object.
//
// It defines the <id> of the new queue type.
type OpTypeQueue uint32

func (c OpTypeQueue) Opcode() uint32 { return 25 }

func bindOpTypeQueue(set *InstructionSet) {
	set.Set(
		OpTypeQueue(0).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 1 {
					return nil, ErrMissingInstructionArgs
				}

				return OpTypeQueue(argv[0]), nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(OpTypeQueue)
				out[0] = EncodeOpcode(2, v.Opcode())
				out[1] = uint32(v)
				return nil
			},
		},
	)
}
