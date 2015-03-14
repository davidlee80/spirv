// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "fmt"

// OpConstantTrue declares a true Boolean-type scalar constant.
type OpConstantTrue struct {
	ResultType Id
	ResultId   Id
}

func (c *OpConstantTrue) Opcode() uint32 { return 27 }
func (c *OpConstantTrue) Optional() bool { return false }
func (c *OpConstantTrue) Verify() error  { return nil }

// OpConstantFalse declares a true Boolean-type scalar constant.
type OpConstantFalse struct {
	// Result Type must be the scalar Boolean type.
	ResultType Id

	// The <id> of the new constant type.
	ResultId Id
}

func (c *OpConstantFalse) Opcode() uint32 { return 28 }
func (c *OpConstantFalse) Optional() bool { return false }
func (c *OpConstantFalse) Verify() error  { return nil }

// OpConstant declares a new Integer-type or Floating-point-type
// scalar constant.
type OpConstant struct {
	ResultType Id
	ResultId   Id

	// Value is the bit pattern for the constant.
	//
	// Types 32 bits wide or smaller take one word. Larger types take
	// multiple words, with low-order words appearing first.
	Value []uint32
}

func (c *OpConstant) Opcode() uint32 { return 29 }
func (c *OpConstant) Optional() bool { return false }
func (c *OpConstant) Verify() error  { return nil }

// OpConstantComposite declares a new composite constant.
type OpConstantComposite struct {
	// Result Type must be a composite type, whose top-level
	// members/elements/components/columns have the same type as the
	// types of the operands
	ResultType Id
	ResultId   Id

	// Constituents will become members of a structure, or elements of an
	// array, or components of a vector, or columns of a matrix. There must
	// be exactly one Constituent for each top-level member/element/component/column
	// of the result.
	//
	// The Constituents must appear in the order needed by the definition of
	// the type of the result. The Constituents must be the <id> of other
	// constant declarations.
	Constituents []Id
}

func (c *OpConstantComposite) Opcode() uint32 { return 30 }
func (c *OpConstantComposite) Optional() bool { return false }
func (c *OpConstantComposite) Verify() error  { return nil }

// FIXME: Specification uses a Literal Number for Mode and Filter, not
// SamplerAddressingMode and SamplerFilterMode, respectively. Probably
// in error.

// OpConstantSampler declares a new null sampler constant.
type OpConstantSampler struct {
	ResultType Id
	ResultId   Id

	// Mode is the addressing Mode.
	Mode SamplerAddressingMode

	// Param is one of:
	//
	//   0: Nonparametric
	//   1: Parametric
	//
	Param uint32

	// Filter is the filter mode.
	Filter SamplerFilterMode
}

func (c *OpConstantSampler) Opcode() uint32 { return 31 }
func (c *OpConstantSampler) Optional() bool { return false }
func (c *OpConstantSampler) Verify() error {
	switch c.Param {
	case 0, 1:
	default:
		return fmt.Errorf("OpConstantSampler: Param must be 0 or 1")
	}

	return nil
}

// OpConstantNullPointer declares a new null pointer constant.
type OpConstantNullPointer struct {
	ResultType Id
	ResultId   Id
}

func (c *OpConstantNullPointer) Opcode() uint32 { return 32 }
func (c *OpConstantNullPointer) Optional() bool { return false }
func (c *OpConstantNullPointer) Verify() error  { return nil }

// OpConstantNullObject declares a new null object constant.
// The objerct can be a queue, event or reservation id.
type OpConstantNullObject struct {
	ResultType Id
	ResultId   Id
}

func (c *OpConstantNullObject) Opcode() uint32 { return 33 }
func (c *OpConstantNullObject) Optional() bool { return false }
func (c *OpConstantNullObject) Verify() error  { return nil }

// OpSpecConstantTrue declares a Boolean-type scalar specialization
// constant with a default value of true.
//
// This instruction can be specialized to become either an OpConstantTrue
// or OpConstantFalse instruction.
type OpSpecConstantTrue struct {
	ResultType Id
	ResultId   Id
}

func (c *OpSpecConstantTrue) Opcode() uint32 { return 34 }
func (c *OpSpecConstantTrue) Optional() bool { return false }
func (c *OpSpecConstantTrue) Verify() error  { return nil }

// OpSpecConstantFalse declares a Boolean-type scalar specialization
// constant with a default value of false.
//
// This instruction can be specialized to become either an OpConstantTrue
// or OpConstantFalse instruction.
type OpSpecConstantFalse struct {
	ResultType Id
	ResultId   Id
}

func (c *OpSpecConstantFalse) Opcode() uint32 { return 35 }
func (c *OpSpecConstantFalse) Optional() bool { return false }
func (c *OpSpecConstantFalse) Verify() error  { return nil }

// OpSpecConstant declares a new Integer-type or Floating-point-type
// scalar specialization constant.
//
// This instruction can be specialized to become either an OpConstantTrue
// or OpConstantFalse instruction.
type OpSpecConstant struct {
	ResultType Id
	ResultId   Id

	// Value is the bit pattern for the default value of the constant.
	// Types 32 bits wide or smaller take one word. Larger types take multiple
	// words, with low-order words appearing first.
	Value []uint32
}

func (c *OpSpecConstant) Opcode() uint32 { return 36 }
func (c *OpSpecConstant) Optional() bool { return false }
func (c *OpSpecConstant) Verify() error  { return nil }

// OpSpecConstantComposite declares a new composite specialization constant.
//
// This instruction will be specialized to an OpConstantComposite instruction.
type OpSpecConstantComposite struct {
	// Result Type must be a composite type, whose top-level
	// members/elements/components/columns have the same type as the types
	// of the operands.
	ResultType Id
	ResultId   Id

	// Constituents will become members of a structure, or elements of an
	// array, or components of a vector, or columns of a matrix. There must be
	// exactly one Constituent for each top-level member/element/component/column
	// of the result.
	//
	// The Constituents must appear in the order needed by the definition of
	// the type of the result. The Constituents must be the <id> of other
	// specialization constant or constant declarations
	Constituents []Id
}

func (c *OpSpecConstantComposite) Opcode() uint32 { return 37 }
func (c *OpSpecConstantComposite) Optional() bool { return false }
func (c *OpSpecConstantComposite) Verify() error  { return nil }

func init() {
	Bind(func() Instruction { return &OpSpecConstantTrue{} })
	Bind(func() Instruction { return &OpConstantTrue{} })
	Bind(func() Instruction { return &OpConstantFalse{} })
	Bind(func() Instruction { return &OpConstant{} })
	Bind(func() Instruction { return &OpConstantComposite{} })
	Bind(func() Instruction { return &OpConstantSampler{} })
	Bind(func() Instruction { return &OpConstantNullPointer{} })
	Bind(func() Instruction { return &OpConstantNullObject{} })
	Bind(func() Instruction { return &OpSpecConstantFalse{} })
	Bind(func() Instruction { return &OpSpecConstant{} })
	Bind(func() Instruction { return &OpSpecConstantComposite{} })
}
