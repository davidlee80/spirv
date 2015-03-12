// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpSource defines the OpSource instruction.
//
// It documents what source language a module was translated from.
// This has no semantic impact and can safely be removed from a module.
type OpSource struct {
	SourceLanguage uint32
	Version        uint32
}

func (c *OpSource) Opcode() uint32 { return 1 }

func init() {
	Bind(
		(&OpSource{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 2 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpSource{
					SourceLanguage: argv[0],
					Version:        argv[1],
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpSource)
				out[0] = EncodeOpcode(3, v.Opcode())
				out[1] = v.SourceLanguage
				out[2] = v.Version
				return nil
			},
		},
	)
}
