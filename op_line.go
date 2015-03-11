// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpLine defines the OpLine instruction.
//
// It adds source-level location information.
// This has no semantic impact and can safely be removed from a module.
type OpLine struct {
	Target uint32
	File   uint32
	Line   uint32
	Column uint32
}

func (c *OpLine) Opcode() uint32 { return 57 }

// NewOpLine creates a new codec for the OpLine instruction.
func NewOpLine() Codec {
	return Codec{
		Decode: func(argv []uint32) (Instruction, error) {
			if len(argv) != 4 {
				return nil, ErrInvalidInstructionSize
			}

			return &OpLine{
				Target: argv[0],
				File:   argv[1],
				Line:   argv[2],
				Column: argv[3],
			}, nil
		},
		Encode: func(i Instruction, out []uint32) error {
			v := i.(*OpLine)
			out[0] = EncodeOpcode(5, v.Opcode())
			out[1] = v.Target
			out[2] = v.File
			out[3] = v.Line
			out[4] = v.Column
			return nil
		},
	}
}
