// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import (
	"fmt"
	"strings"
)

// AccessQualifier defines the access permissions of OpTypeSampler
// and OpTypePipe object. Used by OpTypePipe.
type AccessQualifier uint32

// Known access qualifiers.
const (
	AccessReadOnly  AccessQualifier = 0 // A read-only object.
	AccessWriteOnly AccessQualifier = 1 // A write-only object.
	AccessReadWrite AccessQualifier = 2 // A readable and writable object.
)

func (e AccessQualifier) String() string {
	switch e {
	case AccessReadOnly:
		return "Read Only"
	case AccessWriteOnly:
		return "Write Only"
	case AccessReadWrite:
		return "Read Write"
	}

	return fmt.Sprintf("AccessQualifier(%d)", uint32(e))
}

// AddressingMode defines an existing addressing mode.
type AddressingMode uint32

// Known addressing modes.
const (
	AddressLogical    AddressingMode = 0
	AddressPhysical32 AddressingMode = 1
	AddressPhysical64 AddressingMode = 2
)

func (am AddressingMode) String() string {
	switch am {
	case AddressLogical:
		return "Logical"
	case AddressPhysical32:
		return "Physical32"
	case AddressPhysical64:
		return "Physical64"
	}

	return fmt.Sprintf("AddressingMode(%d)", uint32(am))
}

// Dim defines the dimensionality of a texture.
//
// Used by OpTypeSampler.
type Dim uint32

// Known execution models.
const (
	Dim1D     Dim = 0
	Dim2D     Dim = 1
	Dim3D     Dim = 2
	DimCube   Dim = 3
	DimRect   Dim = 4
	DimBuffer Dim = 5
)

func (d Dim) String() string {
	switch d {
	case Dim1D:
		return "1D"
	case Dim2D:
		return "2D"
	case Dim3D:
		return "3D"
	case DimCube:
		return "Cube"
	case DimRect:
		return "Rect"
	case DimBuffer:
		return "Buffer"
	}

	return fmt.Sprintf("Dim(%d)", uint32(d))
}

// ExecutionMode defines a mode a module’s stage will execute in.
type ExecutionMode uint32

