// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "io"

// Writer writes 32-bit words to an underlying stream.
//
// The stream's expected endianess defaults to LittleEndian but can be
// changed to BigEndian if applicable.
type Writer struct {
	w      io.Writer
	endian Endian
}

// NewWriter creates a new word writer from the given stream.
func NewWriter(w io.Writer) *Writer {
	return &Writer{
		w:      w,
		endian: LittleEndian,
	}
}

// SetEndian sets the endianess for the output stream.
// This defaults to LittleEndian.
func (w *Writer) SetEndian(e Endian) { w.endian = e }
