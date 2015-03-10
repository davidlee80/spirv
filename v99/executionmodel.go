// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package v99

import "fmt"

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
