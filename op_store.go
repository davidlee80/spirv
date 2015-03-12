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
	MemoryAccess []MemoryAccess
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

				op := &OpStore{
					Pointer: argv[0],
					Object:  argv[1],
				}

				if len(argv) > 2 {
					op.MemoryAccess = make([]MemoryAccess, len(argv[2:]))
					for i := range op.MemoryAccess {
						op.MemoryAccess[i] = MemoryAccess(argv[2+i])
					}
				}

				return op, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpStore)
				size := uint32(len(v.MemoryAccess))

				out[0] = EncodeOpcode(3+size, v.Opcode())
				out[1] = v.Pointer
				out[2] = v.Object

				for i, ma := range v.MemoryAccess {
					out[3+i] = uint32(ma)
				}

				return nil
			},
		},
	)
}
