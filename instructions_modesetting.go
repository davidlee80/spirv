// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpMemoryModel represents the OpMemoryModel instruction.
//
// It sets addressing model and memory model for the entire module.
type OpMemoryModel struct {
	AddressingMode uint32 // Selects the module’s addressing model
	MemoryMode     uint32 // Selects the module’s memory mode
}

func (c *OpMemoryModel) Opcode() uint32 { return 5 }

func init() {
	Bind(
		(&OpMemoryModel{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 2 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpMemoryModel{
					AddressingMode: argv[0],
					MemoryMode:     argv[1],
				}, nil
			},
		},
	)
}

// OpEntryPoint represents the OpEntryPoint instruction.
// It declares an entry point and its execution model.
type OpEntryPoint struct {
	ExecutionModel uint32 // Execution model for the entry point and its static call tree.
	ResultId       uint32 // Must the Result <id> of an OpFunction instruction.
}

func (c *OpEntryPoint) Opcode() uint32 { return 6 }

func init() {
	Bind(
		(&OpEntryPoint{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 2 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpEntryPoint{
					ExecutionModel: argv[0],
					ResultId:       argv[1],
				}, nil
			},
		},
	)
}

// OpExecutionMode represents the OpExecutionMode instruction.
// It declares an execution mode for an entry point.
type OpExecutionMode struct {
	EntryPoint    uint32   // Must be the Entry Point <id> operand of an OpEntryPoint instruction.
	ExecutionMode uint32   // The execution mode.
	Argv          []uint32 // Literal arguments.
}

func (c *OpExecutionMode) Opcode() uint32 { return 7 }

func init() {
	Bind(
		(&OpExecutionMode{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 2 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpExecutionMode{
					EntryPoint:    argv[0],
					ExecutionMode: argv[1],
					Argv:          Copy(argv[2:]),
				}, nil
			},
		},
	)
}

// OpCompileFlag represents the OpCompileFlag instruction.
type OpCompileFlag String

func (c OpCompileFlag) Opcode() uint32 { return 218 }

func init() {
	Bind(
		OpCompileFlag("").Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) == 0 {
					return nil, ErrMissingInstructionArgs
				}

				return OpCompileFlag(
					DecodeString(argv),
				), nil
			},
		},
	)
}
