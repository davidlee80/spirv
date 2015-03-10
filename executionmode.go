// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "fmt"

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
