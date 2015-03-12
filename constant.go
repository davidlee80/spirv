// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// Access Qualifiers define the access permissions of OpTypeSampler
// and OpTypePipe object. Used by OpTypePipe.
const (
	AQReadOnly  = 0 // A read-only object.
	AQWriteOnly = 1 // A write-only object.
	AQReadWrite = 2 // A readable and writable object.
)

// Addressing Modes define an existing addressing mode.
const (
	AMLogical    = 0
	AMPhysical32 = 1
	AMPhysical64 = 2
)

// Dimensionalities define the dimensionality of a texture.
//
// Used by OpTypeSampler.
const (
	D1D     = 0
	D2D     = 1
	D3D     = 2
	DCube   = 3
	DRect   = 4
	DBuffer = 5
)

// Execution Modes define a mode a module’s stage will execute in.
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
	EMInvocations = 0

	// Requests the tessellation primitive generator to divide edges into a
	// collection of equal-sized segments. Only valid with one of the
	// tessellation Execution Models.
	EMSpacingEqual = 1

	// Requests the tessellation primitive generator to divide edges into an
	// even number of equal-length segments plus two additional shorter
	// fractional segments. Only valid with one of the tessellation
	// Execution Models.
	EMSpacingFractionalEven = 2

	// Requests the tessellation primitive generator to divide edges into an
	// even number of equal-length segments plus two additional shorter
	// fractional segments. Only valid with one of the tessellation.
	// Execution Models.
	EMSpacingFractionalOdd = 3

	// Requests the tessellation primitive generator to generate triangles in
	// clockwise order. Only valid with one of the tessellation Execution Models.
	EMVertexOrderCw = 4

	// Requests the tessellation primitive generator to generate triangles in
	// counter-clockwise order. Only valid with one of the tessellation
	// Execution Models.
	EMVertexOrderCcw = 5

	// Pixels appear centered on whole-number pixel offsets. E.g., the
	// coordinate (0.5, 0.5) appears to move to (0.0, 0.0). Only valid with
	// the Fragment Execution Model.
	EMPixelCenterInteger = 6

	// Pixel coordinates appear to originate in the upper left, and increase
	// toward the right and downward. Only valid with the Fragment Execution Model.
	EMOriginUpperLeft = 7

	// Fragment tests are to be performed before fragment shader execution.
	// Only valid with the Fragment Execution Model.
	EMEarlyFragmentTests = 8

	// Requests the tessellation primitive generator to generate a point for
	// each distinct vertex in the subdivided primitive, rather than to
	// generate lines or triangles. Only valid with one of the tessellation
	// Execution Models.
	EMPointMode = 9

	// This stage will run in transform feedback-capturing mode and this module
	// is responsible for describing the transform-feedback setup.
	// See the XfbBuffer, Offset, and Stride Decorations.
	EMXFB = 10

	// This mode must be declared if this module potentially changes the
	// fragment’s depth. Only valid with the Fragment Execution Model.
	EMDepthReplacing = 11

	// TODO: this should probably be removed. Depth testing will always be
	// performed after the shader has executed. Only valid with the Fragment
	// Execution Model.
	EMDepthAny = 12

	// External optimizations may assume depth modifications will leave the
	// fragment’s depth as greater than or equal to the fragment’s interpolated
	// depth value (given by the z component of the FragCoord Built-In
	// decorated variable). Only valid with the Fragment Execution Model.
	EMDepthGreater = 13

	// External optimizations may assume depth modifications leave the
	// fragment’s depth less than the fragment’s interpolated depth
	// value, (given by the z component of the FragCoord Built-In decorated
	// variable). Only valid with the Fragment Execution Model.
	EMDepthLess = 14

	// External optimizations may assume this stage did not modify the
	// fragment’s depth. However, DepthReplacing mode must accurately
	// represent depth modification. Only valid with the Fragment Execution Model.
	EMDepthUnchanged = 15

	// Indicates the work-group size in the x, y, and z dimensions. Only valid
	// with the GLCompute or Kernel Execution Models.
	//
	// Arguments:
	//
	//   [0]: x size
	//   [1]: y size
	//   [2]: z size
	//
	EMLocalSize = 16

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
	EMLocalSizeHint = 17

	// Stage input primitive is points. Only valid with the Geometry Execution Model.
	EMInputPoints = 18

	// Stage input primitive is lines. Only valid with the Geometry Execution Model.
	EMInputLines = 19

	// Stage input primitive is lines adjacency. Only valid with the Geometry
	// Execution Model.
	EMInputLinesAdjacency = 20

	// For a geometry stage, input primitive is triangles. For a tessellation
	// stage, requests the tessellation primitive generator to generate
	// triangles. Only valid with the Geometry or one of the tessellation
	// Execution Models.
	EMInputTriangles = 21

	// Geometry stage input primitive is triangles adjacency. Only valid with
	// the Geometry Execution Model.
	EMInputTrianglesAdjacency = 22

	// Requests the tessellation primitive generator to generate quads.
	// Only valid with one of the tessellation Execution Models.
	EMInputQuads = 23

	// Requests the tessellation primitive generator to generate isolines.
	// Only valid with one of the tessellation Execution Models
	EMInputIsolines = 24

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
	EMOutputVertices = 25

	// Stage output primitive is points. Only valid with the Geometry
	// Execution Model.
	EMOutputPoints = 26

	// Stage output primitive is line strip. Only valid with the Geometry
	// Execution Model.
	EMOutputLinestrip = 27

	// Stage output primitive is triangle strip. Only valid with the
	// Geometry Execution Model.
	EMOutputTrianglestrip = 28

	// A hint to the compiler, which indicates that most operations used
	// in the entry point are explicitly vectorized using a particular
	// vector type. Only valid with the Kernel Execution Model.
	//
	// Arguments:
	//
	//   [0]: Vector type
	//
	EMVecTypeHint = 29

	// Indicates that floating-point-expressions contraction is disallowed.
	// Only valid with the Kernel Execution Model.
	EMContractionOff = 30
)

