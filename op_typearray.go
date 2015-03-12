// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpTypeArray declares a new array type: a dynamically-indexable ordered
// aggregate of elements all having the same type.
type OpTypeArray struct {
	// The <id> of the new array type.
	ResultId uint32

	// The type of each element in the array
	ElementType uint32

	// The number of elements in the array. It must be at least 1.
	//
	// Length must come from a constant instruction of an Integer-type
	// scalar whose value is at least 1.
	Length uint32
}

func (c *OpTypeArray) Opcode() uint32 { return 16 }

func bindOpTypeArray(set *InstructionSet) {
	set.Set(
		(&OpTypeArray{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 3 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpTypeArray{
					ResultId:    argv[0],
					ElementType: argv[1],
					Length:      argv[2],
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpTypeArray)
				out[0] = EncodeOpcode(4, v.Opcode())
				out[1] = v.ResultId
				out[2] = v.ElementType
				out[3] = v.Length
				return nil
			},
		},
	)
}
