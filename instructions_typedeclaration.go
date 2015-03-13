// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "errors"

// OpTypeVoid represents the OpTypeVoid instruction.
type OpTypeVoid struct {
	ResultId uint32
}

func (c *OpTypeVoid) Opcode() uint32 { return 8 }
func (c *OpTypeVoid) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTypeVoid{}
	})
}

// OpTypeBool represents the OpTypeBool instruction.
type OpTypeBool struct {
	ResultId uint32
}

func (c *OpTypeBool) Opcode() uint32 { return 9 }
func (c *OpTypeBool) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTypeBool{}
	})
}

// OpTypeInt represents the OpTypeInt instruction.
type OpTypeInt struct {
	// The <id> of the new integer type.
	ResultId uint32

	// Specifies how many bits wide the type is.
	// The bit pattern of a signed integer value is two’s complement.
	Width uint32

	// Signedness specifies whether there are signed semantics to
	// preserve or validate.
	//
	//   0: indicates unsigned, or no signedness semantics.
	//   1: indicates signed semantics.
	//
	// In all cases, the type of operation of an instruction comes from
	// the instruction’s opcode, not the signedness of the operands.
	Signedness uint32
}

func (c *OpTypeInt) Opcode() uint32 { return 10 }
func (c *OpTypeInt) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTypeInt{}
	})
}

// OpTypeFloat represents the OpTypeFloat instruction.
// It declares a new floating point type.
type OpTypeFloat struct {
	// The <id> of the new floating-point type.
	ResultId uint32

	// Specifies how many bits wide the type is. The bit pattern of a
	// floating-point value is as described by the IEEE 754 standard.
	Width uint32
}

func (c *OpTypeFloat) Opcode() uint32 { return 11 }
func (c *OpTypeFloat) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTypeFloat{}
	})
}

// OpTypeVector represents the OpTypeVector instruction.
// It declares a new vector type.
type OpTypeVector struct {
	// The <id> of the new vector type.
	ResultId uint32

	// Specifies the type of each component in the resulting type.
	ComponentType uint32

	// Specifies the number of compononents in the resulting type.
	// It must be at least 2.
	ComponentCount uint32
}

func (c *OpTypeVector) Opcode() uint32 { return 12 }
func (c *OpTypeVector) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTypeVector{}
	})
}

// OpTypeMatrix declares a new matrix type.
type OpTypeMatrix struct {
	// The <id> of the new matrix type
	ResultId uint32

	// The type of each column in the matrix. It must be vector type.
	ColumnType uint32

	// The number of columns in the new matrix type. It must be at least 2.
	ColumnCount uint32
}

func (c *OpTypeMatrix) Opcode() uint32 { return 13 }
func (c *OpTypeMatrix) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTypeMatrix{}
	})
}

// OpTypeSampler declares a new sampler type.
//
// It is consumed, for example, by OpTextureSample.
// This type is opaque: values of this type have no defined physical
// size or bit pattern..
type OpTypeSampler struct {
	// The <id> of the new sampler type.
	ResultId uint32

	// A scalar type, of the type of the components resulting from
	// sampling or loading through this sampler
	SampledType uint32

	// The texture dimensionality.
	Dimensionality uint32

	// Content must be one of the following indicated values:
	//
	//   0: indicates a texture, no filter (no sampling state)
	//   1: indicates an image
	//   2: indicates both a texture and filter (sampling state),
	//      see OpTypeFilter
	//
	Content uint32

	// Arrayed must be one of the following indicated values:
	//
	//   0: indicates non-arrayed content
	//   1: indicates arrayed content
	//
	Arrayed uint32

	// Compare must be one of the following indicated values:
	//
	//   0: indicates depth comparisons are not done
	//   1: indicates depth comparison are done
	//
	Compare uint32

	// MS is multisampled and must be one of the following indicated values:
	//
	//   0: indicates single-sampled content
	//   1: indicates multisampled content
	//
	MS uint32

	// AccessQualifier is an image access qualifier and is optional.
	AccessQualifier uint32 `spirv:"optional"`
}

func (c *OpTypeSampler) Opcode() uint32 { return 14 }

func (c *OpTypeSampler) Verify() error {
	switch c.Content {
	case 0, 1, 2:
	default:
		return errors.New("OpTypeSampler.Compare: expected: 0, 1, 2")
	}

	switch c.Arrayed {
	case 0, 1:
	default:
		return errors.New("OpTypeSampler.Arrayed: expected: 0, 1")
	}

	switch c.Compare {
	case 0, 1:
	default:
		return errors.New("OpTypeSampler.Compare: expected: 0, 1")
	}

	switch c.MS {
	case 0, 1:
	default:
		return errors.New("OpTypeSampler.MS: expected: 0, 1")
	}

	return nil
}

func init() {
	Bind(func() Instruction {
		return &OpTypeSampler{}
	})
}

