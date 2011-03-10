package raw

import "C"
import "reflect"
import "unsafe"

const(
	POINTER = iota
	UINTPTR
	INTERFACE
	BOOLEAN
	BYTE
	INT
	INT8
	INT16
	INT32
	INT64
	UINT
	UINT8
	UINT16
	UINT32
	UINT64
	FLOAT32
	FLOAT64
	COMPLEX64
	COMPLEX128
)

var _a interface{} = 0

var ALIGNMENT = []int{
	unsafe.Alignof(unsafe.Pointer(&_a)),
	unsafe.Alignof(uintptr(0)),
	unsafe.Alignof(_a),
	unsafe.Alignof(true),
	unsafe.Alignof(byte(0)),
	unsafe.Alignof(int(0)),
	unsafe.Alignof(int8(0)),
	unsafe.Alignof(int16(0)),
	unsafe.Alignof(int32(0)),
	unsafe.Alignof(int64(0)),
	unsafe.Alignof(uint(0)),
	unsafe.Alignof(uint8(0)),
	unsafe.Alignof(uint16(0)),
	unsafe.Alignof(uint32(0)),
	unsafe.Alignof(uint64(0)),
	unsafe.Alignof(float32(0.0)),
	unsafe.Alignof(float64(0.0)),
}

var SIZE = []int{
	unsafe.Sizeof(unsafe.Pointer(&_a)),
	unsafe.Sizeof(uintptr(0)),
	unsafe.Sizeof(_a),
	unsafe.Sizeof(true),
	unsafe.Sizeof(byte(0)),
	unsafe.Sizeof(int(0)),
	unsafe.Sizeof(int8(0)),
	unsafe.Sizeof(int16(0)),
	unsafe.Sizeof(int32(0)),
	unsafe.Sizeof(int64(0)),
	unsafe.Sizeof(uint(0)),
	unsafe.Sizeof(uint8(0)),
	unsafe.Sizeof(uint16(0)),
	unsafe.Sizeof(uint32(0)),
	unsafe.Sizeof(uint64(0)),
	unsafe.Sizeof(float32(0.0)),
	unsafe.Sizeof(float64(0.0)),
	unsafe.Sizeof(complex64(0)),
	unsafe.Sizeof(complex128(0)),
}

var SLICE_TYPE = []reflect.Type{
	reflect.Typeof([]unsafe.Pointer{}),
	reflect.Typeof([]uintptr{}),
	reflect.Typeof([]interface{}{}),
	reflect.Typeof([]bool{}),
	reflect.Typeof([]byte{}),
	reflect.Typeof([]int{}),
	reflect.Typeof([]int8{}),
	reflect.Typeof([]int16{}),
	reflect.Typeof([]int32{}),
	reflect.Typeof([]int64{}),
	reflect.Typeof([]uint{}),
	reflect.Typeof([]uint8{}),
	reflect.Typeof([]uint16{}),
	reflect.Typeof([]uint32{}),
	reflect.Typeof([]uint64{}),
	reflect.Typeof([]float32{}),
	reflect.Typeof([]float64{}),
	reflect.Typeof([]complex64{}),
	reflect.Typeof([]complex128{}),
}