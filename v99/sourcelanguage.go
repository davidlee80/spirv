// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package v99

import "fmt"

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
