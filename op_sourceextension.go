// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpSourceExtension defines optional extensions used within the source language.
//
// It documents an extension to the source language. This has no semantic
// impact and can safely be removed from a module.
type OpSourceExtension string

func init() {
	DefaultInstructionSet[opSourceExtension] = Codec{
		Decode: func(argv []uint32) (interface{}, error) {
			return OpSourceExtension(
				DecodeString(argv),
			), nil
		},
		Encode: func(instr interface{}) ([]uint32, error) {
			sext := instr.(OpSourceExtension)
			out := make([]uint32, EncodedStringLen(string(sext)))
			EncodeString(string(sext), out)
			return out, nil
		},
	}
}
