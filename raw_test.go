package raw

import "reflect"
import "testing"
import "unsafe"


func TestByteSliceWithByteSlice(t *testing.T) {
	b := make([]byte, 10, 10)
	buf := ByteSlice(b)
	if len(buf) != len(b) {
		t.Fatalf("byte buffer lengths differ: %v != %v", len(buf), len(b))
	}
	if cap(buf) != cap(b) {
		t.Fatalf("byte buffer capacities differ: %v != %v", cap(buf), cap(b))
	}

	sliceheader := *(*reflect.SliceHeader)(unsafe.Pointer(&b))
	bufheader := *(*reflect.SliceHeader)(unsafe.Pointer(&buf))
	if sliceheader.Data != bufheader.Data {
		t.Fatalf("slice addresses don't match: %v != %v", sliceheader.Data, bufheader.Data)
	}
}

func TestByteSliceWithMap(t *testing.T) {
	m := make(map[int]int)
	defer func() {
		if x := recover(); x != nil {
			if x != m {
				t.Fatalf("panic threw wrong type: %v instead of %v", x, m)
			}
		} else {
			t.Fatalf("should have raised a panic")
		}
	}()
	ByteSlice(m)
}

func TestByteSliceWithChannel(t *testing.T) {
	c := make(chan int)
	defer func() {
		if x := recover(); x != nil {
			if x != c {
				t.Fatalf("panic threw wrong type: %v instead of %v", x, c)
			}
		} else {
			t.Fatalf("should have raised a panic")
		}
	}()
	ByteSlice(c)
}

func TestByteSliceWithInterface(t *testing.T) {
	var i interface{} = make([]byte, 16, 16)
	b := i.([]byte)
	buf := ByteSlice(i)
	if len(buf) != len(b) {
		t.Fatalf("byte buffer lengths differ: %v != %v", len(buf), len(b))
	}
	if cap(buf) != cap(b) {
		t.Fatalf("byte buffer capacities differ: %v != %v", cap(buf), cap(b))
	}

	sliceheader := *(*reflect.SliceHeader)(unsafe.Pointer(&b))
	bufheader := *(*reflect.SliceHeader)(unsafe.Pointer(&buf))
	if sliceheader.Data != bufheader.Data {
		t.Fatalf("slice addresses don't match: %v != %v", sliceheader.Data, bufheader.Data)
	}
}

