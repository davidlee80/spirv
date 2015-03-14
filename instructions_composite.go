// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpVectorExtractDynamic reads a single, dynamically selected, component of
// a vector.
type OpVectorExtractDynamic struct {
	ResultType Id
	ResultId   Id
	Vector     Id
	Index      Id
}

func (c *OpVectorExtractDynamic) Opcode() uint32 { return 58 }
func (c *OpVectorExtractDynamic) Verify() error  { return nil }

// OpVectorInsertDynamic writes a single, variably selected, component
// into a vector.
type OpVectorInsertDynamic struct {
	ResultType Id
	ResultId   Id
	Vector     Id
	Component  Id
	Index      Id
}

func (c *OpVectorInsertDynamic) Opcode() uint32 { return 59 }
func (c *OpVectorInsertDynamic) Verify() error  { return nil }

// OpVectorShuffle selects arbitrary components from two vectors to make
// a new vector.
type OpVectorShuffle struct {
	ResultType Id
	ResultId   Id
	Vector1    Id
	Vector2    Id
	Components []uint32
}

func (c *OpVectorShuffle) Opcode() uint32 { return 60 }
func (c *OpVectorShuffle) Verify() error  { return nil }

// OpCompositeConstruct constructs a new composite object from a set of
// constituent objects that will fully form it
type OpCompositeConstruct struct {
	ResultType   Id
	ResultId     Id
	Constituents []Id
}

func (c *OpCompositeConstruct) Opcode() uint32 { return 61 }
func (c *OpCompositeConstruct) Verify() error  { return nil }

// OpCompositeExtract extracts a part of a composite object.
type OpCompositeExtract struct {
	ResultType Id
	ResultId   Id
	Composite  Id
	Indices    []uint32
}

func (c *OpCompositeExtract) Opcode() uint32 { return 62 }
func (c *OpCompositeExtract) Verify() error  { return nil }

// OpCompositeInsert inserts into a composite object.
type OpCompositeInsert struct {
	ResultType Id
	ResultId   Id
	Object     Id
	Composite  Id
	Indices    []uint32
}

func (c *OpCompositeInsert) Opcode() uint32 { return 63 }
func (c *OpCompositeInsert) Verify() error  { return nil }

// OpCopyObject makes a copy of Operand.
// There are no dereferences involved.
type OpCopyObject struct {
	ResultType Id
	ResultId   Id
	Operand    Id
}

func (c *OpCopyObject) Opcode() uint32 { return 64 }
func (c *OpCopyObject) Verify() error  { return nil }

// OpTranspose transposes a matrix.
type OpTranspose struct {
	ResultType Id
	ResultId   Id
	Matrix     Id
}

func (c *OpTranspose) Opcode() uint32 { return 112 }
func (c *OpTranspose) Verify() error  { return nil }

func init() {
	Bind(func() Instruction { return &OpVectorExtractDynamic{} })
	Bind(func() Instruction { return &OpVectorInsertDynamic{} })
	Bind(func() Instruction { return &OpVectorShuffle{} })
	Bind(func() Instruction { return &OpCompositeConstruct{} })
	Bind(func() Instruction { return &OpCompositeExtract{} })
	Bind(func() Instruction { return &OpCompositeInsert{} })
	Bind(func() Instruction { return &OpCopyObject{} })
	Bind(func() Instruction { return &OpTranspose{} })
}
