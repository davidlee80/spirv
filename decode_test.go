// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import (
	"bytes"
	"math"
	"reflect"
	"testing"
)

type InstructionTest struct {
	in   []byte
	want interface{}
	err  error
}

func TestDecodeInstructions(t *testing.T) {
	for i, st := range []InstructionTest{
		{
			in:   nil,
			want: nil,
			err:  ErrUnexpectedEOF,
		},
		{
			in: []byte{
				0x00, 0x00, 0x00, opSource,
			},
			err: ErrInvalidInstructionSize,
		},
		{
			in: []byte{
				0x00, 0x01, 0x00, opSource,
			},
			err: ErrMissingInstructionArgs,
		},
		{
			in: []byte{
				0x00, 0x03, 0x00, opSource,
				0x00, 0x00, 0x00, byte(SourceGLSL),
				0x00, 0x00, 0x01, 0xc2,
			},
			want: &OpSource{
				Language: SourceGLSL,
				Version:  450,
			},
		},
		{
			in: []byte{
				0x00, 0x07, 0x00, opSourceExtension,
				0x74, 0x73, 0x65, 0x74,
				0x75, 0x6f, 0x73, 0x20,
				0x20, 0x65, 0x63, 0x72,
				0x65, 0x74, 0x78, 0x65,
				0x6f, 0x69, 0x73, 0x6e,
				0x00, 0x00, 0x00, 0x6e,
			},
			want: OpSourceExtension("test source extension"),
		},
		{
			in: []byte{
				0x00, 0x05, 0x00, opExtension,
				0x74, 0x73, 0x65, 0x74,
				0x74, 0x78, 0x65, 0x20,
				0x69, 0x73, 0x6e, 0x65,
				0x00, 0x00, 0x6e, 0x6f,
			},
			want: OpExtension("test extension"),
		},
		{
			in: []byte{
				0x00, 0x06, 0x00, opExtInstImport,
				0x00, 0x00, 0x00, 0x23,
				0x74, 0x73, 0x65, 0x74,
				0x74, 0x78, 0x65, 0x20,
				0x69, 0x73, 0x6e, 0x65,
				0x00, 0x00, 0x6e, 0x6f,
			},
			want: &OpExtInstImport{
				ResultId: 0x23,
				Name:     "test extension",
			},
		},
		{
			in: []byte{
				0x00, 0x08, 0x00, opExtInst,
				0x00, 0x00, 0x00, 0x01,
				0x00, 0x00, 0x00, 0x02,
				0x00, 0x00, 0x00, 0x03,
				0x00, 0x00, 0x00, 0x04,
				0x00, 0x00, 0x00, 0x05,
				0x00, 0x00, 0x00, 0x06,
				0x00, 0x00, 0x00, 0x07,
			},
			want: &OpExtInst{
				ResultType:  0x01,
				ResultId:    0x02,
				Set:         0x03,
				Instruction: 0x04,
				Operands:    []uint32{0x05, 0x06, 0x07},
			},
		},
		{
			in: []byte{
				0x00, 0x03, 0x00, opMemoryModel,
				0x00, 0x00, 0x00, byte(AddressPhysical32),
				0x00, 0x00, 0x00, byte(MemoryGLSL450),
			},
			want: &OpMemoryModel{
				Addressing: AddressPhysical32,
				Memory:     MemoryGLSL450,
			},
		},
		{
			in: []byte{
				0x00, 0x03, 0x00, opEntryPoint,
				0x00, 0x00, 0x00, byte(ExecFragment),
				0x00, 0x00, 0x00, 0x7f,
			},
			want: &OpEntryPoint{
				Execution: ExecFragment,
				Id:        0x7f,
			},
		},
		{
			in: []byte{
				0x00, 0x06, 0x00, opExecutionMode,
				0x00, 0x00, 0x00, 0x7f,
				0x00, 0x00, 0x00, byte(ModeSpacingEqual),
				0x00, 0x00, 0x00, 0x01,
				0x00, 0x00, 0x00, 0x02,
				0x00, 0x00, 0x00, 0x03,
			},
			want: &OpExecutionMode{
				EntryPoint: 0x7f,
				Mode:       ModeSpacingEqual,
				Argv:       []uint32{0x01, 0x02, 0x03},
			},
		},
		{
			in: []byte{
				0x00, 0x05, 0x00, opCompileFlag,
				0x74, 0x73, 0x65, 0x74,
				0x74, 0x78, 0x65, 0x20,
				0x69, 0x73, 0x6e, 0x65,
				0x00, 0x00, 0x6e, 0x6f,
			},
			want: OpCompileFlag("test extension"),
		},
		{
			in: []byte{
				0x00, 0x02, 0x00, opTypeVoid,
				0x00, 0x00, 0x00, 0x32,
			},
			want: OpTypeVoid(0x32),
		},
		{
			in: []byte{
				0x00, 0x02, 0x00, opTypeBool,
				0x00, 0x00, 0x00, 0x32,
			},
			want: OpTypeBool(0x32),
		},
		{
			in: []byte{
				0x00, 0x04, 0x00, opTypeInt,
				0x00, 0x00, 0x00, 0x32,
				0x00, 0x00, 0x00, 0x40,
				0x00, 0x00, 0x00, 0x01,
			},
			want: &OpTypeInt{
				Result:     0x32,
				Width:      64,
				Signedness: 1,
			},
		},
		{
			in: []byte{
				0x00, 0x03, 0x00, opTypeFloat,
				0x00, 0x00, 0x00, 0x32,
				0x00, 0x00, 0xf6, 0x42,
			},
			want: &OpTypeFloat{
				Result: 0x32,
				Width:  reverse(math.Float32bits(123)),
			},
		},
	} {
		wr := NewReader(bytes.NewBuffer(st.in))
		wr.SetEndian(LittleEndian)

		dec := NewDecoder(wr, DefaultInstructionSet)
		have, err := dec.DecodeInstruction()
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
			t.Fatalf("case %d: decode value mismatch: %v\nHave: %T(%v)\nWant: %T(%v)",
				i, st.in, have, have, st.want, st.want)
		}
	}
}

