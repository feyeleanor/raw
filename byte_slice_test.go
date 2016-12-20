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
	FailOnHeaderMismatch(t, unsafe.Pointer(&b), buf)
}

func TestByteSliceWithByteSlice(t *testing.T) {
	b := make([]byte, 10, 10)
	buf := ByteSlice(b)
	FailOnBadBufferSize(t,
		len(buf) == len(b),
		cap(buf) == cap(b),
	)
	FailOnHeaderMismatch(t, unsafe.Pointer(&b), buf)
}

func TestByteSliceWithMap(t *testing.T) {
	FailIfNotCopyable(t, make(map[int] int))
}

func TestByteSliceWithChannel(t *testing.T) {
	FailIfNotCopyable(t, make(chan int))
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
	FailOnHeaderMismatch(t, unsafe.Pointer(&b), buf)
*/}

func TestByteSliceWithString(t *testing.T) {
	s := "hello"
	buf := ByteSlice(s)
	FailOnBadBufferSize(t,
		len(buf) == len(s),
		cap(buf) == len(s),
	)
	FailOnHeaderMismatch(t, unsafe.Pointer(&s), buf)
}

func TestByteSliceWithEmptyStructValue(t *testing.T) {
	s := struct {}{}
	buf := ByteSlice(&s)
	size := int(unsafe.Sizeof(struct {}{}))

	FailOnBadBufferSize(t,
		len(buf) == size,
		cap(buf) == size,
	)
	FailOnAddressMismatch(t, unsafe.Pointer(&s), buf)
}

func TestByteSliceWithStructValue(t *testing.T) {
	point := point{ 3, 4, 5 }
	buf := ByteSlice(&point)
	size := int(unsafe.Sizeof(point))

	FailOnBadBufferSize(t,
		len(buf) == size,
		cap(buf) == size,
	)
	FailOnAddressMismatch(t, unsafe.Pointer(&point), buf)
}

func TestByteSliceWithStructPointer(t *testing.T) {
	point := &point{ 3, 4, 5 }
	buf := ByteSlice(point)
	size := int(unsafe.Sizeof(*point))

	FailOnBadBufferSize(t,
		len(buf) == size,
		cap(buf) == size,
	)
	FailOnAddressMismatch(t, unsafe.Pointer(point), buf)
}

func TestByteSliceWithEmbeddedStructValue(t *testing.T) {
	point := point{ 3, 4, 5 }
	tag := &taggedPoint{ point, "this is a tag" }
	buf := ByteSlice(tag)
	size := int(unsafe.Sizeof(*tag))

	FailOnBadBufferSize(t,
		len(buf) == size,
		cap(buf) == size,
	)
	FailOnAddressMismatch(t, unsafe.Pointer(tag), buf)
}

func TestByteSliceWithEmbeddedStructPointer(t *testing.T) {
	point := point{ 3, 4, 5 }
	tag := &taggedPointReference{ &point, "this is a tag" }
	buf := ByteSlice(tag)
	size := int(unsafe.Sizeof(*tag))

	FailOnBadBufferSize(t,
		len(buf) == size,
		cap(buf) == size,
	)
	FailOnAddressMismatch(t, unsafe.Pointer(tag), buf)
}

func TestByteSliceWithInt32Slice(t *testing.T) {
	i := make([]int32, 10, 20)
	buf := ByteSlice(i)
	is := int(reflect.TypeOf(i).Elem().Size())

	FailOnBadBufferSize(t,
		len(buf) == is * len(i),
		cap(buf) == is * cap(i),
	)
	FailOnHeaderMismatch(t, unsafe.Pointer(&i), buf)
}

func TestByteSliceWithInt64Slice(t *testing.T) {
	i := make([]int64, 10, 20)
	buf := ByteSlice(i)
	is := int(reflect.TypeOf(i).Elem().Size())

	FailOnBadBufferSize(t,
		len(buf) == is * len(i),
		cap(buf) == is * cap(i),
	)
	FailOnHeaderMismatch(t, unsafe.Pointer(&i), buf)
}

func TestByteSliceWithFloat32Slice(t *testing.T) {
	f := make([]float32, 10, 20)
	buf := ByteSlice(f)
	fs := int(reflect.TypeOf(f).Elem().Size())

	FailOnBadBufferSize(t,
		len(buf) == fs * len(f),
		cap(buf) == fs * cap(f),
	)
	FailOnHeaderMismatch(t, unsafe.Pointer(&f), buf)
}

func TestByteSliceWithFloat64Slice(t *testing.T) {
	f := make([]float64, 10, 20)
	buf := ByteSlice(f)
	fs := int(reflect.TypeOf(f).Elem().Size())

	FailOnBadBufferSize(t,
		len(buf) == fs * len(f),
		cap(buf) == fs * cap(f),
	)
	FailOnHeaderMismatch(t, unsafe.Pointer(&f), buf)
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
	values := []interface{}{
		int(0), int8(0), int16(0), int32(0), int64(0),
		uint(0), uint8(0), uint16(0), uint32(0), uint64(0),
		float32(0.0), float64(0.0),
		complex64(0), complex128(0),
	}
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