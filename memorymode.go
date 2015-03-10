// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "fmt"

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
