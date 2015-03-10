// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package v99

import "fmt"

// SamplerAddressingMode defines the filter mode of read image
// extended instructions.
type SamplerFilterMode uint32

// Known addressing modes.
const (
	// Use filter nearset mode when performing a read image operation.
	SFMNearest SamplerFilterMode = 16

	// Use filter linear mode when performing a read image operation.
	SFMLinear SamplerFilterMode = 32
)

func (sfm SamplerFilterMode) String() string {
	switch sfm {
	case SFMNearest:
		return "Nearest"
	case SFMLinear:
		return "Linear"
	}

	return fmt.Sprintf("SamplerFilterMode(%d)", uint32(sfm))
}
