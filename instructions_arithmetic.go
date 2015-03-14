// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpSNegate performs signed-integer subtract of Operand from zero.
type OpSNegate struct {
	ResultType Id
	ResultId   Id
	Operand    Id
}

func (c *OpSNegate) Opcode() uint32 { return 95 }
func (c *OpSNegate) Verify() error  { return nil }

// OpFNegate performs Floating-point subtract of Operand from zero.
type OpFNegate struct {
	ResultType Id
	ResultId   Id
	Operand    Id
}

func (c *OpFNegate) Opcode() uint32 { return 96 }
func (c *OpFNegate) Verify() error  { return nil }

// OpNot complements the bits of Operand.
type OpNot struct {
	ResultType Id
	ResultId   Id
	Operand    Id
}

func (c *OpNot) Opcode() uint32 { return 97 }
func (c *OpNot) Verify() error  { return nil }

// OpIAdd performs Integer addition of Operand 1 and Operand 2.
type OpIAdd struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpIAdd) Opcode() uint32 { return 122 }
func (c *OpIAdd) Verify() error  { return nil }

// OpFAdd performs Floating-point addition of Operand 1 and Operand 2.
type OpFAdd struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpFAdd) Opcode() uint32 { return 123 }
func (c *OpFAdd) Verify() error  { return nil }

// OpISub performs Integer subtraction of Operand 2 from Operand 1.
type OpISub struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpISub) Opcode() uint32 { return 124 }
func (c *OpISub) Verify() error  { return nil }

// OpFSub performs Floating-point subtraction of Operand 2 from Operand 1.
type OpFSub struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpFSub) Opcode() uint32 { return 125 }
func (c *OpFSub) Verify() error  { return nil }

// OpIMul performs Integer multiplication of Operand 1 and Operand 2.
type OpIMul struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpIMul) Opcode() uint32 { return 126 }
func (c *OpIMul) Verify() error  { return nil }

// OpFMul performs Floating-point multiplication of Operand 1 and Operand 2.
type OpFMul struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpFMul) Opcode() uint32 { return 127 }
func (c *OpFMul) Verify() error  { return nil }

// OpUDiv performs Unsigned-integer division of Operand 1 divided by Operand 2.
type OpUDiv struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpUDiv) Opcode() uint32 { return 128 }
func (c *OpUDiv) Verify() error  { return nil }

// OpSDiv performs Signed-integer division of Operand 1 divided by Operand 2.
type OpSDiv struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpSDiv) Opcode() uint32 { return 129 }
func (c *OpSDiv) Verify() error  { return nil }

// OpFDiv performs Floating-point division of Operand 1 divided by Operand 2.
type OpFDiv struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpFDiv) Opcode() uint32 { return 130 }
func (c *OpFDiv) Verify() error  { return nil }

// OpUMod performs Unsigned modulo operation of Operand 1 modulo Operand 2.
type OpUMod struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpUMod) Opcode() uint32 { return 131 }
func (c *OpUMod) Verify() error  { return nil }

// OpSRem performs Signed remainder operation of Operand 1 divided by Operand 2.
type OpSRem struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpSRem) Opcode() uint32 { return 132 }
func (c *OpSRem) Verify() error  { return nil }

// OpSMod performs Signed modulo operation of Operand 1 modulo Operand 2.
type OpSMod struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpSMod) Opcode() uint32 { return 133 }
func (c *OpSMod) Verify() error  { return nil }

// OpFRem performs Floating-point remainder operation of Operand 1 divided by Operand 2.
type OpFRem struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpFRem) Opcode() uint32 { return 134 }
func (c *OpFRem) Verify() error  { return nil }

// OpFMod performs Floating-point modulo operation of Operand 1 modulo Operand 2.
type OpFMod struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpFMod) Opcode() uint32 { return 135 }
func (c *OpFMod) Verify() error  { return nil }

// OpVectorTimesScalar scales a floating-point vector.
type OpVectorTimesScalar struct {
	ResultType Id
	ResultId   Id
	Vector     Id
	Scalar     Id
}

func (c *OpVectorTimesScalar) Opcode() uint32 { return 136 }
func (c *OpVectorTimesScalar) Verify() error  { return nil }

// OpMatrixTimesScalar scales a floating-point matrix.
type OpMatrixTimesScalar struct {
	ResultType Id
	ResultId   Id
	Vector     Id
	Scalar     Id
}

func (c *OpMatrixTimesScalar) Opcode() uint32 { return 137 }
func (c *OpMatrixTimesScalar) Verify() error  { return nil }

// OpVectorTimesMatrix performs Linear-algebraic Vector X Matrix.
type OpVectorTimesMatrix struct {
	ResultType Id
	ResultId   Id
	Vector     Id
	Matrix     Id
}

func (c *OpVectorTimesMatrix) Opcode() uint32 { return 138 }
func (c *OpVectorTimesMatrix) Verify() error  { return nil }

