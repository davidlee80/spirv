// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// AccessQualifier defines the access permissions of OpTypeSampler
// and OpTypePipe object. Used by OpTypePipe.
type AccessQualifier uint32

// Known access qualifiers.
const (
	AQReadOnly  AccessQualifier = 0 // A read-only object.
	AQWriteOnly AccessQualifier = 1 // A write-only object.
	AQReadWrite AccessQualifier = 2 // A readable and writable object.
)

// AddressingMode defines an existing addressing mode.
type AddressingMode uint32

// Known addressing modes.
const (
	AMLogical    AddressingMode = 0
	AMPhysical32 AddressingMode = 1
	AMPhysical64 AddressingMode = 2
)

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

// LinkageType associates a linkage type to functions or global
// variables. By default, functions and global variables are private
// to a module and cannot be accessed by other modules.
type LinkageType uint32

// Known execution models.
const (
	LTExport LinkageType = 0 // Accessible by other modules as well.
	LTImport LinkageType = 1 // Declaration for a global identifier that exists in another module.
)

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

// SourceLanguage defines a source language constant.
type SourceLanguage uint32

// Known source languages.
const (
	SLUnknown SourceLanguage = 0
	SLESSL    SourceLanguage = 1
	SLGLSL    SourceLanguage = 2
	SLOpenCL  SourceLanguage = 3
)

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

// Decoration is used by OpDecorate and OpMemberDecorate
type Decoration uint32

