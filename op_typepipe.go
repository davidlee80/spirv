// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpTypePipe declares an OpenCL pipe object type.
type OpTypePipe struct {
	// The <id> of the new pipe type.
	Result uint32

	// Type is the data type of the pipe.
	Type uint32

	// Qualifier is the pipe access qualifier.
	AccessQualifier uint32
}

func (c *OpTypePipe) Opcode() uint32 { return 26 }

func bindOpTypePipe(set *InstructionSet) {
	set.Set(
		(&OpTypePipe{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 3 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpTypePipe{
					Result:          argv[0],
					Type:            argv[1],
					AccessQualifier: argv[2],
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpTypePipe)
				out[0] = EncodeOpcode(4, v.Opcode())
				out[1] = v.Result
				out[2] = v.Type
				out[3] = v.AccessQualifier
				return nil
			},
		},
	)
}
