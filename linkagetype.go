// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "fmt"

// LinkageType associates a linkage type to functions or global
// variables. By default, functions and global variables are private
// to a module and cannot be accessed by other modules.
type LinkageType uint32

// Known execution models.
const (
	LinkExport LinkageType = 0 // Accessible by other modules as well.
	LinkImport LinkageType = 1 // Declaration for a global identifier that exists in another module.
)

func (e LinkageType) String() string {
	switch e {
	case LinkExport:
		return "Export"
	case LinkImport:
		return "Import"
	}

	return fmt.Sprintf("LinkageType(%d)", uint32(e))
}
