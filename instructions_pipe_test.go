// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "testing"

func TestPipe(t *testing.T) {
	for _, st := range []InstructionTest{
		{
			in: []uint32{0x000500ea, 1, 2, 3, 4},
			want: &OpReadPipe{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				P:          uint32(3),
				Ptr:        uint32(4),
			},
		},
		{
			in: []uint32{0x000500eb, 1, 2, 3, 4},
			want: &OpWritePipe{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				P:          uint32(3),
				Ptr:        uint32(4),
			},
		},
		{
			in: []uint32{0x000700ec, 1, 2, 3, 4, 5, 6},
			want: &OpReservedReadPipe{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				P:          uint32(3),
				ReserveId:  uint32(4),
				Index:      uint32(5),
				Ptr:        uint32(6),
			},
		},
		{
			in: []uint32{0x000700ed, 1, 2, 3, 4, 5, 6},
			want: &OpReservedWritePipe{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				P:          uint32(3),
				ReserveId:  uint32(4),
				Index:      uint32(5),
				Ptr:        uint32(6),
			},
		},
		{
			in: []uint32{0x000500ee, 1, 2, 3, 4},
			want: &OpReserveReadPipePackets{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				P:          uint32(3),
				NumPackets: uint32(4),
			},
		},
		{
			in: []uint32{0x000500ef, 1, 2, 3, 4},
			want: &OpReserveWritePipePackets{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				P:          uint32(3),
				NumPackets: uint32(4),
			},
		},
		{
			in: []uint32{0x000300f0, 1, 2},
			want: &OpCommitReadPipe{
				P:         uint32(1),
				ReserveId: uint32(2),
			},
		},
		{
			in: []uint32{0x000300f1, 1, 2},
			want: &OpCommitWritePipe{
				P:         uint32(1),
				ReserveId: uint32(2),
			},
		},
		{
			in: []uint32{0x000400f2, 1, 2, 3},
			want: &OpIsValidReserveId{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				ReserveId:  uint32(3),
			},
		},
		{
			in: []uint32{0x000400f3, 1, 2, 3},
			want: &OpGetNumPipePackets{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				P:          uint32(3),
			},
		},
		{
			in: []uint32{0x000400f4, 1, 2, 3},
			want: &OpGetMaxPipePackets{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				P:          uint32(3),
			},
		},
		{
			in: []uint32{0x000600f5, 1, 2, 3, 4, 5},
			want: &OpGroupReserveReadPipePackets{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				Scope:      uint32(3),
				P:          uint32(4),
				NumPackets: uint32(5),
			},
		},
		{
			in: []uint32{0x000600f6, 1, 2, 3, 4, 5},
			want: &OpGroupReserveWritePipePackets{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				Scope:      uint32(3),
				P:          uint32(4),
				NumPackets: uint32(5),
			},
		},
		{
			in: []uint32{0x000400f7, 1, 2, 3},
			want: &OpGroupCommitReadPipe{
				Scope:     uint32(1),
				P:         uint32(2),
				ReserveId: uint32(3),
			},
		},
		{
			in: []uint32{0x000400f8, 1, 2, 3},
			want: &OpGroupCommitWritePipe{
				Scope:     uint32(1),
				P:         uint32(2),
				ReserveId: uint32(3),
			},
		},
	} {
		testInstruction(t, st)
	}
}
