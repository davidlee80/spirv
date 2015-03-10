// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package v99

import "fmt"

// StorageClass defines a class of storage for declared variables
// (does not include intermediate values).
//
// Used by: OpTypePointer, OpTypeVariable, OpTypeVariableArray,
// OpTypeGenericCastToPtrExplicit
type StorageClass uint32

// Known storage classes
const (
	// Shared externally, read-only memory, visible across all instantiation
	// or work groups. Graphics uniform memory. OpenCL Constant memory
	StoreUniformConstant StorageClass = 0

	// Input from pipeline. Read only
	StoreInput StorageClass = 1

	// Shared externally, visible across all instantiations or work groups
	StoreUniform StorageClass = 2

	// Output to pipeline.
	StoreOutput StorageClass = 3

	// Shared across all work items within a work group. OpenGL "shared".
	// OpenCL local memory.
	StoreWorkgroupLocal StorageClass = 4

	// Visible across all work items of all work groups. OpenCL global memory.
	StoreWorkgroupGlobal StorageClass = 5

	// Accessible across functions within a module, non-IO (not visible outside
	// the module).
	StorePrivateGlobal StorageClass = 6

	// A variable local to a function.
	StoreFunction StorageClass = 7

	// A generic pointer, which overloads StoragePrivate, StorageLocal,
	// StorageGlobal. not a real storage class.
	StoreGeneric StorageClass = 8

	// Private to a work-item and is not visible to another work-item.
	// OpenCL private memory.
	StorePrivate StorageClass = 9

	// For holding atomic counters.
	StoreAtomicCounter StorageClass = 10
)

func (s StorageClass) String() string {
	switch s {
	case StoreUniformConstant:
		return "Uniform Constant"
	case StoreInput:
		return "Input"
	case StoreUniform:
		return "Uniform"
	case StoreOutput:
		return "Output"
	case StoreWorkgroupLocal:
		return "Workgroup: Local"
	case StoreWorkgroupGlobal:
		return "Workgroup: Global"
	case StorePrivateGlobal:
		return "Private: Global"
	case StoreFunction:
		return "Function"
	case StoreGeneric:
		return "Generic"
	case StorePrivate:
		return "Private"
	case StoreAtomicCounter:
		return "Atomic Counter"
	}

	return fmt.Sprintf("StorageClass(%d)", uint32(s))
}
