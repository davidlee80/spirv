// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import (
	"fmt"
	"reflect"
	"testing"
)

type InstructionTest struct {
	in   []uint32
	want Instruction
	err  error
}

func testInstruction(t *testing.T, st InstructionTest) {
	have, err := Decode(st.in)

	// We have a decoding error. This is only a test failure if
	// we were not expecting an error.
	if err != nil {
		if !reflect.DeepEqual(err, st.err) {
			t.Fatalf("decode error mismatch: %v\nHave: %v\nWant: %v",
				st.in, err, st.err)
		}

		// We got an error as we expected, so no further processing
		// is needed.
		return
	}

	if !reflect.DeepEqual(have, st.want) {
		t.Fatalf("decode value mismatch: %v\nHave: %T(%+v)\nWant: %T(%+v)",
			st.in, have, have, st.want, st.want)
	}

	data, err := Encode(have)
	if !reflect.DeepEqual(data, st.in) {
		t.Fatalf("encode mismatch: %T(%v)\nHave: %v\nWant: %v",
			have, have, data, st.in)
	}
}

func TestCodec(t *testing.T) {
	for _, st := range []InstructionTest{
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
	} {
		testInstruction(t, st)
	}
}
