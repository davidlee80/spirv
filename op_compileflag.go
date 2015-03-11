// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpCompileFlag represents the OpCompileFlag instruction.
type OpCompileFlag String

func (c OpCompileFlag) Opcode() uint32 { return 218 }

func bindOpCompileFlag(set *InstructionSet) {
	set.Set(
		OpCompileFlag("").Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) == 0 {
					return nil, ErrMissingInstructionArgs
				}

				return OpCompileFlag(
					DecodeString(argv),
				), nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(OpCompileFlag)
				size := String(v).EncodedLen()
				out[0] = EncodeOpcode(size+1, v.Opcode())
				String(v).Encode(out[1:])
				return nil
			},
		},
	)
}
