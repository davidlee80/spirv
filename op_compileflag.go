// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpCompileFlag represents the OpCompileFlag instruction.
type OpCompileFlag string

func init() {
	DefaultInstructionSet[opCompileFlag] = Codec{
		Decode: func(argv []uint32) (interface{}, error) {
			return OpCompileFlag(
				DecodeString(argv),
			), nil
		},
		Encode: func(instr interface{}) ([]uint32, error) {
			cf := instr.(OpCompileFlag)
			size := EncodedStringLen(string(cf))
			out := make([]uint32, size)
			EncodeString(string(cf), out)
			return out, nil
		},
	}
}
