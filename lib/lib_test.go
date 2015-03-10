// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package lib

import (
	"math"
	"reflect"
	"testing"

	"github.com/jteeuwen/spirv"
)

var lib InstructionSet

func init() { lib.BindDefaults() }

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
			err:  spirv.ErrUnexpectedEOF,
		},
		{
			in:  []uint32{0x00000001},
			err: spirv.ErrInvalidInstructionSize,
		},
		{
			in:  []uint32{0x00010001},
			err: ErrMissingInstructionArgs,
		},
		{
			in: []uint32{
				0x0030001,
				uint32(SourceGLSL),
				450,
			},
			want: &OpSource{
				Language: SourceGLSL,
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
			in: []uint32{
				0x00030005,
				uint32(AddressPhysical32),
				uint32(MemoryGLSL450),
			},
			want: &OpMemoryModel{
				Addressing: AddressPhysical32,
				Memory:     MemoryGLSL450,
			},
		},
		{
			in: []uint32{0x00030006, uint32(ExecFragment), 0x7f},
			want: &OpEntryPoint{
				Execution: ExecFragment,
				Id:        0x7f,
			},
		},
		{
			in: []uint32{0x00060007, 0x7f, uint32(ModeSpacingEqual), 1, 2, 3},
			want: &OpExecutionMode{
				EntryPoint: 0x7f,
				Mode:       ModeSpacingEqual,
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
				reverse(math.Float32bits(123)),
			},
			want: &OpTypeFloat{
				Result: 0x32,
				Width:  reverse(math.Float32bits(123)),
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
			t.Fatalf("case %d: decode value mismatch: %v\nHave: %#v\nWant: %#v",
				i, st.in, have, st.want)
		}

		data, err := lib.Encode(have)
		if !reflect.DeepEqual(data, st.in) {
			t.Fatalf("case %d: encode mismatch: %T(%v)\nHave: %v\nWant: %v",
				i, have, have, data, st.in)
		}
	}
}

// reverse reverses the bytes in the given integer.
func reverse(v uint32) uint32 {
	a := v & 0xff
	b := (v >> 8) & 0xff
	c := (v >> 16) & 0xff
	d := (v >> 24) & 0xff
	return a<<24 | b<<16 | c<<8 | d
}