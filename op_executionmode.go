// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpExecutionMode represents the OpExecutionMode instruction.
// It declares an execution mode for an entry point.
type OpExecutionMode struct {
	EntryPoint uint32        // Must be the Entry Point <id> operand of an OpEntryPoint instruction.
	Mode       ExecutionMode // The execution mode.
	Argv       []uint32      // Literal arguments.
}

func (c *OpExecutionMode) Opcode() uint32 { return 7 }

// NewOpExecutionMode creates a new codec for the OpExecutionMode instruction.
func NewOpExecutionMode() Codec {
	return Codec{
		Decode: func(argv []uint32) (Instruction, error) {
			if len(argv) < 2 {
				return nil, ErrMissingInstructionArgs
			}

			return &OpExecutionMode{
				EntryPoint: argv[0],
				Mode:       ExecutionMode(argv[1]),
				Argv:       argv[2:],
			}, nil
		},
		Encode: func(i Instruction) ([]uint32, error) {
			em := i.(*OpExecutionMode)
			out := make([]uint32, 2+len(em.Argv))
			out[0] = em.EntryPoint
			out[1] = uint32(em.Mode)
			copy(out[2:], em.Argv)
			return out, nil
		},
	}
}
