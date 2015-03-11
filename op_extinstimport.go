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

func (c *OpExtInstImport) Opcode() uint32 { return 4 }

// NewOpExtInstImport creates a new codec for the OpExtInstImport instruction.
func NewOpExtInstImport() Codec {
	return Codec{
		Decode: func(argv []uint32) (Instruction, error) {
			if len(argv) < 2 {
				return nil, ErrMissingInstructionArgs
			}

			return &OpExtInstImport{
				ResultId: argv[0],
				Name:     DecodeString(argv[1:]),
			}, nil
		},
		Encode: func(i Instruction, out []uint32) error {
			v := i.(*OpExtInstImport)
			size := EncodedStringLen(v.Name)

			out[0] = EncodeOpcode(2+uint32(size), v.Opcode())
			out[1] = v.ResultId
			EncodeString(v.Name, out[2:])
			return nil
		},
	}
}
