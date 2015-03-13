// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "fmt"

// OpEnqueueMarker enqueues a marker command to to the queue object specified by Q.
type OpEnqueueMarker struct {
	ResultType uint32
	ResultId   uint32
	Q          uint32
	NumEvents  uint32
	WaitEvents uint32
	RetEvent   uint32
}

func (c *OpEnqueueMarker) Opcode() uint32 { return 249 }
func (c *OpEnqueueMarker) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpEnqueueMarker{}
	})
}

// OpEnqueueKernel enqueues the function specified by Invoke and NDRange
// for execution to the queue object specified by Q.
type OpEnqueueKernel struct {
	ResultType uint32
	ResultId   uint32
	Q          uint32
	Flags      uint32
	NDRange    uint32
	NumEvents  uint32
	WaitEvents uint32
	RetEvent   uint32
	Invoke     uint32
	Param      uint32
	ParamSize  uint32
	ParamAlign uint32
	LocalSize  []uint32 `spirv:"optional"`
}

func (c *OpEnqueueKernel) Opcode() uint32 { return 250 }
func (c *OpEnqueueKernel) Verify() error {
	if (c.Flags &^ (KernelEnqueueFlagNoWait | KernelEnqueueFlagWaitKernel | KernelEnqueueFlagWaitWorkGroup)) != 0 {
		return fmt.Errorf("OpEnqueueKernel.Flags: expected bits within Kernel Enqueue Flags constants")
	}
	return nil
}

func init() {
	Bind(func() Instruction {
		return &OpEnqueueKernel{
			LocalSize: []uint32{},
		}
	})
}

// OpGetKernelNDrangeSubGroupCount returns the number of subgroups in each
// work-group of the dispatch given the combination of NDRange and
// Invoke.
type OpGetKernelNDrangeSubGroupCount struct {
	ResultType uint32
	ResultId   uint32
	NDRange    uint32
	Invoke     uint32
}

func (c *OpGetKernelNDrangeSubGroupCount) Opcode() uint32 { return 251 }
func (c *OpGetKernelNDrangeSubGroupCount) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpGetKernelNDrangeSubGroupCount{}
	})
}

// OpGetKernelNDrangeMaxSubGroupSize returns the maximum sub-group size for
// Invoke and NDRange.
type OpGetKernelNDrangeMaxSubGroupSize struct {
	ResultType uint32
	ResultId   uint32
	NDRange    uint32
	Invoke     uint32
}

func (c *OpGetKernelNDrangeMaxSubGroupSize) Opcode() uint32 { return 252 }
func (c *OpGetKernelNDrangeMaxSubGroupSize) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpGetKernelNDrangeMaxSubGroupSize{}
	})
}

// OpGetKernelWorkGroupSize returns the maximum work-group size that can be
// used to execute Invoke on the device.
type OpGetKernelWorkGroupSize struct {
	ResultType uint32
	ResultId   uint32
	Invoke     uint32
}

func (c *OpGetKernelWorkGroupSize) Opcode() uint32 { return 253 }
func (c *OpGetKernelWorkGroupSize) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpGetKernelWorkGroupSize{}
	})
}

// OpGetKernelPreferredWorkGroupSizeMultiple returns the preferred multiple of work-group size for Invoke.
type OpGetKernelPreferredWorkGroupSizeMultiple struct {
	ResultType uint32
	ResultId   uint32
	Invoke     uint32
}

func (c *OpGetKernelPreferredWorkGroupSizeMultiple) Opcode() uint32 { return 254 }
func (c *OpGetKernelPreferredWorkGroupSizeMultiple) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpGetKernelPreferredWorkGroupSizeMultiple{}
	})
}

// OpRetainEvent increments the reference count of Event.
type OpRetainEvent struct {
	Event uint32
}

func (c *OpRetainEvent) Opcode() uint32 { return 255 }
func (c *OpRetainEvent) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpRetainEvent{}
	})
}

// OpReleaseEvent decrements the reference count of Event.
type OpReleaseEvent struct {
	Event uint32
}

func (c *OpReleaseEvent) Opcode() uint32 { return 256 }
func (c *OpReleaseEvent) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpReleaseEvent{}
	})
}

// OpCreateUserEvent creates a user event.
type OpCreateUserEvent struct {
	ResultType uint32
	ResultId   uint32
}

func (c *OpCreateUserEvent) Opcode() uint32 { return 257 }
func (c *OpCreateUserEvent) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpCreateUserEvent{}
	})
}

// OpIsValidEvent returns true if Event is a valid event, otherwise returns
// false.
type OpIsValidEvent struct {
	ResultType uint32
	ResultId   uint32
	Event      uint32
}

func (c *OpIsValidEvent) Opcode() uint32 { return 258 }
func (c *OpIsValidEvent) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpIsValidEvent{}
	})
}

// OpSetUserEventStatus sets the execution status of a user event.
type OpSetUserEventStatus struct {
	Event  uint32
	Status uint32
}

func (c *OpSetUserEventStatus) Opcode() uint32 { return 259 }
func (c *OpSetUserEventStatus) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpSetUserEventStatus{}
	})
}

// OpCaptureEventProfilingInfo captures the profiling information specified by info for the command associated with the event specified by event in the memory pointed by value.
type OpCaptureEventProfilingInfo struct {
	Event uint32
	Info  uint32
	Value uint32
}

func (c *OpCaptureEventProfilingInfo) Opcode() uint32 { return 260 }
func (c *OpCaptureEventProfilingInfo) Verify() error {
	switch c.Info {
	case KernelProfilingInfoCmdExecTime:
	default:
		fmt.Errorf("OpCaptureEventProfilingInfo.Info: expected a Kernel Profiling Info constant")
	}
	return nil
}

func init() {
	Bind(func() Instruction {
		return &OpCaptureEventProfilingInfo{}
	})
}

// OpGetDefaultQueue returns the default device queue.
type OpGetDefaultQueue struct {
	ResultType uint32
	ResultId   uint32
}

func (c *OpGetDefaultQueue) Opcode() uint32 { return 261 }
func (c *OpGetDefaultQueue) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpGetDefaultQueue{}
	})
}

// OpBuildNDRange BLAHBLAH.
type OpBuildNDRange struct {
	ResultType       uint32
	ResultId         uint32
	GlobalWorkSize   uint32
	LocalWorkSize    uint32
	GlobalWorkOffset uint32
}

func (c *OpBuildNDRange) Opcode() uint32 { return 262 }
func (c *OpBuildNDRange) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpBuildNDRange{}
	})
}
