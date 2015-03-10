// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "fmt"

// Dim defines the dimensionality of a texture.
//
// Used by OpTypeSampler.
type Dim uint32

// Known execution models.
const (
	Dim1D     Dim = 0
	Dim2D     Dim = 1
	Dim3D     Dim = 2
	DimCube   Dim = 3
	DimRect   Dim = 4
	DimBuffer Dim = 5
)

func (d Dim) String() string {
	switch d {
	case Dim1D:
		return "1D"
	case Dim2D:
		return "2D"
	case Dim3D:
		return "3D"
	case DimCube:
		return "Cube"
	case DimRect:
		return "Rect"
	case DimBuffer:
		return "Buffer"
	}

	return fmt.Sprintf("Dim(%d)", uint32(d))
}
