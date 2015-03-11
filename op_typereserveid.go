// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpTypeReserveId declares an OpenCL reservation id object.
//
// It defines the <id> of the new reservation type.
type OpTypeReserveId uint32

func (c OpTypeReserveId) Opcode() uint32 { return 24 }

func bindOpTypeReserveId(set *InstructionSet) {
	set.Set(
		OpTypeReserveId(0).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 1 {
					return nil, ErrMissingInstructionArgs
				}

				return OpTypeReserveId(argv[0]), nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(OpTypeReserveId)
				out[0] = EncodeOpcode(2, v.Opcode())
				out[1] = uint32(v)
				return nil
			},
		},
	)
}
