// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import (
	"bytes"
	"reflect"
	"testing"
)

type InstructionTest struct {
	in   []byte
	want Instruction
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
				0x00, 0x00, 0x00, OpSource,
			},
			err: ErrInvalidInstructionSize,
		},
		{
			in: []byte{
				0x00, 0x01, 0x00, OpSource,
			},
			err: ErrMissingInstructionArgs,
		},
		{
			in: []byte{
				0x00, 0x03, 0x00, OpSource,
				0x00, 0x00, 0x00, byte(SourceGLSL),
				0x00, 0x00, 0x01, 0xc2,
			},
			want: &Source{
				Language: SourceGLSL,
				Version:  450,
			},
		},
		{
			in: []byte{
				0x00, 0x07, 0x00, OpSourceExtension,
				0x74, 0x73, 0x65, 0x74,
				0x75, 0x6f, 0x73, 0x20,
				0x20, 0x65, 0x63, 0x72,
				0x65, 0x74, 0x78, 0x65,
				0x6f, 0x69, 0x73, 0x6e,
				0x00, 0x00, 0x00, 0x6e,
			},
			want: SourceExtension("test source extension"),
		},
		{
			in: []byte{
				0x00, 0x05, 0x00, OpExtension,
				0x74, 0x73, 0x65, 0x74,
				0x74, 0x78, 0x65, 0x20,
				0x69, 0x73, 0x6e, 0x65,
				0x00, 0x00, 0x6e, 0x6f,
			},
			want: Extension("test extension"),
		},
		{
			in: []byte{
				0x00, 0x06, 0x00, OpExtInstImport,
				0x00, 0x00, 0x00, 0x23,
				0x74, 0x73, 0x65, 0x74,
				0x74, 0x78, 0x65, 0x20,
				0x69, 0x73, 0x6e, 0x65,
				0x00, 0x00, 0x6e, 0x6f,
			},
			want: &ExtInstImport{
				ResultId: 0x23,
				Name:     "test extension",
			},
		},
		{
			in: []byte{
				0x00, 0x08, 0x00, OpExtInst,
				0x00, 0x00, 0x00, 0x01,
				0x00, 0x00, 0x00, 0x02,
				0x00, 0x00, 0x00, 0x03,
				0x00, 0x00, 0x00, 0x04,
				0x00, 0x00, 0x00, 0x05,
				0x00, 0x00, 0x00, 0x06,
				0x00, 0x00, 0x00, 0x07,
			},
			want: &ExtInst{
				ResultType:  0x01,
				ResultId:    0x02,
				Set:         0x03,
				Instruction: 0x04,
				Operands:    []uint32{0x05, 0x06, 0x07},
			},
		},
		{
			in: []byte{
				0x00, 0x03, 0x00, OpMemoryModel,
				0x00, 0x00, 0x00, byte(AddressPhysical32),
				0x00, 0x00, 0x00, byte(MemoryGLSL450),
			},
			want: &MemoryModel{
				Addressing: AddressPhysical32,
				Memory:     MemoryGLSL450,
			},
		},
		{
			in: []byte{
				0x00, 0x03, 0x00, OpEntryPoint,
				0x00, 0x00, 0x00, byte(ExecFragment),
				0x00, 0x00, 0x00, 0x7f,
			},
			want: &EntryPoint{
				Execution: ExecFragment,
				Id:        0x7f,
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