// Known Decoration types
const (
	// Apply as described in the ES Precision section.
	DPrecisionLow Decoration = 0

	// Apply as described in the ES Precision section.
	DPrecisionMedium Decoration = 1

	// Apply as described in the ES Precision section.
	DPrecisionHigh Decoration = 2

	// Apply to a structure type to establish it is a non-SSBO-like
	// shader-interface block.
	//
	// TODO: can this be removed? Probably doesn’t add anything over a
	// nonwritable structure in the UniformConstant or Uniform storage class.
	// With a Binding and DescriptorSet decoration.
	DBlock Decoration = 3

	// Apply to a structure type to establish it is an SSBO-like
	// shader-interface block.
	//
	// TODO: can this be removed? Probably doesn’t add anything over a
	// structure in the UniformConstant or Uniform storage class.
	// With a Binding and DescriptorSet decoration.
	DBufferBlock Decoration = 4

	// Apply to a variable or a member of a structure. Must decorate an
	// entity whose type is a matrix. Indicates that components within a
	// row are contiguous in memory.
	DRowMajor Decoration = 5

	// Apply to a variable or a member of a structure. Must decorate an
	// entity whose type is a matrix. Indicates that components within a
	// column are contiguous in memory.
	DColMajor Decoration = 6

	// Apply to a structure type to get GLSL shared memory layout.
	DGLSLShared Decoration = 7

	// Apply to a structure type to get GLSL std140 memory layout.
	DGLSLStd140 Decoration = 8

	// Apply to a structure type to get GLSL std430 memory layout.
	DGLSLStd430 Decoration = 9

	// Apply to a structure type to get GLSL packed memory layout.
	DGLSLPacked Decoration = 10

	// Apply to a variable or a member of a structure. Indicates that
	// perspective-correct interpolation must be used. Only valid for the
	// Input and Output Storage Classes.
	DSmooth Decoration = 11

	// Apply to a variable or a member of a structure. Indicates that linear,
	// non-perspective correct interpolation must be used. Only valid for
	// the Input and Output Storage Classes.
	DNoperspective Decoration = 12

	// Apply to a variable or a member of a structure. Indicates no
	// interpolation will be done. The non-interpolated value will come
	// from a vertex, as described in the API specification. Only valid
	// for the Input and Output Storage Classes.
	DFlat Decoration = 13

	// Apply to a variable or a member of a structure. Indicates a tessellation
	// patch. Only valid for the Input and Output Storage Classes.
	DPatch Decoration = 14

	// Apply to a variable or a member of a structure. When used with
	// multi-sampling rasterization, allows a single interpolation location
	// for an entire pixel. The interpolation location must lie in both
	// the pixel and in the primitive being rasterized. Only valid for the
	// Input and Output Storage Classes.
	DCentroid Decoration = 15

	// Apply to a variable or a member of a structure. When used with
	// multi-sampling rasterization, requires per-sample interpolation.
	//
	// The interpolation locations must be the locations of the samples
	// lying in both the pixel and in the primitive being rasterized.
	// Only valid for the Input and Output Storage Classes.
	DSample Decoration = 16

	// Apply to a variable, to indicate expressions computing its value
	// be done invariant with respect to other modules computing the
	// same expressions
	DInvariant Decoration = 17

	// Apply to a variable, to indicate the compiler may compile as if there
	// is no aliasing. See the Aliasing section for more detail.
	DRestrict Decoration = 18

	// Apply to a variable, to indicate the compiler is to generate accesses
	// to the variable that work correctly in the presence of aliasing.
	// See the Aliasing section for more detail.
	DAliased Decoration = 19

	// Apply to a variable, to indicate the memory holding the variable is
	// volatile. See the Memory Model section for more detail.
	DVolatile Decoration = 20

	// Indicates that a global variable is constant and will never be modified.
	// Only allowed on global variables
	DConstant Decoration = 21

	// Apply to a variable, to indicate the memory holding the variable is
	// coherent. See the Memory Model section for more detail.
	DCoherent Decoration = 22

	// Apply to a variable, to indicate the memory holding the variable is
	// not writable, and that this module does not write to it.
	DNonwritable Decoration = 23

	// Apply to a variable, to indicate the memory holding the variable is
	// not readable, and that this module does not read from it
	DNonreadable Decoration = 24

	// Apply to a variable or a member of a structure. Asserts that the
	// value backing the decorated <id> is dynamically uniform across all
	// instantiations that might run in parallel.
	DUniform Decoration = 25

	// Apply to a variable to indicate that it is known that this
	// module does not read or write it. Useful for establishing
	// interface.
	//
	// TODO: consider removing this?
	DNoStaticUse Decoration = 26

	// Marks a structure type as "packed", indicating that the alignment
	// of the structure is one and that there is no padding between
	// structure members.
	DCPacked Decoration = 27

	// Indicates that a conversion to an integer type is saturated.
	// Only valid for conversion instructions to integer type.
	DFPSaturatedConversion Decoration = 28

	// Apply to a variable or a member of a structure. Indicates the stream
	// number to put an output on. Only valid for the Output Storage
	// Class and the Geometry Execution Model.
	//
	// Arguments:
	//  - Stream number
	//
	DStream Decoration = 29

	// Apply to a variable or a structure member. Forms the main
	// linkage for Storage Class Input and Output variables:
	//
	//  - between the API and vertex-stage inputs,
	//  - between consecutive programmable stages, or
	//  - between fragment-stage outputs and the API.
	//
	// Only valid for the Input and Output Storage Classes.
	//
	// Arguments:
	//  - Location
	//
	DLocation Decoration = 30

	// Apply to a variable or a member of a structure. Indicates
	// which component within a Location will be taken by the
	// decorated entity. Only valid for the Input and Output
	// Storage Classes.
	//
	// Arguments:
	//  - Component within a vector
	//
	DComponent Decoration = 31

	// Apply to a variable to identify a blend equation input index,
	// used as described in the API specification. Only valid for the
	// Output Storage Class and the Fragment Execution Model.
	//
	// Arguments:
	//  - Index
	//
	DIndex Decoration = 32

	// Apply to a variable. Part of the main linkage between the API
	// and SPIR-V modules for memory buffers, textures, etc. See the
	// API specification for more information.
	//
	// Arguments:
	//  - Binding point
	//
	DBinding Decoration = 33

	// Apply to a variable. Part of the main linkage between the API and
	// SPIR-V modules for memory buffers, textures, etc. See the API
	// specification for more information.
	//
	// Arguments:
	//  - Descriptor set
	//
	DDescriptorSet Decoration = 34

	// Apply to a structure member. This gives the byte offset of the
	// member relative to the beginning of the structure. Can be used,
	// for example, by both uniform and transform-feedback buffers.
	//
	// Arguments:
	//  - Byte offset
	//
	DOffset Decoration = 35

	// TODO: This can probably be removed.
	//
	// Arguments:
	//  - Declared alignment
	//
	DAlignment Decoration = 36

	// Apply to a variable or a member of a structure. Indicates which
	// transform-feedback buffer an output is written to. Only valid for
	// the Output Storage Classes of vertex processing Execution Models.
	//
	// Arguments:
	//  - XFB Buffer number
	//
	DXfbBuffer Decoration = 37

	// The stride, in bytes, of array elements or transform-feedback
	// buffer vertices.
	//
	// Arguments:
	//  - Stride
	//
	DStride Decoration = 38

	// Apply to a variable or a member of a structure.
	// Indicates which built-in variable the entity represents.
	//
	// Arguments:
	//  - See Built-In
	//
	DBuiltIn Decoration = 39

	// Indicates a function return value or parameter attribute.
	//
	// Arguments:
	//  - function parameter attribute
	//
	DFuncParamAttr Decoration = 40

	// Indicates a floating-point rounding mode
	//
	// Arguments:
	//  - floating-point rounding mode
	//
	DFPRoundingMode Decoration = 41

	// Indicates a floating-point fast math flag
	//
	// Arguments:
	//  - fast-math mode
	//
	DFPFastMathMode Decoration = 42

	// Indicate a linkage type. Only valid on an OpFunction or a
	// module scope OpVariable.
	//
	// Arguments:
	//  - linkage type
	//
	DLinkageType Decoration = 43

	// Apply to a specialization constant. Forms the API linkage for
	// setting a specialized value. See specialization.
	//
	// Arguments:
	//  - Literal Number: Specialization Constant ID
	//
	DSpecId Decoration = 44
)

