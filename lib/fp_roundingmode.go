// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package lib

import "fmt"

// ExecutionModel defines a single execution model.
// This is used in the EntryPoint instruction to determine what stage of the
// pipeline a given set of instructions belongs to.
type RoundingMode uint32

// Known execution models.
const (
	RTE RoundingMode = 0 // Round to nearest even.
	RTZ RoundingMode = 1 // Round towards zero.
	RTP RoundingMode = 2 // Round towards positive infinity.
	RTN RoundingMode = 3 // Round towards negative infinity.
)

func (r RoundingMode) String() string {
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

	return fmt.Sprintf("RoundingMode(%d)", uint32(r))
}