// OpTypeFilter declares a filter type. It is consumed by OpSampler.
// This type is opaque: values of this type have no defined
// physical size or bit pattern.
type OpTypeFilter struct {
	ResultId uint32
}

func (c *OpTypeFilter) Opcode() uint32 { return 15 }
func (c *OpTypeFilter) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTypeFilter{}
	})
}

// OpTypeArray declares a new array type: a dynamically-indexable ordered
// aggregate of elements all having the same type.
type OpTypeArray struct {
	// The <id> of the new array type.
	ResultId uint32

	// The type of each element in the array
	ElementType uint32

	// The number of elements in the array. It must be at least 1.
	//
	// Length must come from a constant instruction of an Integer-type
	// scalar whose value is at least 1.
	Length uint32
}

func (c *OpTypeArray) Opcode() uint32 { return 16 }
func (c *OpTypeArray) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTypeArray{}
	})
}

// OpTypeRuntimeArray declares a new run-time array type.
// Its length is not known at compile time.
//
// Objects of this type can only be created with OpVariable
// using the Uniform Storage Class.
type OpTypeRuntimeArray struct {
	// The <id> of the new run-time array type.
	ResultId uint32

	// The type of each element in the array.
	// See OpArrayLength for getting the Length of an array of this type.
	ElementType uint32
}

func (c *OpTypeRuntimeArray) Opcode() uint32 { return 17 }
func (c *OpTypeRuntimeArray) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTypeRuntimeArray{}
	})
}

// OpTypeStruct declares a new structure type: an aggregate
// of heteregeneous members
type OpTypeStruct struct {
	// The <id> of the new array type.
	ResultId uint32

	// Member N type is the type of member N of the structure. The first
	// member is member 0, the next is member 1, . . .
	Members []uint32
}

func (c *OpTypeStruct) Opcode() uint32 { return 18 }
func (c *OpTypeStruct) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTypeStruct{}
	})
}

// OpTypeOpaque declares a named structure type with no body specified.
type OpTypeOpaque struct {
	// The <id> of the new array type.
	ResultId uint32

	// The name of the opaque type.
	Name String
}

func (c *OpTypeOpaque) Opcode() uint32 { return 19 }
func (c *OpTypeOpaque) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTypeOpaque{}
	})
}

// OpTypePointer declares a new pointer type.
type OpTypePointer struct {
	// The <id> of the new integer type.
	ResultId uint32

	// The storage class of the memory holding the object pointed to.
	StorageClass uint32

	// The type of the object pointed to.
	Type uint32
}

func (c *OpTypePointer) Opcode() uint32 { return 20 }
func (c *OpTypePointer) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTypePointer{}
	})
}

// OpTypeFunction declares a new function type.
//
// OpFunction will use this to declare the return type and
// parameter types of a function
type OpTypeFunction struct {
	// The <id> of the new function type.
	ResultId uint32

	// The type of the return value of functions of this type.
	// If the function has no return value, Return Type should
	// be from OpTypeVoid.
	ReturnType uint32

	// Parameter N Type is the type <id> of the type of parameter N
	Parameters []uint32
}

func (c *OpTypeFunction) Opcode() uint32 { return 21 }
func (c *OpTypeFunction) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTypeFunction{}
	})
}

// OpTypeEvent declares an OpenCL event object.
type OpTypeEvent struct {
	ResultId uint32
}

func (c *OpTypeEvent) Opcode() uint32 { return 22 }
func (c *OpTypeEvent) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTypeEvent{}
	})
}

// OpTypeDeviceEvent declares an OpenCL device-side event object.
//
// It defines the <id> of the new device-side-event type.
type OpTypeDeviceEvent struct {
	ResultId uint32
}

func (c *OpTypeDeviceEvent) Opcode() uint32 { return 23 }
func (c *OpTypeDeviceEvent) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTypeDeviceEvent{}
	})
}

// OpTypeReserveId declares an OpenCL reservation id object.
//
// It defines the <id> of the new reservation type.
type OpTypeReserveId struct {
	ResultId uint32
}

func (c *OpTypeReserveId) Opcode() uint32 { return 24 }
func (c *OpTypeReserveId) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTypeReserveId{}
	})
}

// OpTypeQueue declares an OpenCL queue object.
//
// It defines the <id> of the new queue type.
type OpTypeQueue struct {
	ResultId uint32
}

func (c *OpTypeQueue) Opcode() uint32 { return 25 }
func (c *OpTypeQueue) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTypeQueue{}
	})
}

// OpTypePipe declares an OpenCL pipe object type.
type OpTypePipe struct {
	// The <id> of the new pipe type.
	ResultId uint32

	// Type is the data type of the pipe.
	Type uint32

	// Qualifier is the pipe access qualifier.
	AccessQualifier uint32
}

func (c *OpTypePipe) Opcode() uint32 { return 26 }
func (c *OpTypePipe) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTypePipe{}
	})
}