// Builtin defines a builtin operation.
//
// It us when Decoration is Built-In. Apply to either:
//   - The result <id> of the variable declaration of the built-in variable, or
//   - A structure member, if the built-in is a member of a structure.
//
// These have the semantics described by their originating API and
// high-level language environments. TBD: make these native to this
// specification
type Builtin uint32

// Known builtin operations.
const (
	BPosition                  Builtin = 0
	BPointSize                 Builtin = 1
	BClipVertex                Builtin = 2
	BClipDistance              Builtin = 3
	BCullDistance              Builtin = 4
	BVertexId                  Builtin = 5
	BInstanceId                Builtin = 6
	BPrimitiveId               Builtin = 7
	BInvocationId              Builtin = 8
	BLayer                     Builtin = 9
	BViewportIndex             Builtin = 10
	BTessLevelOuter            Builtin = 11
	BTessLevelInner            Builtin = 12
	BTessCoord                 Builtin = 13
	BPatchVertices             Builtin = 14
	BFragCoord                 Builtin = 15
	BPointCoord                Builtin = 16
	BFrontFacing               Builtin = 17
	BSampleId                  Builtin = 18
	BSamplePosition            Builtin = 19
	BSampleMask                Builtin = 20
	BFragColor                 Builtin = 21
	BFragDepth                 Builtin = 22
	BHelperInvocation          Builtin = 23
	BNumWorkgroups             Builtin = 24
	BWorkgroupSize             Builtin = 25
	BWorkgroupId               Builtin = 26
	BLocalInvocationId         Builtin = 27
	BGlobalInvocationId        Builtin = 28
	BLocalInvocationIndex      Builtin = 29
	BWorkDim                   Builtin = 30
	BGlobalSize                Builtin = 31
	BEnqueuedWorkgroupSize     Builtin = 32
	BGlobalOffset              Builtin = 33
	BGlobalLinearId            Builtin = 34
	BWorkgroupLinearId         Builtin = 35
	BSubgroupSize              Builtin = 36
	BSubgroupMaxSize           Builtin = 37
	BNumSubgroups              Builtin = 38
	BNumEnqueuedSubgroups      Builtin = 39
	BSubgroupId                Builtin = 40
	BSubgroupLocalInvocationId Builtin = 41
)

// SelectionControl defines priorities for flattening
// of flow control structures.
//
// It is used by OpSelectionMerge.
type SelectionControl uint32

// Known Selection Control values.
const (
	// No control requested.
	SCNoControl SelectionControl = 0

	// Strong request, to the extent possible, to remove the flow
	// control for this selection.
	SCFlatten SelectionControl = 1

	// Strong request, to the extent possible, to keep this
	// selection as flow control.
	SCDontFlatten SelectionControl = 2
)

// LoopControl defines priorities for unrolling of
// loop constructs.
//
// It is used by OpLoopMerge.
type LoopControl uint32

// Known Loop Control values.
const (
	// No control requested.
	LCNoControl LoopControl = 0

	// Strong request, to the extent possible, to unroll or unwind this loop.
	LCUnroll LoopControl = 1

	// Strong request, to the extent possible, to keep this loop as a loop,
	// without unrolling.
	LCDontUnroll LoopControl = 2
)

// FunctionControlMask defines bitmask hints for function optimisations.
//
// It is used by OpFunction.
type FunctionControlMask uint32

// Known Function Control masks.
const (
	// Strong request, to the extent possible, to inline the function.
	FCMInLine FunctionControlMask = 1

	// Strong request, to the extent possible, to not inline the function.
	FCMDontInline FunctionControlMask = 2

	// Compiler can assume this function has no side effect, but might
	// read global memory or read through dereferenced function parameters.
	// Always computes the same result for the same argument values.
	FCMPure FunctionControlMask = 4

	// Compiler can assume this function has no side effects, and will not
	// access global memory or dereference function parameters. Always
	// computes the same result for the same argument values.
	FCMConst FunctionControlMask = 8
)

// MemorySemantics defines bitflag memory classifications and
// ordering semantics. Used by:
//
// - OpMemoryBarrier
// - OpAtomicLoad
// - OpAtomicStore
// - OpAtomicExchange
// - OpAtomicCompareExchange
// - OpAtomicCompareExchangeWeak
// - OpAtomicIIncrement
// - OpAtomicIDecrement
// - OpAtomicIAdd
// - OpAtomicISub
// - OpAtomicUMin
// - OpAtomicUMax
// - OpAtomicAnd
// - OpAtomicOr
// - OpAtomicXor
//
type MemorySemantics uint32

