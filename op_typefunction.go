// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpTypeFunction declares a new function type.
//
// OpFunction will use this to declare the return type and
// parameter types of a function
type OpTypeFunction struct {
	// The <id> of the new function type.
	ResultId uint32

	// The type of the return value of functions of this type.
	// If the function has no return value, Return Type should
	// be from OpTypeVoid.
	ReturnType uint32

	// Parameter N Type is the type <id> of the type of parameter N
	Parameters []uint32
}

func (c *OpTypeFunction) Opcode() uint32 { return 21 }

func bindOpTypeFunction(set *InstructionSet) {
	set.Set(
		(&OpTypeFunction{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 2 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpTypeFunction{
					ResultId:   argv[0],
					ReturnType: argv[1],
					Parameters: Copy(argv[2:]),
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpTypeFunction)
				size := uint32(len(v.Parameters))

				out[0] = EncodeOpcode(3+size, v.Opcode())
				out[1] = v.ResultId
				out[2] = v.ReturnType
				copy(out[3:], v.Parameters)
				return nil
			},
		},
	)
}
