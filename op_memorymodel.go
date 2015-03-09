// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "fmt"

// MemoryModel defines represents the OpMemoryModel instruction.
//
// It sets addressing model and memory model for the entire module.
type MemoryModel struct {
	Addressing AddressingMode // Selects the module’s addressing model
	Memory     MemoryMode     // Selects the module’s memory model
}

func init() {
	DefaultInstructionSet[OpMemoryModel] = InstructionBuilder{
		Decode: decodeOpMemoryModel,
		Encode: encodeOpMemoryModel,
	}
}

func decodeOpMemoryModel(argv []uint32) (Instruction, error) {
	if len(argv) < 2 {
		return nil, ErrMissingInstructionArgs
	}

	return &MemoryModel{
		Addressing: AddressingMode(argv[0]),
		Memory:     MemoryMode(argv[1]),
	}, nil
}

func encodeOpMemoryModel(instr Instruction) ([]uint32, error) {
	mm := instr.(*MemoryModel)
	return []uint32{
		uint32(mm.Addressing),
		uint32(mm.Memory),
	}, nil
}

// AddressingMode defines an existing addressing mode.
type AddressingMode uint32

// Known addressing modes.
const (
	AddressLogical    AddressingMode = 0
	AddressPhysical32 AddressingMode = 1
	AddressPhysical64 AddressingMode = 2
)

func (am AddressingMode) String() string {
	switch am {
	case AddressLogical:
		return "Logical"
	case AddressPhysical32:
		return "Physical32"
	case AddressPhysical64:
		return "Physical64"
	}

	return fmt.Sprintf("AddressingMode(%d)", uint32(am))
}

// MemoryMode defines an existing memory model.
type MemoryMode uint32

// Known addressing modes.
const (
	MemorySimple   MemoryMode = 0 // No shared memory consistency issues.
	MemoryGLSL450  MemoryMode = 1 // Memory model needed by later versions of GLSL and ESSL. Works across multiple versions.
	MemoryOpenCL12 MemoryMode = 2 // OpenCL 1.2 memory model.
	MemoryOpenCL20 MemoryMode = 3 // OpenCL 2.0 memory model.
	MemoryOpenCL21 MemoryMode = 4 // OpenCL 2.1 memory model.
)

func (mm MemoryMode) String() string {
	switch mm {
	case MemorySimple:
		return "Simple"
	case MemoryGLSL450:
		return "GLSL450"
	case MemoryOpenCL12:
		return "OpenCL1.2"
	case MemoryOpenCL20:
		return "OpenCL2.0"
	case MemoryOpenCL21:
		return "OpenCL2.1"
	}

	return fmt.Sprintf("MemoryMode(%d)", uint32(mm))
}