type HeaderTest struct {
	in   []byte
	want Header
	err  error
}

func TestDecodeHeaders(t *testing.T) {
	for _, st := range []HeaderTest{
		{
			in:  nil,
			err: ErrUnexpectedEOF,
		},
		{
			in: []byte{
				0x01, 0x02, 0x03, 0x04,
			},
			err: ErrInvalidMagicValue,
		},
		{
			in: []byte{
				0x03, 0x02, 0x23, 0x07,
				0x00, 0x00, 0x00, 0x63,
			},
			err: ErrUnexpectedEOF,
		},
		{
			in: []byte{
				0x03, 0x02, 0x23, 0x07,
				0x00, 0x00, 0x00, 0x64,
				0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00,
			},
			err: ErrUnsupportedModuleVersion,
		},
		{
			in: []byte{
				0x03, 0x02, 0x23, 0x07,
				0x00, 0x00, 0x00, 0x63,
				0x00, 0x00, 0x00, 0x01,
				0x00, 0x00, 0x00, 0x10,
				0x00, 0x00, 0x00, 0x00,
			},
			want: Header{
				Magic:     MagicLE,
				Version:   99,
				Generator: 1,
				Bound:     16,
				Reserved:  0,
			},
		},
		{
			in: []byte{
				0x07, 0x23, 0x02, 0x03,
				0x63, 0x00, 0x00, 0x00,
				0x01, 0x00, 0x00, 0x00,
				0x10, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00,
			},
			want: Header{
				Magic:     MagicBE,
				Version:   99,
				Generator: 1,
				Bound:     16,
				Reserved:  0,
			},
		},
	} {
		var have Header

		wr := NewReader(bytes.NewBuffer(st.in))
		dec := NewDecoder(wr, DefaultInstructionSet)

		err := dec.DecodeHeader(&have)
		if err != nil {
			if !reflect.DeepEqual(err, st.err) {
				t.Fatalf("decode error mismatch: %v\nHave: %v\nWant: %v",
					st.in, err, st.err)
			}

			continue
		}

		if !reflect.DeepEqual(have, st.want) {
			t.Fatalf("decode value mismatch: %v\nHave: %v\nWant: %v",
				st.in, have, st.want)
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
