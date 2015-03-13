// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpVariable allocates an object in memory, resulting in a pointer
// to it, which can be used with OpLoad and OpStore.
type OpVariable struct {
	// Result Type is a type from OpTypePointer, where the type pointed to
	// is the type of object in memory.
	ResultType uint32

	ResultId uint32

	// Storage Class is the kind of memory holding the object.
	StorageClass uint32

	// Initializer is optional. If Initializer is present, it will be the
	// initial value of the variable’s memory content. Initializer must
	// be an <id> from a constant instruction. Initializer must have the same
	// type as the type pointed to by Result Type.
	Initializer uint32 `spirv:"optional"`
}

func (c *OpVariable) Opcode() uint32 { return 38 }
func (c *OpVariable) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpVariable{
			Initializer: 0,
		}
	})
}

// OpVariableArray allocates N objects sequentially in memory,
// resulting in a pointer to the first such object.
//
// This is not the same thing as allocating a single object that is an array.
type OpVariableArray struct {
	// Result Type is a type from OpTypePointer whose type pointed to is
	// the type of one of the N objects allocated in memory
	ResultType uint32

	ResultId uint32

	// Storage Class is the kind of memory holding the object.
	StorageClass uint32

	// N is the number of objects to allocate.
	N uint32
}

func (c *OpVariableArray) Opcode() uint32 { return 39 }
func (c *OpVariableArray) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpVariableArray{}
	})
}

// OpLoad loads data through a pointer.
type OpLoad struct {
	// Result Type is a type from OpTypePointer whose type pointed to is
	// the type of one of the N objects allocated in memory
	ResultType uint32

	ResultId uint32

	// Pointer is the pointer to load through. It must have a type of
	// OpTypePointer whose operand is the same as Result Type.
	Pointer uint32

	// MemoryAccess must be one or more Memory Access literals.
	MemoryAccess []uint32
}

func (c *OpLoad) Opcode() uint32 { return 46 }
func (c *OpLoad) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpLoad{}
	})
}

// OpStore stores data through a pointer.
type OpStore struct {
	// Pointer is the pointer to store through. It must have a type
	// of OpTypePointer whose operand is the same as the type of Object.
	Pointer uint32

	// Object is the object to store.
	Object uint32

	// MemoryAccess must be one or more Memory Access literals.
	MemoryAccess []uint32
}

func (c *OpStore) Opcode() uint32 { return 47 }
func (c *OpStore) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpStore{}
	})
}

// OpCopyMemory copies from the memory pointed to by Source to the
// memory pointed to by Target.
//
// Both operands must be non-void pointers of the same type.
// Matching storage class is not required. The amount of memory copied is
// the size of the type pointed to.
type OpCopyMemory struct {
	// The target address.
	Target uint32

	// The source address.
	Source uint32

	// MemoryAccess must be one or more Memory Access literals.
	MemoryAccess []uint32
}

func (c *OpCopyMemory) Opcode() uint32 { return 65 }
func (c *OpCopyMemory) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpCopyMemory{}
	})
}

// OpCopyMemorySized copies from the memory pointed to by Source to the
// memory pointed to by Target.
//
// Both operands must be non-void pointers of the same type.
// Matching storage class is not required.
type OpCopyMemorySized struct {
	// The target address.
	Target uint32

	// The source address.
	Source uint32

	// Size is the number of bytes to copy.
	Size uint32

	// MemoryAccess must be one or more Memory Access literals.
	MemoryAccess []uint32
}

func (c *OpCopyMemorySized) Opcode() uint32 { return 66 }
func (c *OpCopyMemorySized) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpCopyMemorySized{}
	})
}

// OpAccessChain creates a pointer into a composite object that can be
// used with OpLoad and OpStore.
//
// The storage class of the pointer created will be the same as the storage
// class of the base operand.
type OpAccessChain struct {
	ResultType uint32
	ResultId   uint32

	// Base must be a pointer type, pointing to the base of the object.
	Base uint32

	// Indices walk the type hierarchy to the desired depth, potentially
	// down to scalar granularity. The type of the pointer created will be to
	// the type reached by walking the type hierarchy down to the last
	// provided index.
	Indices []uint32
}

func (c *OpAccessChain) Opcode() uint32 { return 93 }
func (c *OpAccessChain) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpAccessChain{}
	})
}

// OpInboundsAccessChain has the same semantics as OpAccessChain, with the
// addition that the resulting pointer is known to point within the base object.
type OpInboundsAccessChain struct {
	ResultType uint32
	ResultId   uint32

	// Base must be a pointer type, pointing to the base of the object.
	Base uint32

	// Indices walk the type hierarchy to the desired depth, potentially
	// down to scalar granularity. The type of the pointer created will be to
	// the type reached by walking the type hierarchy down to the last
	// provided index.
	Indices []uint32
}

func (c *OpInboundsAccessChain) Opcode() uint32 { return 94 }
func (c *OpInboundsAccessChain) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpInboundsAccessChain{}
	})
}

// OpArraylength results in the length of a run-time array.
type OpArraylength struct {
	ResultType uint32
	ResultId   uint32

	// Structure must be an object of type OpTypeStruct that contains
	// a member that is a run-time array.
	Structure uint32

	// Array member is a member number of Structure that must have a
	// type from OpTypeRuntimeArray.
	Member uint32
}

func (c *OpArraylength) Opcode() uint32 { return 121 }
func (c *OpArraylength) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpArraylength{}
	})
}

// OpImagePointer forms a pointer to a texel of an image.
// Use of such a pointer is limited to atomic operations.
//
// TODO: This requires an Image storage class to be added.
type OpImagePointer struct {
	ResultType uint32
	ResultId   uint32

	// Image is a pointer to a variable of type of OpTypeSampler.
	Image uint32

	// Coordinate and Sample specify which texel and sample within
	// the image to form an address of.
	Coordinate uint32
	Sample     uint32
}

func (c *OpImagePointer) Opcode() uint32 { return 190 }
func (c *OpImagePointer) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpImagePointer{}
	})
}

// OpGenericPtrMemSemantics returns a valid Memory Semantics
// value for ptr.
type OpGenericPtrMemSemantics struct {
	ResultType uint32 // Result Type must be a 32-bits wide OpTypeInt value
	ResultId   uint32
	Ptr        uint32 // Ptr must point to Generic.
}

func (c *OpGenericPtrMemSemantics) Opcode() uint32 { return 233 }
func (c *OpGenericPtrMemSemantics) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpGenericPtrMemSemantics{}
	})
}