// Known execution modes.
const (
	// Number of times to invoke the geometry stage for each input primitive
	// received. The default is to run once for each input primitive.
	// If greater than the target-dependent maximum, it will fail to compile.
	// Only valid with the Geometry Execution Model.
	//
	// Arguments:
	//
	//   [0]: Number of invocations.
	//
	ModeInvocations ExecutionMode = 0

	// Requests the tessellation primitive generator to divide edges into a
	// collection of equal-sized segments. Only valid with one of the
	// tessellation Execution Models.
	ModeSpacingEqual ExecutionMode = 1

	// Requests the tessellation primitive generator to divide edges into an
	// even number of equal-length segments plus two additional shorter
	// fractional segments. Only valid with one of the tessellation
	// Execution Models.
	ModeSpacingFractionalEven ExecutionMode = 2

	// Requests the tessellation primitive generator to divide edges into an
	// even number of equal-length segments plus two additional shorter
	// fractional segments. Only valid with one of the tessellation.
	// Execution Models.
	ModeSpacingFractionalOdd ExecutionMode = 3

	// Requests the tessellation primitive generator to generate triangles in
	// clockwise order. Only valid with one of the tessellation Execution Models.
	ModeVertexOrderCw ExecutionMode = 4

	// Requests the tessellation primitive generator to generate triangles in
	// counter-clockwise order. Only valid with one of the tessellation
	// Execution Models.
	ModeVertexOrderCcw ExecutionMode = 5

	// Pixels appear centered on whole-number pixel offsets. E.g., the
	// coordinate (0.5, 0.5) appears to move to (0.0, 0.0). Only valid with
	// the Fragment Execution Model.
	ModePixelCenterInteger ExecutionMode = 6

	// Pixel coordinates appear to originate in the upper left, and increase
	// toward the right and downward. Only valid with the Fragment Execution Model.
	ModeOriginUpperLeft ExecutionMode = 7

	// Fragment tests are to be performed before fragment shader execution.
	// Only valid with the Fragment Execution Model.
	ModeEarlyFragmentTests ExecutionMode = 8

	// Requests the tessellation primitive generator to generate a point for
	// each distinct vertex in the subdivided primitive, rather than to
	// generate lines or triangles. Only valid with one of the tessellation
	// Execution Models.
	ModePointMode ExecutionMode = 9

	// This stage will run in transform feedback-capturing mode and this module
	// is responsible for describing the transform-feedback setup.
	// See the XfbBuffer, Offset, and Stride Decorations.
	ModeXFB ExecutionMode = 10

	// This mode must be declared if this module potentially changes the
	// fragment’s depth. Only valid with the Fragment Execution Model.
	ModeDepthReplacing ExecutionMode = 11

	// TBD: this should probably be removed. Depth testing will always be
	// performed after the shader has executed. Only valid with the Fragment
	// Execution Model.
	ModeDepthAny ExecutionMode = 12

	// External optimizations may assume depth modifications will leave the
	// fragment’s depth as greater than or equal to the fragment’s interpolated
	// depth value (given by the z component of the FragCoord Built-In
	// decorated variable). Only valid with the Fragment Execution Model.
	ModeDepthGreater ExecutionMode = 13

	// External optimizations may assume depth modifications leave the
	// fragment’s depth less than the fragment’s interpolated depth
	// value, (given by the z component of the FragCoord Built-In decorated
	// variable). Only valid with the Fragment Execution Model.
	ModeDepthLess ExecutionMode = 14

	// External optimizations may assume this stage did not modify the
	// fragment’s depth. However, DepthReplacing mode must accurately
	// represent depth modification. Only valid with the Fragment Execution Model.
	ModeDepthUnchanged ExecutionMode = 15

	// Indicates the work-group size in the x, y, and z dimensions. Only valid
	// with the GLCompute or Kernel Execution Models.
	//
	// Arguments:
	//
	//   [0]: x size
	//   [1]: y size
	//   [2]: z size
	//
	ModeLocalSize ExecutionMode = 16

	// A hint to the compiler, which indicates the most likely to be used
	// work-group size in the x, y, and z dimensions. Only valid with the
	// Kernel Execution Model.
	//
	// Arguments:
	//
	//   [0]: x size
	//   [1]: y size
	//   [2]: z size
	//
	ModeLocalSizeHint ExecutionMode = 17

	// Stage input primitive is points. Only valid with the Geometry Execution Model.
	ModeInputPoints ExecutionMode = 18

	// Stage input primitive is lines. Only valid with the Geometry Execution Model.
	ModeInputLines ExecutionMode = 19

	// Stage input primitive is lines adjacency. Only valid with the Geometry
	// Execution Model.
	ModeInputLinesAdjacency ExecutionMode = 20

	// For a geometry stage, input primitive is triangles. For a tessellation
	// stage, requests the tessellation primitive generator to generate
	// triangles. Only valid with the Geometry or one of the tessellation
	// Execution Models.
	ModeInputTriangles ExecutionMode = 21

	// Geometry stage input primitive is triangles adjacency. Only valid with
	// the Geometry Execution Model.
	ModeInputTrianglesAdjacency ExecutionMode = 22

	// Requests the tessellation primitive generator to generate quads.
	// Only valid with one of the tessellation Execution Models.
	ModeInputQuads ExecutionMode = 23

	// Requests the tessellation primitive generator to generate isolines.
	// Only valid with one of the tessellation Execution Models
	ModeInputIsolines ExecutionMode = 24

	// For a geometry stage, the maximum number of vertices the shader will
	// ever emit in a single invocation. For a tessellation-control stage,
	// the number of vertices in the output patch produced by the tessellation
	// control shader, which also specifies the number of times the
	// tessellation control shader is invoked. Only valid with the Geometry
	// or one of the tessellation Execution Models.
	//
	// Arguments:
	//
	//   [0]: Vertex count
	//
	ModeOutputVertices ExecutionMode = 25

	// Stage output primitive is points. Only valid with the Geometry
	// Execution Model.
	ModeOutputPoints ExecutionMode = 26

	// Stage output primitive is line strip. Only valid with the Geometry
	// Execution Model.
	ModeOutputLinestrip ExecutionMode = 27

	// Stage output primitive is triangle strip. Only valid with the
	// Geometry Execution Model.
	ModeOutputTrianglestrip ExecutionMode = 28

	// A hint to the compiler, which indicates that most operations used
	// in the entry point are explicitly vectorized using a particular
	// vector type. Only valid with the Kernel Execution Model.
	//
	// Arguments:
	//
	//   [0]: Vector type
	//
	ModeVecTypeHint ExecutionMode = 29

	// Indicates that floating-point-expressions contraction is disallowed.
	// Only valid with the Kernel Execution Model.
	ModeContractionOff ExecutionMode = 30
)

