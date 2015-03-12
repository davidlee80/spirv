// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpCopyMemorySized copies from the memory pointed to by Source to the
// memory pointed to by Target.
//
// Both operands must be non-void pointers of the same type.
// Matching storage class is not required.
type OpCopyMemorySized struct {
	// The target address.
	Target uint32

	// The source address.
	Source uint32

	// Size is the number of bytes to copy.
	Size uint32

	// MemoryAccess must be one or more Memory Access literals.
	MemoryAccess []uint32
}

func (c *OpCopyMemorySized) Opcode() uint32 { return 66 }

func init() {
	Bind(
		(&OpCopyMemorySized{}).Opcode(),
		Codec{
			Decode: func(argv []uint32) (Instruction, error) {
				if len(argv) < 3 {
					return nil, ErrMissingInstructionArgs
				}

				return &OpCopyMemorySized{
					Target:       argv[0],
					Source:       argv[1],
					Size:         argv[2],
					MemoryAccess: Copy(argv[3:]),
				}, nil
			},
			Encode: func(i Instruction, out []uint32) error {
				v := i.(*OpCopyMemorySized)
				size := uint32(len(v.MemoryAccess))

				out[0] = EncodeOpcode(4+size, v.Opcode())
				out[1] = v.Target
				out[2] = v.Source
				out[3] = v.Size
				copy(out[4:], v.MemoryAccess)
				return nil
			},
		},
	)
}
