// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpTypePointer declares a new pointer type.
type OpTypePointer struct {
	// The <id> of the new integer type.
	Result uint32

	// The storage class of the memory holding the object pointed to.
	StorageClass uint32

	// The type of the object pointed to.
	Type uint32
}

func (c *OpTypePointer) Opcode() uint32 { return 20 }

func bindOpTypePointer(set *InstructionSet) {
	set.Set(
		(&OpTypePointer{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 3 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpTypePointer{
					Result:       argv[0],
					StorageClass: argv[1],
					Type:         argv[2],
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpTypePointer)
				out[0] = EncodeOpcode(4, v.Opcode())
				out[1] = v.Result
				out[2] = v.StorageClass
				out[3] = v.Type
				return nil
			},
		},
	)
}
