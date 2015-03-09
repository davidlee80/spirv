// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// Extension defines the OpExtension instruction.
//
// It declares use of an extension to SPIR-V. This allows validation of
// additional instructions, tokens, semantics, etc
type Extension string

func init() {
	DefaultInstructionSet[OpExtension] = InstructionCodec{
		Decode: decodeOpExtension,
		Encode: encodeOpExtension,
	}
}

func decodeOpExtension(argv []uint32) (Instruction, error) {
	if len(argv) < 1 {
		return nil, ErrMissingInstructionArgs
	}

	return Extension(
		DecodeString(argv),
	), nil
}

func encodeOpExtension(instr Instruction) ([]uint32, error) {
	ext := instr.(Extension)
	out := make([]uint32, EncodedStringLen(string(ext)))
	EncodeString(string(ext), out)
	return out, nil
}
