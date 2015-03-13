// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpAny result is true if any component of Vector is true,
// otherwise result is false.
type OpAny struct {
	ResultType uint32
	ResultId   uint32
	Vector     uint32
}

func (c *OpAny) Opcode() uint32 { return 98 }
func (c *OpAny) Verify() error  { return nil }

// OpAll result is true if all components of Vector are true,
// otherwise result is false
type OpAll struct {
	ResultType uint32
	ResultId   uint32
	Vector     uint32
}

func (c *OpAll) Opcode() uint32 { return 99 }
func (c *OpAll) Verify() error  { return nil }

// OpIsNan result is true if x is an IEEE NaN,
// otherwise result is false.
type OpIsNan struct {
	ResultType uint32
	ResultId   uint32
	X          uint32
}

func (c *OpIsNan) Opcode() uint32 { return 113 }
func (c *OpIsNan) Verify() error  { return nil }

// OpIsInf result is true if x is an IEEE Inf,
// otherwise result is false.
type OpIsInf struct {
	ResultType uint32
	ResultId   uint32
	X          uint32
}

func (c *OpIsInf) Opcode() uint32 { return 114 }
func (c *OpIsInf) Verify() error  { return nil }

// OpIsFinite result is true if x is an IEEE finite number,
// otherwise result is false..
type OpIsFinite struct {
	ResultType uint32
	ResultId   uint32
	X          uint32
}

func (c *OpIsFinite) Opcode() uint32 { return 115 }
func (c *OpIsFinite) Verify() error  { return nil }

// OpIsNormal result is true if x is an IEEE normal number,
// otherwise result is false.
type OpIsNormal struct {
	ResultType uint32
	ResultId   uint32
	X          uint32
}

func (c *OpIsNormal) Opcode() uint32 { return 116 }
func (c *OpIsNormal) Verify() error  { return nil }

// OpSignBitSet result is true if x has its sign bit set,
// otherwise result is false.
type OpSignBitSet struct {
	ResultType uint32
	ResultId   uint32
	X          uint32
}

func (c *OpSignBitSet) Opcode() uint32 { return 117 }
func (c *OpSignBitSet) Verify() error  { return nil }

// OpLessOrGreater result is true if x < y or x > y,
// where IEEE comparisons are used, otherwise result is false.
type OpLessOrGreater struct {
	ResultType uint32
	ResultId   uint32
	X          uint32
}

func (c *OpLessOrGreater) Opcode() uint32 { return 118 }
func (c *OpLessOrGreater) Verify() error  { return nil }

// OpOrdered result is true if both x == x and y == y are true,
// where IEEE comparison is used, otherwise result is false.
type OpOrdered struct {
	ResultType uint32
	ResultId   uint32
	X          uint32
	Y          uint32
}

func (c *OpOrdered) Opcode() uint32 { return 119 }
func (c *OpOrdered) Verify() error  { return nil }

// OpUnordered result is true if either x or y is an IEEE NaN,
// otherwise result is false.
type OpUnordered struct {
	ResultType uint32
	ResultId   uint32
	X          uint32
	Y          uint32
}

func (c *OpUnordered) Opcode() uint32 { return 120 }
func (c *OpUnordered) Verify() error  { return nil }

// OpLogicalOr result is true if either Operand 1 or Operand 2 is true.
// Result is false if both Operand 1 and Operand 2 are false.
type OpLogicalOr struct {
	ResultType uint32
	ResultId   uint32
	Operand1   uint32
	Operand2   uint32
}

func (c *OpLogicalOr) Opcode() uint32 { return 146 }
func (c *OpLogicalOr) Verify() error  { return nil }

// OpLogicalXor result is true if exactly one of Operand 1 or
// Operand 2 is true. Result is false if Operand 1 and Operand 2
// have the same value.
type OpLogicalXor struct {
	ResultType uint32
	ResultId   uint32
	Operand1   uint32
	Operand2   uint32
}

