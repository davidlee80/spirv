// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpExtension defines the OpExtension instruction.
//
// It declares use of an extension to SPIR-V. This allows validation of
// additional instructions, tokens, semantics, etc
type OpExtension struct {
	Name String
}

func (c *OpExtension) Opcode() uint32 { return 3 }
func (c *OpExtension) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpExtension{}
	})
}

// OpExtInstImport defines the OpExtInstImport instruction.
//
// It defines an import of an extended set of instructions.
// It can later be referenced by the Result <id>
type OpExtInstImport struct {
	ResultId uint32
	Name     String
}

func (c *OpExtInstImport) Opcode() uint32 { return 4 }
func (c *OpExtInstImport) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpExtInstImport{}
	})
}

// OpExtInst defines an instruction in an imported set of extended instructions.
type OpExtInst struct {
	ResultType  uint32
	ResultId    uint32
	Set         uint32   // Result of an OpExtInstImport instruction.
	Instruction uint32   // Enumerant of the instruction to execute within the extended instruction Set.
	Operands    []uint32 // Operands to the extended instruction.
}

func (c *OpExtInst) Opcode() uint32 { return 44 }
func (c *OpExtInst) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpExtInst{}
	})
}