// Execution Models define a single execution model.
// This is used in the EntryPoint instruction to determine what stage of the
// pipeline a given set of instructions belongs to.
const (
	EMVertex                 = 0 // Vertex shading stage
	EMTessellationControl    = 1 // Tessellation control (or hull) shading stage.
	EMTessellationEvaluation = 2 // Tessellation evaluation (or domain) shading stage
	EMGeometry               = 3 // Geometry shading stage.
	EMFragment               = 4 // Fragment shading stage.
	EMGLCompute              = 5 // Graphical compute shading stage.
	EMKernel                 = 6 // Compute kernel.
)

// FPFastMathModes define bitflags which enable fast math operations
// which are otherwise unsafe.
//
// Only valid on OpFAdd, OpFSub, OpFMul, OpFDiv, OpFRem
// and OpFMod instructions.
const (
	// Assume parameters and result are not NaN.
	FMMNotNaN = 0

	// Assume parameters and result are not +/- Inf.
	FMMNotInf = 2

	// Treat the sign of a zero parameter or result as insignificant.
	FMMNSZ = 4

	// Allow the usage of reciprocal rather than perform a division.
	FMMAllowRecip = 8

	// Allow algebraic transformations according to real-number associative
	// and distributive algebra. This flag implies all the others.
	FMMFast = 16
)

// FPRoundingModes associate a rounding mode with a floating-point
// conversion instruction.
//
// By default:
//
//    - Conversions from floating-point to integer types use the
//      round-toward-zero rounding mode.
//    - Conversions to floating-point types use the round-to-nearest-even
//      rounding mode.
//
const (
	FPRMRTE = 0 // Round to nearest even.
	FPRMRTZ = 1 // Round towards zero.
	FPRMRTP = 2 // Round towards positive infinity.
	FPRMRTN = 3 // Round towards negative infinity.
)

// LinkageTypes associate a linkage type with functions or global
// variables. By default, functions and global variables are private
// to a module and cannot be accessed by other modules.
const (
	LTExport = 0 // Accessible by other modules as well.
	LTImport = 1 // Declaration for a global identifier that exists in another module.
)

// Memory Models define an existing memory model.
const (
	MMSimple   = 0 // No shared memory consistency issues.
	MMGLSL450  = 1 // Memory model needed by later versions of GLSL and ESSL. Works across multiple versions.
	MMOpenCL12 = 2 // OpenCL 1.2 memory model.
	MMOpenCL20 = 3 // OpenCL 2.0 memory model.
	MMOpenCL21 = 4 // OpenCL 2.1 memory model.
)

// Sampler Addressing Modes define the addressing mode of read image
// extended instructions.
const (
	// The image coordinates used to sample elements of the image refer to a
	// location inside the image, otherwise the results are undefined.
	SAMNone = 0

	// Out-of-range image coordinates are clamped to the extent.
	SAMClampEdge = 2

	// Out-of-range image coordinates will return a border color.
	SAMClamp = 4

	// Out-of-range image coordinates are wrapped to the valid range.
	// Can only be used with normalized coordinates.
	SAMRepeat = 6

	// Flip the image coordinate at every integer junction.
	// Can only be used with normalized coordinates.
	SAMRepeatMirrored = 8
)

