// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// SPIR-V Magic value as stored in a module header.
// The byte order of these lets us determine how to read the rest
// of the word stream.
const (
	MagicLE = 0x03022307
	MagicBE = 0x07230203
)

// Version defines the SPIR-V specification for which this package
// was written. The decoder will refuse to read modules with a version >Version.
const Version = 99

// Header defines the header of a SPIR-V Module.
type Header struct {
	// SPIR-V Magic number: 0x07230203
	//
	// A module is defined as a stream of words, not a stream of bytes.
	// However, if stored as a stream of bytes (e.g., in a file), the magic
	// number can be used to deduce what endianness to apply to convert the
	// byte stream back to a word stream.
	Magic uint32

	// Version number -- The first public version will be 100.
	// Uses 99 for pre-release.
	Version uint32

	// Version number of the tool which generated the module.
	// Its value does not effect any semantics, and is allowed to be 0.
	// Using a non-0 value is encouraged, and can be registered with Khronos.
	Generator uint32

	// All Ids in this module are guaranteed to satisfiy: 0 < id < Bound
	// Bound should be small; smaller is better with all <id> in a module
	// being densely packed and near 0.
	Bound uint32

	// 0 (Reserved for instruction schema, if needed.)
	Reserved uint32
}
