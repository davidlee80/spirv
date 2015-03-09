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
	buf    [4]byte
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

// Write writes exactly len(p) words to the underlying stream.
// It returns an error if this failed.
func (w *Writer) Write(p []uint32) error {
	for _, word := range p {
		if w.endian == LittleEndian {
			w.buf[0] = byte(word)
			w.buf[1] = byte(word >> 8)
			w.buf[2] = byte(word >> 16)
			w.buf[3] = byte(word >> 24)
		} else {
			w.buf[0] = byte(word >> 24)
			w.buf[1] = byte(word >> 16)
			w.buf[2] = byte(word >> 8)
			w.buf[3] = byte(word)
		}

		_, err := w.w.Write(w.buf[:])
		if err != nil {
			return err
		}
	}

	return nil
}