// Sampler Filter Modes define the filter mode of read image
// extended instructions.
const (
	// Use filter nearset mode when performing a read image operation.
	SFMNearest = 16

	// Use filter linear mode when performing a read image operation.
	SFMLinear = 32
)

// Source Languages define a source language constant.
const (
	SLUnknown = 0
	SLESSL    = 1
	SLGLSL    = 2
	SLOpenCL  = 3
)

// Storage Classes define a class of storage for declared variables
// (does not include intermediate values).
//
// Used by: OpTypePointer, OpTypeVariable, OpTypeVariableArray,
// OpTypeGenericCastToPtrExplicit
const (
	// Shared externally, read-only memory, visible across all instantiation
	// or work groups. Graphics uniform memory. OpenCL Constant memory
	SCUniformConstant = 0

	// Input from pipeline. Read only
	SCInput = 1

	// Shared externally, visible across all instantiations or work groups
	SCUniform = 2

	// Output to pipeline.
	SCOutput = 3

	// Shared across all work items within a work group. OpenGL "shared".
	// OpenCL local memory.
	SCWorkgroupLocal = 4

	// Visible across all work items of all work groups. OpenCL global memory.
	SCWorkgroupGlobal = 5

	// Accessible across functions within a module, non-IO (not visible outside
	// the module).
	SCPrivateGlobal = 6

	// A variable local to a function.
	SCFunction = 7

	// A generic pointer, which overloads StoragePrivate, StorageLocal,
	// StorageGlobal. not a real storage class.
	SCGeneric = 8

	// Private to a work-item and is not visible to another work-item.
	// OpenCL private memory.
	SCPrivate = 9

	// For holding atomic counters.
	SCAtomicCounter = 10
)

// Function Parameter Attributes add additional information to the return type
// and to each parameter of a function.
const (
	// Value should be zero extended if needed.
	FPAZext = 0

	// Value should be sign extended if needed.
	FPASext = 1

	// This indicates that the pointer parameter should really be passed by
	// value to the function. Only valid for pointer parameters (not
	// for ret value)
	FPAByVal = 2

	// Indicates that the pointer parameter specifies the address of a
	// structure that is the return value of the function in the source
	// program. Only applicable to the first parameter which must be a
	// pointer parameters.
	FPASret = 3

	// Indicates that the memory pointed by a pointer parameter is not
	// accessed via pointer values which are not derived from this pointer
	// parameter. Only valid for pointer parameters. Not valid on return values
	FPANoAlias = 4

	// The callee does not make a copy of the pointer parameter into a
	// location that is accessible after returning from the callee. Only
	// valid for pointer parameters. Not valid on return values.
	FPANoCapture = 5

	// To be determined.
	FPASVM = 6

	// Can only read the memory pointed by a pointer parameter.
	// Only valid for pointer parameters. Not valid on return values.
	FPANoWrite = 7

	// Cannot dereference the memory pointed by a pointer parameter.
	// Only valid for pointer parameters. Not valid on return values.
	FPANoReadWrite = 8
)