func (em ExecutionMode) String() string {
	switch em {
	case ModeInvocations:
		return "Invocations"
	case ModeSpacingEqual:
		return "Spacing Equal"
	case ModeSpacingFractionalEven:
		return "Spacing Fractional Even"
	case ModeSpacingFractionalOdd:
		return "Spacing Fractional Odd"
	case ModeVertexOrderCw:
		return "Vertex Order Cw"
	case ModeVertexOrderCcw:
		return "Vertex Order Ccw"
	case ModePixelCenterInteger:
		return "Pixel Center Integer"
	case ModeOriginUpperLeft:
		return "Origin Upper Left"
	case ModeEarlyFragmentTests:
		return "Early Fragment Tests"
	case ModePointMode:
		return "Point Mode"
	case ModeXFB:
		return "XFB"
	case ModeDepthReplacing:
		return "Depth Replacing"
	case ModeDepthAny:
		return "Depth Any"
	case ModeDepthGreater:
		return "Depth Greater"
	case ModeDepthLess:
		return "Depth Less"
	case ModeDepthUnchanged:
		return "Depth Unchanged"
	case ModeLocalSize:
		return "Local Size"
	case ModeLocalSizeHint:
		return "Local Size Hint"
	case ModeInputPoints:
		return "Input: Points"
	case ModeInputLines:
		return "Input: Lines"
	case ModeInputLinesAdjacency:
		return "Input: Lines Adjacency"
	case ModeInputTriangles:
		return "Input: Triangles"
	case ModeInputTrianglesAdjacency:
		return "Input: Triangles Adjacency"
	case ModeInputQuads:
		return "Input: Quads"
	case ModeInputIsolines:
		return "Input: Isolines"
	case ModeOutputVertices:
		return "Output: Vertices"
	case ModeOutputPoints:
		return "Output: Points"
	case ModeOutputLinestrip:
		return "Output: Linestrip"
	case ModeOutputTrianglestrip:
		return "Output: Trianglestrip"
	case ModeVecTypeHint:
		return "Vector type hint"
	case ModeContractionOff:
		return "Contraction Off"
	}

	return fmt.Sprintf("ExecutionMode(%d)", uint32(em))
}

// ExecutionModel defines a single execution model.
// This is used in the EntryPoint instruction to determine what stage of the
// pipeline a given set of instructions belongs to.
type ExecutionModel uint32

// Known execution models.
const (
	ExecVertex                 ExecutionModel = 0 // Vertex shading stage
	ExecTessellationControl    ExecutionModel = 1 // Tessellation control (or hull) shading stage.
	ExecTessellationEvaluation ExecutionModel = 2 // Tessellation evaluation (or domain) shading stage
	ExecGeometry               ExecutionModel = 3 // Geometry shading stage.
	ExecFragment               ExecutionModel = 4 // Fragment shading stage.
	ExecGLCompute              ExecutionModel = 5 // Graphical compute shading stage.
	ExecKernel                 ExecutionModel = 6 // Compute kernel.
)

func (e ExecutionModel) String() string {
	switch e {
	case ExecVertex:
		return "Vertex"
	case ExecTessellationControl:
		return "Tessellation Control"
	case ExecTessellationEvaluation:
		return "Tessellation Evaluation"
	case ExecGeometry:
		return "Geometry"
	case ExecFragment:
		return "Fragment"
	case ExecGLCompute:
		return "GL Compute"
	case ExecKernel:
		return "Kernel"
	}

	return fmt.Sprintf("ExecutionModel(%d)", uint32(e))
}

// FPFastMathMode defines bitflags which enable fast math operations
// which are otherwise unsafe.
//
// Only valid on OpFAdd, OpFSub, OpFMul, OpFDiv, OpFRem
// and OpFMod instructions.
type FPFastMathMode uint32

