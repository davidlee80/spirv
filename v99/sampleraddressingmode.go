// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package v99

import "fmt"

// SamplerAddressingMode defines the addressing mode of read image
// extended instructions.
type SamplerAddressingMode uint32

// Known addressing modes.
const (
	// The image coordinates used to sample elements of the image refer to a
	// location inside the image, otherwise the results are undefined.
	SAMNone SamplerAddressingMode = 0

	// Out-of-range image coordinates are clamped to the extent.
	SAMClampEdge SamplerAddressingMode = 2

	// Out-of-range image coordinates will return a border color.
	SAMClamp SamplerAddressingMode = 4

	// Out-of-range image coordinates are wrapped to the valid range.
	// Can only be used with normalized coordinates.
	SAMRepeat SamplerAddressingMode = 6

	// Flip the image coordinate at every integer junction.
	// Can only be used with normalized coordinates.
	SAMRepeatMirrored SamplerAddressingMode = 8
)

func (sam SamplerAddressingMode) String() string {
	switch sam {
	case SAMNone:
		return "None"
	case SAMClampEdge:
		return "Clamp: Edge"
	case SAMClamp:
		return "Clamp"
	case SAMRepeat:
		return "Repeat"
	case SAMRepeatMirrored:
		return "Repeat: Mirrored"
	}

	return fmt.Sprintf("SamplerAddressingMode(%d)", uint32(sam))
}
