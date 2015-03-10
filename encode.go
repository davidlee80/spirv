// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import (
	"errors"
	"fmt"
	"io"
)

var (
	ErrZeroLengthInstruction = errors.New("encoded instruction has zero length")
)

// EncodeOpcode returns the first word of an instruction from the
// given word cound and opcode.
func EncodeOpcode(wordCount, opcode int) uint32 {
	return uint32(wordCount&0xffff)<<16 | uint32(opcode&0xffff)
}

// Encoder defines an encoder for the SPIR-V format.
// It writes SPIR-V data structures into a binary stream.
type Encoder struct {
	lib    InstructionSet
	w      io.Writer
	endian Endian
	ubuf   [64]uint32
	bbuf   [4]byte
}

// NewEncoder creates a new encoder for the given stream and
// instruction set.
func NewEncoder(w io.Writer, lib InstructionSet) *Encoder {
	return &Encoder{
		w:      w,
		lib:    lib,
		endian: LittleEndian,
	}
}

// SetEndian sets the endianess for the output stream.
// This defaults to LittleEndian.
func (e *Encoder) SetEndian(endian Endian) { e.endian = endian }

// EncodeModule writes the SPIR-V encoding of module m to the underlying stream.
func (e *Encoder) EncodeModule(m *Module) error {
	err := e.EncodeHeader(&m.Header)
	if err != nil {
		return err
	}

	for _, instr := range m.Code {
		err := e.EncodeInstruction(instr)
		if err != nil {
			return err
		}
	}

	return nil
}

// EncodeHeader writes the SPIR-V encoding of header h to the
// underlying stream.
func (e *Encoder) EncodeHeader(h *Header) error {
	return nil
}

// EncodeInstruction writes the SPIR-V encoding of the given instruction
// to the underlying stream.
func (e *Encoder) EncodeInstruction(i Instruction) error {
	opcode := i.Opcode()

	codec, ok := e.lib.Get(opcode)
	if !ok {
		return fmt.Errorf("unknown instruction %T", i)
	}

	// If the encoder fails to write a proper first word,
	// we want to know about it.
	e.ubuf[0] = 0

	// Fetch encoded instruction operands.
	err := codec.Encode(i, e.ubuf[:])
	if err != nil {
		return err
	}

	// Get the size of the instruction from the first word.
	wordCount := e.ubuf[0] >> 16
	if wordCount == 0 {
		return ErrZeroLengthInstruction
	}

	// Write the instruction.
	return e.write(e.ubuf[:wordCount])
}

// Write writes exactly len(p) words to the underlying stream.
// It returns an error if this failed.
func (e *Encoder) write(p []uint32) error {
	for _, word := range p {
		if e.endian == LittleEndian {
			e.bbuf[0] = byte(word)
			e.bbuf[1] = byte(word >> 8)
			e.bbuf[2] = byte(word >> 16)
			e.bbuf[3] = byte(word >> 24)
		} else {
			e.bbuf[0] = byte(word >> 24)
			e.bbuf[1] = byte(word >> 16)
			e.bbuf[2] = byte(word >> 8)
			e.bbuf[3] = byte(word)
		}

		_, err := e.w.Write(e.bbuf[:])
		if err != nil {
			return err
		}
	}

	return nil
}
