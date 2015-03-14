// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "fmt"

// TODO: Context-aware validation. Most instructions here require context
// (i.e., module) to properly validate.

// OpPhi is the SSA phi function.
type OpPhi struct {
	ResultType Id
	ResultId   Id
	Operands   []Id
}

func (c *OpPhi) Opcode() uint32 { return 48 }
func (c *OpPhi) Optional() bool { return false }
func (c *OpPhi) Verify() error {
	if len(c.Operands) == 0 {
		return fmt.Errorf("OpPhi: expected operands")
	}

	if len(c.Operands)%2 != 0 {
		return fmt.Errorf("OpPhi: expected array of (Variable, ParentBlock) operand pairs")
	}

	return nil
}

// OpLoopMerge declares and controls a structured control-flow loop construct.
type OpLoopMerge struct {
	Label       Id
	LoopControl LoopControl
}

func (c *OpLoopMerge) Opcode() uint32 { return 206 }
func (c *OpLoopMerge) Optional() bool { return false }
func (c *OpLoopMerge) Verify() error {
	switch c.LoopControl {
	case LoopControlNoControl,
		LoopControlUnroll,
		LoopControlDontUnroll:
	default:
		return fmt.Errorf("OpLoopMerge.LoopControl: expected a Loop Control constant")
	}
	return nil
}

// OpSelectionMerge declares and controls a structured control-flow selection
// construct, used with OpBranchConditional or OpSwitch.
type OpSelectionMerge struct {
	Label            Id
	SelectionControl SelectionControl
}

func (c *OpSelectionMerge) Opcode() uint32 { return 207 }
func (c *OpSelectionMerge) Optional() bool { return false }
func (c *OpSelectionMerge) Verify() error {
	switch c.SelectionControl {
	case SelectionControlNoControl,
		SelectionControlFlatten,
		SelectionControlDontFlatten:
	default:
		return fmt.Errorf("OpSelectionMerge.SelectionControl: expected a Selection Control constant")
	}
	return nil
}

// OpLabel defines a block label.
type OpLabel struct {
	ResultId Id
}

func (c *OpLabel) Opcode() uint32 { return 208 }
func (c *OpLabel) Optional() bool { return false }
func (c *OpLabel) Verify() error  { return nil }

// OpBranch is an unconditional branch to TargetLabel.
type OpBranch struct {
	TargetLabel Id
}

func (c *OpBranch) Opcode() uint32 { return 209 }
func (c *OpBranch) Optional() bool { return false }
func (c *OpBranch) Verify() error  { return nil }

// OpBranchConditional branches to TrueLabel if Condition is true, or to
// False Label if Condition is false.
type OpBranchConditional struct {
	Condition     Id
	TrueLabel     Id
	FalseLabel    Id
	BranchWeights []uint32 `spirv:"optional"`
}

func (c *OpBranchConditional) Opcode() uint32 { return 210 }
func (c *OpBranchConditional) Optional() bool { return false }
func (c *OpBranchConditional) Verify() error {
	if len(c.BranchWeights) != 0 && len(c.BranchWeights) != 2 {
		return fmt.Errorf("OpBranchConditional.BranchWeights: expected 0 or 2 elements")
	}
	return nil
}

// OpSwitch branches to a matching operand label.
type OpSwitch struct {
	Selector Id
	Default  Id
	// TODO: struct for pair (Value, LabelId)
	Target []uint32 `spirv:"optional"`
}

func (c *OpSwitch) Opcode() uint32 { return 211 }
func (c *OpSwitch) Optional() bool { return false }
func (c *OpSwitch) Verify() error {
	if len(c.Target)%2 != 0 {
		return fmt.Errorf("OpSwitch.Target: expected array of (LiteralNumber, Label) pairs")
	}
	for j := 0; j < len(c.Target); j += 2 {
		for k := j + 2; k < len(c.Target); k += 2 {
			if c.Target[j] == c.Target[k] {
				return fmt.Errorf("OpSwitch.Target: literals must be unique")
			}
		}
	}
	return nil
}

// OpKill discards the fragment shader.
type OpKill struct{}

func (c *OpKill) Opcode() uint32 { return 212 }
func (c *OpKill) Optional() bool { return false }
func (c *OpKill) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpKill{}
	})
}

// OpReturn returns with no value from a function with void return type.
type OpReturn struct{}

func (c *OpReturn) Opcode() uint32 { return 213 }
func (c *OpReturn) Optional() bool { return false }
func (c *OpReturn) Verify() error  { return nil }

// OpReturnValue returns with a value from a function.
type OpReturnValue struct {
	Value Id
}

func (c *OpReturnValue) Opcode() uint32 { return 214 }
func (c *OpReturnValue) Optional() bool { return false }
func (c *OpReturnValue) Verify() error  { return nil }

// OpUnreachable declares that this block is not reachable in the Control
// Flow Graph.
type OpUnreachable struct{}

func (c *OpUnreachable) Opcode() uint32 { return 215 }
func (c *OpUnreachable) Optional() bool { return false }
func (c *OpUnreachable) Verify() error  { return nil }

// OpLifetimeStart declares that the content of the object pointed to was
// not defined before this instruction.
type OpLifetimeStart struct {
	Object       Id
	MemoryAmount uint32
}

func (c *OpLifetimeStart) Opcode() uint32 { return 216 }
func (c *OpLifetimeStart) Optional() bool { return false }
func (c *OpLifetimeStart) Verify() error  { return nil }

// OpLifetimeStop declares that the content of the object pointed to is
// dead after this instruction.
type OpLifetimeStop struct {
	Object       Id
	MemoryAmount uint32
}

func (c *OpLifetimeStop) Opcode() uint32 { return 217 }
func (c *OpLifetimeStop) Optional() bool { return false }
func (c *OpLifetimeStop) Verify() error  { return nil }

func init() {
	Bind(func() Instruction { return &OpPhi{} })
	Bind(func() Instruction { return &OpLoopMerge{} })
	Bind(func() Instruction { return &OpSelectionMerge{} })
	Bind(func() Instruction { return &OpLabel{} })
	Bind(func() Instruction { return &OpBranch{} })
	Bind(func() Instruction { return &OpBranchConditional{} })
	Bind(func() Instruction { return &OpSwitch{} })
	Bind(func() Instruction { return &OpReturn{} })
	Bind(func() Instruction { return &OpReturnValue{} })
	Bind(func() Instruction { return &OpUnreachable{} })
	Bind(func() Instruction { return &OpLifetimeStart{} })
	Bind(func() Instruction { return &OpLifetimeStop{} })
}
