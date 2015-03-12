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

func init() {
	Bind(
		(&OpVariable{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 3 {
					return nil, ErrMissingInstructionArgs
				}

				op := &OpVariable{
					ResultType:   argv[0],
					ResultId:     argv[1],
					StorageClass: argv[2],
				}

				if len(argv) > 3 {
					op.Initializer = argv[3]
				}

				return op, nil
			},
			Encode: func(i Instruction, out []uint32) (uint32, error) {
				v := i.(*OpVariable)
				size := uint32(3)

				if v.Initializer != 0 {
					size++
				}

				out[0] = v.ResultType
				out[1] = v.ResultId
				out[2] = v.StorageClass

				if v.Initializer != 0 {
					out[3] = v.Initializer
				}

				return size, nil
			},
		},
	)
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

func init() {
	Bind(
		(&OpVariableArray{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 4 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpVariableArray{
					ResultType:   argv[0],
					ResultId:     argv[1],
					StorageClass: argv[2],
					N:            argv[3],
				}, nil
			},
			Encode: func(i Instruction, out []uint32) (uint32, error) {
				v := i.(*OpVariableArray)
				out[0] = v.ResultType
				out[1] = v.ResultId
				out[2] = v.StorageClass
				out[3] = v.N
				return 4, nil
			},
		},
	)
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

func init() {
	Bind(
		(&OpLoad{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 3 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpLoad{
					ResultType:   argv[0],
					ResultId:     argv[1],
					Pointer:      argv[2],
					MemoryAccess: Copy(argv[3:]),
				}, nil
			},
			Encode: func(i Instruction, out []uint32) (uint32, error) {
				v := i.(*OpLoad)
				size := uint32(len(v.MemoryAccess))

				out[0] = v.ResultType
				out[1] = v.ResultId
				out[2] = v.Pointer
				copy(out[3:], v.MemoryAccess)
				return 3 + size, nil
			},
		},
	)
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

func init() {
	Bind(
		(&OpStore{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 2 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpStore{
					Pointer:      argv[0],
					Object:       argv[1],
					MemoryAccess: Copy(argv[2:]),
				}, nil
			},
			Encode: func(i Instruction, out []uint32) (uint32, error) {
				v := i.(*OpStore)
				size := uint32(len(v.MemoryAccess))

				out[0] = v.Pointer
				out[1] = v.Object
				copy(out[2:], v.MemoryAccess)
				return 2 + size, nil
			},
		},
	)
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

func init() {
	Bind(
		(&OpCopyMemory{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 2 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpCopyMemory{
					Target:       argv[0],
					Source:       argv[1],
					MemoryAccess: Copy(argv[2:]),
				}, nil
			},
			Encode: func(i Instruction, out []uint32) (uint32, error) {
				v := i.(*OpCopyMemory)
				size := uint32(len(v.MemoryAccess))

				out[0] = v.Target
				out[1] = v.Source
				copy(out[2:], v.MemoryAccess)
				return 2 + size, nil
			},
		},
	)
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

func init() {
	Bind(
		(&OpCopyMemorySized{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 3 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpCopyMemorySized{
					Target:       argv[0],
					Source:       argv[1],
					Size:         argv[2],
					MemoryAccess: Copy(argv[3:]),
				}, nil
			},
			Encode: func(i Instruction, out []uint32) (uint32, error) {
				v := i.(*OpCopyMemorySized)
				size := uint32(len(v.MemoryAccess))

				out[0] = v.Target
				out[1] = v.Source
				out[2] = v.Size
				copy(out[3:], v.MemoryAccess)
				return 3 + size, nil
			},
		},
	)
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

func init() {
	Bind(
		(&OpAccessChain{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 3 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpAccessChain{
					ResultType: argv[0],
					ResultId:   argv[1],
					Base:       argv[2],
					Indices:    Copy(argv[3:]),
				}, nil
			},
			Encode: func(i Instruction, out []uint32) (uint32, error) {
				v := i.(*OpAccessChain)
				size := uint32(len(v.Indices))

				out[0] = v.ResultType
				out[1] = v.ResultId
				out[2] = v.Base
				copy(out[3:], v.Indices)
				return 3 + size, nil
			},
		},
	)
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

func init() {
	Bind(
		(&OpInboundsAccessChain{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 3 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpInboundsAccessChain{
					ResultType: argv[0],
					ResultId:   argv[1],
					Base:       argv[2],
					Indices:    Copy(argv[3:]),
				}, nil
			},
			Encode: func(i Instruction, out []uint32) (uint32, error) {
				v := i.(*OpInboundsAccessChain)
				size := uint32(len(v.Indices))

				out[0] = v.ResultType
				out[1] = v.ResultId
				out[2] = v.Base
				copy(out[3:], v.Indices)
				return 3 + size, nil
			},
		},
	)
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

func init() {
	Bind(
		(&OpArraylength{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 4 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpArraylength{
					ResultType: argv[0],
					ResultId:   argv[1],
					Structure:  argv[2],
					Member:     argv[3],
				}, nil
			},
			Encode: func(i Instruction, out []uint32) (uint32, error) {
				v := i.(*OpArraylength)
				out[0] = v.ResultType
				out[1] = v.ResultId
				out[2] = v.Structure
				out[3] = v.Member
				return 4, nil
			},
		},
	)
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

func init() {
	Bind(
		(&OpImagePointer{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 5 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpImagePointer{
					ResultType: argv[0],
					ResultId:   argv[1],
					Image:      argv[2],
					Coordinate: argv[3],
					Sample:     argv[4],
				}, nil
			},
			Encode: func(i Instruction, out []uint32) (uint32, error) {
				v := i.(*OpImagePointer)
				out[0] = v.ResultType
				out[1] = v.ResultId
				out[2] = v.Image
				out[3] = v.Coordinate
				out[4] = v.Sample
				return 5, nil
			},
		},
	)
}

// OpGenericPtrMemSemantics returns a valid Memory Semantics
// value for ptr.
type OpGenericPtrMemSemantics struct {
	ResultType uint32 // Result Type must be a 32-bits wide OpTypeInt value
	ResultId   uint32
	Ptr        uint32 // Ptr must point to Generic.
}

func (c *OpGenericPtrMemSemantics) Opcode() uint32 { return 233 }

func init() {
	Bind(
		(&OpGenericPtrMemSemantics{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 3 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpGenericPtrMemSemantics{
					ResultType: argv[0],
					ResultId:   argv[1],
					Ptr:        argv[2],
				}, nil
			},
			Encode: func(i Instruction, out []uint32) (uint32, error) {
				v := i.(*OpGenericPtrMemSemantics)
				out[0] = v.ResultType
				out[1] = v.ResultId
				out[2] = v.Ptr
				return 3, nil
			},
		},
	)
}
