package raw

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestByteSliceWithNil(t *testing.T) {
	b := ByteSlice(nil)
	buf := ByteSlice(b)
	FailOnBadBufferSize(t,
		len(buf) == 0,
		cap(buf) == cap(b),
	)

	sliceheader := *(*reflect.SliceHeader)(unsafe.Pointer(&b))
	bufheader := *(*reflect.SliceHeader)(unsafe.Pointer(&buf))
	if sliceheader.Data != bufheader.Data {
		t.Fatalf("slice addresses don't match: %v != %v", sliceheader.Data, bufheader.Data)
	}
}

func TestByteSliceWithByteSlice(t *testing.T) {
	b := make([]byte, 10, 10)
	buf := ByteSlice(b)
	FailOnBadBufferSize(t,
		len(buf) == len(b),
		cap(buf) == cap(b),
	)

	sliceheader := *(*reflect.SliceHeader)(unsafe.Pointer(&b))
	bufheader := *(*reflect.SliceHeader)(unsafe.Pointer(&buf))
	if sliceheader.Data != bufheader.Data {
		t.Fatalf("slice addresses don't match: %v != %v", sliceheader.Data, bufheader.Data)
	}
}

func TestByteSliceWithMap(t *testing.T) {
	m := make(map[int]int)
	defer func() {
		if x := recover(); x == nil {
			t.Fatalf("should have raised a panic")
		}
	}()
	ByteSlice(m)
}

func TestByteSliceWithChannel(t *testing.T) {
	c := make(chan int)
	defer func() {
		if x := recover(); x == nil {
			t.Fatalf("should have raised a panic")
		}
	}()
	ByteSlice(c)
}

func TestByteSliceWithInterface(t *testing.T) {
	t.Log("Awaiting bug fix for incorrect reporting of interface{} value size with unsafe.Sizeof()")
/*	var i interface{} = make([]byte, INTERFACE.Size(), INTERFACE.Size())
	b := i.([]byte)
	buf := ByteSlice(i)

	FailOnBadBufferSize(t,
		len(buf) == len(b),
		cap(buf) == cap(b),
	)

	sliceheader := *(*reflect.SliceHeader)(unsafe.Pointer(&b))
	bufheader := *(*reflect.SliceHeader)(unsafe.Pointer(&buf))
	if sliceheader.Data != bufheader.Data {
		t.Fatalf("slice addresses don't match: %v != %v", sliceheader.Data, bufheader.Data)
	}
*/}

func TestByteSliceWithString(t *testing.T) {
	s := "hello"
	buf := ByteSlice(s)
	FailOnBadBufferSize(t,
		len(buf) == len(s),
		cap(buf) == len(s),
	)

	stringheader := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	bufheader := *(*reflect.SliceHeader)(unsafe.Pointer(&buf))
	if stringheader.Data != bufheader.Data {
		t.Fatalf("slice addresses don't match: %v != %v", stringheader.Data, bufheader.Data)
	}
}

func TestByteSliceWithEmptyStructValue(t *testing.T) {
	s := struct {}{}
	buf := ByteSlice(&s)
	size := int(unsafe.Sizeof(struct {}{}))

	FailOnBadBufferSize(t,
		len(buf) == size,
		cap(buf) == size,
	)

	base_address := uintptr(unsafe.Pointer(&s))
	bufheader := *(*reflect.SliceHeader)(unsafe.Pointer(&buf))
	if base_address != bufheader.Data {
		t.Fatalf("slice addresses don't match: %v != %v", base_address, bufheader.Data)
	}
}

type Point struct {
	x			int32
	y			int32
	z			int32
}

func TestByteSliceWithStructValue(t *testing.T) {
	point := Point{ 3, 4, 5 }
	buf := ByteSlice(&point)
	size := int(unsafe.Sizeof(point))

	FailOnBadBufferSize(t,
		len(buf) == size,
		cap(buf) == size,
	)

	base_address := uintptr(unsafe.Pointer(&point))
	bufheader := *(*reflect.SliceHeader)(unsafe.Pointer(&buf))
	if base_address != bufheader.Data {
		t.Fatalf("slice addresses don't match: %v != %v", base_address, bufheader.Data)
	}
}

