// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "fmt"

// AddressingMode defines an existing addressing mode.
type AddressingMode uint32

// Known addressing modes.
const (
	AddressLogical    AddressingMode = 0
	AddressPhysical32 AddressingMode = 1
	AddressPhysical64 AddressingMode = 2
)

func (am AddressingMode) String() string {
	switch am {
	case AddressLogical:
		return "Logical"
	case AddressPhysical32:
		return "Physical32"
	case AddressPhysical64:
		return "Physical64"
	}

	return fmt.Sprintf("AddressingMode(%d)", uint32(am))
}
