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

func TestBufferWithEmbeddedStruct(t *testing.T) {
//	p := &Point{ 3, 4, 5 }
//	t := &TaggedPoint{ p, "this is a tag" }
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

