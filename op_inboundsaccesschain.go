// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpInboundsAccessChain has the same semantics as OpAccessChain, with the
// addition that the resulting pointer is known to point within the base object.
type OpInboundsAccessChain struct {
	ResultType uint32
	ResultId   uint32

	// Base must be a pointer type, pointing to the base of the object.
	Base uint32

	// Indices walk the type hierarchy to the desired depth, potentially
	// down to scalar granularity. The type of the pointer created will be to
	// the type reached by walking the type hierarchy down to the last
	// provided index.
	Indices []uint32
}

func (c *OpInboundsAccessChain) Opcode() uint32 { return 94 }

func init() {
	Bind(
		(&OpInboundsAccessChain{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 3 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpInboundsAccessChain{
					ResultType: argv[0],
					ResultId:   argv[1],
					Base:       argv[2],
					Indices:    Copy(argv[3:]),
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpInboundsAccessChain)
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
