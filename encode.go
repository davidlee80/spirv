// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import (
	"fmt"
	"io"
)

// Encoder defines an encoder for the SPIR-V format.
// It writes SPIR-V sequences of words into a binary stream.
type Encoder struct {
	w      io.Writer
	buf    []uint32
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
		h.GeneratorMagic,
		h.Bound,
		h.Reserved,
	})
}

// EncodeInstructionWords writes the SPIR-V encoding of the given instruction
// to the underlying stream. The first word in the set defines the opcode
// and word count. The word count must be <= len(data).
func (e *Encoder) EncodeInstructionWords(data []uint32) error {
	if len(data) == 0 {
		return nil
	}

	size := int(data[0] >> 16)
	if len(data) < size {
		return ErrInvalidInstructionSize
	}

	return e.write(data[:size])
}

// Encode encodes the given instruction into a list of words and
// writes them to the underlying stream.
func (e *Encoder) EncodeInstruction(i Instruction) error {
	instructions.RLock()
	defer instructions.RUnlock()

	opcode := i.Opcode()
	codec, ok := instructions.data[opcode]
	if !ok {
		return fmt.Errorf("unknown instruction: %08x", opcode)
	}

	// Make sure the word buffer has the correct size.
	size := EncodedLen(i)
	if size > len(e.buf) {
		e.buf = make([]uint32, size)
	}

	// Encode the instruction arguments.
	argc, err := codec.Encode(i, e.buf[1:])
	if err != nil {
		return err
	}

	argc++

	// Set the first instruction word.
	e.buf[0] = EncodeOpcode(argc, i.Opcode())
	return e.EncodeInstructionWords(e.buf[:argc])
}

// Write writes exactly len(p) words to the underlying stream.
// It returns an error if this failed.
func (e *Encoder) write(p []uint32) error {
	var buf [4]byte

	for _, word := range p {
		if e.endian == LittleEndian {
			buf[0] = byte(word)
			buf[1] = byte(word >> 8)
			buf[2] = byte(word >> 16)
			buf[3] = byte(word >> 24)
		} else {
			buf[0] = byte(word >> 24)
			buf[1] = byte(word >> 16)
			buf[2] = byte(word >> 8)
			buf[3] = byte(word)
		}

		_, err := e.w.Write(buf[:])
		if err != nil {
			return err
		}
	}

	return nil
}
