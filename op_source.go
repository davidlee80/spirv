// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpSource defines the OpSource instruction.
//
// It documents what source language a module was translated from.
// This has no semantic impact and can safely be removed from a module.
type OpSource struct {
	Language SourceLanguage
	Version  uint32
}

func (c *OpSource) Opcode() uint32 { return 1 }

// NewOpSource creates a new codec for the OpSource instruction.
func NewOpSource() Codec {
	return Codec{
		Decode: func(argv []uint32) (Instruction, error) {
			if len(argv) < 2 {
				return nil, ErrMissingInstructionArgs
			}

			return &OpSource{
				Language: SourceLanguage(argv[0]),
				Version:  argv[1],
			}, nil
		},
		Encode: func(i Instruction) ([]uint32, error) {
			src := i.(*OpSource)
			return []uint32{
				uint32(src.Language),
				src.Version,
			}, nil
		},
	}
}
