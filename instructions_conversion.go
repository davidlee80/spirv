// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpConvertFToU converts (value preserving) Float Value from floating
// point to unsigned integer, with rounding toward 0
type OpConvertFToU struct {
	ResultType Id
	ResultId   Id
	Value      Id
}

func (c *OpConvertFToU) Opcode() uint32 { return 100 }
func (c *OpConvertFToU) Optional() bool { return false }
func (c *OpConvertFToU) Verify() error  { return nil }

// OpConvertFToS converts (value preserving) Float Value from floating
// point to signed integer, with round toward 0
type OpConvertFToS struct {
	ResultType Id
	ResultId   Id
	Value      Id
}

func (c *OpConvertFToS) Opcode() uint32 { return 101 }
func (c *OpConvertFToS) Optional() bool { return false }
func (c *OpConvertFToS) Verify() error  { return nil }

// OpConvertSToF converts (value preserving) Signed Value from signed integer
// to floating point.
type OpConvertSToF struct {
	ResultType Id
	ResultId   Id
	Value      Id
}

func (c *OpConvertSToF) Opcode() uint32 { return 102 }
func (c *OpConvertSToF) Optional() bool { return false }
func (c *OpConvertSToF) Verify() error  { return nil }

// OpConvertUToF converts (value preserving) Unsigned value from unsigned
// integer to floating point
type OpConvertUToF struct {
	ResultType Id
	ResultId   Id
	Value      Id
}

func (c *OpConvertUToF) Opcode() uint32 { return 103 }
func (c *OpConvertUToF) Optional() bool { return false }
func (c *OpConvertUToF) Verify() error  { return nil }

// OpUConvert converts (value preserving) the width of Unsigned value.
// This is either a truncate or a zero extend.
type OpUConvert struct {
	ResultType Id
	ResultId   Id
	Value      Id
}

func (c *OpUConvert) Opcode() uint32 { return 104 }
func (c *OpUConvert) Optional() bool { return false }
func (c *OpUConvert) Verify() error  { return nil }

// OpSConvert converts (value preserving) the width of Signed Value.
// This is either a truncate or a sign extend.
type OpSConvert struct {
	ResultType Id
	ResultId   Id
	Value      Id
}

func (c *OpSConvert) Opcode() uint32 { return 105 }
func (c *OpSConvert) Optional() bool { return false }
func (c *OpSConvert) Verify() error  { return nil }

// OpFConvert converts (value preserving) the width of Float Value.
//
// Results are computed per component. The operand’s type and Result Type must
// have the same number of components. The widths of the components of the
// operand and the Result Type must be different.
type OpFConvert struct {
	ResultType Id
	ResultId   Id
	Value      Id
}

func (c *OpFConvert) Opcode() uint32 { return 106 }
func (c *OpFConvert) Optional() bool { return false }
func (c *OpFConvert) Verify() error  { return nil }

// OpConvertPtrToU converts Pointer to an unsigned integer type. A Result Type
// width larger than the width of Pointer will zero extend.
type OpConvertPtrToU struct {
	ResultType Id
	ResultId   Id
	Value      Id
}

func (c *OpConvertPtrToU) Opcode() uint32 { return 107 }
func (c *OpConvertPtrToU) Optional() bool { return false }
func (c *OpConvertPtrToU) Verify() error  { return nil }

// OpConvertUToPtr converts Integer value to a pointer. A Result Type width
// smaller than the width of Integer value pointer will truncate.
type OpConvertUToPtr struct {
	ResultType Id
	ResultId   Id
	Value      Id
}

func (c *OpConvertUToPtr) Opcode() uint32 { return 108 }
func (c *OpConvertUToPtr) Optional() bool { return false }
func (c *OpConvertUToPtr) Verify() error  { return nil }

// OpPtrCastToGeneric converts Source pointer to a pointer value pointing to
// storage class Generic. Source pointer must point to storage class
// WorkgroupLocal, WorkgroupGlobal or Private.
type OpPtrCastToGeneric struct {
	ResultType Id
	ResultId   Id
	Source     Id
}

func (c *OpPtrCastToGeneric) Opcode() uint32 { return 109 }
func (c *OpPtrCastToGeneric) Optional() bool { return false }
func (c *OpPtrCastToGeneric) Verify() error  { return nil }

// OpGenericCastToPtr converts Source pointer to a non-Generic storage-class
// pointer value. Source pointer must point to Generic.
type OpGenericCastToPtr struct {
	ResultType Id
	ResultId   Id
	Source     Id
}

func (c *OpGenericCastToPtr) Opcode() uint32 { return 110 }
func (c *OpGenericCastToPtr) Optional() bool { return false }
func (c *OpGenericCastToPtr) Verify() error  { return nil }

// OpBitcast defines a Bit-pattern preserving type conversion for
// Numerical-type or pointer-type vectors and scalars.
type OpBitcast struct {
	ResultType Id
	ResultId   Id
	Operand    Id // Operand is the bit pattern whose type will change
}

func (c *OpBitcast) Opcode() uint32 { return 111 }
func (c *OpBitcast) Optional() bool { return false }
func (c *OpBitcast) Verify() error  { return nil }

// OpGenericCastToPtrExplicit attempts to explicitly convert Source pointer
// to storage storage-class pointer value.
type OpGenericCastToPtrExplicit struct {
	ResultType Id
	ResultId   Id
	SourcePtr  Id
	Storage    StorageClass
}

func (c *OpGenericCastToPtrExplicit) Opcode() uint32 { return 232 }
func (c *OpGenericCastToPtrExplicit) Optional() bool { return false }
func (c *OpGenericCastToPtrExplicit) Verify() error  { return nil }

func init() {
	Bind(func() Instruction { return &OpConvertFToU{} })
	Bind(func() Instruction { return &OpConvertFToS{} })
	Bind(func() Instruction { return &OpConvertSToF{} })
	Bind(func() Instruction { return &OpConvertUToF{} })
	Bind(func() Instruction { return &OpUConvert{} })
	Bind(func() Instruction { return &OpSConvert{} })
	Bind(func() Instruction { return &OpFConvert{} })
	Bind(func() Instruction { return &OpConvertPtrToU{} })
	Bind(func() Instruction { return &OpConvertUToPtr{} })
	Bind(func() Instruction { return &OpPtrCastToGeneric{} })
	Bind(func() Instruction { return &OpGenericCastToPtr{} })
	Bind(func() Instruction { return &OpBitcast{} })
	Bind(func() Instruction { return &OpGenericCastToPtrExplicit{} })
}
