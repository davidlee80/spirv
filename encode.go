// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "io"

// Encoder defines an encoder for the SPIR-V format.
// It writes a SPIR-V data structure into a binary stream.
type Encoder struct {
	w   *Writer
	lib InstructionSet
}

// NewEncoder creates a new encoder for the given stream and
// instruction set.
func NewEncoder(w io.Writer, lib InstructionSet) *Encoder {
	return &Encoder{
		w:   NewWriter(w),
		lib: lib,
	}
}

// Encode writes the SPIR-V encoding of module m to the underlying stream.
func (e *Encoder) Encode(m *Module) error {
	return nil
}

// EncodeHeader writes the SPIR-V encoding of header h to the
// underlying stream.
func (e *Encoder) EncodeHeader(h *Header) error {
	return nil
}

// EncodeInstruction writes the SPIR-V encoding of the given instruction
// to the underlying stream.
func (e *Encoder) EncodeInstruction(instr interface{}) error {
	return nil
}
