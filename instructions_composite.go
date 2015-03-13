// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpVectorExtractDynamic reads a single, dynamically selected, component of
// a vector.
type OpVectorExtractDynamic struct {
	ResultType uint32
	ResultId   uint32
	Vector     uint32
	Index      uint32
}

func (c *OpVectorExtractDynamic) Opcode() uint32 { return 58 }
func (c *OpVectorExtractDynamic) Verify() error  { return nil }

// OpVectorInsertDynamic writes a single, variably selected, component
// into a vector.
type OpVectorInsertDynamic struct {
	ResultType uint32
	ResultId   uint32
	Vector     uint32
	Component  uint32
	Index      uint32
}

func (c *OpVectorInsertDynamic) Opcode() uint32 { return 59 }
func (c *OpVectorInsertDynamic) Verify() error  { return nil }

// OpVectorShuffle selects arbitrary components from two vectors to make
// a new vector.
type OpVectorShuffle struct {
	ResultType uint32
	ResultId   uint32
	Vector1    uint32
	Vector2    uint32
	Components []uint32
}

func (c *OpVectorShuffle) Opcode() uint32 { return 60 }
func (c *OpVectorShuffle) Verify() error  { return nil }

// OpCompositeConstruct constructs a new composite object from a set of
// constituent objects that will fully form it
type OpCompositeConstruct struct {
	ResultType   uint32
	ResultId     uint32
	Constituents []uint32
}

func (c *OpCompositeConstruct) Opcode() uint32 { return 61 }
func (c *OpCompositeConstruct) Verify() error  { return nil }

// OpCompositeExtract extracts a part of a composite object.
type OpCompositeExtract struct {
	ResultType uint32
	ResultId   uint32
	Composite  uint32
	Indices    []uint32
}

func (c *OpCompositeExtract) Opcode() uint32 { return 62 }
func (c *OpCompositeExtract) Verify() error  { return nil }

// OpCompositeInsert inserts into a composite object.
type OpCompositeInsert struct {
	ResultType uint32
	ResultId   uint32
	Object     uint32
	Composite  uint32
	Indices    []uint32
}

func (c *OpCompositeInsert) Opcode() uint32 { return 63 }
func (c *OpCompositeInsert) Verify() error  { return nil }

// OpCopyObject makes a copy of Operand.
// There are no dereferences involved.
type OpCopyObject struct {
	ResultType uint32
	ResultId   uint32
	Operand    uint32
}

func (c *OpCopyObject) Opcode() uint32 { return 64 }
func (c *OpCopyObject) Verify() error  { return nil }

// OpTranspose transposes a matrix.
type OpTranspose struct {
	ResultType uint32
	ResultId   uint32
	Matrix     uint32
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
