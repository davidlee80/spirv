// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpTypeFilter declares a filter type. It is consumed by OpSampler.
// This type is opaque: values of this type have no defined
// physical size or bit pattern.
type OpTypeFilter uint32

func (c OpTypeFilter) Opcode() uint32 { return 15 }

func bindOpTypeFilter(set *InstructionSet) {
	set.Set(
		OpTypeFilter(0).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 1 {
					return nil, ErrMissingInstructionArgs
				}

				return OpTypeFilter(argv[0]), nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(OpTypeFilter)
				out[0] = EncodeOpcode(2, v.Opcode())
				out[1] = uint32(v)
				return nil
			},
		},
	)
}
