package raw

import "C"
import "reflect"
import "unsafe"

var _SIZE_PTR = unsafe.Sizeof(uintptr(0))

var _SIZE_INTERFACE = unsafe.Sizeof(interface{}(0))

var _SIZE_BYTE = unsafe.Sizeof(byte(0))

var _SIZE_INT = unsafe.Sizeof(int(0))
var _SIZE_INT8 = unsafe.Sizeof(int8(0))
var _SIZE_INT16 = unsafe.Sizeof(int16(0))
var _SIZE_INT32 = unsafe.Sizeof(int32(0))
var _SIZE_INT64 = unsafe.Sizeof(int64(0))

var _SIZE_UINT = unsafe.Sizeof(uint(0))
var _SIZE_UINT8 = unsafe.Sizeof(uint8(0))
var _SIZE_UINT16 = unsafe.Sizeof(uint16(0))
var _SIZE_UINT32 = unsafe.Sizeof(uint32(0))
var _SIZE_UINT64 = unsafe.Sizeof(uint64(0))

var _SIZE_FLOAT32 = unsafe.Sizeof(float32(0.0))
var _SIZE_FLOAT64 = unsafe.Sizeof(float64(0.0))

var _SIZE_COMPLEX64 = unsafe.Sizeof(complex64(0))
var _SIZE_COMPLEX128 = unsafe.Sizeof(complex128(0))

var _SLICE_PTR = reflect.Typeof([]uintptr{})

var _SLICE_INTERFACE = reflect.Typeof([]interface{}{})

var _SLICE_BYTE = reflect.Typeof([]byte{})

var _SLICE_INT = reflect.Typeof([]int{})
var _SLICE_INT8 = reflect.Typeof([]int8{})
var _SLICE_INT16 = reflect.Typeof([]int16{})
var _SLICE_INT32 = reflect.Typeof([]int32{})
var _SLICE_INT64 = reflect.Typeof([]int64{})

var _SLICE_UINT = reflect.Typeof([]uint{})
var _SLICE_UINT8 = reflect.Typeof([]uint8{})
var _SLICE_UINT16 = reflect.Typeof([]uint16{})
var _SLICE_UINT32 = reflect.Typeof([]uint32{})
var _SLICE_UINT64 = reflect.Typeof([]uint64{})

var _SLICE_FLOAT32 = reflect.Typeof([]float32{})
var _SLICE_FLOAT64 = reflect.Typeof([]float64{})

var _SLICE_COMPLEX64 = reflect.Typeof([]complex64{})
var _SLICE_COMPLEX128 = reflect.Typeof([]complex128{})
