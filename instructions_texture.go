// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpSampler creates a sampler containing both a filter and texture.
type OpSampler struct {
	ResultType uint32
	ResultId   uint32
	Sampler    uint32
	Filter     uint32
}

func (c *OpSampler) Opcode() uint32 { return 67 }
func (c *OpSampler) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpSampler{}
	})
}

// OpTextureSample samples a texture with an implicit level of detail.
type OpTextureSample struct {
	ResultType uint32
	ResultId   uint32
	Sampler    uint32
	Coordinate uint32
	Bias       uint32 `spirv:"optional"`
}

func (c *OpTextureSample) Opcode() uint32 { return 68 }
func (c *OpTextureSample) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTextureSample{
			Bias: 0,
		}
	})
}

// OpTextureSampleDref samples a cube-map-array texture with depth
// comparison using an implicit level of detail.
type OpTextureSampleDref struct {
	ResultType uint32
	ResultId   uint32
	Sampler    uint32
	Coordinate uint32
	Dref       uint32
}

func (c *OpTextureSampleDref) Opcode() uint32 { return 69 }
func (c *OpTextureSampleDref) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTextureSampleDref{}
	})
}

// OpTextureSampleLod samples a texture using an explicit level of detail.
type OpTextureSampleLod struct {
	ResultType    uint32
	ResultId      uint32
	Sampler       uint32
	Coordinate    uint32
	LevelofDetail uint32
}

func (c *OpTextureSampleLod) Opcode() uint32 { return 70 }
func (c *OpTextureSampleLod) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTextureSampleLod{}
	})
}

// OpTextureSampleProj sample a texture with a projective coordinate
// using an implicit level of detail.
type OpTextureSampleProj struct {
	ResultType uint32
	ResultId   uint32
	Sampler    uint32
	Coordinate uint32
	Bias       uint32 `spirv:"optional"`
}

func (c *OpTextureSampleProj) Opcode() uint32 { return 71 }
func (c *OpTextureSampleProj) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTextureSampleProj{
			Bias: 0,
		}
	})
}

// OpTextureSampleGrad samples a texture with an explicit gradient.
type OpTextureSampleGrad struct {
	ResultType uint32
	ResultId   uint32
	Sampler    uint32
	Coordinate uint32
	Dx         uint32
	Dy         uint32
}

func (c *OpTextureSampleGrad) Opcode() uint32 { return 72 }
func (c *OpTextureSampleGrad) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTextureSampleGrad{}
	})
}

// OpTextureSampleOffset samples a texture with an offset from a
// coordinate using an implicit level of detail.
type OpTextureSampleOffset struct {
	ResultType uint32
	ResultId   uint32
	Sampler    uint32
	Coordinate uint32
	Offset     uint32
	Bias       uint32 `spirv:"optional"`
}

func (c *OpTextureSampleOffset) Opcode() uint32 { return 73 }
func (c *OpTextureSampleOffset) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTextureSampleOffset{
			Bias: 0,
		}
	})
}

// OpTextureSampleProjLod samples a texture with a projective
// coordinate using an explicit level of detail.
type OpTextureSampleProjLod struct {
	ResultType    uint32
	ResultId      uint32
	Sampler       uint32
	Coordinate    uint32
	LevelofDetail uint32
}

func (c *OpTextureSampleProjLod) Opcode() uint32 { return 74 }
func (c *OpTextureSampleProjLod) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTextureSampleProjLod{}
	})
}

// OpTextureSampleProjGrad sample a texture with a projective
// coordinate using an explicit gradient.
type OpTextureSampleProjGrad struct {
	ResultType uint32
	ResultId   uint32
	Sampler    uint32
	Coordinate uint32
	Dx         uint32
	Dy         uint32
}

func (c *OpTextureSampleProjGrad) Opcode() uint32 { return 75 }
func (c *OpTextureSampleProjGrad) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTextureSampleProjGrad{}
	})
}

// OpTextureSampleLodOffset samples a texture with explicit level of
// detail using an offset from a coordinate.
type OpTextureSampleLodOffset struct {
	ResultType    uint32
	ResultId      uint32
	Sampler       uint32
	Coordinate    uint32
	LevelofDetail uint32
	Offset        uint32
}

func (c *OpTextureSampleLodOffset) Opcode() uint32 { return 76 }
func (c *OpTextureSampleLodOffset) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTextureSampleLodOffset{}
	})
}

// OpTextureSampleProjOffset samples a texture with an offset from a
// projective coordinate using an implicit level of detail.
type OpTextureSampleProjOffset struct {
	ResultType uint32
	ResultId   uint32
	Sampler    uint32
	Coordinate uint32
	Offset     uint32
	Bias       uint32 `spirv:"optional"`
}

func (c *OpTextureSampleProjOffset) Opcode() uint32 { return 77 }
func (c *OpTextureSampleProjOffset) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTextureSampleProjOffset{
			Bias: 0,
		}
	})
}

// OpTextureSampleGradOffset samples a texture with an offset
// coordinate and an explicit gradient.
type OpTextureSampleGradOffset struct {
	ResultType uint32
	ResultId   uint32
	Sampler    uint32
	Coordinate uint32
	Dx         uint32
	Dy         uint32
	Offset     uint32
}

func (c *OpTextureSampleGradOffset) Opcode() uint32 { return 78 }
func (c *OpTextureSampleGradOffset) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTextureSampleGradOffset{}
	})
}

// OpTextureSampleProjLodOffset samples a texture with an offset from
// a projective coordinate and an explicit level of detail.
type OpTextureSampleProjLodOffset struct {
	ResultType    uint32
	ResultId      uint32
	Sampler       uint32
	Coordinate    uint32
	LevelofDetail uint32
	Offset        uint32
}

