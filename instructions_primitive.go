// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpEmitVertex emits the current values of all output variables to the
// current output primitive. After execution, the values of all output
// variables are undefined.
type OpEmitVertex struct{}

func (c *OpEmitVertex) Opcode() uint32 { return 184 }
func (c *OpEmitVertex) Verify() error  { return nil }

// OpEndPrimitive finishes the current primitive and starts a new one.
// No vertex is emitted.
//
// This instruction can only be used when only one stream is present.
type OpEndPrimitive struct{}

func (c *OpEndPrimitive) Opcode() uint32 { return 185 }
func (c *OpEndPrimitive) Verify() error  { return nil }

// OpEmitStreamVertex emits the current values of all output variables to
// the current output primitive. After execution, the values of all output
// variables are undefined.
//
// This instruction can only be used when multiple streams are present
type OpEmitStreamVertex struct {
	Stream uint32
}

func (c *OpEmitStreamVertex) Opcode() uint32 { return 186 }
func (c *OpEmitStreamVertex) Verify() error  { return nil }

// OpEndStreamPrimitive finishes the current primitive and starts a new one.
// No vertex is emitted
//
// This instruction can only be used when multiple streams are present
type OpEndStreamPrimitive struct {
	Stream uint32
}

func (c *OpEndStreamPrimitive) Opcode() uint32 { return 187 }
func (c *OpEndStreamPrimitive) Verify() error  { return nil }

func init() {
	Bind(func() Instruction { return &OpEmitVertex{} })
	Bind(func() Instruction { return &OpEndPrimitive{} })
	Bind(func() Instruction { return &OpEmitStreamVertex{} })
	Bind(func() Instruction { return &OpEndStreamPrimitive{} })
}
