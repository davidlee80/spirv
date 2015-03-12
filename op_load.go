// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpLoad loads data through a pointer.
type OpLoad struct {
	// Result Type is a type from OpTypePointer whose type pointed to is
	// the type of one of the N objects allocated in memory
	ResultType uint32

	ResultId uint32

	// Pointer is the pointer to load through. It must have a type of
	// OpTypePointer whose operand is the same as Result Type.
	Pointer uint32

	// MemoryAccess must be one or more Memory Access literals.
	MemoryAccess []uint32
}

func (c *OpLoad) Opcode() uint32 { return 46 }

func init() {
	Bind(
		(&OpLoad{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 3 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpLoad{
					ResultType:   argv[0],
					ResultId:     argv[1],
					Pointer:      argv[2],
					MemoryAccess: Copy(argv[3:]),
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpLoad)
				size := uint32(len(v.MemoryAccess))

				out[0] = EncodeOpcode(4+size, v.Opcode())
				out[1] = v.ResultType
				out[2] = v.ResultId
				out[3] = v.Pointer
				copy(out[4:], v.MemoryAccess)
				return nil
			},
		},
	)
}
