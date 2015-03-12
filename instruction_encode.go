// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import (
	"fmt"
	"reflect"
)

// EncodedLen returns the number of words the given instruction
// will occupy once encoded.
func EncodedLen(i Instruction) int {
	rv := reflect.ValueOf(i)
	rv = reflect.Indirect(rv)
	return encodedValueLen(rv) + 1
}

func encodedValueLen(rv reflect.Value) int {
	switch rv.Kind() {
	case reflect.Struct:
		return encodedStructLen(rv)
	case reflect.Slice, reflect.Array:
		return encodedSliceLen(rv)
	case reflect.Uint32:
		return 1
	case reflect.String:
		return int(String(rv.String()).EncodedLen())
	}
	return 0
}

func encodedStructLen(rv reflect.Value) int {
	var len int

	for i := 0; i < rv.NumField(); i++ {
		len += encodedValueLen(rv.Field(i))
	}

	return len
}

func encodedSliceLen(rv reflect.Value) int {
	var len int

	for i := 0; i < rv.Len(); i++ {
		len += encodedValueLen(rv.Index(i))
	}

	return len
}

// encode uses reflection to encode the given instruction into
// a sequence of words.
func encode(i Instruction, out []uint32) (uint32, error) {
	rv := reflect.ValueOf(i)
	rv = reflect.Indirect(rv)
	return encodeValue(rv, out)
}

func encodeValue(rv reflect.Value, out []uint32) (uint32, error) {
	switch rv.Kind() {
	case reflect.Struct:
		return encodeStruct(rv, out)
	case reflect.Uint32:
		return encodeUint32(rv, out)
	case reflect.String:
		return encodeString(rv, out)
	case reflect.Slice, reflect.Array:
		return encodeSlice(rv, out)
	}

	return 0, fmt.Errorf("unsupported type: %v", rv.Kind())
}

// encodeStruct encodes the given struct.
func encodeStruct(rv reflect.Value, out []uint32) (uint32, error) {
	var index uint32

	for i := 0; i < rv.NumField(); i++ {
		fld := rv.Field(i)

		argc, err := encodeValue(fld, out[index:])
		if err != nil {
			return 0, err
		}

		index += argc
	}

	return index, nil
}

// encodeUint32 encodes the given uint32.
func encodeUint32(rv reflect.Value, out []uint32) (uint32, error) {
	out[0] = uint32(rv.Uint())
	return 1, nil
}

// encodeString encodes the given string.
func encodeString(rv reflect.Value, out []uint32) (uint32, error) {
	str := String(rv.String())
	size := str.EncodedLen()
	str.Encode(out)
	return size, nil
}

// encodeSlice encodes the given slice
func encodeSlice(rv reflect.Value, out []uint32) (uint32, error) {
	var index uint32

	for i := 0; i < rv.Len(); i++ {
		fld := rv.Index(i)

		argc, err := encodeValue(fld, out[index:])
		if err != nil {
			return 0, err
		}

		index += argc
	}

	return index, nil
}
