// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package v99

import "github.com/jteeuwen/spirv"

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
func NewOpSource() spirv.Codec {
	return spirv.Codec{
		Decode: func(argv []uint32) (spirv.Instruction, error) {
			if len(argv) < 2 {
				return nil, spirv.ErrMissingInstructionArgs
			}

			return &OpSource{
				Language: SourceLanguage(argv[0]),
				Version:  argv[1],
			}, nil
		},
		Encode: func(i spirv.Instruction, out []uint32) error {
			v := i.(*OpSource)
			out[0] = spirv.EncodeOpcode(3, 1)
			out[1] = uint32(v.Language)
			out[2] = uint32(v.Version)
			return nil
		},
	}
}
