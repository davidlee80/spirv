// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpSourceExtension defines optional extensions used within the source language.
//
// It documents an extension to the source language. This has no semantic
// impact and can safely be removed from a module.
type OpSourceExtension string

func (c OpSourceExtension) Opcode() uint32 { return 2 }

// NewOpSourceExtension creates a new codec for the OpSourceExtension instruction.
func NewOpSourceExtension() Codec {
	return Codec{
		Decode: func(argv []uint32) (Instruction, error) {
			return OpSourceExtension(
				DecodeString(argv),
			), nil
		},
		Encode: func(i Instruction, out []uint32) error {
			v := i.(OpSourceExtension)
			size := EncodedStringLen(string(v))
			out[0] = EncodeOpcode(uint32(size)+1, v.Opcode())
			EncodeString(string(v), out[1:])
			return nil
		},
	}
}
