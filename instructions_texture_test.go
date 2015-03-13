// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "testing"

func TestTexture(t *testing.T) {
	for _, st := range []InstructionTest{
		{
			in: []uint32{0x00050043, 1, 2, 3, 4},
			want: &OpSampler{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				Sampler:    uint32(3),
				Filter:     uint32(4),
			},
		},
		{
			in: []uint32{0x00050044, 1, 2, 3, 4},
			want: &OpTextureSample{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				Sampler:    uint32(3),
				Coordinate: uint32(4),
				Bias:       uint32(0), // optional
			},
		},
		{
			in: []uint32{0x00060044, 1, 2, 3, 4, 5},
			want: &OpTextureSample{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				Sampler:    uint32(3),
				Coordinate: uint32(4),
				Bias:       uint32(5), // optional
			},
		},
		{
			in: []uint32{0x00060045, 1, 2, 3, 4, 5},
			want: &OpTextureSampleDref{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				Sampler:    uint32(3),
				Coordinate: uint32(4),
				Dref:       uint32(5),
			},
		},
		{
			in: []uint32{0x00060046, 1, 2, 3, 4, 5},
			want: &OpTextureSampleLod{
				ResultType:    uint32(1),
				ResultId:      uint32(2),
				Sampler:       uint32(3),
				Coordinate:    uint32(4),
				LevelofDetail: uint32(5),
			},
		},
		{
			in: []uint32{0x00050047, 1, 2, 3, 4},
			want: &OpTextureSampleProj{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				Sampler:    uint32(3),
				Coordinate: uint32(4),
				Bias:       uint32(0), // optional
			},
		},
		{
			in: []uint32{0x00060047, 1, 2, 3, 4, 5},
			want: &OpTextureSampleProj{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				Sampler:    uint32(3),
				Coordinate: uint32(4),
				Bias:       uint32(5), // optional
			},
		},
		{
			in: []uint32{0x00070048, 1, 2, 3, 4, 5, 6},
			want: &OpTextureSampleGrad{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				Sampler:    uint32(3),
				Coordinate: uint32(4),
				Dx:         uint32(5),
				Dy:         uint32(6),
			},
		},
		{
			in: []uint32{0x00060049, 1, 2, 3, 4, 5},
			want: &OpTextureSampleOffset{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				Sampler:    uint32(3),
				Coordinate: uint32(4),
				Offset:     uint32(5),
				Bias:       uint32(0), // optional
			},
		},
		{
			in: []uint32{0x00070049, 1, 2, 3, 4, 5, 6},
			want: &OpTextureSampleOffset{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				Sampler:    uint32(3),
				Coordinate: uint32(4),
				Offset:     uint32(5),
				Bias:       uint32(6), // optional
			},
		},
		{
			in: []uint32{0x0006004a, 1, 2, 3, 4, 5},
			want: &OpTextureSampleProjLod{
				ResultType:    uint32(1),
				ResultId:      uint32(2),
				Sampler:       uint32(3),
				Coordinate:    uint32(4),
				LevelofDetail: uint32(5),
			},
		},
		{
			in: []uint32{0x0007004b, 1, 2, 3, 4, 5, 6},
			want: &OpTextureSampleProjGrad{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				Sampler:    uint32(3),
				Coordinate: uint32(4),
				Dx:         uint32(5),
				Dy:         uint32(6),
			},
		},
		{
			in: []uint32{0x0007004c, 1, 2, 3, 4, 5, 6},
			want: &OpTextureSampleLodOffset{
				ResultType:    uint32(1),
				ResultId:      uint32(2),
				Sampler:       uint32(3),
				Coordinate:    uint32(4),
				LevelofDetail: uint32(5),
				Offset:        uint32(6),
			},
		},
		{
			in: []uint32{0x0006004d, 1, 2, 3, 4, 5},
			want: &OpTextureSampleProjOffset{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				Sampler:    uint32(3),
				Coordinate: uint32(4),
				Offset:     uint32(5),
				Bias:       uint32(0), // optional
			},
		},
		{
			in: []uint32{0x0007004d, 1, 2, 3, 4, 5, 6},
			want: &OpTextureSampleProjOffset{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				Sampler:    uint32(3),
				Coordinate: uint32(4),
				Offset:     uint32(5),
				Bias:       uint32(6), // optional
			},
		},
		{
			in: []uint32{0x0008004e, 1, 2, 3, 4, 5, 6, 7},
			want: &OpTextureSampleGradOffset{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				Sampler:    uint32(3),
				Coordinate: uint32(4),
				Dx:         uint32(5),
				Dy:         uint32(6),
				Offset:     uint32(7),
			},
		},
		{
			in: []uint32{0x0007004f, 1, 2, 3, 4, 5, 6},
			want: &OpTextureSampleProjLodOffset{
				ResultType:    uint32(1),
				ResultId:      uint32(2),
				Sampler:       uint32(3),
				Coordinate:    uint32(4),
				LevelofDetail: uint32(5),
				Offset:        uint32(6),
			},
		},
		{
			in: []uint32{0x00080050, 1, 2, 3, 4, 5, 6, 7},
			want: &OpTextureSampleProjGradOffset{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				Sampler:    uint32(3),
				Coordinate: uint32(4),
				Dx:         uint32(5),
				Dy:         uint32(6),
				Offset:     uint32(7),
			},
		},
		{
			in: []uint32{0x00060051, 1, 2, 3, 4, 5},
			want: &OpTextureFetchTexel{
				ResultType:    uint32(1),
				ResultId:      uint32(2),
				Sampler:       uint32(3),
				Coordinate:    uint32(4),
				LevelofDetail: uint32(5),
			},
		},
		{
			in: []uint32{0x00060052, 1, 2, 3, 4, 5},
			want: &OpTextureFetchTexelOffset{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				Sampler:    uint32(3),
				Coordinate: uint32(4),
				Offset:     uint32(5),
			},
		},
		{
			in: []uint32{0x00060053, 1, 2, 3, 4, 5},
			want: &OpTextureFetchSample{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				Sampler:    uint32(3),
				Coordinate: uint32(4),
				Sample:     uint32(5),
			},
		},
		{
			in: []uint32{0x00050054, 1, 2, 3, 4},
			want: &OpTextureFetchBuffer{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				Sampler:    uint32(3),
				Element:    uint32(4),
			},
		},
		{
			in: []uint32{0x00060055, 1, 2, 3, 4, 5},
			want: &OpTextureGather{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				Sampler:    uint32(3),
				Coordinate: uint32(4),
				Component:  uint32(5),
			},
		},
		{
			in: []uint32{0x00070056, 1, 2, 3, 4, 5, 6},
			want: &OpTextureGatherOffset{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				Sampler:    uint32(3),
				Coordinate: uint32(4),
				Component:  uint32(5),
				Offset:     uint32(6),
			},
		},
		{
			in: []uint32{0x00070057, 1, 2, 3, 4, 5, 6},
			want: &OpTextureGatherOffsets{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				Sampler:    uint32(3),
				Coordinate: uint32(4),
				Component:  uint32(5),
				Offsets:    uint32(6),
			},
		},
		{
			in: []uint32{0x00050058, 1, 2, 3, 4},
			want: &OpTextureQuerySizeLod{
				ResultType:    uint32(1),
				ResultId:      uint32(2),
				Sampler:       uint32(3),
				LevelofDetail: uint32(4),
			},
		},
		{
			in: []uint32{0x00040059, 1, 2, 3},
			want: &OpTextureQuerySize{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				Sampler:    uint32(3),
			},
		},
		{
			in: []uint32{0x0005005a, 1, 2, 3, 4},
			want: &OpTextureQueryLod{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				Sampler:    uint32(3),
				Coordinate: uint32(4),
			},
		},
		{
			in: []uint32{0x0004005b, 1, 2, 3},
			want: &OpTextureQueryLevels{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				Sampler:    uint32(3),
			},
		},
		{
			in: []uint32{0x0004005c, 1, 2, 3},
			want: &OpTextureQuerySamples{
				ResultType: uint32(1),
				ResultId:   uint32(2),
				Sampler:    uint32(3),
			},
		},
	} {
		testInstruction(t, st)
	}
}