// Decorations are used by OpDecorate and OpMemberDecorate
const (
	// Apply as described in the ES Precision section.
	DPrecisionLow = 0

	// Apply as described in the ES Precision section.
	DPrecisionMedium = 1

	// Apply as described in the ES Precision section.
	DPrecisionHigh = 2

	// Apply to a structure type to establish it is a non-SSBO-like
	// shader-interface block.
	//
	// TODO: can this be removed? Probably doesn’t add anything over a
	// nonwritable structure in the UniformConstant or Uniform storage class.
	// With a Binding and DescriptorSet decoration.
	DBlock = 3

	// Apply to a structure type to establish it is an SSBO-like
	// shader-interface block.
	//
	// TODO: can this be removed? Probably doesn’t add anything over a
	// structure in the UniformConstant or Uniform storage class.
	// With a Binding and DescriptorSet decoration.
	DBufferBlock = 4

	// Apply to a variable or a member of a structure. Must decorate an
	// entity whose type is a matrix. Indicates that components within a
	// row are contiguous in memory.
	DRowMajor = 5

	// Apply to a variable or a member of a structure. Must decorate an
	// entity whose type is a matrix. Indicates that components within a
	// column are contiguous in memory.
	DColMajor = 6

	// Apply to a structure type to get GLSL shared memory layout.
	DGLSLShared = 7

	// Apply to a structure type to get GLSL std140 memory layout.
	DGLSLStd140 = 8

	// Apply to a structure type to get GLSL std430 memory layout.
	DGLSLStd430 = 9

	// Apply to a structure type to get GLSL packed memory layout.
	DGLSLPacked = 10

	// Apply to a variable or a member of a structure. Indicates that
	// perspective-correct interpolation must be used. Only valid for the
	// Input and Output Storage Classes.
	DSmooth = 11

	// Apply to a variable or a member of a structure. Indicates that linear,
	// non-perspective correct interpolation must be used. Only valid for
	// the Input and Output Storage Classes.
	DNoperspective = 12

	// Apply to a variable or a member of a structure. Indicates no
	// interpolation will be done. The non-interpolated value will come
	// from a vertex, as described in the API specification. Only valid
	// for the Input and Output Storage Classes.
	DFlat = 13

	// Apply to a variable or a member of a structure. Indicates a tessellation
	// patch. Only valid for the Input and Output Storage Classes.
	DPatch = 14

	// Apply to a variable or a member of a structure. When used with
	// multi-sampling rasterization, allows a single interpolation location
	// for an entire pixel. The interpolation location must lie in both
	// the pixel and in the primitive being rasterized. Only valid for the
	// Input and Output Storage Classes.
	DCentroid = 15

	// Apply to a variable or a member of a structure. When used with
	// multi-sampling rasterization, requires per-sample interpolation.
	//
	// The interpolation locations must be the locations of the samples
	// lying in both the pixel and in the primitive being rasterized.
	// Only valid for the Input and Output Storage Classes.
	DSample = 16

	// Apply to a variable, to indicate expressions computing its value
	// be done invariant with respect to other modules computing the
	// same expressions
	DInvariant = 17

	// Apply to a variable, to indicate the compiler may compile as if there
	// is no aliasing. See the Aliasing section for more detail.
	DRestrict = 18

	// Apply to a variable, to indicate the compiler is to generate accesses
	// to the variable that work correctly in the presence of aliasing.
	// See the Aliasing section for more detail.
	DAliased = 19

	// Apply to a variable, to indicate the memory holding the variable is
	// volatile. See the Memory Model section for more detail.
	DVolatile = 20

	// Indicates that a global variable is constant and will never be modified.
	// Only allowed on global variables
	DConstant = 21

	// Apply to a variable, to indicate the memory holding the variable is
	// coherent. See the Memory Model section for more detail.
	DCoherent = 22

	// Apply to a variable, to indicate the memory holding the variable is
	// not writable, and that this module does not write to it.
	DNonwritable = 23

	// Apply to a variable, to indicate the memory holding the variable is
	// not readable, and that this module does not read from it
	DNonreadable = 24

	// Apply to a variable or a member of a structure. Asserts that the
	// value backing the decorated <id> is dynamically uniform across all
	// instantiations that might run in parallel.
	DUniform = 25

	// Apply to a variable to indicate that it is known that this
	// module does not read or write it. Useful for establishing
	// interface.
	//
	// TODO: consider removing this?
	DNoStaticUse = 26

	// Marks a structure type as "packed", indicating that the alignment
	// of the structure is one and that there is no padding between
	// structure members.
	DCPacked = 27

	// Indicates that a conversion to an integer type is saturated.
	// Only valid for conversion instructions to integer type.
	DFPSaturatedConversion = 28

	// Apply to a variable or a member of a structure. Indicates the stream
	// number to put an output on. Only valid for the Output Storage
	// Class and the Geometry Execution Model.
	//
	// Arguments:
	//  - Stream number
	//
	DStream = 29

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
	DLocation = 30

	// Apply to a variable or a member of a structure. Indicates
	// which component within a Location will be taken by the
	// decorated entity. Only valid for the Input and Output
	// Storage Classes.
	//
	// Arguments:
	//  - Component within a vector
	//
	DComponent = 31

	// Apply to a variable to identify a blend equation input index,
	// used as described in the API specification. Only valid for the
	// Output Storage Class and the Fragment Execution Model.
	//
	// Arguments:
	//  - Index
	//
	DIndex = 32

	// Apply to a variable. Part of the main linkage between the API
	// and SPIR-V modules for memory buffers, textures, etc. See the
	// API specification for more information.
	//
	// Arguments:
	//  - Binding point
	//
	DBinding = 33

	// Apply to a variable. Part of the main linkage between the API and
	// SPIR-V modules for memory buffers, textures, etc. See the API
	// specification for more information.
	//
	// Arguments:
	//  - Descriptor set
	//
	DDescriptorSet = 34

	// Apply to a structure member. This gives the byte offset of the
	// member relative to the beginning of the structure. Can be used,
	// for example, by both uniform and transform-feedback buffers.
	//
	// Arguments:
	//  - Byte offset
	//
	DOffset = 35

	// TODO: This can probably be removed.
	//
	// Arguments:
	//  - Declared alignment
	//
	DAlignment = 36

	// Apply to a variable or a member of a structure. Indicates which
	// transform-feedback buffer an output is written to. Only valid for
	// the Output Storage Classes of vertex processing Execution Models.
	//
	// Arguments:
	//  - XFB Buffer number
	//
	DXfbBuffer = 37

	// The stride, in bytes, of array elements or transform-feedback
	// buffer vertices.
	//
	// Arguments:
	//  - Stride
	//
	DStride = 38

	// Apply to a variable or a member of a structure.
	// Indicates which built-in variable the entity represents.
	//
	// Arguments:
	//  - See Built-In
	//
	DBuiltIn = 39

	// Indicates a function return value or parameter attribute.
	//
	// Arguments:
	//  - function parameter attribute
	//
	DFuncParamAttr = 40

	// Indicates a floating-point rounding mode
	//
	// Arguments:
	//  - floating-point rounding mode
	//
	DFPRoundingMode = 41

	// Indicates a floating-point fast math flag
	//
	// Arguments:
	//  - fast-math mode
	//
	DFPFastMathMode = 42

	// Indicate a linkage type. Only valid on an OpFunction or a
	// module scope OpVariable.
	//
	// Arguments:
	//  - linkage type
	//
	DLinkageType = 43

	// Apply to a specialization constant. Forms the API linkage for
	// setting a specialized value. See specialization.
	//
	// Arguments:
	//  - Literal Number: Specialization Constant ID
	//
	DSpecId = 44
)

