// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpDPdx is equivalent to either OpDPdxFine or OpDPdxCoarse on P.
// Selection of which one is based on external factors.
type OpDPdx struct {
	ResultType uint32
	ResultId   uint32
	P          uint32
}

func (c *OpDPdx) Opcode() uint32 { return 175 }
func (c *OpDPdx) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpDPdx{}
	})
}

// OpDPdy is equivalent to either OpDPdyFine or OpDPdyCoarse on P.
// Selection of which one is based on external factors.
type OpDPdy struct {
	ResultType uint32
	ResultId   uint32
	P          uint32
}

func (c *OpDPdy) Opcode() uint32 { return 176 }
func (c *OpDPdy) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpDPdy{}
	})
}

// OpFwidth is equivalent to computing the sum of the absolute values of
// OpDPdx and OpDPdy on P.
type OpFwidth struct {
	ResultType uint32
	ResultId   uint32
	P          uint32
}

func (c *OpFwidth) Opcode() uint32 { return 177 }
func (c *OpFwidth) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpFwidth{}
	})
}

// OpDPdxFine calculates the partial derivative of P with respect to the
// window x coordinate.
type OpDPdxFine struct {
	ResultType uint32
	ResultId   uint32
	P          uint32
}

func (c *OpDPdxFine) Opcode() uint32 { return 178 }
func (c *OpDPdxFine) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpDPdxFine{}
	})
}

// OpDPdyFine calculates the partial derivative of P with respect to the
// window y coordinate.
type OpDPdyFine struct {
	ResultType uint32
	ResultId   uint32
	P          uint32
}

func (c *OpDPdyFine) Opcode() uint32 { return 179 }
func (c *OpDPdyFine) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpDPdyFine{}
	})
}

// OpFwidthFine is equivalent to computing the sum of the absolute values
// of OpDPdxFine and OpDPdyFine on P.
type OpFwidthFine struct {
	ResultType uint32
	ResultId   uint32
	P          uint32
}

func (c *OpFwidthFine) Opcode() uint32 { return 180 }
func (c *OpFwidthFine) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpFwidthFine{}
	})
}

// OpDPdxCoarse calculates the partial derivative of P with respect to the
// window x coordinate.
type OpDPdxCoarse struct {
	ResultType uint32
	ResultId   uint32
	P          uint32
}

func (c *OpDPdxCoarse) Opcode() uint32 { return 181 }
func (c *OpDPdxCoarse) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpDPdxCoarse{}
	})
}

// OpDPdyCoarse calculates the partial derivative of P with respect to the
// window y coordinate.
type OpDPdyCoarse struct {
	ResultType uint32
	ResultId   uint32
	P          uint32
}

func (c *OpDPdyCoarse) Opcode() uint32 { return 182 }
func (c *OpDPdyCoarse) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpDPdyCoarse{}
	})
}

// OpFwidthCoarse is equivalent to computing the sum of the absolute values
// of OpDPdxCoarse and OpDPdyCoarse on P.
type OpFwidthCoarse struct {
	ResultType uint32
	ResultId   uint32
	P          uint32
}

func (c *OpFwidthCoarse) Opcode() uint32 { return 183 }
func (c *OpFwidthCoarse) Verify() error  { return nil }

func init() {
	Bind(func() Instruction {
		return &OpFwidthCoarse{}
	})
}
