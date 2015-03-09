// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "io"

// Reader reads 32-bit words from an underlying stream.
//
// The stream's expected endianess defaults to LittleEndian but can be
// changed to BigEndian if applicable. When decoding an entire module,
// the endianess can be inferred from the module's magic number.
type Reader struct {
	r      io.Reader
	buf    [4]byte
	endian Endian
}

// NewReader creates a new word reader from the given stream.
func NewReader(r io.Reader) *Reader {
	return &Reader{
		r:      r,
		endian: LittleEndian,
	}
}

// SetEndian sets the endianess for the input stream.
// This defaults to LittleEndian.
func (r *Reader) SetEndian(e Endian) { r.endian = e }

// Next reads exactly len(p) words from the stream.
// Returns an error if there is either not enough data or something else
// went wrong.
func (r *Reader) Read(p []uint32) error {
	for i := range p {
		_, err := io.ReadFull(r.r, r.buf[:])
		if err != nil {
			if err == io.EOF {
				// This particular EOF is unexpected.
				// We should be able to read at least len(p) words.
				return ErrUnexpectedEOF
			}
			return err
		}

		if r.endian == LittleEndian {
			p[i] = uint32(r.buf[3]) | uint32(r.buf[2])<<8 | uint32(r.buf[1])<<16 |
				uint32(r.buf[0])<<24
		} else {
			p[i] = uint32(r.buf[0]) | uint32(r.buf[1])<<8 | uint32(r.buf[2])<<16 |
				uint32(r.buf[3])<<24
		}
	}

	return nil
}
