// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// ExtInstImport defines the OpExtInstImport instruction.
//
// It defines an import of an extended set of instructions.
// It can later be referenced by the Result <id>
type ExtInstImport struct {
	Name     string
	ResultId uint32
}

func init() {
	DefaultInstructionSet[OpExtInstImport] = InstructionCodec{
		Decode: decodeOpExtInstImport,
		Encode: encodeOpExtInstImport,
	}
}

func decodeOpExtInstImport(argv []uint32) (Instruction, error) {
	if len(argv) < 2 {
		return nil, ErrMissingInstructionArgs
	}

	return &ExtInstImport{
		ResultId: argv[0],
		Name:     DecodeString(argv[1:]),
	}, nil
}

func encodeOpExtInstImport(instr Instruction) ([]uint32, error) {
	ext := instr.(*ExtInstImport)
	out := make([]uint32, 1+EncodedStringLen(ext.Name))
	out[0] = ext.ResultId
	EncodeString(ext.Name, out[1:])
	return out, nil
}