func TestByteSliceWithString(t *testing.T) {
	s := "hello"
	buf := ByteSlice(s)
	if len(buf) != len(s) {
		t.Fatalf("byte buffer lengths differ: %v != %v", len(buf), len(s))
	}
	if cap(buf) != len(s) {
		t.Fatalf("byte buffer capacities differ: %v != %v", cap(buf), len(s))
	}

	stringheader := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	bufheader := *(*reflect.SliceHeader)(unsafe.Pointer(&buf))
	if stringheader.Data != bufheader.Data {
		t.Fatalf("slice addresses don't match: %v != %v", stringheader.Data, bufheader.Data)
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

	size := unsafe.Sizeof(point)
	if len(buf) != size {
		t.Fatalf("byte buffer lengths differ: %v != %v", len(buf), size)
	}
	if cap(buf) != size {
		t.Fatalf("byte buffer capacities differ: %v != %v", cap(buf), size)
	}

	base_address := uintptr(unsafe.Pointer(&point))
	bufheader := *(*reflect.SliceHeader)(unsafe.Pointer(&buf))
	if base_address != bufheader.Data {
		t.Fatalf("slice addresses don't match: %v != %v", base_address, bufheader.Data)
	}
}

func TestByteSliceWithStructPointer(t *testing.T) {
	point := &Point{ 3, 4, 5 }
	buf := ByteSlice(point)

	size := unsafe.Sizeof(*point)
	if len(buf) != size {
		t.Fatalf("byte buffer lengths differ: %v != %v", len(buf), size)
	}
	if cap(buf) != size {
		t.Fatalf("byte buffer capacities differ: %v != %v", cap(buf), size)
	}

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

	size := unsafe.Sizeof(*tag)
	if len(buf) != size {
		t.Fatalf("byte buffer lengths differ: %v != %v", len(buf), size)
	}
	if cap(buf) != size {
		t.Fatalf("byte buffer capacities differ: %v != %v", cap(buf), size)
	}

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

	size := unsafe.Sizeof(*tag)
	if len(buf) != size {
		t.Fatalf("byte buffer lengths differ: %v != %v", len(buf), size)
	}
	if cap(buf) != size {
		t.Fatalf("byte buffer capacities differ: %v != %v", cap(buf), size)
	}

	base_address := uintptr(unsafe.Pointer(tag))
	bufheader := *(*reflect.SliceHeader)(unsafe.Pointer(&buf))
	if base_address != bufheader.Data {
		t.Fatalf("slice addresses don't match: %v != %v", base_address, bufheader.Data)
	}
}

func TestByteSliceWithInt32Slice(t *testing.T) {
	i := make([]int32, 10, 10)
	buf := ByteSlice(i)
	if len(buf) != len(i) * 4 {
		t.Fatalf("byte buffer lengths differ: %v != %v", len(buf), len(i) * 4)
	}
	if cap(buf) != cap(i) * 4 {
		t.Fatalf("byte buffer capacities differ: %v != %v", cap(buf), cap(i) * 4)
	}

	sliceheader := *(*reflect.SliceHeader)(unsafe.Pointer(&i))
	bufheader := *(*reflect.SliceHeader)(unsafe.Pointer(&buf))
	if sliceheader.Data != bufheader.Data {
		t.Fatalf("slice addresses don't match: %v != %v", sliceheader.Data, bufheader.Data)
	}
}

func TestByteSliceWithInt64Slice(t *testing.T) {
	i := make([]int64, 10, 10)
	buf := ByteSlice(i)
	if len(buf) != len(i) * 8 {
		t.Fatalf("byte buffer lengths differ: %v != %v", len(buf), len(i) * 8)
	}
	if cap(buf) != cap(i) * 8 {
		t.Fatalf("byte buffer capacities differ: %v != %v", cap(buf), cap(i) * 8)
	}

	sliceheader := *(*reflect.SliceHeader)(unsafe.Pointer(&i))
	bufheader := *(*reflect.SliceHeader)(unsafe.Pointer(&buf))
	if sliceheader.Data != bufheader.Data {
		t.Fatalf("slice addresses don't match: %v != %v", sliceheader.Data, bufheader.Data)
	}
}

func TestByteSliceWithFloat32Slice(t *testing.T) {
	f := make([]float32, 10, 10)
	buf := ByteSlice(f)
	if len(buf) != len(f) * 4 {
		t.Fatalf("byte buffer lengths differ: %v != %v", len(buf), len(f) * 4)
	}
	if cap(buf) != cap(f) * 4 {
		t.Fatalf("byte buffer capacities differ: %v != %v", cap(buf), cap(f) * 4)
	}

	sliceheader := *(*reflect.SliceHeader)(unsafe.Pointer(&f))
	bufheader := *(*reflect.SliceHeader)(unsafe.Pointer(&buf))
	if sliceheader.Data != bufheader.Data {
		t.Fatalf("slice addresses don't match: %v != %v", sliceheader.Data, bufheader.Data)
	}
}

func TestByteSliceWithFloat64Slice(t *testing.T) {
	f := make([]float64, 10, 10)
	buf := ByteSlice(f)
	if len(buf) != len(f) * 8 {
		t.Fatalf("byte buffer lengths differ: %v != %v", len(buf), len(f) * 8)
	}
	if cap(buf) != cap(f) * 8 {
		t.Fatalf("byte buffer capacities differ: %v != %v", cap(buf), cap(f) * 8)
	}

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

	v := reflect.NewValue(value)
	switch v := v.(type) {
	case *reflect.PtrValue:
		if e := v.Elem(); e != nil {
			size = int(e.Type().Size())
			addr = e.Addr()
			numtype = e.Type()
		}
	default:
		size = int(v.Type().Size())
		addr = v.Addr()
		numtype = v.Type()
	}

	buf := ByteSlice(value)
	if len(buf) != size {
		t.Fatalf("%v: byte buffer lengths differ: %v != %v", numtype, len(buf), size)
	}

	if cap(buf) != size {
		t.Fatalf("%v: byte buffer capacities differ: %v != %v", numtype, cap(buf), size)
	}

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
	values := []interface{}{ int(0), int8(0), int16(0), int32(0), int64(0), uint(0), uint8(0), uint16(0), uint32(0), uint64(0), float32(0.0), float64(0.0) }
	for _, n := range values {
		ValidateNumericByteSlice(t, n)
	}
}