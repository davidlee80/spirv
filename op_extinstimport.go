// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpExtInstImport defines the OpExtInstImport instruction.
//
// It defines an import of an extended set of instructions.
// It can later be referenced by the Result <id>
type OpExtInstImport struct {
	Name     string
	ResultId uint32
}

func init() {
	DefaultInstructionSet[opExtInstImport] = Codec{
		Decode: func(argv []uint32) (interface{}, error) {
			if len(argv) < 2 {
				return nil, ErrMissingInstructionArgs
			}

			return &OpExtInstImport{
				ResultId: argv[0],
				Name:     DecodeString(argv[1:]),
			}, nil
		},
		Encode: func(instr interface{}) ([]uint32, error) {
			ext := instr.(*OpExtInstImport)
			out := make([]uint32, 1+EncodedStringLen(ext.Name))
			out[0] = ext.ResultId
			EncodeString(ext.Name, out[1:])
			return out, nil
		},
	}
}
