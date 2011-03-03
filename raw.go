package raw

//	THIS SHOULD BE IN A SEPARATE PACKAGE WHERE IT CAN BE REUSED MORE EASILY
//	THE PACKAGE SHOULD BE INSTALLABLE VIA GOINSTALL
//	AND NEEDS EXTENSIVE TESTS!!!!

import "C"
import "reflect"
import "unsafe"

var _BYTE_SLICE reflect.Type
var _STRING		reflect.Type

func init() {
	_BYTE_SLICE = reflect.Typeof([]byte(nil))
	_STRING = reflect.Typeof("")
}

/*
	A Buffered object can present itself as a byteslice.
	This byteslice can then be manipulated directly to modify the contents of memory.
	Use with extreme caution.
*/
type MemoryBlock interface {
	ByteSlice() []byte
}

func ByteSlice(i interface{}) []byte {
	/*
		A byteslice is by definition its own buffer.
		Any type which implements the Buffer interface will generate its result using that method.
	*/
	switch b := i.(type) {
	case []byte:					return b
	case MemoryBlock:				return b.ByteSlice()
	default:
	}

	/*
		There are cerain types which cannot be cast as a buffer and instead raise a panic.
		In the rare case of the interface itself containing another interface we recursively query.
		When given a pointer we use its target address and the size of the type it points to construct a SliceHeader.
		For SliceValues we can do a simple conversion of the SliceHeader to a byteslice.
		For StringValues we treat them as a fixed capacity byte slice.
	*/
	var header *reflect.SliceHeader

	value := reflect.NewValue(i)
	switch value := value.(type) {
	case nil:
		panic(i)
	case reflect.Type:
		panic(i)
	case *reflect.MapValue:
		panic(i)
	case *reflect.ChanValue:
		panic(i)
	case *reflect.PtrValue:
		if value := value.Elem(); value == nil {
			panic(i)
		} else {
			size := int(value.Type().Size())
			header = &reflect.SliceHeader{ value.Addr(), size, size }
		}
	case *reflect.InterfaceValue:
		return ByteSlice(value.Elem())
	case *reflect.SliceValue:
		address := value.Elem(0).Addr()
		bytes := int(value.Get() - address)
		length := (bytes / value.Cap()) * value.Len()
		header = &reflect.SliceHeader{ address, length, bytes }
	case *reflect.StringValue:
		s := value.Get()
		stringheader := *(*reflect.StringHeader)(unsafe.Pointer(&s))
		header = &reflect.SliceHeader{ stringheader.Data, stringheader.Len, stringheader.Len }
	default:
		/*
			For every other type the value gives us an address for the data
			Given this and the size of the underlying allocated memory we can
			then create a []byte sliceheader and return a valid slice
		*/
		size := int(value.Type().Size())
		header = &reflect.SliceHeader{ value.Addr(), size, size }
	}
	return unsafe.Unreflect(_BYTE_SLICE, unsafe.Pointer(header)).([]byte)
}

func Address(i interface{}) (addr unsafe.Pointer) {
	switch b := i.(type) {
	case []byte:					addr = unsafe.Pointer(&(b[0]))
	case MemoryBlock:				addr = unsafe.Pointer(&(b.ByteSlice()[0]))
	default:						addr = unsafe.Pointer(&(ByteSlice(i)[0]))
	}
	return
}

func Range(i interface{}, increment int, f func(int, unsafe.Pointer)) {
	b := ByteSlice(i)
	items := len(b) / increment
	for i := 0; i < items; i++ {
		f(i, unsafe.Pointer((&b[:increment][0])))
		b = b[increment:]
	}
}

func Overwrite(d interface{}, s interface{}) {
	copy(ByteSlice(d), ByteSlice(s))
}