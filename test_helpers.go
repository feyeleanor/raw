package raw

import (
	"reflect"
	"unsafe"
)

type Fatal interface {
	Fatal(...interface{})
	Fatalf(string, ...interface{})
}

func FailOnBadBufferSize(f Fatal, good_length, good_capacity bool) {
	switch {
	case !good_length:
		f.Fatalf("byte buffer lengths differ")
	case !good_capacity:
		f.Fatalf("byte buffer capacities differ")
	}
}

func FailOnHeaderMismatch(f Fatal, slice unsafe.Pointer, buf []byte) {
	sliceheader := *(*reflect.SliceHeader)(slice)
	bufheader := *(*reflect.SliceHeader)(unsafe.Pointer(&buf))
	if sliceheader.Data != bufheader.Data {
		f.Fatalf("slice addresses don't match: %v != %v", sliceheader.Data, bufheader.Data)
	}
}

func FailOnAddressMismatch(f Fatal, slice unsafe.Pointer, buf []byte) {
	base_address := uintptr(slice)
	bufheader := *(*reflect.SliceHeader)(unsafe.Pointer(&buf))
	if base_address != bufheader.Data {
		f.Fatalf("slice addresses don't match: %v != %v", base_address, bufheader.Data)
	}
}

func FailIfNotCopyable(f Fatal, v interface{}) {
	defer func() {
		if x := recover(); x == nil {
			f.Fatalf("should have raised a panic")
		}
	}()
	ByteSlice(v)
}

func ValidateNumericByteSlice(f Fatal, value interface{}) {
	v := reflect.Indirect(reflect.ValueOf(value))
	size := int(v.Type().Size())
	addr := v.UnsafeAddr()
	buf := ByteSlice(value)

	FailOnBadBufferSize(f,
		len(buf) == size,
		cap(buf) == size,
	)
	FailOnAddressMismatch(f, unsafe.Pointer(addr), buf)
}