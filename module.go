// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "io"

// Module defines a complete SPIR-V module.
type Module struct {
	Header Header
	Code   []Instruction
}

// NewModule creates a new, default module.
func NewModule() *Module {
	return &Module{
		Header: Header{
			Magic:          MagicLE,
			Version:        SpecificationVersion,
			GeneratorMagic: 0,
			Bound:          0,
			Reserved:       0,
		},
	}
}

// Load loads a full module from the given input stream.
func Load(r io.Reader) (*Module, error) {
	var mod Module
	var err error
	var instr Instruction

	dec := NewDecoder(r)

	// Load the module header.
	mod.Header, err = dec.DecodeHeader()
	if err != nil {
		return nil, err
	}

	// Load all instructions.
	mod.Code = make([]Instruction, 0, 128)

	for {
		instr, err = dec.DecodeInstruction()
		if err != nil {
			if err == io.EOF {
				break // Not an error -- just end of stream.
			}

			return nil, err
		}

		mod.Code = append(mod.Code, instr)
	}

	return &mod, nil
}

// Save writes the module to the given stream.
func (m *Module) Save(w io.Writer) error {
	enc := NewEncoder(w)

	// Write the header.
	err := enc.EncodeHeader(m.Header)
	if err != nil {
		return err
	}

	// Write all instructions.
	for _, instr := range m.Code {
		err := enc.EncodeInstruction(instr)
		if err != nil {
			return err
		}
	}

	return nil
}

// Verify returns an error if the module contains invalid data.
// This level of verification applies semantic validation as defined in
// the specification.
func (m *Module) Verify() error {
	// Check the header for structural validity.
	err := m.Header.Verify()
	if err != nil {
		return err
	}

	// Perform structural validity checks on each instruction, before
	// we proceed to the semantic checks.
	//
	// This uses reflection to call Verify() on all relevant struct fields
	// and then on the instruction itself. The latter is used by some
	// instructions to validate parts which can not be caught by the field
	// types themselves.
	for _, instr := range m.Code {
		err := verifyInstruction(instr)
		if err != nil {
			return err
		}
	}

	// TODO Implement all semantic validation.
	err = m.verifyLogicalLayout()
	if err != nil {
		return err
	}

	return nil
}

// Strip removes all instructions which have no semantic impact on the code.
// This includes debug symbols like source context and names.
func (m *Module) Strip() {
	for i := 0; i < len(m.Code); i++ {
		if !m.Code[i].Optional() {
			continue
		}

		copy(m.Code[i:], m.Code[i+1:])
		m.Code[len(m.Code)-1] = nil
		m.Code = m.Code[:len(m.Code)-1]
		i--
	}
}

// verifyLogicalLayout ensures the module meets the Logical Layout
// requirements as defined in the spec chapter 2.4.
//
// FIXME: We're using the standard regexp package here and we treat the
// instruction set as a sequence of runes. We may want to investigate writing a
// custom regex parser; specifically for the limited subset of functionality we
// need and which operates on ints instead of strings.
// Refer to http://swtch.com/~rsc/regexp/regexp2.html for inspiration.
//
// FIXME: This approach can only tell us if the structure is valid or not.
// It will not give us any context as to which instructions are wrong when
// the check fails. This is not useful at all.
func (m *Module) verifyLogicalLayout() error {
	// We must have one and only one OpmemoryModel.
	//
	// This will be caught by the regex match below, but here we can be
	// more specific with our error message.
	if count(m.Code, opcodeMemoryModel) != 1 {
		return ErrMemoryModel
	}

	// We must have at least one OpEntryPoint.
	//
	// This will be caught by the regex match below, but here we can be
	// more specific with our error message.
	if count(m.Code, opcodeEntryPoint) == 0 {
		return ErrEntrypoint
	}

	// We must have at least one OpExecutionMode.
	//
	// This will be caught by the regex match below, but here we can be
	// more specific with our error message.
	if count(m.Code, opcodeExecutionMode) == 0 {
		return ErrExecutionMode
	}

	// Test instruction order.
	err := verifyLayoutPattern(m.Code)
	if err != nil {
		return err
	}

	// Some instructions have requirements beyond what can be tested
	// with a regular expression pattern.

	// Global Variables must not have StorageClassFunction.
	err = verifyGlobalVariables(m.Code)
	if err != nil {
		return err
	}

	// Local Variables must have StorageClassFunction.
	err = verifyLocalVariables(m.Code)
	if err != nil {
		return err
	}

	// All local variables must be the first instructions in the first block.
	return verifyFunctionStructure(m.Code)
}
