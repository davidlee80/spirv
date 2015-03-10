// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "io"

// Encoder defines an encoder for the SPIR-V format.
// It writes SPIR-V sequences of words into a binary stream.
type Encoder struct {
	w      io.Writer
	buf    [4]byte
	endian Endian
}

// NewEncoder creates a new encoder for the given stream and
// instruction set.
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		w:      w,
		endian: LittleEndian,
	}
}

// EncodeHeader writes the SPIR-V encoding of header h to the
// underlying stream. The magic value (first element in the set), determines
// the byte order for all remaining data being written in this- and subsequent
// calls to the Encoder.
func (e *Encoder) EncodeHeader(h Header) error {
	// The magic value should be written byte-for-byte, regardless
	// of the endianess. Infact, its byte order defines the byte order
	// for the remaining stream.
	switch h.Magic {
	case MagicLE:
		e.endian = LittleEndian
	case MagicBE:
		e.endian = BigEndian
	default:
		return ErrInvalidMagicValue
	}

	_, err := e.w.Write([]byte{
		byte(h.Magic),
		byte(h.Magic >> 8),
		byte(h.Magic >> 16),
		byte(h.Magic >> 24),
	})

	if err != nil {
		return err
	}

	if h.Version != 99 {
		return ErrInvalidVersion
	}

	return e.write([]uint32{
		h.Version,
		h.Generator,
		h.Bound,
		h.Reserved,
	})
}

// EncodeInstruction writes the SPIR-V encoding of the given instruction
// to the underlying stream. The first word in the set defines the opcode
// and word count. The word count must be <= len(data).
func (e *Encoder) EncodeInstruction(data []uint32) error {
	if len(data) == 0 {
		return nil
	}

	size := int(data[0] >> 16)
	if len(data) < size {
		return ErrInvalidInstructionSize
	}

	return e.write(data[:size])
}

// Write writes exactly len(p) words to the underlying stream.
// It returns an error if this failed.
func (e *Encoder) write(p []uint32) error {
	for _, word := range p {
		if e.endian == LittleEndian {
			e.buf[0] = byte(word)
			e.buf[1] = byte(word >> 8)
			e.buf[2] = byte(word >> 16)
			e.buf[3] = byte(word >> 24)
		} else {
			e.buf[0] = byte(word >> 24)
			e.buf[1] = byte(word >> 16)
			e.buf[2] = byte(word >> 8)
			e.buf[3] = byte(word)
		}

		_, err := e.w.Write(e.buf[:])
		if err != nil {
			return err
		}
	}

	return nil
}
