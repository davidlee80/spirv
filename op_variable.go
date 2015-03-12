// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpVariable allocates an object in memory, resulting in a pointer
// to it, which can be used with OpLoad and OpStore.
type OpVariable struct {
	// Result Type is a type from OpTypePointer, where the type pointed to
	// is the type of object in memory.
	ResultType uint32

	ResultId uint32

	// Storage Class is the kind of memory holding the object.
	Storage StorageClass

	// Initializer is optional. If Initializer is present, it will be the
	// initial value of the variable’s memory content. Initializer must
	// be an <id> from a constant instruction. Initializer must have the same
	// type as the type pointed to by Result Type.
	Initializer uint32
}

func (c *OpVariable) Opcode() uint32 { return 38 }

func bindOpVariable(set *InstructionSet) {
	set.Set(
		(&OpVariable{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 3 {
					return nil, ErrMissingInstructionArgs
				}

				op := &OpVariable{
					ResultType: argv[0],
					ResultId:   argv[1],
					Storage:    StorageClass(argv[2]),
				}

				if len(argv) > 3 {
					op.Initializer = argv[3]
				}

				return op, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpVariable)
				size := uint32(4)

				if v.Initializer != 0 {
					size++
				}

				out[0] = EncodeOpcode(size, v.Opcode())
				out[1] = v.ResultType
				out[2] = v.ResultId
				out[3] = uint32(v.Storage)

				if v.Initializer != 0 {
					out[4] = v.Initializer
				}

				return nil
			},
		},
	)
}
