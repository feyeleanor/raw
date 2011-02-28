package raw

import "reflect"
import "testing"
import "unsafe"


func TestBufferWithByteSlice(t *testing.T) {
	b := make([]byte, 10, 10)
	buf := Buffer(b)
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

func TestBufferWithMap(t *testing.T) {
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
	Buffer(m)
}

func TestBufferWithChannel(t *testing.T) {
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
	Buffer(c)
}

func TestBufferWithInterface(t *testing.T) {
	var i interface{} = make([]byte, 16, 16)
	b := i.([]byte)
	buf := Buffer(i)
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

func TestBufferWithString(t *testing.T) {
	s := "hello"
	buf := Buffer(s)
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

func TestBufferWithStructValue(t *testing.T) {
	point := Point{ 3, 4, 5 }
	buf := Buffer(&point)

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

func TestBufferWithStructPointer(t *testing.T) {
	point := &Point{ 3, 4, 5 }
	buf := Buffer(point)

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

func TestBufferWithEmbeddedStructValue(t *testing.T) {
	point := Point{ 3, 4, 5 }
	tag := &TaggedPoint{ point, "this is a tag" }
	buf := Buffer(tag)

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

func TestBufferWithEmbeddedStructPointer(t *testing.T) {
	point := Point{ 3, 4, 5 }
	tag := &TaggedPointReference{ &point, "this is a tag" }
	buf := Buffer(tag)

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

func TestBufferWithInt32Slice(t *testing.T) {
	i := make([]int32, 10, 10)
	buf := Buffer(i)
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

func TestBufferWithInt64Slice(t *testing.T) {
	i := make([]int64, 10, 10)
	buf := Buffer(i)
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

func TestBufferWithFloat32Slice(t *testing.T) {
	f := make([]float32, 10, 10)
	buf := Buffer(f)
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

func TestBufferWithFloat64Slice(t *testing.T) {
	f := make([]float64, 10, 10)
	buf := Buffer(f)
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

func numericBufferTest(t *testing.T, value interface{}) {
	var size	int
	var addr	uintptr

	if v, ok := reflect.NewValue(value).(*reflect.PtrValue); ok {
		if v := v.Elem(); v != nil {
			size = int(v.Type().Size())
			addr = v.Addr()
		}
	} else {
		size = int(v.Type().Size())
		addr = v.Addr()
	}

	buf := Buffer(value)
	if len(buf) != size {
		t.Fatalf("byte buffer lengths differ: %v != %v", len(buf), size)
	}

	if cap(buf) != size {
		t.Fatalf("byte buffer capacities differ: %v != %v", cap(buf), size)
	}

	bufheader := *(*reflect.SliceHeader)(unsafe.Pointer(&buf))
	if addr != bufheader.Data {
		t.Fatalf("addresses don't match: %v != %v", addr, bufheader.Data)
	}
}

func TestBufferWithNumbers(t *testing.T) {
	var i8		int8
	var i16		int16
	var i32 	int32
	var i64		int64
	var u8		uint8
	var u16		uint16
	var u32 	uint32
	var u64		uint64
	var f32 	float32
	var f64		float64
	var c64		complex64
	var c128	complex64

	numericBufferTest(t, &i8)
	numericBufferTest(t, &i16)
	numericBufferTest(t, &i32)
	numericBufferTest(t, &i64)
	numericBufferTest(t, &u8)
	numericBufferTest(t, &u16)
	numericBufferTest(t, &u32)
	numericBufferTest(t, &u64)
	numericBufferTest(t, &f32)
	numericBufferTest(t, &f64)
	numericBufferTest(t, &c64)
	numericBufferTest(t, &c128)
}