func (c *OpLogicalXor) Opcode() uint32 { return 147 }
func (c *OpLogicalXor) Verify() error  { return nil }

// OpLogicalAnd result is true if both Operand 1 and Operand 2 are true.
// Result is false if either Operand 1 or Operand 2 are false
type OpLogicalAnd struct {
	ResultType uint32
	ResultId   uint32
	Operand1   uint32
	Operand2   uint32
}

func (c *OpLogicalAnd) Opcode() uint32 { return 148 }
func (c *OpLogicalAnd) Verify() error  { return nil }

// OpSelect select between two objects.
// Results are computed per component
type OpSelect struct {
	ResultType uint32
	ResultId   uint32
	Condition  uint32
	Object1    uint32
	Object2    uint32
}

func (c *OpSelect) Opcode() uint32 { return 152 }
func (c *OpSelect) Verify() error  { return nil }

// OpIEqual performs Integer comparison for equality.
type OpIEqual struct {
	ResultType uint32
	ResultId   uint32
	Object1    uint32
	Object2    uint32
}

func (c *OpIEqual) Opcode() uint32 { return 153 }
func (c *OpIEqual) Verify() error  { return nil }

// OpFOrdEqual performs Floating-point comparison for being
// ordered and equal.
type OpFOrdEqual struct {
	ResultType uint32
	ResultId   uint32
	Object1    uint32
	Object2    uint32
}

func (c *OpFOrdEqual) Opcode() uint32 { return 154 }
func (c *OpFOrdEqual) Verify() error  { return nil }

// OpFUnordEqual performs Floating-point comparison for being
// unordered or equal.
type OpFUnordEqual struct {
	ResultType uint32
	ResultId   uint32
	Object1    uint32
	Object2    uint32
}

func (c *OpFUnordEqual) Opcode() uint32 { return 155 }
func (c *OpFUnordEqual) Verify() error  { return nil }

// OpINotEqual performs Integer comparison for inequality
type OpINotEqual struct {
	ResultType uint32
	ResultId   uint32
	Object1    uint32
	Object2    uint32
}

func (c *OpINotEqual) Opcode() uint32 { return 156 }
func (c *OpINotEqual) Verify() error  { return nil }

// OpFOrdNotEqual performs Floating-point comparison for being
// ordered and not equal.
type OpFOrdNotEqual struct {
	ResultType uint32
	ResultId   uint32
	Object1    uint32
	Object2    uint32
}

func (c *OpFOrdNotEqual) Opcode() uint32 { return 157 }
func (c *OpFOrdNotEqual) Verify() error  { return nil }

// OpFUnordNotEqual performs Floating-point comparison for
// being unordered or not equal.
type OpFUnordNotEqual struct {
	ResultType uint32
	ResultId   uint32
	Object1    uint32
	Object2    uint32
}

func (c *OpFUnordNotEqual) Opcode() uint32 { return 158 }
func (c *OpFUnordNotEqual) Verify() error  { return nil }

// OpULessThan performs Unsigned-integer comparison if Operand 1
// is less than Operand 2.
type OpULessThan struct {
	ResultType uint32
	ResultId   uint32
	Object1    uint32
	Object2    uint32
}

func (c *OpULessThan) Opcode() uint32 { return 159 }
func (c *OpULessThan) Verify() error  { return nil }

// OpSLessThan performs Signed-integer comparison if Operand 1
// is less than Operand 2.
type OpSLessThan struct {
	ResultType uint32
	ResultId   uint32
	Object1    uint32
	Object2    uint32
}

func (c *OpSLessThan) Opcode() uint32 { return 160 }
func (c *OpSLessThan) Verify() error  { return nil }

// OpFOrdLessThan performs Floating-point comparison if operands are
// ordered and Operand 1 is less than Operand 2.
type OpFOrdLessThan struct {
	ResultType uint32
	ResultId   uint32
	Object1    uint32
	Object2    uint32
}

func (c *OpFOrdLessThan) Opcode() uint32 { return 161 }
func (c *OpFOrdLessThan) Verify() error  { return nil }

