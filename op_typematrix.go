// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpTypeMatrix declares a new matrix type.
type OpTypeMatrix struct {
	// The <id> of the new matrix type
	Result uint32

	// The type of each column in the matrix. It must be vector type.
	ColumnType uint32

	// The number of columns in the new matrix type. It must be at least 2.
	ColumnCount uint32
}

func (c *OpTypeMatrix) Opcode() uint32 { return 13 }

func bindOpTypeMatrix(set *InstructionSet) {
	set.Set(
		(&OpTypeMatrix{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 3 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpTypeMatrix{
					Result:      argv[0],
					ColumnType:  argv[1],
					ColumnCount: argv[2],
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpTypeMatrix)
				out[0] = EncodeOpcode(4, v.Opcode())
				out[1] = v.Result
				out[2] = v.ColumnType
				out[3] = v.ColumnCount
				return nil
			},
		},
	)
}
