// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpVariableArray allocates N objects sequentially in memory,
// resulting in a pointer to the first such object.
//
// This is not the same thing as allocating a single object that is an array.
type OpVariableArray struct {
	// Result Type is a type from OpTypePointer whose type pointed to is
	// the type of one of the N objects allocated in memory
	ResultType uint32

	ResultId uint32

	// Storage Class is the kind of memory holding the object.
	StorageClass uint32

	// N is the number of objects to allocate.
	N uint32
}

func (c *OpVariableArray) Opcode() uint32 { return 39 }

func init() {
	Bind(
		(&OpVariableArray{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 4 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpVariableArray{
					ResultType:   argv[0],
					ResultId:     argv[1],
					StorageClass: argv[2],
					N:            argv[3],
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpVariableArray)
				out[0] = EncodeOpcode(5, v.Opcode())
				out[1] = v.ResultType
				out[2] = v.ResultId
				out[3] = v.StorageClass
				out[4] = v.N
				return nil
			},
		},
	)
}
