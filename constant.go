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
	AQReadOnly  AccessQualifier = 0 // A read-only object.
	AQWriteOnly AccessQualifier = 1 // A write-only object.
	AQReadWrite AccessQualifier = 2 // A readable and writable object.
)

func (e AccessQualifier) String() string {
	switch e {
	case AQReadOnly:
		return "Read Only"
	case AQWriteOnly:
		return "Write Only"
	case AQReadWrite:
		return "Read Write"
	}

	return fmt.Sprintf("AccessQualifier(%d)", uint32(e))
}

// AddressingMode defines an existing addressing mode.
type AddressingMode uint32

// Known addressing modes.
const (
	AMLogical    AddressingMode = 0
	AMPhysical32 AddressingMode = 1
	AMPhysical64 AddressingMode = 2
)

func (am AddressingMode) String() string {
	switch am {
	case AMLogical:
		return "Logical"
	case AMPhysical32:
		return "Physical32"
	case AMPhysical64:
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
	D1D     Dim = 0
	D2D     Dim = 1
	D3D     Dim = 2
	DCube   Dim = 3
	DRect   Dim = 4
	DBuffer Dim = 5
)

func (d Dim) String() string {
	switch d {
	case D1D:
		return "1D"
	case D2D:
		return "2D"
	case D3D:
		return "3D"
	case DCube:
		return "Cube"
	case DRect:
		return "Rect"
	case DBuffer:
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
	EMInvocations ExecutionMode = 0

	// Requests the tessellation primitive generator to divide edges into a
	// collection of equal-sized segments. Only valid with one of the
	// tessellation Execution Models.
	EMSpacingEqual ExecutionMode = 1

	// Requests the tessellation primitive generator to divide edges into an
	// even number of equal-length segments plus two additional shorter
	// fractional segments. Only valid with one of the tessellation
	// Execution Models.
	EMSpacingFractionalEven ExecutionMode = 2

	// Requests the tessellation primitive generator to divide edges into an
	// even number of equal-length segments plus two additional shorter
	// fractional segments. Only valid with one of the tessellation.
	// Execution Models.
	EMSpacingFractionalOdd ExecutionMode = 3

	// Requests the tessellation primitive generator to generate triangles in
	// clockwise order. Only valid with one of the tessellation Execution Models.
	EMVertexOrderCw ExecutionMode = 4

	// Requests the tessellation primitive generator to generate triangles in
	// counter-clockwise order. Only valid with one of the tessellation
	// Execution Models.
	EMVertexOrderCcw ExecutionMode = 5

	// Pixels appear centered on whole-number pixel offsets. E.g., the
	// coordinate (0.5, 0.5) appears to move to (0.0, 0.0). Only valid with
	// the Fragment Execution Model.
	EMPixelCenterInteger ExecutionMode = 6

	// Pixel coordinates appear to originate in the upper left, and increase
	// toward the right and downward. Only valid with the Fragment Execution Model.
	EMOriginUpperLeft ExecutionMode = 7

	// Fragment tests are to be performed before fragment shader execution.
	// Only valid with the Fragment Execution Model.
	EMEarlyFragmentTests ExecutionMode = 8

	// Requests the tessellation primitive generator to generate a point for
	// each distinct vertex in the subdivided primitive, rather than to
	// generate lines or triangles. Only valid with one of the tessellation
	// Execution Models.
	EMPointMode ExecutionMode = 9

	// This stage will run in transform feedback-capturing mode and this module
	// is responsible for describing the transform-feedback setup.
	// See the XfbBuffer, Offset, and Stride Decorations.
	EMXFB ExecutionMode = 10

	// This mode must be declared if this module potentially changes the
	// fragment’s depth. Only valid with the Fragment Execution Model.
	EMDepthReplacing ExecutionMode = 11

	// TBD: this should probably be removed. Depth testing will always be
	// performed after the shader has executed. Only valid with the Fragment
	// Execution Model.
	EMDepthAny ExecutionMode = 12

	// External optimizations may assume depth modifications will leave the
	// fragment’s depth as greater than or equal to the fragment’s interpolated
	// depth value (given by the z component of the FragCoord Built-In
	// decorated variable). Only valid with the Fragment Execution Model.
	EMDepthGreater ExecutionMode = 13

	// External optimizations may assume depth modifications leave the
	// fragment’s depth less than the fragment’s interpolated depth
	// value, (given by the z component of the FragCoord Built-In decorated
	// variable). Only valid with the Fragment Execution Model.
	EMDepthLess ExecutionMode = 14

	// External optimizations may assume this stage did not modify the
	// fragment’s depth. However, DepthReplacing mode must accurately
	// represent depth modification. Only valid with the Fragment Execution Model.
	EMDepthUnchanged ExecutionMode = 15

	// Indicates the work-group size in the x, y, and z dimensions. Only valid
	// with the GLCompute or Kernel Execution Models.
	//
	// Arguments:
	//
	//   [0]: x size
	//   [1]: y size
	//   [2]: z size
	//
	EMLocalSize ExecutionMode = 16

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
	EMLocalSizeHint ExecutionMode = 17

	// Stage input primitive is points. Only valid with the Geometry Execution Model.
	EMInputPoints ExecutionMode = 18

	// Stage input primitive is lines. Only valid with the Geometry Execution Model.
	EMInputLines ExecutionMode = 19

	// Stage input primitive is lines adjacency. Only valid with the Geometry
	// Execution Model.
	EMInputLinesAdjacency ExecutionMode = 20

	// For a geometry stage, input primitive is triangles. For a tessellation
	// stage, requests the tessellation primitive generator to generate
	// triangles. Only valid with the Geometry or one of the tessellation
	// Execution Models.
	EMInputTriangles ExecutionMode = 21

	// Geometry stage input primitive is triangles adjacency. Only valid with
	// the Geometry Execution Model.
	EMInputTrianglesAdjacency ExecutionMode = 22

	// Requests the tessellation primitive generator to generate quads.
	// Only valid with one of the tessellation Execution Models.
	EMInputQuads ExecutionMode = 23

	// Requests the tessellation primitive generator to generate isolines.
	// Only valid with one of the tessellation Execution Models
	EMInputIsolines ExecutionMode = 24

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
	EMOutputVertices ExecutionMode = 25

	// Stage output primitive is points. Only valid with the Geometry
	// Execution Model.
	EMOutputPoints ExecutionMode = 26

	// Stage output primitive is line strip. Only valid with the Geometry
	// Execution Model.
	EMOutputLinestrip ExecutionMode = 27

	// Stage output primitive is triangle strip. Only valid with the
	// Geometry Execution Model.
	EMOutputTrianglestrip ExecutionMode = 28

	// A hint to the compiler, which indicates that most operations used
	// in the entry point are explicitly vectorized using a particular
	// vector type. Only valid with the Kernel Execution Model.
	//
	// Arguments:
	//
	//   [0]: Vector type
	//
	EMVecTypeHint ExecutionMode = 29

	// Indicates that floating-point-expressions contraction is disallowed.
	// Only valid with the Kernel Execution Model.
	EMContractionOff ExecutionMode = 30
)

func (em ExecutionMode) String() string {
	switch em {
	case EMInvocations:
		return "Invocations"
	case EMSpacingEqual:
		return "Spacing Equal"
	case EMSpacingFractionalEven:
		return "Spacing Fractional Even"
	case EMSpacingFractionalOdd:
		return "Spacing Fractional Odd"
	case EMVertexOrderCw:
		return "Vertex Order Cw"
	case EMVertexOrderCcw:
		return "Vertex Order Ccw"
	case EMPixelCenterInteger:
		return "Pixel Center Integer"
	case EMOriginUpperLeft:
		return "Origin Upper Left"
	case EMEarlyFragmentTests:
		return "Early Fragment Tests"
	case EMPointMode:
		return "Point Mode"
	case EMXFB:
		return "XFB"
	case EMDepthReplacing:
		return "Depth Replacing"
	case EMDepthAny:
		return "Depth Any"
	case EMDepthGreater:
		return "Depth Greater"
	case EMDepthLess:
		return "Depth Less"
	case EMDepthUnchanged:
		return "Depth Unchanged"
	case EMLocalSize:
		return "Local Size"
	case EMLocalSizeHint:
		return "Local Size Hint"
	case EMInputPoints:
		return "Input: Points"
	case EMInputLines:
		return "Input: Lines"
	case EMInputLinesAdjacency:
		return "Input: Lines Adjacency"
	case EMInputTriangles:
		return "Input: Triangles"
	case EMInputTrianglesAdjacency:
		return "Input: Triangles Adjacency"
	case EMInputQuads:
		return "Input: Quads"
	case EMInputIsolines:
		return "Input: Isolines"
	case EMOutputVertices:
		return "Output: Vertices"
	case EMOutputPoints:
		return "Output: Points"
	case EMOutputLinestrip:
		return "Output: Linestrip"
	case EMOutputTrianglestrip:
		return "Output: Trianglestrip"
	case EMVecTypeHint:
		return "Vector type hint"
	case EMContractionOff:
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
	EMVertex                 ExecutionModel = 0 // Vertex shading stage
	EMTessellationControl    ExecutionModel = 1 // Tessellation control (or hull) shading stage.
	EMTessellationEvaluation ExecutionModel = 2 // Tessellation evaluation (or domain) shading stage
	EMGeometry               ExecutionModel = 3 // Geometry shading stage.
	EMFragment               ExecutionModel = 4 // Fragment shading stage.
	EMGLCompute              ExecutionModel = 5 // Graphical compute shading stage.
	EMKernel                 ExecutionModel = 6 // Compute kernel.
)

func (e ExecutionModel) String() string {
	switch e {
	case EMVertex:
		return "Vertex"
	case EMTessellationControl:
		return "Tessellation Control"
	case EMTessellationEvaluation:
		return "Tessellation Evaluation"
	case EMGeometry:
		return "Geometry"
	case EMFragment:
		return "Fragment"
	case EMGLCompute:
		return "GL Compute"
	case EMKernel:
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
	FMMNotNaN FPFastMathMode = 0

	// Assume parameters and result are not +/- Inf.
	FMMNotInf FPFastMathMode = 2

	// Treat the sign of a zero parameter or result as insignificant.
	FMMNSZ FPFastMathMode = 4

	// Allow the usage of reciprocal rather than perform a division.
	FMMAllowRecip FPFastMathMode = 8

	// Allow algebraic transformations according to real-number associative
	// and distributive algebra. This flag implies all the others.
	FMMFast FPFastMathMode = 16
)

func (fm FPFastMathMode) String() string {
	set := make([]string, 0, 5)

	if fm&FMMNotNaN != 0 {
		set = append(set, "Not NaN")
	}

	if fm&FMMNotInf != 0 {
		set = append(set, "Not Inf")
	}

	if fm&FMMNSZ != 0 {
		set = append(set, "Non-Significant Sign")
	}

	if fm&FMMAllowRecip != 0 {
		set = append(set, "Allow Reciprocal")
	}

	if fm&FMMFast != 0 {
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
	FPRMRTE FPRoundingMode = 0 // Round to nearest even.
	FPRMRTZ FPRoundingMode = 1 // Round towards zero.
	FPRMRTP FPRoundingMode = 2 // Round towards positive infinity.
	FPRMRTN FPRoundingMode = 3 // Round towards negative infinity.
)

func (r FPRoundingMode) String() string {
	switch r {
	case FPRMRTE:
		return "Nearest Even"
	case FPRMRTZ:
		return "Zero"
	case FPRMRTP:
		return "Positive Infinity"
	case FPRMRTN:
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
	LTExport LinkageType = 0 // Accessible by other modules as well.
	LTImport LinkageType = 1 // Declaration for a global identifier that exists in another module.
)

func (e LinkageType) String() string {
	switch e {
	case LTExport:
		return "Export"
	case LTImport:
		return "Import"
	}

	return fmt.Sprintf("LinkageType(%d)", uint32(e))
}

// MemoryMode defines an existing memory model.
type MemoryMode uint32

// Known addressing modes.
const (
	MMSimple   MemoryMode = 0 // No shared memory consistency issues.
	MMGLSL450  MemoryMode = 1 // Memory model needed by later versions of GLSL and ESSL. Works across multiple versions.
	MMOpenCL12 MemoryMode = 2 // OpenCL 1.2 memory model.
	MMOpenCL20 MemoryMode = 3 // OpenCL 2.0 memory model.
	MMOpenCL21 MemoryMode = 4 // OpenCL 2.1 memory model.
)

func (mm MemoryMode) String() string {
	switch mm {
	case MMSimple:
		return "Simple"
	case MMGLSL450:
		return "GLSL450"
	case MMOpenCL12:
		return "OpenCL1.2"
	case MMOpenCL20:
		return "OpenCL2.0"
	case MMOpenCL21:
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
	SLUnknown SourceLanguage = 0
	SLESSL    SourceLanguage = 1
	SLGLSL    SourceLanguage = 2
	SLOpenCL  SourceLanguage = 3
)

func (sl SourceLanguage) String() string {
	switch sl {
	case SLUnknown:
		return "Unknown"
	case SLESSL:
		return "ESSL"
	case SLGLSL:
		return "GLSL"
	case SLOpenCL:
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
	SCUniformConstant StorageClass = 0

	// Input from pipeline. Read only
	SCInput StorageClass = 1

	// Shared externally, visible across all instantiations or work groups
	SCUniform StorageClass = 2

	// Output to pipeline.
	SCOutput StorageClass = 3

	// Shared across all work items within a work group. OpenGL "shared".
	// OpenCL local memory.
	SCWorkgroupLocal StorageClass = 4

	// Visible across all work items of all work groups. OpenCL global memory.
	SCWorkgroupGlobal StorageClass = 5

	// Accessible across functions within a module, non-IO (not visible outside
	// the module).
	SCPrivateGlobal StorageClass = 6

	// A variable local to a function.
	SCFunction StorageClass = 7

	// A generic pointer, which overloads StoragePrivate, StorageLocal,
	// StorageGlobal. not a real storage class.
	SCGeneric StorageClass = 8

	// Private to a work-item and is not visible to another work-item.
	// OpenCL private memory.
	SCPrivate StorageClass = 9

	// For holding atomic counters.
	SCAtomicCounter StorageClass = 10
)

func (s StorageClass) String() string {
	switch s {
	case SCUniformConstant:
		return "Uniform Constant"
	case SCInput:
		return "Input"
	case SCUniform:
		return "Uniform"
	case SCOutput:
		return "Output"
	case SCWorkgroupLocal:
		return "Workgroup: Local"
	case SCWorkgroupGlobal:
		return "Workgroup: Global"
	case SCPrivateGlobal:
		return "Private: Global"
	case SCFunction:
		return "Function"
	case SCGeneric:
		return "Generic"
	case SCPrivate:
		return "Private"
	case SCAtomicCounter:
		return "Atomic Counter"
	}

	return fmt.Sprintf("StorageClass(%d)", uint32(s))
}

// FunctionParamAttr adds additional information to the return type
// and to each parameter of a function.
type FunctionParamAttr uint32

// Known function parameter attributes.
const (
	// Value should be zero extended if needed.
	FPAZext FunctionParamAttr = 0

	// Value should be sign extended if needed.
	FPASext FunctionParamAttr = 1

	// This indicates that the pointer parameter should really be passed by
	// value to the function. Only valid for pointer parameters (not
	// for ret value)
	FPAByVal FunctionParamAttr = 2

	// Indicates that the pointer parameter specifies the address of a
	// structure that is the return value of the function in the source
	// program. Only applicable to the first parameter which must be a
	// pointer parameters.
	FPASret FunctionParamAttr = 3

	// Indicates that the memory pointed by a pointer parameter is not
	// accessed via pointer values which are not derived from this pointer
	// parameter. Only valid for pointer parameters. Not valid on return values
	FPANoAlias FunctionParamAttr = 4

	// The callee does not make a copy of the pointer parameter into a
	// location that is accessible after returning from the callee. Only
	// valid for pointer parameters. Not valid on return values.
	FPANoCapture FunctionParamAttr = 5

	// To be determined.
	FPASVM FunctionParamAttr = 6

	// Can only read the memory pointed by a pointer parameter.
	// Only valid for pointer parameters. Not valid on return values.
	FPANoWrite = 7

	// Cannot dereference the memory pointed by a pointer parameter.
	// Only valid for pointer parameters. Not valid on return values.
	FPANoReadWrite = 8
)

func (fpa FunctionParamAttr) String() string {
	switch fpa {
	case FPAZext:
		return "Zero Extend"
	case FPASext:
		return "Sign Extend"
	case FPAByVal:
		return "By Value"
	case FPASret:
		return "Struct Address"
	case FPANoAlias:
		return "No Alias"
	case FPANoCapture:
		return "No Capture"
	case FPASVM:
		return "SVM"
	case FPANoWrite:
		return "No Write"
	case FPANoReadWrite:
		return "No Read/Write"
	}

	return fmt.Sprintf("FunctionParamAttr(%d)", uint32(fpa))
}
