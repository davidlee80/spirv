// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpAtomicInit initializes atomic memory to Value.
// This is not done atomically with respect to anything.
type OpAtomicInit struct {
	Pointer uint32
	Value   uint32
}

func (c *OpAtomicInit) Opcode() uint32 { return 191 }
func (c *OpAtomicInit) Verify() error  { return nil }

// OpAtomicLoad atomically loads through Pointer using the given Semantics.
//
// All subparts of the value that is loaded will be read atomically with
// respect to all other atomic accesses to it within Scope.
type OpAtomicLoad struct {
	ResultType      uint32
	ResultId        uint32
	Pointer         uint32
	ExecutionScope  uint32
	MemorySemantics uint32
}

func (c *OpAtomicLoad) Opcode() uint32 { return 192 }
func (c *OpAtomicLoad) Verify() error  { return nil }

// OpAtomicStore atomically stores through Pointer using the given Semantics.
//
// All subparts of Value will be written atomically with respect to all
// other atomic accesses to it within Scope.
type OpAtomicStore struct {
	Pointer         uint32
	ExecutionScope  uint32
	MemorySemantics uint32
	Value           uint32
}

func (c *OpAtomicStore) Opcode() uint32 { return 193 }
func (c *OpAtomicStore) Verify() error  { return nil }

// OpAtomicExchange performs the following steps atomically with respect to any
// other atomic accesses within Scope to the same location:
//
//   1) load through Pointer to get an Original Value,
//   2) get a New Value from copying Value, and
//   3) store the New Value back through Pointer.
//
// The instruction’s result is the Original Value.
type OpAtomicExchange struct {
	ResultType      uint32
	ResultId        uint32
	Pointer         uint32
	ExecutionScope  uint32
	MemorySemantics uint32
	Value           uint32
}

func (c *OpAtomicExchange) Opcode() uint32 { return 194 }
func (c *OpAtomicExchange) Verify() error  { return nil }

// OpAtomicCompareExchange performs the following steps atomically with respect
// to any other atomic accesses within Scope to the same location:
//
//   1) load through Pointer to get an Original Value,
//   2) get a New Value by selecting Value if Original Value equals Comparator
//      or selecting Original Value otherwise, and
//   3) store the New Value back through Pointer.
//
// The instruction’s result is the Original Value.
type OpAtomicCompareExchange struct {
	ResultType      uint32
	ResultId        uint32
	Pointer         uint32
	ExecutionScope  uint32
	MemorySemantics uint32
	Value           uint32
	Comparator      uint32
}

func (c *OpAtomicCompareExchange) Opcode() uint32 { return 195 }
func (c *OpAtomicCompareExchange) Verify() error  { return nil }

// OpAtomicCompareExchangeWeak performs the following steps atomically
// with respect to any other atomic accesses within Scope to the same location:
//
//   1) load through Pointer to get an Original Value,
//   2) get a New Value by selecting Value if Original Value equals Comparator
//      or selecting Original Value otherwise, and
//   3) store the New Value back through Pointer.
//
type OpAtomicCompareExchangeWeak struct {
	ResultType      uint32
	ResultId        uint32
	Pointer         uint32
	ExecutionScope  uint32
	MemorySemantics uint32
	Value           uint32
	Comparator      uint32
}

func (c *OpAtomicCompareExchangeWeak) Opcode() uint32 { return 196 }
func (c *OpAtomicCompareExchangeWeak) Verify() error  { return nil }

// OpAtomicIIncrement performs the following steps atomically with respect
// to any other atomic accesses within Scope to the same location
//
//   1) load through Pointer to get an Original Value,
//   2) get a New Value through integer addition of 1 to Original Value, and
//   3) store the New Value back through Pointer.
//
type OpAtomicIIncrement struct {
	ResultType      uint32
	ResultId        uint32
	Pointer         uint32
	ExecutionScope  uint32
	MemorySemantics uint32
}

func (c *OpAtomicIIncrement) Opcode() uint32 { return 197 }
func (c *OpAtomicIIncrement) Verify() error  { return nil }

// OpAtomicIDecrement performs the following steps atomically with respect
// to any other atomic accesses within Scope to the same location
//
//   1) load through Pointer to get an Original Value,
//   2) get a New Value through integer subtraction of 1 from Original Value, and
//   3) store the New Value back through Pointer.
//
type OpAtomicIDecrement struct {
	ResultType      uint32
	ResultId        uint32
	Pointer         uint32
	ExecutionScope  uint32
	MemorySemantics uint32
}

func (c *OpAtomicIDecrement) Opcode() uint32 { return 198 }
func (c *OpAtomicIDecrement) Verify() error  { return nil }

// OpAtomicIAdd performs the following steps atomically with respect to
// any other atomic accesses within Scope to the same location:
//
//   1) load through Pointer to get an Original Value,
//   2) get a New Value by integer addition of Original Value and Value, and
//   3) store the New Value back through Pointer.
//
type OpAtomicIAdd struct {
	ResultType      uint32
	ResultId        uint32
	Pointer         uint32
	ExecutionScope  uint32
	MemorySemantics uint32
	Value           uint32
}

