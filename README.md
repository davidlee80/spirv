## SPIR-V

_Note_: This is a work in progress. Not all instructions have yet
been implemented. They will be in the course of the next few days.
Expect breaking code changes to appear regularly.

Package SPIR-V is a Go encoder/decoder for the Vulkan SPIR-V format.
This is based on the [preliminary specificationn][1] (pdf) and is subject to
change as the specification matures.

Additional SPIR-V information can be found [here][2] (pdf).

[1]: https://www.khronos.org/registry/spir-v/specs/1.0/SPIRV.pdf
[2]: https://www.khronos.org/registry/spir-v/

SPIR-V is a binary intermediate language for representing graphical-shader
stages and compute kernels for multiple Khronos APIs. Each function in a SPIR-V
module contains a control-flow graph (CFG) of basic blocks, with additional
instructions and constraints to retain source-code structured flow control.

SPIR-V has the following goals:

* Provide a simple binary intermediate language for all functionality appearing in Khronos shaders/kernels.
* Have a short, transparent, self-contained specification (sections Specification and Binary Form).
* Map easily to other IRs, including LLVM IR.
* Be the form passed by an API into a driver to set shaders/kernels.
* Can be targeted by new front ends for novel high-level languages.
* Allow the first parts of the source compilation and reflection process to be done offline.
* Be low-level enough to require a reverse-engineering step to reconstruct source code.
* Improve portability by enabling shared tools to generate or operate on it.
* Allow separation of core specification from source language-specific sets of built-in functions.
* Reduce compile time during application run time. (Eliminating most of the compile time during application run time is not a goal of this IL. Target-specific register allocation and scheduling are expected to still take significant time.)
* Allow some optimizations to be done offline.


### Vulkan

Vulkan is a new open standard API by Khronos that offers low-level control of
GPUs for graphics and general purpose computation. It has been designed from
the ground up around the capabilities of modern hardware.

* Direct GPU control with minimal driver overhead. For example, data is written
  directly to GPU memory instead of using calls equivalent to glUniform.
  Applications can implement their own memory allocation strategies. 

* Another feature is the render pass, which offers control over loading of
  render targets at the start and end of renders, which is very useful for
  tiling architectures.

* Multi-threading friendly architecture.
  Multiple threads can create and populate command buffers at the same time,
  which can then be submitted to the GPU by a separate thread.

* Unified API for desktop, mobile and embedded platforms.
  There is no equivalent of OpenGL ES, Vulkan offers the same API on all platforms.

* Intermediate bytecode for shaders. The driver accepts shaders in the SPIR-V
  bytecode format. Khronos will supply a separate compiler for GLSL that
  targets this intermediate format. Third-party developers will be able to
  write their own compilers for other languages.


### Install

    go get github.com/jteeuwen/spirv


### Usage

By itself, this package is not very useful. All it does is decode SPIR-V
binary into sets of 32-bit words data structure and vice-versa. It is intended
as a tool to facilitate the creation of SPIR-V debugging tools, compilers,
and whatever else you may require.

Here is an sample on how to decode a module from a file.
The `Decoder` methods return either a `Header`, `[]uint32` or `Instruction`.

	// Create a decoder.
	dec := spirv.NewDecoder(fd)
	hdr, err := dec.DecodeHeader()
	...

	for {
		instructionWords, err := dec.DecodeInstructionWords()
		...
	}

Or to get a higher level abstraction of the instruction data.
This performs some copying and allocations, but gives a typed instruction
one can work with:

	for {
		instruction, err := dec.DecodeInstruction()
		...
	}


And the same to encode a module to a file:

	var hdr spirv.Header
	...

	enc := spirv.NewEncoder(fd)
	err := enc.EncodeHeader(hdr)
	...

	for _, data := [][]uint32{...} {
		err := enc.EncodeInstructionWords(data)
		...
	}

or:

	for _, data := []Instruction{...} {
		err := enc.EncodeInstruction(words)
		...
	}


### License

Unless otherwise stated, all of the work in this project is subject to a
1-clause BSD license. Its contents can be found in the enclosed LICENSE file.

