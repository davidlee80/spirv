// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "fmt"

// FPRoundingMode associates a rounding mode to a floating-point
// conversion instruction.
//
// By default:
//
//    - Conversions from floating-point to integer types use the
//      round-toward-zero rounding mode.
//    - Conversions to floating-point types use the round-to-nearest-even
//      rounding mode.
//
type FPRoundingMode uint32

// Known rounding modes.
const (
	RTE FPRoundingMode = 0 // Round to nearest even.
	RTZ FPRoundingMode = 1 // Round towards zero.
	RTP FPRoundingMode = 2 // Round towards positive infinity.
	RTN FPRoundingMode = 3 // Round towards negative infinity.
)

func (r FPRoundingMode) String() string {
	switch r {
	case RTE:
		return "Nearest Even"
	case RTZ:
		return "Zero"
	case RTP:
		return "Positive Infinity"
	case RTN:
		return "Negative Infinity"
	}

	return fmt.Sprintf("FPRoundingMode(%d)", uint32(r))
}
