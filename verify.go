// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import (
	"reflect"
	"regexp"
)

// flags defining regular expression properties.
//
// They are defined in the upper 16 bit range, as not to
// be confused for actual opcodes, which are all set in the
// lower 16-bit range.
const (
	regSingleOptional = (1 << 16) + iota
	regMultiOptional
	regAtleastOne
	regBeginGroup
	regEndGroup
	regOr
	regAny
)

// layoutPattern defines the order in which instructions must appear in a valid
// SPIR-V module. We use a regex pattern to match against sequences of
// instruction opcodes, which are encoded in the pattern as 32-bit unicode
// code points.
var regLayoutPattern *regexp.Regexp

// construct the layout regex pattern.
func init() {
	pattern := []rune{
		opcodeSource, regSingleOptional,
		opcodeSourceExtension, regMultiOptional,
		opcodeCompileFlag, regMultiOptional,
		opcodeExtension, regMultiOptional,
		opcodeExtInstImport, regMultiOptional,

		opcodeMemoryModel,

		opcodeEntryPoint, regAtleastOne,
		opcodeExecutionMode, regAtleastOne,
		opcodeString, regMultiOptional,
		opcodeName, regMultiOptional,
		opcodeMemberName, regMultiOptional,
		opcodeLine, regMultiOptional,

		regBeginGroup,
		opcodeDecorate, regOr,
		opcodeMemberDecorate, regOr,
		opcodeGroupDecorate, regOr,
		opcodeGroupMemberDecorate, regOr,
		opcodeDecorationGroup,
		regEndGroup, regMultiOptional,

		regBeginGroup,
		opcodeTypeArray, regOr,
		opcodeTypeBool, regOr,
		opcodeTypeDeviceEvent, regOr,
		opcodeTypeEvent, regOr,
		opcodeTypeFilter, regOr,
		opcodeTypeFloat, regOr,
		opcodeTypeFunction, regOr,
		opcodeTypeInt, regOr,
		opcodeTypeMatrix, regOr,
		opcodeTypeOpaque, regOr,
		opcodeTypePipe, regOr,
		opcodeTypePointer, regOr,
		opcodeTypeQueue, regOr,
		opcodeTypeReserveId, regOr,
		opcodeTypeRuntimeArray, regOr,
		opcodeTypeSampler, regOr,
		opcodeTypeStruct, regOr,
		opcodeTypeVector, regOr,
		opcodeTypeVoid, regOr,
		opcodeConstant, regOr,
		opcodeConstantComposite, regOr,
		opcodeConstantFalse, regOr,
		opcodeConstantNullObject, regOr,
		opcodeConstantNullPointer, regOr,
		opcodeConstantSampler, regOr,
		opcodeConstantTrue, regOr,
		opcodeVariable, // Can be only global -- Storage Class != Function
		regEndGroup, regMultiOptional,

		// Function
		regBeginGroup,
		opcodeFunction,
		opcodeFunctionParameter, regMultiOptional,

		// Block
		regBeginGroup,
		opcodeLabel,
		regAny, regMultiOptional,
		opcodeBranch,
		regEndGroup, regMultiOptional,

		opcodeFunctionEnd,
		regEndGroup, regAtleastOne,
	}

	// Assemble the pattern into a list of unicode code points,
	// ready to be read by the regex parser.
	out := make([]rune, 0, len(pattern))

	for _, value := range pattern {
		switch value {
		case regSingleOptional:
			out = append(out, '?')

		case regMultiOptional:
			out = append(out, '*')

		case regAtleastOne:
			out = append(out, '+')

		case regBeginGroup:
			out = append(out, '(')

		case regEndGroup:
			out = append(out, ')')

		case regOr:
			out = append(out, '|')

		case regAny:
			out = append(out, '.')

		default:
			// We add 0xff to each opcode, because we want to prevent the
			// possibility that an opcode is treated as an ASCII character.
			// It could be interpreted as a special regex marker like ?, +, *, etc,
			// which is not ideal.
			out = append(out, 0xff+value)
		}
	}

	// The pattern has to be applied to the entire input slice, not just a subset.
	regLayoutPattern = regexp.MustCompile("^" + string(out) + "$")
}

// Verifiable defines any type which implements verification semantics.
// This may entail simple range checks on numeric fields and constants, or
// as complex as semantic rule validation in a whole module.
type Verifiable interface {
	Verify() error
}

// verifyInstruction iterates over the instruction fields and calls
// Verify() on all of those which implement it. We then call Verify
// on the instruction itself.
func verifyInstruction(i Instruction) error {
	rv := reflect.ValueOf(i)
	rv = reflect.Indirect(rv)

	err := verifyValue(rv)
	if err != nil {
		return err
	}

	return i.Verify()
}

func verifyValue(rv reflect.Value) error {
	switch rv.Kind() {
	case reflect.Struct:
		return verifyStruct(rv)
	case reflect.Slice:
		return verifySlice(rv)
	}

	data := rv.Interface()
	verifiable, ok := data.(Verifiable)
	if ok {
		return verifiable.Verify()
	}

	return nil
}

func verifyStruct(rv reflect.Value) error {
	for i := 0; i < rv.NumField(); i++ {
		err := verifyValue(rv.Field(i))
		if err != nil {
			return err
		}
	}

	return nil
}

func verifySlice(rv reflect.Value) error {
	for i := 0; i < rv.Len(); i++ {
		err := verifyValue(rv.Index(i))
		if err != nil {
			return err
		}
	}

	return nil
}
