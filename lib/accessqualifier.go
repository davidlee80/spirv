// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package lib

import "fmt"

// AccessQualifier defines the access permissions of OpTypeSampler
// and OpTypePipe object. Used by OpTypePipe.
type AccessQualifier uint32

// Known execution models.
const (
	AccessReadOnly  AccessQualifier = 0 // A read-only object.
	AccessWriteOnly AccessQualifier = 1 // A write-only object.
	AccessReadWrite AccessQualifier = 2 // A readable and writable object.
)

func (e AccessQualifier) String() string {
	switch e {
	case AccessReadOnly:
		return "Read Only"
	case AccessWriteOnly:
		return "Write Only"
	case AccessReadWrite:
		return "Read Write"
	}

	return fmt.Sprintf("AccessQualifier(%d)", uint32(e))
}
