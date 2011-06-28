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
	case reflect.Slice:			if r := r.Type(); r.Kind() == reflect.Slice {
									b = l.Elem() == r.Elem()
								}

	case reflect.Map:			if r := r.Type(); r.Kind() == reflect.Map {
									b = l.Key() == r.Key() && l.Elem() == r.Elem()
								}
	}
	return
}

func Type(name string, v, s interface{}) (r BasicType) {
	r.name = name
	t := reflect.TypeOf(v)
	r.size = int(t.Size())
	r.alignment = int(t.Align())
	r.slice_type = reflect.TypeOf(s)
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

var POINTER		= Type("unsafe.Pointer", unsafe.Pointer(&_a), []unsafe.Pointer{})
var UINTPTR		= Type("uintptr", uintptr(0), []uintptr{})
var INTERFACE	= Type("interface{}", _a, []interface{}{})
var BOOLEAN		= Type("bool", true, []bool{})
var BYTE		= Type("byte", byte(0), []byte{})
var INT			= Type("int", int(0), []int{})
var INT8		= Type("int8", int8(0), []int8{})
var INT16		= Type("int16", int16(0), []int16{})
var INT32		= Type("int32", int32(0), []int32{})
var INT64		= Type("int64", int64(0), []int64{})
var UINT		= Type("uint", uint(0), []uint{})
var UINT8		= Type("uint8", uint8(0), []uint8{})
var UINT16		= Type("uint16", uint16(0), []uint16{})
var UINT32		= Type("uint32", uint32(0), []uint32{})
var UINT64		= Type("uint64", uint64(0), []uint64{})
var FLOAT32		= Type("float32", float32(0.0), []float32{})
var FLOAT64		= Type("float64", float64(0.0), []float64{})
var COMPLEX64	= Type("complex64", complex64(0), []complex64{})
var COMPLEX128	= Type("complex128", complex128(0), []complex128{})