// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpAccessChain creates a pointer into a composite object that can be
// used with OpLoad and OpStore.
//
// The storage class of the pointer created will be the same as the storage
// class of the base operand.
type OpAccessChain struct {
	ResultType uint32
	ResultId   uint32

	// Base must be a pointer type, pointing to the base of the object.
	Base uint32

	// Indexes walk the type hierarchy to the desired depth, potentially
	// down to scalar granularity. The type of the pointer created will be to
	// the type reached by walking the type hierarchy down to the last
	// provided index.
	Indices []uint32
}

func (c *OpAccessChain) Opcode() uint32 { return 93 }

func bindOpAccessChain(set *InstructionSet) {
	set.Set(
		(&OpAccessChain{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 3 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpAccessChain{
					ResultType: argv[0],
					ResultId:   argv[1],
					Base:       argv[2],
					Indices:    Copy(argv[3:]),
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpAccessChain)
				size := uint32(len(v.Indices))

				out[0] = EncodeOpcode(4+size, v.Opcode())
				out[1] = v.ResultType
				out[2] = v.ResultId
				out[3] = v.Base
				copy(out[4:], v.Indices)
				return nil
			},
		},
	)
}
