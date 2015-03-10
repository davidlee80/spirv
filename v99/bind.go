// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package v99

import "github.com/jteeuwen/spirv"

// Bind adds all instruction codecs to the given set.
func Bind(set *spirv.InstructionSet) {
	set.Set(1, NewOpSource())
	set.Set(2, NewOpSourceExtension())
	set.Set(3, NewOpExtension())
	set.Set(4, NewOpExtInstImport())
	set.Set(5, NewOpMemoryModel())
	set.Set(6, NewOpEntryPoint())
	set.Set(7, NewOpExecutionMode())
	set.Set(8, NewOpTypeVoid())
	set.Set(9, NewOpTypeBool())
	set.Set(10, NewOpTypeInt())
	set.Set(11, NewOpTypeFloat())
	set.Set(12, NewOpTypeVector())
	set.Set(44, NewOpExtInst())
	set.Set(218, NewOpCompileFlag())
}
