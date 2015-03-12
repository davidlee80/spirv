// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpImagePointer forms a pointer to a texel of an image.
// Use of such a pointer is limited to atomic operations.
//
// TODO: This requires an Image storage class to be added.
type OpImagePointer struct {
	ResultType uint32
	ResultId   uint32

	// Image is a pointer to a variable of type of OpTypeSampler.
	Image uint32

	// Coordinate and Sample specify which texel and sample within
	// the image to form an address of.
	Coordinate uint32
	Sample     uint32
}

func (c *OpImagePointer) Opcode() uint32 { return 190 }

func bindOpImagePointer(set *InstructionSet) {
	set.Set(
		(&OpImagePointer{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 5 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpImagePointer{
					ResultType: argv[0],
					ResultId:   argv[1],
					Image:      argv[2],
					Coordinate: argv[3],
					Sample:     argv[4],
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpImagePointer)
				out[0] = EncodeOpcode(6, v.Opcode())
				out[1] = v.ResultType
				out[2] = v.ResultId
				out[3] = v.Image
				out[4] = v.Coordinate
				out[5] = v.Sample
				return nil
			},
		},
	)
}
