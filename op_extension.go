// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpExtension defines the OpExtension instruction.
//
// It declares use of an extension to SPIR-V. This allows validation of
// additional instructions, tokens, semantics, etc
type OpExtension String

func (c OpExtension) Opcode() uint32 { return 3 }

// NewOpExtension creates a new codec for the OpExtension instruction.
func NewOpExtension() Codec {
	return Codec{
		Decode: func(argv []uint32) (Instruction, error) {
			if len(argv) < 1 {
				return nil, ErrMissingInstructionArgs
			}

			return OpExtension(
				DecodeString(argv),
			), nil
		},
		Encode: func(i Instruction, out []uint32) error {
			v := i.(OpExtension)
			size := String(v).EncodedLen()
			out[0] = EncodeOpcode(uint32(size)+1, v.Opcode())
			String(v).Encode(out[1:])
			return nil
		},
	}
}