// OpFUnordLessThan performs Floating-point comparison if operands
// are unordered or Operand 1 is less than Operand 2.
type OpFUnordLessThan struct {
	ResultType uint32
	ResultId   uint32
	Object1    uint32
	Object2    uint32
}

func (c *OpFUnordLessThan) Opcode() uint32 { return 162 }
func (c *OpFUnordLessThan) Verify() error  { return nil }

// OpUGreaterThan performs Unsigned-integer comparison if Operand 1
// is greater than Operand 2.
type OpUGreaterThan struct {
	ResultType uint32
	ResultId   uint32
	Object1    uint32
	Object2    uint32
}

func (c *OpUGreaterThan) Opcode() uint32 { return 163 }
func (c *OpUGreaterThan) Verify() error  { return nil }

// OpSGreaterThan performs Signed-integer comparison if Operand 1
// is greater than Operand 2.
type OpSGreaterThan struct {
	ResultType uint32
	ResultId   uint32
	Object1    uint32
	Object2    uint32
}

func (c *OpSGreaterThan) Opcode() uint32 { return 164 }
func (c *OpSGreaterThan) Verify() error  { return nil }

// OpFOrdGreaterThan performs Floating-point comparison if operands
// are ordered and Operand 1 is greater than Operand 2.
type OpFOrdGreaterThan struct {
	ResultType uint32
	ResultId   uint32
	Object1    uint32
	Object2    uint32
}

func (c *OpFOrdGreaterThan) Opcode() uint32 { return 165 }
func (c *OpFOrdGreaterThan) Verify() error  { return nil }

// OpFUnordGreaterThan performs Floating-point comparison if
// operands are unordered or Operand 1 is greater than Operand 2.
type OpFUnordGreaterThan struct {
	ResultType uint32
	ResultId   uint32
	Object1    uint32
	Object2    uint32
}

func (c *OpFUnordGreaterThan) Opcode() uint32 { return 166 }
func (c *OpFUnordGreaterThan) Verify() error  { return nil }

// OpULessThanEqual performs Unsigned-integer comparison if Operand 1 is
// less than or equal to Operand 2.
type OpULessThanEqual struct {
	ResultType uint32
	ResultId   uint32
	Object1    uint32
	Object2    uint32
}

func (c *OpULessThanEqual) Opcode() uint32 { return 167 }
func (c *OpULessThanEqual) Verify() error  { return nil }

// OpSLessThanEqual performs Signed-integer comparison if Operand 1 is
// less than or equal to Operand 2.
type OpSLessThanEqual struct {
	ResultType uint32
	ResultId   uint32
	Object1    uint32
	Object2    uint32
}

func (c *OpSLessThanEqual) Opcode() uint32 { return 168 }
func (c *OpSLessThanEqual) Verify() error  { return nil }

// OpFOrdLessThanEqual performs Floating-point comparison if operands
// are ordered and Operand 1 is less than or equal to Operand 2.
type OpFOrdLessThanEqual struct {
	ResultType uint32
	ResultId   uint32
	Object1    uint32
	Object2    uint32
}

func (c *OpFOrdLessThanEqual) Opcode() uint32 { return 169 }
func (c *OpFOrdLessThanEqual) Verify() error  { return nil }

// OpFUnordLessThanEqual performs Floating-point comparison if
// operands are unordered or Operand 1 is less than or equal to Operand 2.
type OpFUnordLessThanEqual struct {
	ResultType uint32
	ResultId   uint32
	Object1    uint32
	Object2    uint32
}

func (c *OpFUnordLessThanEqual) Opcode() uint32 { return 170 }
func (c *OpFUnordLessThanEqual) Verify() error  { return nil }

// OpUGreaterThanEqual performs Unsigned-integer comparison if Operand 1 is
// greater than or equal to Operand 2.
type OpUGreaterThanEqual struct {
	ResultType uint32
	ResultId   uint32
	Object1    uint32
	Object2    uint32
}

