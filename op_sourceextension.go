// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// SourceExtension defines optional extensions used within the source language.
//
// It documents an extension to the source language. This has no semantic
// impact and can safely be removed from a module.
type SourceExtension string

func init() {
	DefaultInstructionSet[OpSourceExtension] = InstructionCodec{
		Decode: decodeOpSourceExtension,
		Encode: encodeOpSourceExtension,
	}
}

func decodeOpSourceExtension(argv []uint32) (Instruction, error) {
	return SourceExtension(
		DecodeString(argv),
	), nil
}

func encodeOpSourceExtension(instr Instruction) ([]uint32, error) {
	sext := instr.(SourceExtension)
	out := make([]uint32, EncodedStringLen(string(sext)))
	EncodeString(string(sext), out)
	return out, nil
}
