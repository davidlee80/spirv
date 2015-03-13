// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "fmt"

// OpConstantTrue declares a true Boolean-type scalar constant.
type OpConstantTrue struct {
	// Result Type must be the scalar Boolean type.
	ResultType uint32

	// The <id> of the new constant type.
	ResultId uint32
}

func (c *OpConstantTrue) Opcode() uint32 { return 27 }
func (c *OpConstantTrue) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpConstantTrue{}
	})
}

// OpConstantFalse declares a true Boolean-type scalar constant.
type OpConstantFalse struct {
	// Result Type must be the scalar Boolean type.
	ResultType uint32

	// The <id> of the new constant type.
	ResultId uint32
}

func (c *OpConstantFalse) Opcode() uint32 { return 28 }
func (c *OpConstantFalse) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpConstantFalse{}
	})
}

// OpConstant declares a new Integer-type or Floating-point-type
// scalar constant.
type OpConstant struct {
	// Result Type must be a scalar Integer type or Floating-point type.
	ResultType uint32

	// The <id> of the new constant type.
	ResultId uint32

	// Value is the bit pattern for the constant.
	//
	// Types 32 bits wide or smaller take one word. Larger types take
	// multiple words, with low-order words appearing first.
	Value []uint32
}

func (c *OpConstant) Opcode() uint32 { return 29 }
func (c *OpConstant) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpConstant{}
	})
}

// OpConstantComposite declares a new composite constant.
type OpConstantComposite struct {
	// Result Type must be a composite type, whose top-level
	// members/elements/components/columns have the same type as the
	// types of the operands
	ResultType uint32

	// The <id> of the new composite type.
	ResultId uint32

	// Constituents will become members of a structure, or elements of an
	// array, or components of a vector, or columns of a matrix. There must
	// be exactly one Constituent for each top-level member/element/component/column
	// of the result.
	//
	// The Constituents must appear in the order needed by the definition of
	// the type of the result. The Constituents must be the <id> of other
	// constant declarations.
	Constituents []uint32
}

func (c *OpConstantComposite) Opcode() uint32 { return 30 }
func (c *OpConstantComposite) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpConstantComposite{}
	})
}

// OpConstantSampler declares a new null sampler constant.
type OpConstantSampler struct {
	ResultType uint32
	ResultId   uint32

	// Addressing is the addressing Mode.
	Addressing uint32

	// Param is one of:
	//
	//   0: Nonparametric
	//   1: Parametric
	//
	Param uint32

	// Filter is the filter mode.
	Filter uint32
}

func (c *OpConstantSampler) Opcode() uint32 { return 31 }
func (c *OpConstantSampler) Verify() error {
	switch c.Param {
	case 0, 1:
	default:
		return fmt.Errorf("OpConstantSampler.Param: expected: 0, 1")
	}

	return nil
}

func init() {
	Bind(func() Instruction {
		return &OpConstantSampler{}
	})
}

// OpConstantNullPointer declares a new null pointer constant.
type OpConstantNullPointer struct {
	// Result Type must be the scalar Boolean type.
	ResultType uint32

	// The <id> of the new constant type.
	ResultId uint32
}

func (c *OpConstantNullPointer) Opcode() uint32 { return 32 }
func (c *OpConstantNullPointer) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpConstantNullPointer{}
	})
}

// OpConstantNullObject declares a new null object constant.
// The objerct can be a queue, event or reservation id.
type OpConstantNullObject struct {
	// Result Type must be the scalar Boolean type.
	ResultType uint32

	// The <id> of the new constant type.
	ResultId uint32
}

func (c *OpConstantNullObject) Opcode() uint32 { return 33 }
func (c *OpConstantNullObject) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpConstantNullObject{}
	})
}

// OpSpecConstantTrue declares a Boolean-type scalar specialization
// constant with a default value of true.
//
// This instruction can be specialized to become either an OpConstantTrue
// or OpConstantFalse instruction.
type OpSpecConstantTrue struct {
	// Result Type must be the scalar Boolean type.
	ResultType uint32

	// The <id> of the new constant type.
	ResultId uint32
}

func (c *OpSpecConstantTrue) Opcode() uint32 { return 34 }
func (c *OpSpecConstantTrue) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpSpecConstantTrue{}
	})
}

// OpSpecConstantFalse declares a Boolean-type scalar specialization
// constant with a default value of false.
//
// This instruction can be specialized to become either an OpConstantTrue
// or OpConstantFalse instruction.
type OpSpecConstantFalse struct {
	// Result Type must be the scalar Boolean type.
	ResultType uint32

	// The <id> of the new constant type.
	ResultId uint32
}

func (c *OpSpecConstantFalse) Opcode() uint32 { return 35 }
func (c *OpSpecConstantFalse) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpSpecConstantFalse{}
	})
}

// OpSpecConstant declares a new Integer-type or Floating-point-type
// scalar specialization constant.
//
// This instruction can be specialized to become either an OpConstantTrue
// or OpConstantFalse instruction.
type OpSpecConstant struct {
	// Result Type must be a scalar Integer type or Floating-point type.
	ResultType uint32

	// The <id> of the new constant type.
	ResultId uint32

	// Value is the bit pattern for the default value of the constant.
	// Types 32 bits wide or smaller take one word. Larger types take multiple
	// words, with low-order words appearing first.
	Value []uint32
}

func (c *OpSpecConstant) Opcode() uint32 { return 36 }
func (c *OpSpecConstant) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpSpecConstant{}
	})
}

// OpSpecConstantComposite declares a new composite specialization constant.
//
// This instruction will be specialized to an OpConstantComposite instruction.
type OpSpecConstantComposite struct {
	// Result Type must be a composite type, whose top-level
	// members/elements/components/columns have the same type as the types
	// of the operands.
	ResultType uint32

	// The <id> of the new constant type.
	ResultId uint32

	// Constituents will become members of a structure, or elements of an
	// array, or components of a vector, or columns of a matrix. There must be
	// exactly one Constituent for each top-level member/element/component/column
	// of the result.
	//
	// The Constituents must appear in the order needed by the definition of
	// the type of the result. The Constituents must be the <id> of other
	// specialization constant or constant declarations
	Constituents []uint32
}

func (c *OpSpecConstantComposite) Opcode() uint32 { return 37 }
func (c *OpSpecConstantComposite) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpSpecConstantComposite{}
	})
}
