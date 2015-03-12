// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpArraylength results in the length of a run-time array.
type OpArraylength struct {
	ResultType uint32
	ResultId   uint32

	// Structure must be an object of type OpTypeStruct that contains
	// a member that is a run-time array.
	Structure uint32

	// Array member is a member number of Structure that must have a
	// type from OpTypeRuntimeArray.
	Member uint32
}

func (c *OpArraylength) Opcode() uint32 { return 121 }

func bindOpArraylength(set *InstructionSet) {
	set.Set(
		(&OpArraylength{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 4 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpArraylength{
					ResultType: argv[0],
					ResultId:   argv[1],
					Structure:  argv[2],
					Member:     argv[3],
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpArraylength)
				out[0] = EncodeOpcode(5, v.Opcode())
				out[1] = v.ResultType
				out[2] = v.ResultId
				out[3] = v.Structure
				out[4] = v.Member
				return nil
			},
		},
	)
}
