// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import (
	"math"
	"testing"
)

func TestTypeDeclaration(t *testing.T) {
	for _, st := range []InstructionTest{
		{
			in:   []uint32{0x00020008, 0x32},
			want: OpTypeVoid(0x32),
		},
		{
			in:   []uint32{0x00020009, 0x32},
			want: OpTypeBool(0x32),
		},
		{
			in: []uint32{0x0004000a, 0x32, 64, 1},
			want: &OpTypeInt{
				ResultId:   0x32,
				Width:      64,
				Signedness: 1,
			},
		},
		{
			in: []uint32{
				0x0003000b, 0x32,
				math.Float32bits(123),
			},
			want: &OpTypeFloat{
				ResultId: 0x32,
				Width:    math.Float32bits(123),
			},
		},
		{
			in: []uint32{0x0004000c, 0x32, 0x12, 0x03},
			want: &OpTypeVector{
				ResultId:       0x32,
				ComponentType:  0x12,
				ComponentCount: 0x03,
			},
		},
		{
			in: []uint32{0x0004000d, 0x32, 0x12, 0x04},
			want: &OpTypeMatrix{
				ResultId:    0x32,
				ColumnType:  0x12,
				ColumnCount: 0x04,
			},
		},
		{
			in: []uint32{0x008000e, 1, 2, Dim3D, 2, 1, 0, 1},
			want: &OpTypeSampler{
				ResultId:       1,
				SampledType:    2,
				Dimensionality: Dim3D,
				Content:        2,
				Arrayed:        1,
				Compare:        0,
				MS:             1,
			},
		},
		{
			in: []uint32{0x009000e, 1, 2, Dim3D, 2, 1, 0, 1, AccessQualifierWriteOnly},
			want: &OpTypeSampler{
				ResultId:        1,
				SampledType:     2,
				Dimensionality:  Dim3D,
				Content:         2,
				Arrayed:         1,
				Compare:         0,
				MS:              1,
				AccessQualifier: AccessQualifierWriteOnly,
			},
		},
		{
			in:   []uint32{0x002000f, 1},
			want: OpTypeFilter(1),
		},
		{
			in: []uint32{0x0040010, 1, 2, 3},
			want: &OpTypeArray{
				ResultId:    1,
				ElementType: 2,
				Length:      3,
			},
		},
		{
			in: []uint32{0x0030011, 1, 2},
			want: &OpTypeRuntimeArray{
				ResultId:    1,
				ElementType: 2,
			},
		},
		{
			in: []uint32{0x0060012, 1, 2, 3, 4, 5},
			want: &OpTypeStruct{
				ResultId: 1,
				Members:  []uint32{2, 3, 4, 5},
			},
		},
		{
			in: []uint32{0x0050013, 1, 0x74736574, 0x72747320, 0x676e69},
			want: &OpTypeOpaque{
				ResultId: 1,
				Name:     "test string",
			},
		},
		{
			in: []uint32{0x0040014, 1, StorageClassAtomicCounter, 2},
			want: &OpTypePointer{
				ResultId:     1,
				StorageClass: StorageClassAtomicCounter,
				Type:         2,
			},
		},
		{
			in: []uint32{0x0030015, 1, 2},
			want: &OpTypeFunction{
				ResultId:   1,
				ReturnType: 2,
			},
		},
		{
			in: []uint32{0x0060015, 1, 2, 3, 4, 5},
			want: &OpTypeFunction{
				ResultId:   1,
				ReturnType: 2,
				Parameters: []uint32{3, 4, 5},
			},
		},
		{
			in:   []uint32{0x0020016, 123},
			want: OpTypeEvent(123),
		},
		{
			in:   []uint32{0x0020017, 4321},
			want: OpTypeDeviceEvent(4321),
		},
		{
			in:   []uint32{0x0020018, 123},
			want: OpTypeReserveId(123),
		},
		{
			in:   []uint32{0x0020019, 123},
			want: OpTypeQueue(123),
		},
		{
			in: []uint32{0x004001a, 1, 2, AccessQualifierReadWrite},
			want: &OpTypePipe{
				ResultId:        1,
				Type:            2,
				AccessQualifier: AccessQualifierReadWrite,
			},
		},
	} {
		testInstruction(t, st)
	}
}
