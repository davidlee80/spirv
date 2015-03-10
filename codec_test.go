// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import (
	"bytes"
	"math"
	"reflect"
	"testing"
)

var lib InstructionSet

func init() {
	BindDefault(&lib)
}

type InstructionTest struct {
	in   []byte
	want interface{}
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
			in: []byte{
				0x01, 0x00, 0x00, 0x00,
			},
			err: ErrInvalidInstructionSize,
		},
		{
			in: []byte{
				0x01, 0x00, 0x01, 0x00,
			},
			err: ErrMissingInstructionArgs,
		},
		{
			in: []byte{
				0x01, 0x00, 0x03, 0x00,
				byte(SourceGLSL), 0x00, 0x00, 0x00,
				0xc2, 0x01, 0x00, 0x00,
			},
			want: &OpSource{
				Language: SourceGLSL,
				Version:  450,
			},
		},
		{
			in: []byte{
				0x02, 0x00, 0x07, 0x00,
				0x74, 0x65, 0x73, 0x74,
				0x20, 0x73, 0x6f, 0x75,
				0x72, 0x63, 0x65, 0x20,
				0x65, 0x78, 0x74, 0x65,
				0x6e, 0x73, 0x69, 0x6f,
				0x6e, 0x00, 0x00, 0x00,
			},
			want: OpSourceExtension("test source extension"),
		},
		{
			in: []byte{
				0x03, 0x00, 0x05, 0x00,
				0x74, 0x65, 0x73, 0x74,
				0x20, 0x65, 0x78, 0x74,
				0x65, 0x6e, 0x73, 0x69,
				0x6f, 0x6e, 0x00, 0x00,
			},
			want: OpExtension("test extension"),
		},
		{
			in: []byte{
				0x04, 0x00, 0x06, 0x00,
				0x23, 0x00, 0x00, 0x00,
				0x74, 0x65, 0x73, 0x74,
				0x20, 0x65, 0x78, 0x74,
				0x65, 0x6e, 0x73, 0x69,
				0x6f, 0x6e, 0x00, 0x00,
			},
			want: &OpExtInstImport{
				ResultId: 0x23,
				Name:     "test extension",
			},
		},
		{
			in: []byte{
				0x2c, 0x00, 0x08, 0x00,
				0x01, 0x00, 0x00, 0x00,
				0x02, 0x00, 0x00, 0x00,
				0x03, 0x00, 0x00, 0x00,
				0x04, 0x00, 0x00, 0x00,
				0x05, 0x00, 0x00, 0x00,
				0x06, 0x00, 0x00, 0x00,
				0x07, 0x00, 0x00, 0x00,
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
				0x05, 0x00, 0x03, 0x00,
				byte(AddressPhysical32), 0x00, 0x00, 0x00,
				byte(MemoryGLSL450), 0x00, 0x00, 0x00,
			},
			want: &OpMemoryModel{
				Addressing: AddressPhysical32,
				Memory:     MemoryGLSL450,
			},
		},
		{
			in: []byte{
				0x06, 0x00, 0x03, 0x00,
				byte(ExecFragment), 0x00, 0x00, 0x00,
				0x7f, 0x00, 0x00, 0x00,
			},
			want: &OpEntryPoint{
				Execution: ExecFragment,
				Id:        0x7f,
			},
		},
		{
			in: []byte{
				0x07, 0x00, 0x06, 0x00,
				0x7f, 0x00, 0x00, 0x00,
				byte(ModeSpacingEqual), 0x00, 0x00, 0x00,
				0x01, 0x00, 0x00, 0x00,
				0x02, 0x00, 0x00, 0x00,
				0x03, 0x00, 0x00, 0x00,
			},
			want: &OpExecutionMode{
				EntryPoint: 0x7f,
				Mode:       ModeSpacingEqual,
				Argv:       []uint32{0x01, 0x02, 0x03},
			},
		},
		{
			in: []byte{
				0xda, 0x00, 0x05, 0x00,
				0x74, 0x65, 0x73, 0x74,
				0x20, 0x65, 0x78, 0x74,
				0x65, 0x6e, 0x73, 0x69,
				0x6f, 0x6e, 0x00, 0x00,
			},
			want: OpCompileFlag("test extension"),
		},
		{
			in: []byte{
				0x08, 0x00, 0x02, 0x00,
				0x32, 0x00, 0x00, 0x00,
			},
			want: OpTypeVoid(0x32),
		},
		{
			in: []byte{
				0x09, 0x00, 0x02, 0x00,
				0x32, 0x00, 0x00, 0x00,
			},
			want: OpTypeBool(0x32),
		},
		{
			in: []byte{
				0x0a, 0x00, 0x04, 0x00,
				0x32, 0x00, 0x00, 0x00,
				0x40, 0x00, 0x00, 0x00,
				0x01, 0x00, 0x00, 0x00,
			},
			want: &OpTypeInt{
				Result:     0x32,
				Width:      64,
				Signedness: 1,
			},
		},
		{
			in: []byte{
				0x0b, 0x00, 0x03, 0x00,
				0x32, 0x00, 0x00, 0x00,
				0x42, 0xf6, 0x00, 0x00,
			},
			want: &OpTypeFloat{
				Result: 0x32,
				Width:  reverse(math.Float32bits(123)),
			},
		},
		{
			in: []byte{
				0x0c, 0x00, 0x04, 0x00,
				0x32, 0x00, 0x00, 0x00,
				0x12, 0x00, 0x00, 0x00,
				0x03, 0x00, 0x00, 0x00,
			},
			want: &OpTypeVector{
				Result:         0x32,
				ComponentType:  0x12,
				ComponentCount: 0x03,
			},
		},
	} {
		// Test the decoder. Its output must match the structure
		// defined in the test case.
		dec := NewDecoder(bytes.NewBuffer(st.in), lib)
		have, err := dec.DecodeInstruction()

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
			t.Fatalf("case %d: decode value mismatch: %v\nHave: %T(%v)\nWant: %T(%v)",
				i, st.in, have, have, st.want, st.want)
		}

		// Test encoding roundtrip.
		// Its output must match the test case input.
		var buf bytes.Buffer

		enc := NewEncoder(&buf, lib)
		err = enc.EncodeInstruction(have)
		if err != nil {
			t.Fatalf("case %d: encode error: %T(%v)\n%v",
				i, have, have, err)
		}

		if !bytes.Equal(buf.Bytes(), st.in) {
			t.Fatalf("case %d: encode mismatch: %T(%v)\nHave: %v\nWant: %v",
				i, have, have, buf.Bytes(), st.in)
		}
	}
}

