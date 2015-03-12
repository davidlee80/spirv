// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpTypeStruct declares a new structure type: an aggregate
// of heteregeneous members
type OpTypeStruct struct {
	// The <id> of the new array type.
	ResultId uint32

	// Member N type is the type of member N of the structure. The first
	// member is member 0, the next is member 1, . . .
	Members []uint32
}

func (c *OpTypeStruct) Opcode() uint32 { return 18 }

func bindOpTypeStruct(set *InstructionSet) {
	set.Set(
		(&OpTypeStruct{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 1 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpTypeStruct{
					ResultId: argv[0],
					Members:  Copy(argv[1:]),
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpTypeStruct)
				size := uint32(len(v.Members))
				out[0] = EncodeOpcode(2+size, v.Opcode())
				out[1] = v.ResultId
				copy(out[2:], v.Members)
				return nil
			},
		},
	)
}
