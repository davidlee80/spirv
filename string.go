// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "math"

// EncodedStringLen returns the number of words occupied by the given
// string, once encoded.
func EncodedStringLen(v string) int {
	if len(v) == 0 {
		// An empty string still needs a nul terminator.
		return 1
	}

	return int(math.Ceil(float64(len(v)+1) / 4))
}

// EncodeString converts string v to a SPIR-V string literal and writes
// the result into out. Out must have enough space to hold the entire string.
// To provide a buffer with the correct size, EncodedStringLen() may be
// used to initialize it.
//
// A SPIR-V string literal is a nul-terminated stream of characters consuming
// an integral number of words. The character set is Unicode in the UTF-8
// encoding scheme.
//
// The UTF-8 octets (8-bit bytes) are packed four per word, following the
// little-endian convention (i.e., the first octet is in the lowest-order
// 8-bits of the word). The final word contains the string’s nul-termination
// character (0), and all contents past the end of the string in the final
// word are padded with 0.
//
// It returns the number of words written.
func EncodeString(v string, out []uint32) int {
	if len(v) == 0 {
		out[0] = 0
		return 1
	}

	var index int

	// Write whole blocks from string.
	for len(v) >= 4 {
		out[index] = uint32(v[0]) | uint32(v[1])<<8 |
			uint32(v[2])<<16 | uint32(v[3])<<24
		index++
		v = v[4:]
	}

	// Write non-whole tail block with nul-terminator.
	// If the tail is empty, it simply consists of the nul-terminator.
	word := uint32(0)

	switch len(v) {
	case 3:
		word |= uint32(v[2]) << 16
		fallthrough
	case 2:
		word |= uint32(v[1]) << 8
		fallthrough
	case 1:
		word |= uint32(v[0])
	}

	out[index] = word
	return index + 1
}

// DecodeString reads a Go string from the given slice of words.
// Refer to fromGoString() documentation for details on the expected encoding.
func DecodeString(words []uint32) string {
	out := make([]byte, 0, len(words)/4)

	// Read words as bytes.
	for _, w := range words {
		out = append(
			out,
			byte(w),
			byte(w>>8),
			byte(w>>16),
			byte(w>>24),
		)
	}

	// Remove any trailing nul bytes.
	sz := len(out) - 1
	for ; sz >= 0; sz-- {
		if out[sz] != 0 {
			sz++
			break
		}
	}

	if sz <= 0 {
		return ""
	}

	return string(out[:sz])
}
