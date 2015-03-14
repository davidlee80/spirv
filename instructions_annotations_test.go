// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "testing"

func TestAnnotations(t *testing.T) {
	for _, st := range []InstructionTest{
		{
			in: []uint32{0x00020031, 0xff},
			want: &OpDecorationGroup{
				ResultId: 0xff,
			},
		},
		{
			in: []uint32{0x00060032, 1, DecorationNoStaticUse, 2, 3, 4},
			want: &OpDecorate{
				Target:     1,
				Decoration: DecorationNoStaticUse,
				Argv:       []uint32{2, 3, 4},
			},
		},
		{
			in: []uint32{0x00070033, 1, 2, DecorationNoStaticUse, 3, 4, 5},
			want: &OpMemberDecorate{
				StructType: 1,
				Member:     2,
				Decoration: DecorationNoStaticUse,
				Argv:       []uint32{3, 4, 5},
			},
		},
		{
			in: []uint32{0x00060034, 1, 2, 3, 4, 5},
			want: &OpGroupDecorate{
				Group:   1,
				Targets: []Id{2, 3, 4, 5},
			},
		},
		{
			in: []uint32{0x00060035, 1, 2, 3, 4, 5},
			want: &OpGroupMemberDecorate{
				Group:   1,
				Targets: []Id{2, 3, 4, 5},
			},
		},
	} {
		testInstruction(t, st)
	}
}
