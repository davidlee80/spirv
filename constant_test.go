// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "testing"

type BitTest struct {
	in   uint32
	none bool
	mask uint32
	want bool
}

func TestBitflag(t *testing.T) {
	for _, bt := range []BitTest{
		{0, true, 0, true},
		{0, false, 0, false},
		{0, true, 7, true},
		{2, true, 7, true},
		{5, true, 7, true},
	} {
		testBitflag(t, bt.in, bt.none, bt.mask, bt.want)
	}
}

func testBitflag(t *testing.T, in uint32, none bool, mask uint32, want bool) {
	have := verifyBitFlag(in, none, mask)
	if want != have {
		t.Fatalf("mismatch: in: %x, none: %v, mask: %x, want: %v, have: %v",
			in, none, mask, want, have)
	}
}
