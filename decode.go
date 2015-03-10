// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import (
	"errors"
	"io"
)

var (
	ErrUnexpectedEOF          = errors.New("unexpected EOF")
	ErrInvalidMagicValue      = errors.New("invalid magic value")
	ErrInvalidInstructionSize = errors.New("instruction has invalid size")
)

// Decoder defines a decoder for the SPIR-V format.
// It reads binary data from a stream and yields sequences
// of 32-bit words.
type Decoder struct {
	r      io.Reader
	ubuf   []uint32 // Scratch buffer for instruction decoding.
	bbuf   [4]byte  // Scratch buffer for the word reader.
	endian Endian
}

// NewDecoder creates a new decoder for the given stream and instruction set.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		r:      r,
		endian: LittleEndian,
		ubuf:   make([]uint32, 32),
	}
}

// DecodeHeader reads a module header from the underlying stream.
//
// The magic value's byte order will be used to determine the byte order
// for the entire stream.
//
// Returns an error if the magic value is invalid.
func (d *Decoder) DecodeHeader() (Header, error) {
	var hdr Header

	// Read the magic value. This is the one value we read as separate bytes.
	// The order will tell us the byte order of the rest of the stream.
	_, err := io.ReadFull(d.r, d.bbuf[:])
	if err != nil {
		if err == io.EOF {
			return hdr, ErrUnexpectedEOF
		}
		return hdr, err
	}

	hdr.Magic = uint32(d.bbuf[0]) | uint32(d.bbuf[1])<<8 |
		uint32(d.bbuf[2])<<16 | uint32(d.bbuf[3])<<24

	// Make sure it's a valid number. The order of the magic bytes lets us
	// determine the endianness of the stream's remaining data.
	switch hdr.Magic {
	case MagicLE:
		d.endian = LittleEndian
	case MagicBE:
		d.endian = BigEndian
	default:
		return hdr, ErrInvalidMagicValue
	}

	// Read remaining header.
	err = d.read(d.ubuf[:4])
	if err != nil {
		return hdr, nil
	}

	hdr.Version = d.ubuf[0]
	hdr.Generator = d.ubuf[1]
	hdr.Bound = d.ubuf[2]
	hdr.Reserved = d.ubuf[3]
	return hdr, nil
}

// DecodeInstruction decodes the next instruction from the underlying stream.
// The returned slice of words contains all data for the entire instruction.
//
// The data remains valid until the next call to DecodeHeader or
// DecodeInstruction.
func (d *Decoder) DecodeInstruction() ([]uint32, error) {
	// Read the first word: word count + opcode.
	err := d.read(d.ubuf[:1])
	if err != nil {
		return nil, err
	}

	words := int(d.ubuf[0] >> 16)
	if words < 1 {
		return nil, ErrInvalidInstructionSize
	}

	if words > 1 {
		if words >= len(d.ubuf) {
			// Resize read buffer if necessary.
			tmp := d.ubuf[0]
			d.ubuf = make([]uint32, words)
			d.ubuf[0] = tmp
		}

		// words defines the number of words for the entire instruction.
		// This includes the first one we just read. So remaining number of
		// operands are words - 1.
		err = d.read(d.ubuf[1:words])
		if err != nil {
			return nil, err
		}
	}

	return d.ubuf[:words:words], nil
}

// Next reads exactly len(p) words from the stream.
// Returns an error if there is either not enough data or something else
// went wrong.
func (d *Decoder) read(p []uint32) error {
	for i := range p {
		_, err := io.ReadFull(d.r, d.bbuf[:])
		if err != nil {
			if err == io.EOF {
				// This particular EOF is unexpected.
				// We should be able to read at least len(p) words.
				return ErrUnexpectedEOF
			}
			return err
		}

		if d.endian == LittleEndian {
			p[i] = uint32(d.bbuf[0]) | uint32(d.bbuf[1])<<8 | uint32(d.bbuf[2])<<16 |
				uint32(d.bbuf[3])<<24
		} else {
			p[i] = uint32(d.bbuf[3]) | uint32(d.bbuf[2])<<8 | uint32(d.bbuf[1])<<16 |
				uint32(d.bbuf[0])<<24
		}
	}

	return nil
}
