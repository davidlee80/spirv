// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "fmt"

// EntryPoint represents the OpEntryPoint instruction.
// It declares an entry point and its execution model.
type EntryPoint struct {
	Execution ExecutionModel // Execution model for the entry point and its static call tree.
	Id        uint32         // Must the Result <id> of an OpFunction instruction.
}

func init() {
	DefaultInstructionSet[OpEntryPoint] = InstructionCodec{
		Decode: decodeOpEntryPoint,
		Encode: encodeOpEntryPoint,
	}
}

func decodeOpEntryPoint(argv []uint32) (Instruction, error) {
	if len(argv) < 2 {
		return nil, ErrMissingInstructionArgs
	}

	return &EntryPoint{
		Execution: ExecutionModel(argv[0]),
		Id:        argv[1],
	}, nil
}

func encodeOpEntryPoint(instr Instruction) ([]uint32, error) {
	ep := instr.(*EntryPoint)
	return []uint32{
		uint32(ep.Execution),
		ep.Id,
	}, nil
}

// ExecutionModel defines a single execution model.
// This is used in the EntryPoint instruction to determine what stage of the
// pipeline a given set of instructions belongs to.
type ExecutionModel uint32

// Known execution models.
const (
	ExecVertex                 ExecutionModel = 0 // Vertex shading stage
	ExecTessellationControl    ExecutionModel = 1 // Tessellation control (or hull) shading stage.
	ExecTessellationEvaluation ExecutionModel = 2 // Tessellation evaluation (or domain) shading stage
	ExecGeometry               ExecutionModel = 3 // Geometry shading stage.
	ExecFragment               ExecutionModel = 4 // Fragment shading stage.
	ExecGLCompute              ExecutionModel = 5 // Graphical compute shading stage.
	ExecKernel                 ExecutionModel = 6 // Compute kernel.
)

func (e ExecutionModel) String() string {
	switch e {
	case ExecVertex:
		return "Vertex"
	case ExecTessellationControl:
		return "Tessellation Control"
	case ExecTessellationEvaluation:
		return "Tessellation Evaluation"
	case ExecGeometry:
		return "Geometry"
	case ExecFragment:
		return "Fragment"
	case ExecGLCompute:
		return "GL Compute"
	case ExecKernel:
		return "Kernel"
	}

	return fmt.Sprintf("ExecutionModel(%d)", uint32(e))
}
