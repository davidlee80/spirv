// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package lib

import "github.com/jteeuwen/spirv"

// Module defines a single compilation unit of SPIR-V. Corresponds to one
// full stage of the graphical pipeline. Corresponds to a fully or partially
// linked OpenCL kernel module with one or more entry points.
type Module struct {
	Header spirv.Header
	Code   []Instruction
}
