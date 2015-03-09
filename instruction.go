// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// Codec defines handlers used to encode or decode
// a the operand list for a specific type of instruction
type Codec struct {
	// Decoder for an instruction's arguments.
	//
	// The provided set of arguments is guaranteed to have the size
	// defined in the instruction's word count. However, this does not
	// mean it is the amount actually expected by the instruction. So
	// a size check on this slice is warrented. ErrMissingInstructionArgs
	// should returned if this check fails.
	Decode func(argv []uint32) (interface{}, error)

	// Encoder for an instruction's arguments.
	//
	// The word set being returned must not include the instruction's
	// first word with opcode or word count. This will be generated
	// by the module encoder.
	Encode func(interface{}) ([]uint32, error)
}

// InstructionSet maps opcodes to an instruction encoder/decoder.
type InstructionSet map[uint32]Codec

// Known opcode values as defined in the specification.
const (
	opNop             = 0
	opSource          = 1
	opSourceExtension = 2
	opExtension       = 3
	opExtInstImport   = 4
	opMemoryModel     = 5
	opEntryPoint      = 6
	opExecutionMode   = 7
	opTypeVoid        = 8
	opTypeBool        = 9
	opTypeInt         = 10
	opExtInst         = 44
	opCompileFlag     = 218
)

// DefaultInstructionSet defines a map with the default instruction set.
// It is filled with all known operations at init time and can be extended
// with custom instructions if needed.
//
// It can be passed into an encoder or decoder.
var DefaultInstructionSet = make(InstructionSet)
