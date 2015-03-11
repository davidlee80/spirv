// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpTypeVoid represents the OpTypeVoid instruction.
type OpTypeVoid uint32

func (c OpTypeVoid) Opcode() uint32 { return 8 }

func bindOpTypeVoid(set *InstructionSet) {
	set.Set(
		OpTypeVoid(0).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 1 {
					return nil, ErrMissingInstructionArgs
				}

				return OpTypeVoid(argv[0]), nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(OpTypeVoid)
				out[0] = EncodeOpcode(2, v.Opcode())
				out[1] = uint32(v)
				return nil
			},
		},
	)
}
