// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// InstructionBuilder creates a specific instruction from the given
// set of arguments.
type InstructionBuilder struct {
	// Decoder for an instruction's arguments.
	//
	// The provided set of arguments is guaranteed to have the size
	// defined in the instruction's word count. However, this does not
	// mean it is the amount actually expected by the instruction. So
	// a size check on this slice is warrented. ErrMissingInstructionArgs
	// should returned if this check fails.
	Decode func(argv []uint32) (Instruction, error)

	// Encoder for an instruction's arguments.
	//
	// The word set being returned must not include the instruction's
	// first word with opcode or word count. This will be generated
	// by the module encoder.
	Encode func(Instruction) ([]uint32, error)
}

// InstructionSet maps opcodes to an insutrction constructor.
type InstructionSet map[uint32]InstructionBuilder

// Instruction defines a generic instruction.
type Instruction interface{}

// Known opcode values as defined in the specification.
const (
	OpNop             = 0x0000
	OpSource          = 0x0001
	OpSourceExtension = 0x0002
	OpExtension       = 0x0003
	OpExtInstImport   = 0x0004
	OpExtInst         = 0x0044
	OpMemoryModel     = 0x0005
	OpEntryPoint      = 0x0006
)

// DefaultInstructionSet defines a map with the default instruction set.
// It is filled with all known operations at init time and can be extended
// with custom instructions if needed.
//
// It can be passed into an encoder or decoder.
var DefaultInstructionSet = make(InstructionSet)
