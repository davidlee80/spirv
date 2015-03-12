// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpStore stores data through a pointer.
type OpStore struct {
	// Pointer is the pointer to store through. It must have a type
	// of OpTypePointer whose operand is the same as the type of Object.
	Pointer uint32

	// Object is the object to store.
	Object uint32

	// MemoryAccess must be one or more Memory Access literals.
	MemoryAccess []uint32
}

func (c *OpStore) Opcode() uint32 { return 47 }

func bindOpStore(set *InstructionSet) {
	set.Set(
		(&OpStore{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 2 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpStore{
					Pointer:      argv[0],
					Object:       argv[1],
					MemoryAccess: Copy(argv[2:]),
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpStore)
				size := uint32(len(v.MemoryAccess))

				out[0] = EncodeOpcode(3+size, v.Opcode())
				out[1] = v.Pointer
				out[2] = v.Object
				copy(out[3:], v.MemoryAccess)
				return nil
			},
		},
	)
}