// Known memory semantics.
const (
	// TODO: ...
	MSRelaxed MemorySemantics = 1

	// All observers will see this memory access in the same order WRT to
	// other sequentially-consistent memory accesses from this invocation.
	MSSequentiallyConsistent MemorySemantics = 2

	// All memory operations provided in program order after this memory
	// operation will execute after this memory operation.
	MSAcquire MemorySemantics = 4

	// All memory operations provided in program order before this memory
	// operation will execute before this memory operation.
	MSRelease MemorySemantics = 8

	// Filter the memory operations being constrained to just those
	// accessing Uniform Storage Class memory.
	MSUniformMemory MemorySemantics = 16

	// The memory semantics only have to be correct WRT to this invocation’s
	// subgroup memory
	MSSubgroupMemory MemorySemantics = 32

	// The memory semantics only have to be correct WRT to this invocation’s
	// local workgroup memory.
	MSWorkgroupLocalMemory MemorySemantics = 64

	// The memory semantics only have to be correct WRT to this invocation’s
	// global workgroup memory.
	MSWorkgroupGlobalMemory MemorySemantics = 128

	// Filter the memory operations being constrained to just those
	// accessing AtomicCounter Storage Class memory.
	MSAtomicCounterMemory MemorySemantics = 256

	// Filter the memory operations being constrained to just those
	// accessing images (see OpTypeSampler Content).
	MSImageMemory MemorySemantics = 512
)

// MemoryAccess defines memory access semantics.
type MemoryAccess uint32

// Known memory access values.
const (
	// This access cannot be optimized away; it has to be executed.
	MAVolatile MemoryAccess = 1

	// This access has a known alignment, provided as a literal in
	// the next operand.
	MAAligned MemoryAccess = 2
)

// ExecutionScope defines the scope of execution.
// It is used by:
//
//  - OpControlBarrier
//  - OpMemoryBarrier
//  - OpAtomicLoad
//  - OpAtomicStore
//  - OpAtomicExchange
//  - OpAtomicCompareExchange
//  - OpAtomicCompareExchangeWeak
//  - OpAtomicIIncrement
//  - OpAtomicIDecrement
//  - OpAtomicIAdd
//  - OpAtomicISub
//  -  OpAtomicUMin
//  - OpAtomicUMax
//  - OpAtomicAnd
//  - OpAtomicOr
//  - OpAtomicXor
//  - OpAsyncGroupCopy
//  - OpWaitGroupEvents
//  - OpGroupAll
//  - OpGroupAny
//  - OpGroupBroadcast
//  - OpGroupIAdd
//  - OpGroupFAdd
//  - OpGroupFMin
//  - OpGroupUMin
//  - OpGroupSMin
//  - OpGroupFMax
//  - OpGroupUMax
//  - OpGroupSMax
//  - OpGroupReserveReadPipePackets
//  - OpGroupReserveWritePipePackets
//  - OpGroupCommitReadPipe
//  - OpGroupCommitWritePipe
//
type ExecutionScope uint32

// Known execution scopes.
const (
	// Everything executing on all the execution devices in the system.
	ESCrossDevice ExecutionScope = 0

	// Everything executing on the device executing this invocation
	ESDevice ExecutionScope = 1

	// All invocations for the invoking workgroup.
	ESWorkgroup ExecutionScope = 2

	// All invocations in the currently executing subgroup.
	ESSubgroup ExecutionScope = 3
)

// GroupOperation defines the class of workgroup or subgroup operation.
// It is used by:
type GroupOperation uint32

// Known group operations.
//
//  - OpGroupIAdd
//  - OpGroupFAdd
//  - OpGroupFMin
//  - OpGroupUMin
//  - OpGroupSMin
//  - OpGroupFMax
//  - OpGroupUMax
//  - OpGroupSMax
//
const (
	// Returns the result of a reduction operation for all values of a
	// specific value X specified by workitems within a workgroup.
	GOReduce GroupOperation = 0

	// The inclusive scan performs a binary operation with an identity
	// I and n (where n is the size of the workgroup) elements[a0, a1, . . . an-1]
	// and returns [a0, (a0 op a1), . . . (a0 op a1 op . . . op an-1)]
	GOInclusiveScan GroupOperation = 1

	// The exclusive scan performs a binary operation with an identity
	// I and n (where n is the size of the workgroup) elements[a0, a1, . . . an-1]
	// and returns [I, a0, (a0 op a1), . . . (a0 op a1 op . . . op an-2)].
	GOExclusiveScan GroupOperation = 2
)