type HeaderTest struct {
	in   []byte
	want Header
	err  error
}

func TestCodecHeaders(t *testing.T) {
	for i, st := range []HeaderTest{
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
				0x63, 0x00, 0x00, 0x00,
				0x01, 0x00, 0x00, 0x00,
				0x10, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00,
			},
			want: Header{
				Magic:          MagicLE,
				Version:        99,
				GeneratorMagic: 1,
				Bound:          16,
				Reserved:       0,
			},
		},
		{
			in: []byte{
				0x07, 0x23, 0x02, 0x03,
				0x00, 0x00, 0x00, 0x63,
				0x00, 0x00, 0x00, 0x01,
				0x00, 0x00, 0x00, 0x10,
				0x00, 0x00, 0x00, 0x00,
			},
			want: Header{
				Magic:          MagicBE,
				Version:        99,
				GeneratorMagic: 1,
				Bound:          16,
				Reserved:       0,
			},
		},
	} {
		var have Header

		dec := NewDecoder(bytes.NewBuffer(st.in), lib)
		err := dec.DecodeHeader(&have)
		if err != nil {
			if !reflect.DeepEqual(err, st.err) {
				t.Fatalf("case %d: decode error mismatch: %v\nHave: %v\nWant: %v",
					i, st.in, err, st.err)
			}

			continue
		}

		if !reflect.DeepEqual(have, st.want) {
			t.Fatalf("case %d: decode value mismatch: %v\nHave: %v\nWant: %v",
				i, st.in, have, st.want)
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