func (c *OpTextureSampleProjLodOffset) Opcode() uint32 { return 79 }
func (c *OpTextureSampleProjLodOffset) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTextureSampleProjLodOffset{}
	})
}

// OpTextureSampleProjGradOffset samples a texture with an offset from
// a projective coordinate and an explicit gradient.
type OpTextureSampleProjGradOffset struct {
	ResultType uint32
	ResultId   uint32
	Sampler    uint32
	Coordinate uint32
	Dx         uint32
	Dy         uint32
	Offset     uint32
}

func (c *OpTextureSampleProjGradOffset) Opcode() uint32 { return 80 }
func (c *OpTextureSampleProjGradOffset) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTextureSampleProjGradOffset{}
	})
}

// OpTextureFetchTexel fetches a single texel from a texture.
type OpTextureFetchTexel struct {
	ResultType    uint32
	ResultId      uint32
	Sampler       uint32
	Coordinate    uint32
	LevelofDetail uint32
}

func (c *OpTextureFetchTexel) Opcode() uint32 { return 81 }
func (c *OpTextureFetchTexel) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTextureFetchTexel{}
	})
}

// OpTextureFetchTexelOffset fetches a single offset texel from a texture.
type OpTextureFetchTexelOffset struct {
	ResultType uint32
	ResultId   uint32
	Sampler    uint32
	Coordinate uint32
	Offset     uint32
}

func (c *OpTextureFetchTexelOffset) Opcode() uint32 { return 82 }
func (c *OpTextureFetchTexelOffset) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTextureFetchTexelOffset{}
	})
}

// OpTextureFetchSample fetches a single sample from a multi-sample texture.
type OpTextureFetchSample struct {
	ResultType uint32
	ResultId   uint32
	Sampler    uint32
	Coordinate uint32
	Sample     uint32
}

func (c *OpTextureFetchSample) Opcode() uint32 { return 83 }
func (c *OpTextureFetchSample) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTextureFetchSample{}
	})
}

// OpTextureFetchBuffer fetches an element out of a buffer texture.
type OpTextureFetchBuffer struct {
	ResultType uint32
	ResultId   uint32
	Sampler    uint32
	Element    uint32
}

func (c *OpTextureFetchBuffer) Opcode() uint32 { return 84 }
func (c *OpTextureFetchBuffer) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTextureFetchBuffer{}
	})
}

// OpTextureGather gathers the requested component from four sampled texels.
type OpTextureGather struct {
	ResultType uint32
	ResultId   uint32
	Sampler    uint32
	Coordinate uint32
	Component  uint32
}

func (c *OpTextureGather) Opcode() uint32 { return 85 }
func (c *OpTextureGather) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTextureGather{}
	})
}

// OpTextureGatherOffset gathers the requested component from four
// offset sampled texels.
type OpTextureGatherOffset struct {
	ResultType uint32
	ResultId   uint32
	Sampler    uint32
	Coordinate uint32
	Component  uint32
	Offset     uint32
}

func (c *OpTextureGatherOffset) Opcode() uint32 { return 86 }
func (c *OpTextureGatherOffset) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTextureGatherOffset{}
	})
}

// OpTextureGatherOffsets gathers the requested component from four
// offset sampled texels.
type OpTextureGatherOffsets struct {
	ResultType uint32
	ResultId   uint32
	Sampler    uint32
	Coordinate uint32
	Component  uint32
	Offsets    uint32
}

func (c *OpTextureGatherOffsets) Opcode() uint32 { return 87 }
func (c *OpTextureGatherOffsets) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTextureGatherOffsets{}
	})
}

// OpTextureQuerySizeLod queries the dimensions of the texture for
// Sampler for mipmap level for Level of Detail.
type OpTextureQuerySizeLod struct {
	ResultType    uint32
	ResultId      uint32
	Sampler       uint32
	LevelofDetail uint32
}

func (c *OpTextureQuerySizeLod) Opcode() uint32 { return 88 }
func (c *OpTextureQuerySizeLod) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTextureQuerySizeLod{}
	})
}

// OpTextureQuerySize queries the dimensions of the texture for
// Sampler, with no level of detail.
type OpTextureQuerySize struct {
	ResultType uint32
	ResultId   uint32
	Sampler    uint32
}

func (c *OpTextureQuerySize) Opcode() uint32 { return 89 }
func (c *OpTextureQuerySize) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTextureQuerySize{}
	})
}

// OpTextureQueryLod queries the mipmap level and the level of detail
// for a hypothetical sampling of Sampler at Coordinate using an
// implicit level of detail.
type OpTextureQueryLod struct {
	ResultType uint32
	ResultId   uint32
	Sampler    uint32
	Coordinate uint32
}

func (c *OpTextureQueryLod) Opcode() uint32 { return 90 }
func (c *OpTextureQueryLod) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTextureQueryLod{}
	})
}

// OpTextureQueryLevels queries the number of mipmap levels accessible
// through Sampler.
type OpTextureQueryLevels struct {
	ResultType uint32
	ResultId   uint32
	Sampler    uint32
}

func (c *OpTextureQueryLevels) Opcode() uint32 { return 91 }
func (c *OpTextureQueryLevels) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTextureQueryLevels{}
	})
}

// OpTextureQuerySamples queries the number of samples available per
// texel fetch in a multisample texture.
type OpTextureQuerySamples struct {
	ResultType uint32
	ResultId   uint32
	Sampler    uint32
}

func (c *OpTextureQuerySamples) Opcode() uint32 { return 92 }
func (c *OpTextureQuerySamples) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpTextureQuerySamples{}
	})
}
