// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpTypeRuntimeArray declares a new run-time array type.
// Its length is not known at compile time.
//
// Objects of this type can only be created with OpVariable
// using the Uniform Storage Class.
type OpTypeRuntimeArray struct {
	// The <id> of the new run-time array type.
	ResultId uint32

	// The type of each element in the array.
	// See OpArrayLength for getting the Length of an array of this type.
	ElementType uint32
}

func (c *OpTypeRuntimeArray) Opcode() uint32 { return 17 }

func init() {
	Bind(
		(&OpTypeRuntimeArray{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 2 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpTypeRuntimeArray{
					ResultId:    argv[0],
					ElementType: argv[1],
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpTypeRuntimeArray)
				out[0] = EncodeOpcode(3, v.Opcode())
				out[1] = v.ResultId
				out[2] = v.ElementType
				return nil
			},
		},
	)
}