// Builtins define a builtin operation.
//
// Used when Decoration is Built-In. Apply to either:
//   - The result <id> of the variable declaration of the built-in variable, or
//   - A structure member, if the built-in is a member of a structure.
//
// These have the semantics described by their originating API and
// high-level language environments.
//
// TODO: make these native to this specification
const (
	BPosition                  = 0
	BPointSize                 = 1
	BClipVertex                = 2
	BClipDistance              = 3
	BCullDistance              = 4
	BVertexId                  = 5
	BInstanceId                = 6
	BPrimitiveId               = 7
	BInvocationId              = 8
	BLayer                     = 9
	BViewportIndex             = 10
	BTessLevelOuter            = 11
	BTessLevelInner            = 12
	BTessCoord                 = 13
	BPatchVertices             = 14
	BFragCoord                 = 15
	BPointCoord                = 16
	BFrontFacing               = 17
	BSampleId                  = 18
	BSamplePosition            = 19
	BSampleMask                = 20
	BFragColor                 = 21
	BFragDepth                 = 22
	BHelperInvocation          = 23
	BNumWorkgroups             = 24
	BWorkgroupSize             = 25
	BWorkgroupId               = 26
	BLocalInvocationId         = 27
	BGlobalInvocationId        = 28
	BLocalInvocationIndex      = 29
	BWorkDim                   = 30
	BGlobalSize                = 31
	BEnqueuedWorkgroupSize     = 32
	BGlobalOffset              = 33
	BGlobalLinearId            = 34
	BWorkgroupLinearId         = 35
	BSubgroupSize              = 36
	BSubgroupMaxSize           = 37
	BNumSubgroups              = 38
	BNumEnqueuedSubgroups      = 39
	BSubgroupId                = 40
	BSubgroupLocalInvocationId = 41
)

// Selection Controls define priorities for flattening
// of flow control structures.
//
// These are used by OpSelectionMerge.
const (
	// No control requested.
	SCNoControl = 0

	// Strong request, to the extent possible, to remove the flow
	// control for this selection.
	SCFlatten = 1

	// Strong request, to the extent possible, to keep this
	// selection as flow control.
	SCDontFlatten = 2
)

// Loop Controls define priorities for unrolling of
// loop constructs.
//
// They are used by OpLoopMerge.
const (
	// No control requested.
	LCNoControl = 0

	// Strong request, to the extent possible, to unroll or unwind this loop.
	LCUnroll = 1

	// Strong request, to the extent possible, to keep this loop as a loop,
	// without unrolling.
	LCDontUnroll = 2
)

