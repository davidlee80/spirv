// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import (
	"fmt"
	"io"
)

// Encoder defines an encoder for the SPIR-V format.
// It writes SPIR-V data structures into a binary stream.
type Encoder struct {
	lib    InstructionSet
	w      io.Writer
	endian Endian
	buf    [4]byte
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

	// Fetch encoded instruction operands.
	argv, err := codec.Encode(i)
	if err != nil {
		return err
	}

	// Write opcode and word count.
	err = e.write(uint32(len(argv)+1)<<16 | opcode&0xffff)
	if err != nil {
		return err
	}

	// Write operands.
	return e.write(argv...)
}

// Write writes exactly len(p) words to the underlying stream.
// It returns an error if this failed.
func (e *Encoder) write(p ...uint32) error {
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