// Known fast math modes.
const (
	// Assume parameters and result are not NaN.
	FMNotNaN FPFastMathMode = 0

	// Assume parameters and result are not +/- Inf.
	FMNotInf FPFastMathMode = 2

	// Treat the sign of a zero parameter or result as insignificant.
	FMNSZ FPFastMathMode = 4

	// Allow the usage of reciprocal rather than perform a division.
	FMAllowRecip FPFastMathMode = 8

	// Allow algebraic transformations according to real-number associative
	// and distributive algebra. This flag implies all the others.
	FMFast FPFastMathMode = 16
)

func (fm FPFastMathMode) String() string {
	set := make([]string, 0, 5)

	if fm&FMNotNaN != 0 {
		set = append(set, "Not NaN")
	}

	if fm&FMNotInf != 0 {
		set = append(set, "Not Inf")
	}

	if fm&FMNSZ != 0 {
		set = append(set, "Non-Significant Sign")
	}

	if fm&FMAllowRecip != 0 {
		set = append(set, "Allow Reciprocal")
	}

	if fm&FMFast != 0 {
		set = append(set, "Allow Reciprocal")
	}

	if len(set) == 0 {
		return fmt.Sprintf("FPFastMathMode(%d)", uint32(fm))
	}

	return strings.Join(set, ", ")
}

// FPRoundingMode associates a rounding mode to a floating-point
// conversion instruction.
//
// By default:
//
//    - Conversions from floating-point to integer types use the
//      round-toward-zero rounding mode.
//    - Conversions to floating-point types use the round-to-nearest-even
//      rounding mode.
//
type FPRoundingMode uint32

// Known rounding modes.
const (
	RTE FPRoundingMode = 0 // Round to nearest even.
	RTZ FPRoundingMode = 1 // Round towards zero.
	RTP FPRoundingMode = 2 // Round towards positive infinity.
	RTN FPRoundingMode = 3 // Round towards negative infinity.
)

func (r FPRoundingMode) String() string {
	switch r {
	case RTE:
		return "Nearest Even"
	case RTZ:
		return "Zero"
	case RTP:
		return "Positive Infinity"
	case RTN:
		return "Negative Infinity"
	}

	return fmt.Sprintf("FPRoundingMode(%d)", uint32(r))
}

// LinkageType associates a linkage type to functions or global
// variables. By default, functions and global variables are private
// to a module and cannot be accessed by other modules.
type LinkageType uint32

// Known execution models.
const (
	LinkExport LinkageType = 0 // Accessible by other modules as well.
	LinkImport LinkageType = 1 // Declaration for a global identifier that exists in another module.
)

func (e LinkageType) String() string {
	switch e {
	case LinkExport:
		return "Export"
	case LinkImport:
		return "Import"
	}

	return fmt.Sprintf("LinkageType(%d)", uint32(e))
}

// MemoryMode defines an existing memory model.
type MemoryMode uint32

// Known addressing modes.
const (
	MemorySimple   MemoryMode = 0 // No shared memory consistency issues.
	MemoryGLSL450  MemoryMode = 1 // Memory model needed by later versions of GLSL and ESSL. Works across multiple versions.
	MemoryOpenCL12 MemoryMode = 2 // OpenCL 1.2 memory model.
	MemoryOpenCL20 MemoryMode = 3 // OpenCL 2.0 memory model.
	MemoryOpenCL21 MemoryMode = 4 // OpenCL 2.1 memory model.
)

func (mm MemoryMode) String() string {
	switch mm {
	case MemorySimple:
		return "Simple"
	case MemoryGLSL450:
		return "GLSL450"
	case MemoryOpenCL12:
		return "OpenCL1.2"
	case MemoryOpenCL20:
		return "OpenCL2.0"
	case MemoryOpenCL21:
		return "OpenCL2.1"
	}

	return fmt.Sprintf("MemoryMode(%d)", uint32(mm))
}

// SamplerAddressingMode defines the addressing mode of read image
// extended instructions.
type SamplerAddressingMode uint32

// Known addressing modes.
const (
	// The image coordinates used to sample elements of the image refer to a
	// location inside the image, otherwise the results are undefined.
	SAMNone SamplerAddressingMode = 0

	// Out-of-range image coordinates are clamped to the extent.
	SAMClampEdge SamplerAddressingMode = 2

	// Out-of-range image coordinates will return a border color.
	SAMClamp SamplerAddressingMode = 4

	// Out-of-range image coordinates are wrapped to the valid range.
	// Can only be used with normalized coordinates.
	SAMRepeat SamplerAddressingMode = 6

	// Flip the image coordinate at every integer junction.
	// Can only be used with normalized coordinates.
	SAMRepeatMirrored SamplerAddressingMode = 8
)

