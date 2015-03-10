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
			cf := i.(OpCompileFlag)
			size := EncodedStringLen(string(cf))
			out[0] = EncodeOpcode(size+1, 218)
			EncodeString(string(cf), out[1:])
			return nil
		},
	}
}
