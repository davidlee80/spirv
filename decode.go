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
	r      io.Reader
	endian Endian
	lib    InstructionSet // Collection of known instruction decoders.
	ubuf   [16]uint32     // Scratch buffer for instruction decoding.
	bbuf   [4]byte        // Scratch buffer for the word reader.
}

// NewDecoder creates a new decoder for the given stream and instruction set.
func NewDecoder(r io.Reader, lib InstructionSet) *Decoder {
	return &Decoder{
		r:      r,
		lib:    lib,
		endian: LittleEndian,
	}
}

// SetEndian sets the endianess for the input stream.
// This defaults to LittleEndian.
func (d *Decoder) SetEndian(e Endian) { d.endian = e }

// DecodeModule reads a SPIR-V module from the input stream.
func (d *Decoder) DecodeModule() (*Module, error) {
	var mod Module
	mod.Code = make([]Instruction, 0, 128)

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
	// Read the magic value. This is the one value we
	// read as separate bytes. The order will tell us the byte
	// order of the rest of the stream.
	_, err := io.ReadFull(d.r, d.bbuf[:])
	if err != nil {
		if err == io.EOF {
			return ErrUnexpectedEOF
		}
		return err
	}

	h.Magic = uint32(d.bbuf[0]) | uint32(d.bbuf[1])<<8 | uint32(d.bbuf[2])<<16 |
		uint32(d.bbuf[3])<<24

	// Make sure it's a valid number. The order of the magic bytes lets us
	// determine the endianness of the stream's data.
	switch h.Magic {
	case MagicLE:
		d.endian = LittleEndian
	case MagicBE:
		d.endian = BigEndian
	default:
		return ErrInvalidMagicValue
	}

	// Read remaining header.
	err = d.read(d.ubuf[:4])
	if err != nil {
		return err
	}

	h.Version = d.ubuf[0]

	// Make sure we have a suitable version number.
	if h.Version > Version {
		return ErrUnsupportedModuleVersion
	}

	h.GeneratorMagic = d.ubuf[1]
	h.Bound = d.ubuf[2]
	h.Reserved = d.ubuf[3]
	return nil
}

// DecodeInstruction decodes the next instruction from the given stream.
func (d *Decoder) DecodeInstruction() (Instruction, error) {
	// Read the first word: word count + opcode.
	err := d.read(d.ubuf[:1])
	if err != nil {
		return nil, err
	}

	wordCount := (d.ubuf[0] >> 16)
	opcode := d.ubuf[0] & 0xffff

	if wordCount < 1 {
		return nil, ErrInvalidInstructionSize
	}

	var argv []uint32

	if wordCount > 1 {
		// wordCount defines the number of words for the entire instruction.
		// This includes the first one we just read. So remaining number of
		// operands are wordCount - 1.
		err = d.read(d.ubuf[:wordCount-1])
		if err != nil {
			return nil, err
		}

		argv = d.ubuf[:wordCount-1]
	}

	// Find the instruction-specific decoder and call it.
	codec, ok := d.lib.Get(opcode)
	if !ok {
		return nil, fmt.Errorf("unknown opcode: 0x%x", opcode)
	}

	return codec.Decode(argv)
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
