// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpConstantTrue declares a true Boolean-type scalar constant.
type OpConstantTrue struct {
	// Result Type must be the scalar Boolean type.
	ResultType uint32

	// The <id> of the new constant type.
	ResultId uint32
}

func (c *OpConstantTrue) Opcode() uint32 { return 27 }

func init() {
	Bind(
		(&OpConstantTrue{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 2 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpConstantTrue{
					ResultType: argv[0],
					ResultId:   argv[1],
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpConstantTrue)
				out[0] = EncodeOpcode(3, v.Opcode())
				out[1] = v.ResultType
				out[2] = v.ResultId
				return nil
			},
		},
	)
}

// OpConstantFalse declares a true Boolean-type scalar constant.
type OpConstantFalse struct {
	// Result Type must be the scalar Boolean type.
	ResultType uint32

	// The <id> of the new constant type.
	ResultId uint32
}

func (c *OpConstantFalse) Opcode() uint32 { return 28 }

func init() {
	Bind(
		(&OpConstantFalse{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 2 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpConstantFalse{
					ResultType: argv[0],
					ResultId:   argv[1],
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpConstantFalse)
				out[0] = EncodeOpcode(3, v.Opcode())
				out[1] = v.ResultType
				out[2] = v.ResultId
				return nil
			},
		},
	)
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

func init() {
	Bind(
		(&OpConstant{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 2 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpConstant{
					ResultType: argv[0],
					ResultId:   argv[1],
					Value:      Copy(argv[2:]),
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpConstant)
				size := uint32(len(v.Value))
				out[0] = EncodeOpcode(3+size, v.Opcode())
				out[1] = v.ResultType
				out[2] = v.ResultId
				copy(out[3:], v.Value)
				return nil
			},
		},
	)
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

func init() {
	Bind(
		(&OpConstantComposite{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 2 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpConstantComposite{
					ResultType:   argv[0],
					ResultId:     argv[1],
					Constituents: Copy(argv[2:]),
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpConstantComposite)
				size := uint32(len(v.Constituents))

				out[0] = EncodeOpcode(3+size, v.Opcode())
				out[1] = v.ResultType
				out[2] = v.ResultId
				copy(out[3:], v.Constituents)
				return nil
			},
		},
	)
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

func init() {
	Bind(
		(&OpConstantSampler{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 5 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpConstantSampler{
					ResultType: argv[0],
					ResultId:   argv[1],
					Addressing: argv[2],
					Param:      argv[3],
					Filter:     argv[4],
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpConstantSampler)
				out[0] = EncodeOpcode(6, v.Opcode())
				out[1] = v.ResultType
				out[2] = v.ResultId
				out[3] = v.Addressing
				out[4] = v.Param
				out[5] = v.Filter
				return nil
			},
		},
	)
}

// OpConstantNullPointer declares a new null pointer constant.
type OpConstantNullPointer struct {
	// Result Type must be the scalar Boolean type.
	ResultType uint32

	// The <id> of the new constant type.
	ResultId uint32
}

func (c *OpConstantNullPointer) Opcode() uint32 { return 32 }

func init() {
	Bind(
		(&OpConstantNullPointer{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) != 2 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpConstantNullPointer{
					ResultType: argv[0],
					ResultId:   argv[1],
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpConstantNullPointer)
				out[0] = EncodeOpcode(3, v.Opcode())
				out[1] = v.ResultType
				out[2] = v.ResultId
				return nil
			},
		},
	)
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

func init() {
	Bind(
		(&OpConstantNullObject{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 2 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpConstantNullObject{
					ResultType: argv[0],
					ResultId:   argv[1],
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpConstantNullObject)
				out[0] = EncodeOpcode(3, v.Opcode())
				out[1] = v.ResultType
				out[2] = v.ResultId
				return nil
			},
		},
	)
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

func init() {
	Bind(
		(&OpSpecConstantTrue{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 2 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpSpecConstantTrue{
					ResultType: argv[0],
					ResultId:   argv[1],
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpSpecConstantTrue)
				out[0] = EncodeOpcode(3, v.Opcode())
				out[1] = v.ResultType
				out[2] = v.ResultId
				return nil
			},
		},
	)
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

func init() {
	Bind(
		(&OpSpecConstantFalse{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 2 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpSpecConstantFalse{
					ResultType: argv[0],
					ResultId:   argv[1],
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpSpecConstantFalse)
				out[0] = EncodeOpcode(3, v.Opcode())
				out[1] = v.ResultType
				out[2] = v.ResultId
				return nil
			},
		},
	)
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

func init() {
	Bind(
		(&OpSpecConstant{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 2 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpSpecConstant{
					ResultType: argv[0],
					ResultId:   argv[1],
					Value:      Copy(argv[2:]),
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpSpecConstant)
				size := uint32(len(v.Value))
				out[0] = EncodeOpcode(3+size, v.Opcode())
				out[1] = v.ResultType
				out[2] = v.ResultId
				copy(out[3:], v.Value)
				return nil
			},
		},
	)
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

func init() {
	Bind(
		(&OpSpecConstantComposite{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 2 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpSpecConstantComposite{
					ResultType:   argv[0],
					ResultId:     argv[1],
					Constituents: Copy(argv[2:]),
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpSpecConstantComposite)
				size := uint32(len(v.Constituents))
				out[0] = EncodeOpcode(3+size, v.Opcode())
				out[1] = v.ResultType
				out[2] = v.ResultId
				copy(out[3:], v.Constituents)
				return nil
			},
		},
	)
}
