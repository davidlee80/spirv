// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import (
	"fmt"
	"strings"
)

// FPFastMathMode defines bitflags which enable fast math operations
// which are otherwise unsafe.
//
// Only valid on OpFAdd, OpFSub, OpFMul, OpFDiv, OpFRem
// and OpFMod instructions.
type FPFastMathMode uint32

// Known fast math modes.
const (
	// Assume parameters and result are not NaN.
	FMNotNaN FPFastMathMode = 0

	// Assume parameters and result are not +/- Inf.
	FMNotInf FPFastMathMode = 2

	// Treat the sign of a zero parameter or result as insignificant.
	FMNSZ FPFastMathMode = 4

	// Allow the usage of reciprocal rather than perform a division.
	FMAllowRecip FPFastMathMode = 8

	// Allow algebraic transformations according to real-number associative
	// and distributive algebra. This flag implies all the others.
	FMFast FPFastMathMode = 16
)

func (fm FPFastMathMode) String() string {
	set := make([]string, 0, 5)

	if fm&FMNotNaN != 0 {
		set = append(set, "Not NaN")
	}

	if fm&FMNotInf != 0 {
		set = append(set, "Not Inf")
	}

	if fm&FMNSZ != 0 {
		set = append(set, "Non-Significant Sign")
	}

	if fm&FMAllowRecip != 0 {
		set = append(set, "Allow Reciprocal")
	}

	if fm&FMFast != 0 {
		set = append(set, "Allow Reciprocal")
	}

	if len(set) == 0 {
		return fmt.Sprintf("FPFastMathMode(%d)", uint32(fm))
	}

	return strings.Join(set, ", ")
}
