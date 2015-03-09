// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "io"

// Encoder defines an encoder for the SPIR-V format.
// It writes a SPIR-V data structure into a binary stream.
type Encoder struct {
	w *Writer
}

// NewEncoder creates a new encoder for the given stream.
func NewEncoder(w io.Writer) *Encoder { return &Encoder{w: NewWriter(w)} }

// Encode writes the SPIR-V encoding of module m to the underlying stream.
func (e *Encoder) Encode(m *Module) error {
	return nil
}