func (sam SamplerAddressingMode) String() string {
	switch sam {
	case SAMNone:
		return "None"
	case SAMClampEdge:
		return "Clamp: Edge"
	case SAMClamp:
		return "Clamp"
	case SAMRepeat:
		return "Repeat"
	case SAMRepeatMirrored:
		return "Repeat: Mirrored"
	}

	return fmt.Sprintf("SamplerAddressingMode(%d)", uint32(sam))
}

// SamplerAddressingMode defines the filter mode of read image
// extended instructions.
type SamplerFilterMode uint32

// Known addressing modes.
const (
	// Use filter nearset mode when performing a read image operation.
	SFMNearest SamplerFilterMode = 16

	// Use filter linear mode when performing a read image operation.
	SFMLinear SamplerFilterMode = 32
)

func (sfm SamplerFilterMode) String() string {
	switch sfm {
	case SFMNearest:
		return "Nearest"
	case SFMLinear:
		return "Linear"
	}

	return fmt.Sprintf("SamplerFilterMode(%d)", uint32(sfm))
}

// SourceLanguage defines a source language constant.
type SourceLanguage uint32

// Known source languages.
const (
	SourceUnknown SourceLanguage = 0
	SourceESSL    SourceLanguage = 1
	SourceGLSL    SourceLanguage = 2
	SourceOpenCL  SourceLanguage = 3
)

func (sl SourceLanguage) String() string {
	switch sl {
	case SourceUnknown:
		return "Unknown"
	case SourceESSL:
		return "ESSL"
	case SourceGLSL:
		return "GLSL"
	case SourceOpenCL:
		return "OpenCL"
	}
	return fmt.Sprintf("SourceLanguage(%d)", uint32(sl))
}

// StorageClass defines a class of storage for declared variables
// (does not include intermediate values).
//
// Used by: OpTypePointer, OpTypeVariable, OpTypeVariableArray,
// OpTypeGenericCastToPtrExplicit
type StorageClass uint32

// Known storage classes
const (
	// Shared externally, read-only memory, visible across all instantiation
	// or work groups. Graphics uniform memory. OpenCL Constant memory
	StoreUniformConstant StorageClass = 0

	// Input from pipeline. Read only
	StoreInput StorageClass = 1

	// Shared externally, visible across all instantiations or work groups
	StoreUniform StorageClass = 2

	// Output to pipeline.
	StoreOutput StorageClass = 3

	// Shared across all work items within a work group. OpenGL "shared".
	// OpenCL local memory.
	StoreWorkgroupLocal StorageClass = 4

	// Visible across all work items of all work groups. OpenCL global memory.
	StoreWorkgroupGlobal StorageClass = 5

	// Accessible across functions within a module, non-IO (not visible outside
	// the module).
	StorePrivateGlobal StorageClass = 6

	// A variable local to a function.
	StoreFunction StorageClass = 7

	// A generic pointer, which overloads StoragePrivate, StorageLocal,
	// StorageGlobal. not a real storage class.
	StoreGeneric StorageClass = 8

	// Private to a work-item and is not visible to another work-item.
	// OpenCL private memory.
	StorePrivate StorageClass = 9

	// For holding atomic counters.
	StoreAtomicCounter StorageClass = 10
)

func (s StorageClass) String() string {
	switch s {
	case StoreUniformConstant:
		return "Uniform Constant"
	case StoreInput:
		return "Input"
	case StoreUniform:
		return "Uniform"
	case StoreOutput:
		return "Output"
	case StoreWorkgroupLocal:
		return "Workgroup: Local"
	case StoreWorkgroupGlobal:
		return "Workgroup: Global"
	case StorePrivateGlobal:
		return "Private: Global"
	case StoreFunction:
		return "Function"
	case StoreGeneric:
		return "Generic"
	case StorePrivate:
		return "Private"
	case StoreAtomicCounter:
		return "Atomic Counter"
	}

	return fmt.Sprintf("StorageClass(%d)", uint32(s))
}