func TestByteSliceWithStructPointer(t *testing.T) {
	point := &Point{ 3, 4, 5 }
	buf := ByteSlice(point)
	size := int(unsafe.Sizeof(*point))

	FailOnBadBufferSize(t,
		len(buf) == size,
		cap(buf) == size,
	)

	base_address := uintptr(unsafe.Pointer(point))
	bufheader := *(*reflect.SliceHeader)(unsafe.Pointer(&buf))
	if base_address != bufheader.Data {
		t.Fatalf("slice addresses don't match: %v != %v", base_address, bufheader.Data)
	}
}

type TaggedPoint struct {
	Point
	tag			string
}

func TestByteSliceWithEmbeddedStructValue(t *testing.T) {
	point := Point{ 3, 4, 5 }
	tag := &TaggedPoint{ point, "this is a tag" }
	buf := ByteSlice(tag)
	size := int(unsafe.Sizeof(*tag))

	FailOnBadBufferSize(t,
		len(buf) == size,
		cap(buf) == size,
	)

	base_address := uintptr(unsafe.Pointer(tag))
	bufheader := *(*reflect.SliceHeader)(unsafe.Pointer(&buf))
	if base_address != bufheader.Data {
		t.Fatalf("slice addresses don't match: %v != %v", base_address, bufheader.Data)
	}
}


type TaggedPointReference struct {
	*Point
	tag			string
}

func TestByteSliceWithEmbeddedStructPointer(t *testing.T) {
	point := Point{ 3, 4, 5 }
	tag := &TaggedPointReference{ &point, "this is a tag" }
	buf := ByteSlice(tag)
	size := int(unsafe.Sizeof(*tag))

	FailOnBadBufferSize(t,
		len(buf) == size,
		cap(buf) == size,
	)

	base_address := uintptr(unsafe.Pointer(tag))
	bufheader := *(*reflect.SliceHeader)(unsafe.Pointer(&buf))
	if base_address != bufheader.Data {
		t.Fatalf("slice addresses don't match: %v != %v", base_address, bufheader.Data)
	}
}

func TestByteSliceWithInt32Slice(t *testing.T) {
	i := make([]int32, 10, 20)
	buf := ByteSlice(i)
	is := int(reflect.TypeOf(i).Elem().Size())

	FailOnBadBufferSize(t,
		len(buf) == is * len(i),
		cap(buf) == is * cap(i),
	)

	sliceheader := *(*reflect.SliceHeader)(unsafe.Pointer(&i))
	bufheader := *(*reflect.SliceHeader)(unsafe.Pointer(&buf))
	if sliceheader.Data != bufheader.Data {
		t.Fatalf("slice addresses don't match: %v != %v", sliceheader.Data, bufheader.Data)
	}
}

func TestByteSliceWithInt64Slice(t *testing.T) {
	i := make([]int64, 10, 20)
	buf := ByteSlice(i)
	is := int(reflect.TypeOf(i).Elem().Size())

	FailOnBadBufferSize(t,
		len(buf) == is * len(i),
		cap(buf) == is * cap(i),
	)

	sliceheader := *(*reflect.SliceHeader)(unsafe.Pointer(&i))
	bufheader := *(*reflect.SliceHeader)(unsafe.Pointer(&buf))
	if sliceheader.Data != bufheader.Data {
		t.Fatalf("slice addresses don't match: %v != %v", sliceheader.Data, bufheader.Data)
	}
}

