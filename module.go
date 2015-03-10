// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// Module defines a single compilation unit of SPIR-V. Corresponds to one
// full stage of the graphical pipeline. Corresponds to a fully or partially
// linked OpenCL kernel module with one or more entry points.
type Module struct {
	Header Header
	Code   []Instruction
}