// Function Control Masks define bitmask hints for function optimisations.
//
// These are used by OpFunction.
const (
	// Strong request, to the extent possible, to inline the function.
	FCMInLine = 1

	// Strong request, to the extent possible, to not inline the function.
	FCMDontInline = 2

	// Compiler can assume this function has no side effect, but might
	// read global memory or read through dereferenced function parameters.
	// Always computes the same result for the same argument values.
	FCMPure = 4

	// Compiler can assume this function has no side effects, and will not
	// access global memory or dereference function parameters. Always
	// computes the same result for the same argument values.
	FCMConst = 8
)

// Memory Semantics define bitflag memory classifications and
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
const (
	// TODO: ...
	MSRelaxed = 1

	// All observers will see this memory access in the same order WRT to
	// other sequentially-consistent memory accesses from this invocation.
	MSSequentiallyConsistent = 2

	// All memory operations provided in program order after this memory
	// operation will execute after this memory operation.
	MSAcquire = 4

	// All memory operations provided in program order before this memory
	// operation will execute before this memory operation.
	MSRelease = 8

	// Filter the memory operations being constrained to just those
	// accessing Uniform Storage Class memory.
	MSUniformMemory = 16

	// The memory semantics only have to be correct WRT to this invocation’s
	// subgroup memory
	MSSubgroupMemory = 32

	// The memory semantics only have to be correct WRT to this invocation’s
	// local workgroup memory.
	MSWorkgroupLocalMemory = 64

	// The memory semantics only have to be correct WRT to this invocation’s
	// global workgroup memory.
	MSWorkgroupGlobalMemory = 128

	// Filter the memory operations being constrained to just those
	// accessing AtomicCounter Storage Class memory.
	MSAtomicCounterMemory = 256

	// Filter the memory operations being constrained to just those
	// accessing images (see OpTypeSampler Content).
	MSImageMemory = 512
)

// Memory Access defines memory access semantics.
const (
	// This access cannot be optimized away; it has to be executed.
	MAVolatile = 1

	// This access has a known alignment, provided as a literal in
	// the next operand.
	MAAligned = 2
)

// Execution Scopes define the scope of execution.
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
const (
	// Everything executing on all the execution devices in the system.
	ESCrossDevice = 0

	// Everything executing on the device executing this invocation
	ESDevice = 1

	// All invocations for the invoking workgroup.
	ESWorkgroup = 2

	// All invocations in the currently executing subgroup.
	ESSubgroup = 3
)

// Group Operations define the class of workgroup or subgroup operation.
// It is used by:
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
	GOReduce = 0

	// The inclusive scan performs a binary operation with an identity
	// I and n (where n is the size of the workgroup) elements[a0, a1, . . . an-1]
	// and returns [a0, (a0 op a1), . . . (a0 op a1 op . . . op an-1)]
	GOInclusiveScan = 1

	// The exclusive scan performs a binary operation with an identity
	// I and n (where n is the size of the workgroup) elements[a0, a1, . . . an-1]
	// and returns [I, a0, (a0 op a1), . . . (a0 op a1 op . . . op an-2)].
	GOExclusiveScan = 2
)

// Kernel Enqueue Flags specify when the child kernel begins execution.
//
// Note: Implementations are not required to honor this flag. Implementations
// may not schedule kernel launch earlier than the point specified by this
// flag, however.
//
// They are used by OpEnqueueKernel.
const (
	// Indicates that the enqueued kernels do not need to wait for the
	// parent kernel to finish execution before they begin execution.
	KEFNoWait = 0

	// Indicates that all work-items of the parent kernel must finish
	// executing and all immediate side effects committed before the
	// enqueued child kernel may begin execution.
	//
	// Note: Immediate meaning not side effects resulting from child
	// kernels. The side effects would include stores to global memory
	// and pipe reads and writes.
	KEFWaitKernel = 1

	// Indicates that the enqueued kernels wait only for the workgroup that
	// enqueued the kernels to finish before they begin execution.
	//
	// Note: This acts as a memory synchronization point between work-items
	// in a work-group and child kernels enqueued by work-items in the
	// work-group.
	KEFWaitWorkGroup = 2
)

// Kernel Profiling Info specifies the profiling information to be queried.
// Used by OpCaptureEventProfilingInfo.
const (
	KPICmdExecTime = 1
)
