// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package v99

import "github.com/jteeuwen/spirv"

// OpEntryPoint represents the OpEntryPoint instruction.
// It declares an entry point and its execution model.
type OpEntryPoint struct {
	Execution ExecutionModel // Execution model for the entry point and its static call tree.
	Id        uint32         // Must the Result <id> of an OpFunction instruction.
}

func (c *OpEntryPoint) Opcode() uint32 { return 6 }

// NewOpEntryPoint creates a new codec for the OpEntryPoint instruction.
func NewOpEntryPoint() spirv.Codec {
	return spirv.Codec{
		Decode: func(argv []uint32) (spirv.Instruction, error) {
			if len(argv) < 2 {
				return nil, spirv.ErrMissingInstructionArgs
			}

			return &OpEntryPoint{
				Execution: ExecutionModel(argv[0]),
				Id:        argv[1],
			}, nil
		},
		Encode: func(i spirv.Instruction, out []uint32) error {
			v := i.(*OpEntryPoint)
			out[0] = spirv.EncodeOpcode(3, 6)
			out[1] = uint32(v.Execution)
			out[2] = v.Id
			return nil
		},
	}
}
