// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpCompileFlag represents the OpCompileFlag instruction.
type OpCompileFlag string

func (c OpCompileFlag) Opcode() uint32 { return 218 }

// NewOpCompileFlag creates a new codec for the OpCompileFlag instruction.
func NewOpCompileFlag() Codec {
	return Codec{
		Decode: func(argv []uint32) (Instruction, error) {
			return OpCompileFlag(
				DecodeString(argv),
			), nil
		},
		Encode: func(i Instruction, out []uint32) error {
			v := i.(OpCompileFlag)
			size := EncodedStringLen(string(v))
			out[0] = EncodeOpcode(uint32(size+1), v.Opcode())
			EncodeString(string(v), out[1:])
			return nil
		},
	}
}
