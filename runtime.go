package raw

import "fmt"
import "reflect"
import "unsafe"


func Throw() {
	panic(nil)
}

func Catch(f func()) {
	defer func() {
		if x := recover(); x != nil {
			panic(x)
		}
	}()
	f()
}

func CatchAll(f func()) {
	defer func() {
		recover()
	}()
	f()
}

type Typed interface {
	Type() reflect.Type
}

func Compatible(l, r Typed) (b bool) {
	switch l := l.Type(); l.Kind() {
	case reflect.Slice:
		if r := r.Type(); r.Kind() == reflect.Slice {
			b = l.Elem() == r.Elem()
		}

	case reflect.Map:
		if r := r.Type(); r.Kind() == reflect.Map {
			b = l.Key() == r.Key() && l.Elem() == r.Elem()
		}
	}
	return
}

type BasicType struct {
	name			string
	size			int
	alignment		int
	slice_type		reflect.Type
}

func (b BasicType) String() string {
	return fmt.Sprintf("%v: %v bytes aligned at %v byte", b.name, b.size, b.alignment)
}

var _a interface{} = 0

var POINTER		= BasicType{ "unsafe.Pointer", unsafe.Sizeof(unsafe.Pointer(&_a)), unsafe.Alignof(unsafe.Pointer(&_a)), reflect.TypeOf([]unsafe.Pointer{}) }
var UINTPTR		= BasicType{ "uintptr", unsafe.Sizeof(uintptr(0)), unsafe.Alignof(uintptr(0)), reflect.TypeOf([]uintptr{}) }
var INTERFACE	= BasicType{ "interface{}", unsafe.Sizeof(_a), unsafe.Alignof(_a), reflect.TypeOf([]interface{}{}) }
var BOOLEAN		= BasicType{ "bool", unsafe.Sizeof(true), unsafe.Alignof(true), reflect.TypeOf([]bool{}) }
var BYTE		= BasicType{ "byte", unsafe.Sizeof(byte(0)), unsafe.Alignof(byte(0)), reflect.TypeOf([]byte{}) }
var INT			= BasicType{ "int", unsafe.Sizeof(int(0)), unsafe.Alignof(int(0)), reflect.TypeOf([]int{}) }
var INT8		= BasicType{ "int8", unsafe.Sizeof(int8(0)), unsafe.Alignof(int8(0)), reflect.TypeOf([]int8{}) }
var INT16		= BasicType{ "int16", unsafe.Sizeof(int16(0)), unsafe.Alignof(int16(0)), reflect.TypeOf([]int16{}) }
var INT32		= BasicType{ "int32", unsafe.Sizeof(int32(0)), unsafe.Alignof(int32(0)), reflect.TypeOf([]int32{}) }
var INT64		= BasicType{ "int64", unsafe.Sizeof(int64(0)), unsafe.Alignof(int64(0)), reflect.TypeOf([]int64{}) }
var UINT		= BasicType{ "uint", unsafe.Sizeof(uint(0)), unsafe.Alignof(uint(0)), reflect.TypeOf([]uint{}) }
var UINT8		= BasicType{ "uint8", unsafe.Sizeof(uint8(0)), unsafe.Alignof(uint8(0)), reflect.TypeOf([]uint8{}) }
var UINT16		= BasicType{ "uint16", unsafe.Sizeof(uint16(0)), unsafe.Alignof(uint16(0)), reflect.TypeOf([]uint16{}) }
var UINT32		= BasicType{ "uint32", unsafe.Sizeof(uint32(0)), unsafe.Alignof(uint32(0)), reflect.TypeOf([]uint32{}) }
var UINT64		= BasicType{ "uint64", unsafe.Sizeof(uint64(0)), unsafe.Alignof(uint64(0)), reflect.TypeOf([]uint64{}) }
var FLOAT32		= BasicType{ "float32", unsafe.Sizeof(float32(0.0)), unsafe.Alignof(float32(0.0)), reflect.TypeOf([]float32{}) }
var FLOAT64		= BasicType{ "float64", unsafe.Sizeof(float64(0.0)), unsafe.Alignof(float64(0.0)), reflect.TypeOf([]float64{}) }
var COMPLEX64	= BasicType{ "complex64", unsafe.Sizeof(complex64(0)), unsafe.Alignof(complex64(0)), reflect.TypeOf([]complex64{}) }
var COMPLEX128	= BasicType{ "complex128", unsafe.Sizeof(complex128(0)), unsafe.Alignof(complex128(0)), reflect.TypeOf([]complex128{}) }