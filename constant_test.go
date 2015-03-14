// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "testing"

type BitTest struct {
	in   uint32
	low  uint32
	high uint32
	want bool
}

func TestBitflag(t *testing.T) {
	for _, bt := range []BitTest{
		{0, 0, 0, true},
		{0, 2, 0, false},
		{0, 2, 4, false},
		{2, 0, 4, true},
		{5, 0, 4, true},
	} {
		testBitflag(t, bt.in, bt.low, bt.high, bt.want)
	}
}

func testBitflag(t *testing.T, in, low, high uint32, want bool) {
	have := verifyBitFlag(in, low, high)
	if want != have {
		t.Fatalf("%x (%x - %x) mismatch: Want %v, Have %v",
			in, low, high, want, have)
	}
}
