// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import (
	"errors"
	"fmt"
	"io"
)

var (
	ErrUnexpectedEOF            = errors.New("unexpected EOF")
	ErrInvalidMagicValue        = errors.New("invalid magic value")
	ErrInvalidInstructionSize   = errors.New("instruction word count is zero")
	ErrMissingInstructionArgs   = errors.New("insufficient instruction arguments")
	ErrUnsupportedModuleVersion = fmt.Errorf("unsupported module version: >%d", Version)
)

// Decoder defines a decoder for the SPIR-V format.
// It reads binary data from a stream and creates a SPIR-V data structure.
type Decoder struct {
	r   *Reader        // Underlying input stream.
	lib InstructionSet // Collection of known instruction decoders.
	buf [16]uint32     // Scratch buffer for instruction decoding.
}

// NewDecoder creates a new decoder for the given stream and instruction set.
func NewDecoder(wr *Reader, lib InstructionSet) *Decoder {
	return &Decoder{r: wr, lib: lib}
}

// Decode reads a SPIR-V module from the input stream.
func (d *Decoder) Decode() (*Module, error) {
	var mod Module
	mod.Code = make([]interface{}, 0, 128)

	err := d.DecodeHeader(&mod.Header)
	if err != nil {
		return nil, err
	}

	// Decode all instructions for as long as we can find some.
	for {
		instr, err := d.DecodeInstruction()

		if err != nil {
			if err == io.EOF {
				break // Nothing wrong here -- just done reading.
			}
			return nil, err
		}

		mod.Code = append(mod.Code, instr)
	}

	return &mod, nil
}

// DecodeHeader reads the given header from the underlying stream.
//
// The magic value's byte order will be used to determine the byte order
// for the entire stream and will be stored in the decoder's WordReader.
//
// Returns an error if the magic value is invalid or the version number
// does not match our expectations.
func (d *Decoder) DecodeHeader(h *Header) error {
	err := d.r.Read(d.buf[:1])
	if err != nil {
		return err
	}

	h.Magic = d.buf[0]

	// Make sure it's a valid one. The order of the magic bytes lets us
	// determine the endianness of the stream's data.
	switch h.Magic {
	case MagicLE:
		d.r.SetEndian(LittleEndian)
	case MagicBE:
		d.r.SetEndian(BigEndian)
	default:
		return ErrInvalidMagicValue
	}

	// Read remaining header.
	err = d.r.Read(d.buf[:4])
	if err != nil {
		return err
	}

	h.Version = d.buf[0]

	// Make sure we have a suitable version number.
	if h.Version > Version {
		return ErrUnsupportedModuleVersion
	}

	h.GeneratorMagic = d.buf[1]
	h.Bound = d.buf[2]
	h.Reserved = d.buf[3]
	return nil
}

// DecodeInstruction decodes the next instruction from the given stream.
func (d *Decoder) DecodeInstruction() (interface{}, error) {
	// Read the first word: word count + opcode.
	err := d.r.Read(d.buf[:1])
	if err != nil {
		return nil, err
	}

	wordCount := (d.buf[0] >> 16)
	opcode := d.buf[0] & 0xffff

	if wordCount < 1 {
		return nil, ErrInvalidInstructionSize
	}

	var argv []uint32

	if wordCount > 1 {
		// wordCount defines the number of words for the entire instruction.
		// This includes the first one we just read. So remaining number of
		// operands are wordCount - 1.
		err = d.r.Read(d.buf[:wordCount-1])
		if err != nil {
			return nil, err
		}

		argv = d.buf[:wordCount-1]
	}

	// Find the instruction-specific decoder and call it.
	instr, ok := d.lib[opcode]
	if !ok {
		return nil, fmt.Errorf("unknown opcode: 0x%x", opcode)
	}

	return instr.Decode(argv)
}
