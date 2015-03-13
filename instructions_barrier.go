// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpControlBarrier waits for other invocations of this module to reach this
// same point of execution. All invocations of this module within Scope must
// reach this point of execution before any will proceed beyond it.
//
// This instruction is only guaranteed to work correctly if placed strictly
// within dynamically uniform control flow within Scope. This ensures that if
// any invocation executes it, all invocations will execute it. If placed
// elsewhere, an invocation may stall indefinitely.
type OpControlBarrier struct {
	ExecutionScope uint32
}

func (c *OpControlBarrier) Opcode() uint32 { return 188 }
func (c *OpControlBarrier) Verify() error  { return nil }

// OpMemoryBarrier controls the order that memory accesses are observed.
//
// Ensures that memory accesses issued before this instruction will be observed
// before memory accesses issued after this instruction. This control is
// ensured only for memory accesses issued by this invocation and observed by
// another invocation executing within Scope.
type OpMemoryBarrier struct {
	ExecutionScope  uint32
	MemorySemantics uint32
}

func (c *OpMemoryBarrier) Opcode() uint32 { return 189 }
func (c *OpMemoryBarrier) Verify() error  { return nil }

func init() {
	Bind(func() Instruction { return &OpControlBarrier{} })
	Bind(func() Instruction { return &OpMemoryBarrier{} })
}
