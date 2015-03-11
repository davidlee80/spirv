// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpTypeDeviceEvent declares an OpenCL device-side event object.
//
// It defines the <id> of the new device-side-event type.
type OpTypeDeviceEvent uint32

func (c OpTypeDeviceEvent) Opcode() uint32 { return 23 }

func bindOpTypeDeviceEvent(set *InstructionSet) {
	set.Set(
		OpTypeDeviceEvent(0).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 1 {
					return nil, ErrMissingInstructionArgs
				}

				return OpTypeDeviceEvent(argv[0]), nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(OpTypeDeviceEvent)
				out[0] = EncodeOpcode(2, v.Opcode())
				out[1] = uint32(v)
				return nil
			},
		},
	)
}
