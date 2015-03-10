// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package lib

import "github.com/jteeuwen/spirv"

// OpMemoryModel represents the OpMemoryModel instruction.
//
// It sets addressing model and memory model for the entire module.
type OpMemoryModel struct {
	Addressing AddressingMode // Selects the module’s addressing model
	Memory     MemoryMode     // Selects the module’s memory model
}

func (c *OpMemoryModel) Opcode() uint32 { return 5 }

// NewOpMemoryModel creates a new codec for the OpMemoryModel instruction.
func NewOpMemoryModel() Codec {
	return Codec{
		Decode: func(argv []uint32) (Instruction, error) {
			if len(argv) < 2 {
				return nil, ErrMissingInstructionArgs
			}

			return &OpMemoryModel{
				Addressing: AddressingMode(argv[0]),
				Memory:     MemoryMode(argv[1]),
			}, nil
		},
		Encode: func(i Instruction, out []uint32) error {
			v := i.(*OpMemoryModel)
			out[0] = spirv.EncodeOpcode(3, 5)
			out[1] = uint32(v.Addressing)
			out[2] = uint32(v.Memory)
			return nil
		},
	}
}
