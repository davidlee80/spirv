// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpMemoryModel represents the OpMemoryModel instruction.
//
// It sets addressing model and memory model for the entire module.
type OpMemoryModel struct {
	AddressingModel AddressingModel // Selects the module’s addressing model
	MemoryModel     MemoryModel     // Selects the module’s memory mode
}

func (c *OpMemoryModel) Opcode() uint32 { return 5 }
func (c *OpMemoryModel) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpMemoryModel{}
	})
}

// OpEntryPoint represents the OpEntryPoint instruction.
// It declares an entry point and its execution model.
type OpEntryPoint struct {
	ExecutionModel ExecutionModel // Execution model for the entry point and its static call tree.
	ResultId       Id             // Must the Result <id> of an OpFunction instruction.
}

func (c *OpEntryPoint) Opcode() uint32 { return 6 }
func (c *OpEntryPoint) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpEntryPoint{}
	})
}

// OpExecutionMode represents the OpExecutionMode instruction.
// It declares an execution mode for an entry point.
type OpExecutionMode struct {
	EntryPoint Id            // Must be the Entry Point <id> operand of an OpEntryPoint instruction.
	Mode       ExecutionMode // The execution mode.
	Argv       []uint32      // Literal arguments.
}

func (c *OpExecutionMode) Opcode() uint32 { return 7 }
func (c *OpExecutionMode) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpExecutionMode{}
	})
}

// OpCompileFlag represents the OpCompileFlag instruction.
type OpCompileFlag struct {
	Flag String
}

func (c *OpCompileFlag) Opcode() uint32 { return 218 }
func (c *OpCompileFlag) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpCompileFlag{}
	})
}
