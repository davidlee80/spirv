// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpGenericPtrMemSemantics returns a valid Memory Semantics
// value for ptr.
type OpGenericPtrMemSemantics struct {
	ResultType uint32 // Result Type must be a 32-bits wide OpTypeInt value
	ResultId   uint32
	Ptr        uint32 // Ptr must point to Generic.
}

func (c *OpGenericPtrMemSemantics) Opcode() uint32 { return 233 }

func bindOpGenericPtrMemSemantics(set *InstructionSet) {
	set.Set(
		(&OpGenericPtrMemSemantics{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 3 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpGenericPtrMemSemantics{
					ResultType: argv[0],
					ResultId:   argv[1],
					Ptr:        argv[2],
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpGenericPtrMemSemantics)
				out[0] = EncodeOpcode(4, v.Opcode())
				out[1] = v.ResultType
				out[2] = v.ResultId
				out[3] = v.Ptr
				return nil
			},
		},
	)
}
