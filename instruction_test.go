// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

var lib = NewInstructionSet()

type InstructionTest struct {
	in   []uint32
	want Instruction
	err  error
}

func TestCodecInstructions(t *testing.T) {
	for i, st := range []InstructionTest{
		{
			in:   nil,
			want: nil,
			err:  ErrUnexpectedEOF,
		},
		{
			in:  []uint32{0x00000001},
			err: ErrInvalidInstructionSize,
		},
		{
			in:  []uint32{0x00010001},
			err: ErrMissingInstructionArgs,
		},
		{
			in:  []uint32{0x00010000},
			err: ErrUnacceptable,
		},
		{
			in:  []uint32{0x0001ffff},
			err: fmt.Errorf("unknown instruction: 0000ffff"),
		},
		{
			in: []uint32{0x0030001, uint32(SLGLSL), 450},
			want: &OpSource{
				Language: SLGLSL,
				Version:  450,
			},
		},
		{
			in: []uint32{
				0x00070002,
				0x74736574,
				0x756f7320,
				0x20656372,
				0x65747865,
				0x6f69736e,
				0x0000006e,
			},
			want: OpSourceExtension("test source extension"),
		},
		{
			in: []uint32{
				0x00070003,
				0x74736574,
				0x756f7320,
				0x20656372,
				0x65747865,
				0x6f69736e,
				0x0000006e,
			},
			want: OpExtension("test source extension"),
		},
		{
			in: []uint32{
				0x00080004,
				0x00000023,
				0x74736574,
				0x756f7320,
				0x20656372,
				0x65747865,
				0x6f69736e,
				0x0000006e,
			},
			want: &OpExtInstImport{
				ResultId: 0x23,
				Name:     "test source extension",
			},
		},
		{
			in: []uint32{0x0008002c, 1, 2, 3, 4, 5, 6, 7},
			want: &OpExtInst{
				ResultType:  1,
				ResultId:    2,
				Set:         3,
				Instruction: 4,
				Operands:    []uint32{5, 6, 7},
			},
		},
		{
			in: []uint32{0x0003002d, 1, 2},
			want: &OpUndef{
				ResultType: 1,
				ResultId:   2,
			},
		},
		{
			in: []uint32{0x00050036, 1, 0x74736574, 0x6d616e5f, 0x00000065},
			want: &OpName{
				Target: 1,
				Name:   "test_name",
			},
		},
		{
			in: []uint32{0x00060037, 1, 2, 0x74736574, 0x6d616e5f, 0x00000065},
			want: &OpMemberName{
				Type:   1,
				Member: 2,
				Name:   "test_name",
			},
		},
		{
			in: []uint32{0x00050038, 1, 0x74736574, 0x72747320, 0x676e69},
			want: &OpString{
				ResultId: 1,
				String:   "test string",
			},
		},
		{
			in: []uint32{0x00050039, 1, 2, 3, 4},
			want: &OpLine{
				Target: 1,
				File:   2,
				Line:   3,
				Column: 4,
			},
		},
		{
			in: []uint32{
				0x00030005,
				uint32(AMPhysical32),
				uint32(MMGLSL450),
			},
			want: &OpMemoryModel{
				Addressing: AMPhysical32,
				Memory:     MMGLSL450,
			},
		},
		{
			in: []uint32{0x00030006, uint32(EMFragment), 0x7f},
			want: &OpEntryPoint{
				Execution: EMFragment,
				Id:        0x7f,
			},
		},
		{
			in: []uint32{0x00060007, 0x7f, uint32(EMSpacingEqual), 1, 2, 3},
			want: &OpExecutionMode{
				EntryPoint: 0x7f,
				Mode:       EMSpacingEqual,
				Argv:       []uint32{1, 2, 3},
			},
		},
		{
			in: []uint32{
				0x000700da,
				0x74736574,
				0x756f7320,
				0x20656372,
				0x65747865,
				0x6f69736e,
				0x0000006e,
			},
			want: OpCompileFlag("test source extension"),
		},
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
				Result:     0x32,
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
				Result: 0x32,
				Width:  math.Float32bits(123),
			},
		},
		{
			in: []uint32{0x0004000c, 0x32, 0x12, 0x03},
			want: &OpTypeVector{
				Result:         0x32,
				ComponentType:  0x12,
				ComponentCount: 0x03,
			},
		},
		{
			in: []uint32{0x0004000d, 0x32, 0x12, 0x04},
			want: &OpTypeMatrix{
				Result:      0x32,
				ColumnType:  0x12,
				ColumnCount: 0x04,
			},
		},
		{
			in:   []uint32{0x00020031, 0xff},
			want: OpDecorationGroup(0xff),
		},
		{
			in: []uint32{0x00060032, 1, uint32(DNoStaticUse), 2, 3, 4},
			want: &OpDecorate{
				Target:     1,
				Decoration: DNoStaticUse,
				Argv:       []uint32{2, 3, 4},
			},
		},
		{
			in: []uint32{0x00070033, 1, 2, uint32(DNoStaticUse), 3, 4, 5},
			want: &OpMemberDecorate{
				StructType: 1,
				Member:     2,
				Decoration: DNoStaticUse,
				Argv:       []uint32{3, 4, 5},
			},
		},
		{
			in: []uint32{0x00060034, 1, 2, 3, 4, 5},
			want: &OpGroupDecorate{
				Group:   1,
				Targets: []uint32{2, 3, 4, 5},
			},
		},
		{
			in: []uint32{0x00060035, 1, 2, 3, 4, 5},
			want: &OpGroupMemberDecorate{
				Group:   1,
				Targets: []uint32{2, 3, 4, 5},
			},
		},
		{
			in: []uint32{0x008000e, 1, 2, uint32(D3D), 2, 1, 0, 1},
			want: &OpTypeSampler{
				Result:      1,
				SampledType: 2,
				Dim:         D3D,
				Content:     2,
				Arrayed:     1,
				Compare:     0,
				MS:          1,
			},
		},
		{
			in: []uint32{0x009000e, 1, 2, uint32(D3D), 2, 1, 0, 1, uint32(AQWriteOnly)},
			want: &OpTypeSampler{
				Result:      1,
				SampledType: 2,
				Dim:         D3D,
				Content:     2,
				Arrayed:     1,
				Compare:     0,
				MS:          1,
				Qualifier:   AQWriteOnly,
			},
		},
		{
			in:   []uint32{0x002000f, 1},
			want: OpTypeFilter(1),
		},
		{
			in: []uint32{0x0040010, 1, 2, 3},
			want: &OpTypeArray{
				Result:      1,
				ElementType: 2,
				Length:      3,
			},
		},
		{
			in: []uint32{0x0030011, 1, 2},
			want: &OpTypeRuntimeArray{
				Result:      1,
				ElementType: 2,
			},
		},
		{
			in: []uint32{0x0060012, 1, 2, 3, 4, 5},
			want: &OpTypeStruct{
				Result:  1,
				Members: []uint32{2, 3, 4, 5},
			},
		},
		{
			in: []uint32{0x0050013, 1, 0x74736574, 0x72747320, 0x676e69},
			want: &OpTypeOpaque{
				Result: 1,
				Name:   "test string",
			},
		},
		{
			in: []uint32{0x0040014, 1, uint32(SCAtomicCounter), 2},
			want: &OpTypePointer{
				Result:  1,
				Storage: SCAtomicCounter,
				Type:    2,
			},
		},
		{
			in: []uint32{0x0030015, 1, 2},
			want: &OpTypeFunction{
				Result:     1,
				ReturnType: 2,
			},
		},
		{
			in: []uint32{0x0060015, 1, 2, 3, 4, 5},
			want: &OpTypeFunction{
				Result:     1,
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
	} {
		have, err := lib.Decode(st.in)

		// We have a decoding error. This is only a test failure if
		// we were not expecting an error.
		if err != nil {
			if !reflect.DeepEqual(err, st.err) {
				t.Fatalf("case %d: decode error mismatch: %v\nHave: %v\nWant: %v",
					i, st.in, err, st.err)
			}

			// We got an error as we expected, so no further processing
			// is needed.
			continue
		}

		if !reflect.DeepEqual(have, st.want) {
			t.Fatalf("case %d: decode value mismatch: %v\nHave: %T(%+v)\nWant: %T(%+v)",
				i, st.in, have, have, st.want, st.want)
		}

		data, err := lib.Encode(have)
		if !reflect.DeepEqual(data, st.in) {
			t.Fatalf("case %d: encode mismatch: %T(%v)\nHave: %v\nWant: %v",
				i, have, have, data, st.in)
		}
	}
}