// OpMatrixTimesVector performs Linear-algebraic Vector X Matrix.
type OpMatrixTimesVector struct {
	ResultType Id
	ResultId   Id
	Matrix     Id
	Vector     Id
}

func (c *OpMatrixTimesVector) Opcode() uint32 { return 139 }
func (c *OpMatrixTimesVector) Verify() error  { return nil }

// OpMatrixTimesMatrix performs Linear-algebraic multiply of Left X Right.
type OpMatrixTimesMatrix struct {
	ResultType Id
	ResultId   Id
	Left       Id
	Right      Id
}

func (c *OpMatrixTimesMatrix) Opcode() uint32 { return 140 }
func (c *OpMatrixTimesMatrix) Verify() error  { return nil }

// OpOuterProduct performs Linear-algebraic outer product of Vector 1 and Vector 2.
type OpOuterProduct struct {
	ResultType Id
	ResultId   Id
	Vector1    Id
	Vector2    Id
}

func (c *OpOuterProduct) Opcode() uint32 { return 141 }
func (c *OpOuterProduct) Verify() error  { return nil }

// OpDot performs Dot product of Vector 1 and Vector 2
type OpDot struct {
	ResultType Id
	ResultId   Id
	Vector1    Id
	Vector2    Id
}

func (c *OpDot) Opcode() uint32 { return 142 }
func (c *OpDot) Verify() error  { return nil }

// OpShiftRightLogical shifts the bits in Operand 1 right by the number
// of bits specified in Operand 2
type OpShiftRightLogical struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpShiftRightLogical) Opcode() uint32 { return 143 }
func (c *OpShiftRightLogical) Verify() error  { return nil }

// OpShiftRightArithmetic shifts the bits in Operand 1 right by the number of
// bits specified in Operand 2
type OpShiftRightArithmetic struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpShiftRightArithmetic) Opcode() uint32 { return 144 }
func (c *OpShiftRightArithmetic) Verify() error  { return nil }

// OpShiftLeftLogical shifts the bits in Operand 1 left by the number
// of bits specified in Operand 2
type OpShiftLeftLogical struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpShiftLeftLogical) Opcode() uint32 { return 145 }
func (c *OpShiftLeftLogical) Verify() error  { return nil }

// OpBitwiseOr result is 1 if either Operand 1 or Operand 2 is 1.
// Result is 0 if both Operand 1 and Operand 2 are 0.
type OpBitwiseOr struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpBitwiseOr) Opcode() uint32 { return 149 }
func (c *OpBitwiseOr) Verify() error  { return nil }

// OpBitwiseXor result is 1 if exactly one of Operand 1 or Operand 2 is 1.
// Result is 0 if Operand 1 and Operand 2 have the same value
type OpBitwiseXor struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpBitwiseXor) Opcode() uint32 { return 150 }
func (c *OpBitwiseXor) Verify() error  { return nil }

// OpBitwiseAnd result is 1 if both Operand 1 and Operand 2 are 1.
// Result is 0 if either Operand 1 or Operand 2 are 0
type OpBitwiseAnd struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpBitwiseAnd) Opcode() uint32 { return 151 }
func (c *OpBitwiseAnd) Verify() error  { return nil }

func init() {
	Bind(func() Instruction { return &OpSNegate{} })
	Bind(func() Instruction { return &OpFNegate{} })
	Bind(func() Instruction { return &OpNot{} })
	Bind(func() Instruction { return &OpIAdd{} })
	Bind(func() Instruction { return &OpFAdd{} })
	Bind(func() Instruction { return &OpISub{} })
	Bind(func() Instruction { return &OpFSub{} })
	Bind(func() Instruction { return &OpIMul{} })
	Bind(func() Instruction { return &OpFMul{} })
	Bind(func() Instruction { return &OpUDiv{} })
	Bind(func() Instruction { return &OpSDiv{} })
	Bind(func() Instruction { return &OpFDiv{} })
	Bind(func() Instruction { return &OpUMod{} })
	Bind(func() Instruction { return &OpSRem{} })
	Bind(func() Instruction { return &OpSMod{} })
	Bind(func() Instruction { return &OpFRem{} })
	Bind(func() Instruction { return &OpFMod{} })
	Bind(func() Instruction { return &OpVectorTimesScalar{} })
	Bind(func() Instruction { return &OpMatrixTimesScalar{} })
	Bind(func() Instruction { return &OpVectorTimesMatrix{} })
	Bind(func() Instruction { return &OpMatrixTimesVector{} })
	Bind(func() Instruction { return &OpMatrixTimesMatrix{} })
	Bind(func() Instruction { return &OpOuterProduct{} })
	Bind(func() Instruction { return &OpDot{} })
	Bind(func() Instruction { return &OpShiftRightLogical{} })
	Bind(func() Instruction { return &OpShiftRightArithmetic{} })
	Bind(func() Instruction { return &OpShiftLeftLogical{} })
	Bind(func() Instruction { return &OpBitwiseOr{} })
	Bind(func() Instruction { return &OpBitwiseXor{} })
	Bind(func() Instruction { return &OpBitwiseAnd{} })
}
