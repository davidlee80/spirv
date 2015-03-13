// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "testing"

func TestDeviceSideEnqueue(t *testing.T) {
	for _, st := range []InstructionTest{
		{
			in: []uint32{0x000700f9, 1, 2, 3, 4, 5, 6},
			want: &OpEnqueueMarker{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				Q:          uint32(3),
				NumEvents:  uint32(4),
				WaitEvents: uint32(5),
				RetEvent:   uint32(6),
			},
		},
		{
			in: []uint32{0x000d00fa, 1, 2, 3, KernelEnqueueFlagNoWait, 5, 6, 7, 8, 9, 10, 11, 12},
			want: &OpEnqueueKernel{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				Q:          uint32(3),
				Flags:      KernelEnqueueFlagNoWait,
				NDRange:    uint32(5),
				NumEvents:  uint32(6),
				WaitEvents: uint32(7),
				RetEvent:   uint32(8),
				Invoke:     uint32(9),
				Param:      uint32(10),
				ParamSize:  uint32(11),
				ParamAlign: uint32(12),
				LocalSize:  []uint32{}, // optional
			},
		},
		{
			in: []uint32{0x000e00fa, 1, 2, 3, KernelEnqueueFlagWaitWorkGroup, 5, 6, 7, 8, 9, 10, 11, 12, 13},
			want: &OpEnqueueKernel{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				Q:          uint32(3),
				Flags:      KernelEnqueueFlagWaitWorkGroup,
				NDRange:    uint32(5),
				NumEvents:  uint32(6),
				WaitEvents: uint32(7),
				RetEvent:   uint32(8),
				Invoke:     uint32(9),
				Param:      uint32(10),
				ParamSize:  uint32(11),
				ParamAlign: uint32(12),
				LocalSize:  []uint32{13}, // optional
			},
		},
		{
			in: []uint32{0x000500fb, 1, 2, 3, 4},
			want: &OpGetKernelNDrangeSubGroupCount{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				NDRange:    uint32(3),
				Invoke:     uint32(4),
			},
		},
		{
			in: []uint32{0x000500fc, 1, 2, 3, 4},
			want: &OpGetKernelNDrangeMaxSubGroupSize{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				NDRange:    uint32(3),
				Invoke:     uint32(4),
			},
		},
		{
			in: []uint32{0x000400fd, 1, 2, 3},
			want: &OpGetKernelWorkGroupSize{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				Invoke:     uint32(3),
			},
		},
		{
			in: []uint32{0x000400fe, 1, 2, 3},
			want: &OpGetKernelPreferredWorkGroupSizeMultiple{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				Invoke:     uint32(3),
			},
		},
		{
			in: []uint32{0x000200ff, 1},
			want: &OpRetainEvent{
				Event: uint32(1),
			},
		},
		{
			in: []uint32{0x00020100, 1},
			want: &OpReleaseEvent{
				Event: uint32(1),
			},
		},
		{
			in: []uint32{0x00030101, 1, 2},
			want: &OpCreateUserEvent{
				ResultType: uint32(1),
				ResultId:   uint32(2),
			},
		},
		{
			in: []uint32{0x00040102, 1, 2, 3},
			want: &OpIsValidEvent{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				Event:      uint32(3),
			},
		},
		{
			in: []uint32{0x00030103, 1, 2},
			want: &OpSetUserEventStatus{
				Event:  uint32(1),
				Status: uint32(2),
			},
		},
		{
			in: []uint32{0x00040104, 1, 2, 3},
			want: &OpCaptureEventProfilingInfo{
				Event: uint32(1),
				Info:  uint32(2),
				Value: uint32(3),
			},
		},
		{
			in: []uint32{0x00030105, 1, 2},
			want: &OpGetDefaultQueue{
				ResultType: uint32(1),
				ResultId:   uint32(2),
			},
		},
		{
			in: []uint32{0x00060106, 1, 2, 3, 4, 5},
			want: &OpBuildNDRange{
				ResultType:       uint32(1),
				ResultId:         uint32(2),
				GlobalWorkSize:   uint32(3),
				LocalWorkSize:    uint32(4),
				GlobalWorkOffset: uint32(5),
			},
		},
	} {
		testInstruction(t, st)
	}
}