func (c *OpUGreaterThanEqual) Opcode() uint32 { return 171 }
func (c *OpUGreaterThanEqual) Verify() error  { return nil }

// OpSGreaterThanEqual performs Signed-integer comparison if Operand 1 is
// greater than or equal to Operand 2.
type OpSGreaterThanEqual struct {
	ResultType uint32
	ResultId   uint32
	Object1    uint32
	Object2    uint32
}

func (c *OpSGreaterThanEqual) Opcode() uint32 { return 172 }
func (c *OpSGreaterThanEqual) Verify() error  { return nil }

// OpFOrdGreaterThanEqual performs Floating-point comparison if
// operands are ordered and Operand 1 is greater than or equal to Operand 2.
type OpFOrdGreaterThanEqual struct {
	ResultType uint32
	ResultId   uint32
	Object1    uint32
	Object2    uint32
}

func (c *OpFOrdGreaterThanEqual) Opcode() uint32 { return 173 }
func (c *OpFOrdGreaterThanEqual) Verify() error  { return nil }

// OpFUnordGreaterThanEqual performs Floating-point comparison if
// operands are unordered or Operand 1 is greater than or equal
// to Operand 2.
type OpFUnordGreaterThanEqual struct {
	ResultType uint32
	ResultId   uint32
	Object1    uint32
	Object2    uint32
}

func (c *OpFUnordGreaterThanEqual) Opcode() uint32 { return 174 }
func (c *OpFUnordGreaterThanEqual) Verify() error  { return nil }

func init() {
	Bind(func() Instruction { return &OpAny{} })
	Bind(func() Instruction { return &OpAll{} })
	Bind(func() Instruction { return &OpIsNan{} })
	Bind(func() Instruction { return &OpIsInf{} })
	Bind(func() Instruction { return &OpIsFinite{} })
	Bind(func() Instruction { return &OpIsNormal{} })
	Bind(func() Instruction { return &OpSignBitSet{} })
	Bind(func() Instruction { return &OpLessOrGreater{} })
	Bind(func() Instruction { return &OpOrdered{} })
	Bind(func() Instruction { return &OpUnordered{} })
	Bind(func() Instruction { return &OpLogicalOr{} })
	Bind(func() Instruction { return &OpLogicalXor{} })
	Bind(func() Instruction { return &OpLogicalAnd{} })
	Bind(func() Instruction { return &OpSelect{} })
	Bind(func() Instruction { return &OpIEqual{} })
	Bind(func() Instruction { return &OpFOrdEqual{} })
	Bind(func() Instruction { return &OpFUnordEqual{} })
	Bind(func() Instruction { return &OpINotEqual{} })
	Bind(func() Instruction { return &OpFOrdNotEqual{} })
	Bind(func() Instruction { return &OpFUnordNotEqual{} })
	Bind(func() Instruction { return &OpULessThan{} })
	Bind(func() Instruction { return &OpSLessThan{} })
	Bind(func() Instruction { return &OpFOrdLessThan{} })
	Bind(func() Instruction { return &OpFUnordLessThan{} })
	Bind(func() Instruction { return &OpUGreaterThan{} })
	Bind(func() Instruction { return &OpSGreaterThan{} })
	Bind(func() Instruction { return &OpFOrdGreaterThan{} })
	Bind(func() Instruction { return &OpFUnordGreaterThan{} })
	Bind(func() Instruction { return &OpULessThanEqual{} })
	Bind(func() Instruction { return &OpSLessThanEqual{} })
	Bind(func() Instruction { return &OpFOrdLessThanEqual{} })
	Bind(func() Instruction { return &OpFUnordLessThanEqual{} })
	Bind(func() Instruction { return &OpUGreaterThanEqual{} })
	Bind(func() Instruction { return &OpSGreaterThanEqual{} })
	Bind(func() Instruction { return &OpFOrdGreaterThanEqual{} })
	Bind(func() Instruction { return &OpFUnordGreaterThanEqual{} })
}
