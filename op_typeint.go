// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpTypeInt represents the OpTypeInt instruction.
type OpTypeInt struct {
	// The <id> of the new integer type.
	Result uint32

	// Specifies how many bits wide the type is.
	// The bit pattern of a signed integer value is two’s complement.
	Width uint32

	// Signedness specifies whether there are signed semantics to
	// preserve or validate.
	//
	//   0: indicates unsigned, or no signedness semantics.
	//   1: indicates signed semantics.
	//
	// In all cases, the type of operation of an instruction comes from
	// the instruction’s opcode, not the signedness of the operands.
	Signedness uint32
}

func (c *OpTypeInt) Opcode() uint32 { return 10 }

// NewOpTypeInt creates a new codec for the OpTypeInt instruction.
func NewOpTypeInt() Codec {
	return Codec{
		Decode: func(argv []uint32) (Instruction, error) {
			if len(argv) < 3 {
				return nil, ErrMissingInstructionArgs
			}

			return &OpTypeInt{
				Result:     argv[0],
				Width:      argv[1],
				Signedness: argv[2],
			}, nil
		},
		Encode: func(i Instruction, out []uint32) error {
			v := i.(*OpTypeInt)
			out[0] = EncodeOpcode(4, v.Opcode())
			out[1] = v.Result
			out[2] = v.Width
			out[3] = v.Signedness
			return nil
		},
	}
}
