// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpCopyMemory copies from the memory pointed to by Source to the
// memory pointed to by Target.
//
// Both operands must be non-void pointers of the same type.
// Matching storage class is not required. The amount of memory copied is
// the size of the type pointed to.
type OpCopyMemory struct {
	// The target address.
	Target uint32

	// The source address.
	Source uint32

	// MemoryAccess must be one or more Memory Access literals.
	MemoryAccess []uint32
}

func (c *OpCopyMemory) Opcode() uint32 { return 65 }

func bindOpCopyMemory(set *InstructionSet) {
	set.Set(
		(&OpCopyMemory{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 2 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpCopyMemory{
					Target:       argv[0],
					Source:       argv[1],
					MemoryAccess: Copy(argv[2:]),
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpCopyMemory)
				size := uint32(len(v.MemoryAccess))

				out[0] = EncodeOpcode(3+size, v.Opcode())
				out[1] = v.Target
				out[2] = v.Source
				copy(out[3:], v.MemoryAccess)
				return nil
			},
		},
	)
}
