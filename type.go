package raw

import (
	"reflect"
	"unsafe"
)

//	Repeatedly uses *reflect.Indirect* to step through a series of *reflect.Ptr* values
//	and arrive at the concrete value they point to.
//
func DereferenceAll(v interface{}) interface{} {
	var r reflect.Value
	for r = reflect.ValueOf(v); r.Kind() == reflect.Ptr; r = reflect.Indirect(r) {}
	return r.Interface()
}

//	Uses reflect.Set to create a *shallow* copy of an existing value.
//
func shallowCopy(v reflect.Value) (r reflect.Value) {
	r = reflect.New(v.Type())
	r.Elem().Set(v)
	return
}

//	Uses reflection to create a *shallow* copy of an existing value.
//	This will only work if the Type of value being copied is *settable* according to the language specification.
//
func ShallowCopy(v interface{}) (r interface{}) {
	var rv reflect.Value
	switch v := v.(type) {
	case reflect.Value:
		rv = shallowCopy(v)
	default:
		rv = shallowCopy(reflect.ValueOf(v))
	}
	return rv.Interface()
}

func MakeAddressable(v reflect.Value) reflect.Value {
	if !v.CanAddr() {
		v = shallowCopy(v)
	}
	return v
}

var _a interface{} = 0

var POINTER		= reflect.TypeOf(unsafe.Pointer(&_a))
var UINTPTR		= reflect.TypeOf(uintptr(0))
var INTERFACE	= reflect.TypeOf(_a)
var BOOLEAN		= reflect.TypeOf(true)
var BYTE		= reflect.TypeOf(byte(0))
var INT			= reflect.TypeOf(int(0))
var INT8		= reflect.TypeOf(int8(0))
var INT16		= reflect.TypeOf(int16(0))
var INT32		= reflect.TypeOf(int32(0))
var INT64		= reflect.TypeOf(int64(0))
var UINT		= reflect.TypeOf(uint(0))
var UINT8		= reflect.TypeOf(uint8(0))
var UINT16		= reflect.TypeOf(uint16(0))
var UINT32		= reflect.TypeOf(uint32(0))
var UINT64		= reflect.TypeOf(uint64(0))
var FLOAT32		= reflect.TypeOf(float32(0.0))
var FLOAT64		= reflect.TypeOf(float64(0.0))
var COMPLEX64	= reflect.TypeOf(complex64(0))
var COMPLEX128	= reflect.TypeOf(complex128(0))