// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "fmt"

// Source defines the OpSource instruction.
//
// It documents what source language a module was translated from.
// This has no semantic impact and can safely be removed from a module.
type Source struct {
	Language SourceLanguage
	Version  uint32
}

func init() {
	DefaultInstructionSet[OpSource] = InstructionBuilder{
		Decode: decodeOpSource,
		Encode: encodeOpSource,
	}
}

func decodeOpSource(argv []uint32) (Instruction, error) {
	if len(argv) < 2 {
		return nil, ErrMissingInstructionArgs
	}

	return &Source{
		Language: SourceLanguage(argv[0]),
		Version:  argv[1],
	}, nil
}

func encodeOpSource(instr Instruction) ([]uint32, error) {
	src := instr.(*Source)
	return []uint32{
		uint32(src.Language),
		src.Version,
	}, nil
}

// SourceLanguage defines a source language constant.
type SourceLanguage uint32

// Known source languages.
const (
	SourceUnknown SourceLanguage = 0
	SourceESSL    SourceLanguage = 1
	SourceGLSL    SourceLanguage = 2
	SourceOpenCL  SourceLanguage = 3
)

func (sl SourceLanguage) String() string {
	switch sl {
	case SourceUnknown:
		return "Unknown"
	case SourceESSL:
		return "ESSL"
	case SourceGLSL:
		return "GLSL"
	case SourceOpenCL:
		return "OpenCL"
	}
	return fmt.Sprintf("SourceLanguage(%d)", uint32(sl))
}
