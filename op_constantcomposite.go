// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

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

func bindOpConstantComposite(set *InstructionSet) {
	set.Set(
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
