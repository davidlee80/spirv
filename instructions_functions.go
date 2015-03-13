// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpFunction defines a function body. This instruction must be
// immediately followed by one OpFunctionParameter instruction per each
// formal parameter of this function.
//
// This function’s body will terminate with the next OpFunctionEnd
// instruction.
type OpFunction struct {
	ResultType  uint32
	ResultId    uint32
	ControlMask uint32
	Type        uint32
}

func (c *OpFunction) Opcode() uint32 { return 40 }
func (c *OpFunction) Verify() error  { return nil }

// OpFunctionParameter declares the <id> for a formal parameter belonging
// to the current function.
//
// This instruction must immediately follow an OpFunction or OpFunctionParameter
// instruction. The order of contiguous OpFunctionParameter instructions is
// the same order arguments will be listed in an OpFunctionCall instruction to
// this function. It is also the same order in which Parameter Type operands
// are listed in the OpTypeFunction of the Function Type operand for this
// function’s OpFunction instruction.
type OpFunctionParameter struct {
	ResultType uint32
	ResultId   uint32
}

func (c *OpFunctionParameter) Opcode() uint32 { return 41 }
func (c *OpFunctionParameter) Verify() error  { return nil }

// OpFunctionParameter is the last instruction of a function definition.
type OpFunctionEnd struct{}

func (c *OpFunctionEnd) Opcode() uint32 { return 42 }
func (c *OpFunctionEnd) Verify() error  { return nil }

// OpFunctionCall defines a function call.
//
// Note: A forward call is possible because there is no missing type
// information: Result Type must match the Return Type of the function, and
// the calling argument types must match the formal parameter types.
type OpFunctionCall struct {
	ResultType uint32
	ResultId   uint32
	FunctionId uint32
	Argv       []uint32
}

func (c *OpFunctionCall) Opcode() uint32 { return 43 }
func (c *OpFunctionCall) Verify() error  { return nil }

func init() {
	Bind(func() Instruction { return &OpFunction{} })
	Bind(func() Instruction { return &OpFunctionParameter{} })
	Bind(func() Instruction { return &OpFunctionEnd{} })
	Bind(func() Instruction { return &OpFunctionCall{} })
}