func (c *OpAtomicIAdd) Opcode() uint32 { return 199 }
func (c *OpAtomicIAdd) Verify() error  { return nil }

// OpAtomicISub performs the following steps atomically with respect to
// any other atomic accesses within Scope to the same location:
//
//   1) load through Pointer to get an Original Value,
//   2) get a New Value by integer subtraction of Value from Original Value, and
//   3) store the New Value back through Pointer.
//
type OpAtomicISub struct {
	ResultType      uint32
	ResultId        uint32
	Pointer         uint32
	ExecutionScope  uint32
	MemorySemantics uint32
	Value           uint32
}

func (c *OpAtomicISub) Opcode() uint32 { return 200 }
func (c *OpAtomicISub) Verify() error  { return nil }

// OpAtomicUMin performs the following steps atomically with respect to
// any other atomic accesses within Scope to the same location:
//
//   1) load through Pointer to get an Original Value,
//   2) get a New Value by finding the smallest unsigned integer of
//      Original Value and Value, and
//   3) store the New Value back through Pointer.
//
type OpAtomicUMin struct {
	ResultType      uint32
	ResultId        uint32
	Pointer         uint32
	ExecutionScope  uint32
	MemorySemantics uint32
	Value           uint32
}

func (c *OpAtomicUMin) Opcode() uint32 { return 201 }
func (c *OpAtomicUMin) Verify() error  { return nil }

// OpAtomicUMax performs the following steps atomically with respect to
// any other atomic accesses within Scope to the same location:
//
//   1) load through Pointer to get an Original Value,
//   2) get a New Value by finding the largest unsigned integer of Original
//      Value and Value, and
//   3) store the New Value back through Pointer.
//
type OpAtomicUMax struct {
	ResultType      uint32
	ResultId        uint32
	Pointer         uint32
	ExecutionScope  uint32
	MemorySemantics uint32
	Value           uint32
}

func (c *OpAtomicUMax) Opcode() uint32 { return 202 }
func (c *OpAtomicUMax) Verify() error  { return nil }

// OpAtomicAnd performs the following steps atomically with respect to
// any other atomic accesses within Scope to the same location:
//
//   1) load through Pointer to get an Original Value,
//   2) get a New Value by the bitwise AND of Original Value and Value, and
//   3) store the New Value back through Pointer.
//
type OpAtomicAnd struct {
	ResultType      uint32
	ResultId        uint32
	Pointer         uint32
	ExecutionScope  uint32
	MemorySemantics uint32
	Value           uint32
}

func (c *OpAtomicAnd) Opcode() uint32 { return 203 }
func (c *OpAtomicAnd) Verify() error  { return nil }

// OpAtomicOr performs the following steps atomically with respect to
// any other atomic accesses within Scope to the same location:
//
//   1) load through Pointer to get an Original Value,
//   2) get a New Value by the bitwise OR of Original Value and Value, and
//   3) store the New Value back through Pointer.
//
type OpAtomicOr struct {
	ResultType      uint32
	ResultId        uint32
	Pointer         uint32
	ExecutionScope  uint32
	MemorySemantics uint32
	Value           uint32
}

func (c *OpAtomicOr) Opcode() uint32 { return 204 }
func (c *OpAtomicOr) Verify() error  { return nil }

// OpAtomicXor performs the following steps atomically with respect to
// any other atomic accesses within Scope to the same location:
//
//   1) load through Pointer to get an Original Value,
//   2) get a New Value by the bitwise exclusive OR of Original
//      Value and Value, and
//   3) store the New Value back through Pointer.
//
type OpAtomicXor struct {
	ResultType      uint32
	ResultId        uint32
	Pointer         uint32
	ExecutionScope  uint32
	MemorySemantics uint32
	Value           uint32
}

func (c *OpAtomicXor) Opcode() uint32 { return 205 }
func (c *OpAtomicXor) Verify() error  { return nil }

func init() {
	Bind(func() Instruction { return &OpAtomicInit{} })
	Bind(func() Instruction { return &OpAtomicLoad{} })
	Bind(func() Instruction { return &OpAtomicStore{} })
	Bind(func() Instruction { return &OpAtomicExchange{} })
	Bind(func() Instruction { return &OpAtomicCompareExchange{} })
	Bind(func() Instruction { return &OpAtomicCompareExchangeWeak{} })
	Bind(func() Instruction { return &OpAtomicIIncrement{} })
	Bind(func() Instruction { return &OpAtomicIDecrement{} })
	Bind(func() Instruction { return &OpAtomicIAdd{} })
	Bind(func() Instruction { return &OpAtomicISub{} })
	Bind(func() Instruction { return &OpAtomicUMin{} })
	Bind(func() Instruction { return &OpAtomicUMax{} })
	Bind(func() Instruction { return &OpAtomicAnd{} })
	Bind(func() Instruction { return &OpAtomicOr{} })
	Bind(func() Instruction { return &OpAtomicXor{} })
}
