// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpExtension defines the OpExtension instruction.
//
// It declares use of an extension to SPIR-V. This allows validation of
// additional instructions, tokens, semantics, etc
type OpExtension string

func init() {
	DefaultInstructionSet[opExtension] = Codec{
		Decode: func(argv []uint32) (interface{}, error) {
			if len(argv) < 1 {
				return nil, ErrMissingInstructionArgs
			}

			return OpExtension(
				DecodeString(argv),
			), nil
		},
		Encode: func(instr interface{}) ([]uint32, error) {
			ext := instr.(OpExtension)
			out := make([]uint32, EncodedStringLen(string(ext)))
			EncodeString(string(ext), out)
			return out, nil
		},
	}
}