func TestByteSliceWithFloat32Slice(t *testing.T) {
	f := make([]float32, 10, 10)
	buf := ByteSlice(f)
	fs := int(reflect.TypeOf(f).Elem().Size())

	FailOnBadBufferSize(t,
		len(buf) == fs * len(f),
		cap(buf) == fs * cap(f),
	)

	sliceheader := *(*reflect.SliceHeader)(unsafe.Pointer(&f))
	bufheader := *(*reflect.SliceHeader)(unsafe.Pointer(&buf))
	if sliceheader.Data != bufheader.Data {
		t.Fatalf("slice addresses don't match: %v != %v", sliceheader.Data, bufheader.Data)
	}
}

func TestByteSliceWithFloat64Slice(t *testing.T) {
	f := make([]float64, 10, 10)
	buf := ByteSlice(f)
	fs := int(reflect.TypeOf(f).Elem().Size())

	FailOnBadBufferSize(t,
		len(buf) == fs * len(f),
		cap(buf) == fs * cap(f),
	)

	sliceheader := *(*reflect.SliceHeader)(unsafe.Pointer(&f))
	bufheader := *(*reflect.SliceHeader)(unsafe.Pointer(&buf))
	if sliceheader.Data != bufheader.Data {
		t.Fatalf("slice addresses don't match: %v != %v", sliceheader.Data, bufheader.Data)
	}
}

func ValidateNumericByteSlice(t *testing.T, value interface{}) {
	var size	int
	var addr	uintptr
	var numtype	reflect.Type

	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	size = int(v.Type().Size())
	addr = v.UnsafeAddr()
	numtype = v.Type()
	buf := ByteSlice(value)

	FailOnBadBufferSize(t,
		len(buf) == size,
		cap(buf) == size,
	)

	bufheader := *(*reflect.SliceHeader)(unsafe.Pointer(&buf))
	if addr != bufheader.Data {
		t.Fatalf("%v: addresses don't match: %v != %v", numtype, addr, bufheader.Data)
	}
}

func TestByteSliceWithNumbers(t *testing.T) {
	var i		int
	var i8		int8
	var i16		int16
	var i32 	int32
	var i64		int64
	var u		uint
	var u8		uint8
	var u16		uint16
	var u32 	uint32
	var u64		uint64
	var f32 	float32
	var f64		float64
	var c64		complex64
	var c128	complex64

	ValidateNumericByteSlice(t, &i)
	ValidateNumericByteSlice(t, &i8)
	ValidateNumericByteSlice(t, &i16)
	ValidateNumericByteSlice(t, &i32)
	ValidateNumericByteSlice(t, &i64)
	ValidateNumericByteSlice(t, &u)
	ValidateNumericByteSlice(t, &u8)
	ValidateNumericByteSlice(t, &u16)
	ValidateNumericByteSlice(t, &u32)
	ValidateNumericByteSlice(t, &u64)
	ValidateNumericByteSlice(t, &f32)
	ValidateNumericByteSlice(t, &f64)
	ValidateNumericByteSlice(t, &c64)
	ValidateNumericByteSlice(t, &c128)
}

func TestByteSliceWithNumbersInSlice(t *testing.T) {
	values := []interface{}{	int(0), int8(0), int16(0), int32(0), int64(0),
								uint(0), uint8(0), uint16(0), uint32(0), uint64(0),
								float32(0.0), float64(0.0),
								complex64(0), complex128(0) }
	for i, _ := range values {
		ValidateNumericByteSlice(t, &values[i])
	}
}

func TestByteSliceDataAddress(t *testing.T) {
	DataAddress([]byte{})
	DataAddress(make([]byte, 0, 0))
}

func TestSliceByteCopy(t *testing.T) {
	s := []int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	for i := 0; i < len(s); i++ {
		c := cap(s) / 2
		if c < i {
			c = i
		}
		d := make([]int, i, c)
		ByteCopy(d, s)
		if len(d) != i {
			t.Fatalf("destination buffer length changed: %v != %v", len(d), i)
		}
		if !reflect.DeepEqual(d, s[:i]) {
			t.Fatalf("buffer contents differ: %v != %v", d, s[:i])
		}
	}
}