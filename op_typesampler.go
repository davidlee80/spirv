// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpTypeSampler declares a new sampler type.
//
// It is consumed, for example, by OpTextureSample.
// This type is opaque: values of this type have no defined physical
// size or bit pattern..
type OpTypeSampler struct {
	// The <id> of the new sampler type.
	Result uint32

	// A scalar type, of the type of the components resulting from
	// sampling or loading through this sampler
	SampledType uint32

	// The texture dimensionality.
	Dim Dimensionality

	// Content must be one of the following indicated values:
	//
	//   0: indicates a texture, no filter (no sampling state)
	//   1: indicates an image
	//   2: indicates both a texture and filter (sampling state),
	//      see OpTypeFilter
	//
	Content uint32

	// Arrayed must be one of the following indicated values:
	//
	//   0: indicates non-arrayed content
	//   1: indicates arrayed content
	//
	Arrayed uint32

	// Compare must be one of the following indicated values:
	//
	//   0: indicates depth comparisons are not done
	//   1: indicates depth comparison are done
	//
	Compare uint32

	// MS is multisampled and must be one of the following indicated values:
	//
	//   0: indicates single-sampled content
	//   1: indicates multisampled content
	//
	MS uint32

	// Qualifier is an image access qualifier.
	Qualifier AccessQualifier
}

func (c *OpTypeSampler) Opcode() uint32 { return 14 }

func bindOpTypeSampler(set *InstructionSet) {
	set.Set(
		(&OpTypeSampler{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 7 {
					return nil, ErrMissingInstructionArgs
				}

				op := &OpTypeSampler{
					Result:      argv[0],
					SampledType: argv[1],
					Dim:         Dimensionality(argv[2]),
					Content:     argv[3],
					Arrayed:     argv[4],
					Compare:     argv[5],
					MS:          argv[6],
				}

				// The qualifier is optional.
				if len(argv) > 7 {
					op.Qualifier = AccessQualifier(argv[7])
				}

				return op, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpTypeSampler)
				size := uint32(8)

				if v.Qualifier != 0 {
					size++
				}

				out[0] = EncodeOpcode(size, v.Opcode())
				out[1] = v.Result
				out[2] = v.SampledType
				out[3] = uint32(v.Dim)
				out[4] = v.Content
				out[5] = v.Arrayed
				out[6] = v.Compare
				out[7] = v.MS

				if v.Qualifier != 0 {
					out[8] = uint32(v.Qualifier)
				}

				return nil
			},
		},
	)
}
