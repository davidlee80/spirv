// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpTypeEvent declares an OpenCL event object.
type OpTypeEvent uint32

func (c OpTypeEvent) Opcode() uint32 { return 22 }

func bindOpTypeEvent(set *InstructionSet) {
	set.Set(
		OpTypeEvent(0).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 1 {
					return nil, ErrMissingInstructionArgs
				}

				return OpTypeEvent(argv[0]), nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(OpTypeEvent)
				out[0] = EncodeOpcode(2, v.Opcode())
				out[1] = uint32(v)
				return nil
			},
		},
	)
}
