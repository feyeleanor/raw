package raw

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)


type HeaderMismatch struct{
	message			string
	x, y, z			interface{}
}

func (h HeaderMismatch) Error() string {
	return fmt.Sprintf(h.message, h.x, h.y, h.z)
}

func MatchHeaders(b reflect.Type, slice interface{}, header reflect.SliceHeader) (e error) {
	h, s, al := SliceHeader(slice)
	switch {
	case s != int(b.Size()):			e = HeaderMismatch{ "%v: slice element size should be %v not %v", b.Name(), b.Size(), s }
	case al != b.Align():		e = HeaderMismatch{ "%v: slice element alignment should be %v not %v", b.Name(), b.Align(), al }
	case h.Data != header.Data:	e = HeaderMismatch{ "%v: slice headers should point to the same memory: %v - %v", b.Name(), h.Data, header.Data }
	case h.Len != header.Len:	e = HeaderMismatch{ "%v: slice header lengths should be the same: %v - %v", b.Name(), h.Len, header.Len }
	case h.Cap != header.Cap:	e = HeaderMismatch{ "%v: slice header capacities should be the same: %v - %v", b.Name(), h.Cap, header.Cap }
	}
	return
}

func TestSliceHeaderByte(t *testing.T) {
	slice := []byte{ 0, 1, 2, 3, 4, 5 }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	if m := MatchHeaders(BYTE, slice, header); m != nil {
		t.Fatal(m)
	}
}

func TestSliceHeaderPointer(t *testing.T) {
	a, b, c, d, e, f := 0, 1, 2, 3, 4, 5
	slice := []unsafe.Pointer{ unsafe.Pointer(&a), unsafe.Pointer(&b), unsafe.Pointer(&c), unsafe.Pointer(&d), unsafe.Pointer(&e), unsafe.Pointer(&f) }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	if m := MatchHeaders(POINTER, slice, header); m != nil {
		t.Fatal(m)
	}
}

func TestSliceHeaderUintptr(t *testing.T) {
	slice := []uintptr{ 0, 1, 2, 3, 4, 5 }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	if m := MatchHeaders(UINTPTR, slice, header); m != nil {
		t.Fatal(m)
	}
}

func TestSliceHeaderInterface(t *testing.T) {
	t.Log("Awaiting bug fix for incorrect reporting of interface{} value size with unsafe.Sizeof()")
/*
	slice := []interface{}{ 0, 1, 2, 3, 4, 5 }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	if m := MatchHeaders(INTERFACE, slice, header); m != nil {
		t.Fatal(m)
	}
*/
}

func TestSliceHeaderBool(t *testing.T) {
	slice := []bool{ true, false, true, false, false, true }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	if m := MatchHeaders(BOOLEAN, slice, header); m != nil {
		t.Fatal(m)
	}
}

func TestSliceHeaderUint(t *testing.T) {
	slice := []uint{ 0, 1, 2, 3, 4, 5 }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	if m := MatchHeaders(UINT, slice, header); m != nil {
		t.Fatal(m)
	}
}

func TestSliceHeaderUint8(t *testing.T) {
	slice := []uint8{ 0, 1, 2, 3, 4, 5 }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	if m := MatchHeaders(UINT8, slice, header); m != nil {
		t.Fatal(m)
	}
}

func TestSliceHeaderUint16(t *testing.T) {
	slice := []uint16{ 0, 1, 2, 3, 4, 5 }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	if m := MatchHeaders(UINT16, slice, header); m != nil {
		t.Fatal(m)
	}
}

func TestSliceHeaderUint32(t *testing.T) {
	slice := []uint32{ 0, 1, 2, 3, 4, 5 }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	if m := MatchHeaders(UINT32, slice, header); m != nil {
		t.Fatal(m)
	}
}

func TestSliceHeaderInt(t *testing.T) {
	slice := []int{ 0, 1, 2, 3, 4, 5 }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	if m := MatchHeaders(INT, slice, header); m != nil {
		t.Fatal(m)
	}
}

func TestSliceHeaderInt8(t *testing.T) {
	slice := []int8{ 0, 1, 2, 3, 4, 5 }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	if m := MatchHeaders(INT8, slice, header); m != nil {
		t.Fatal(m)
	}
}

func TestSliceHeaderInt16(t *testing.T) {
	slice := []int16{ 0, 1, 2, 3, 4, 5 }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	if m := MatchHeaders(INT16, slice, header); m != nil {
		t.Fatal(m)
	}
}

func TestSliceHeaderInt32(t *testing.T) {
	slice := []int32{ 0, 1, 2, 3, 4, 5 }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	if m := MatchHeaders(INT32, slice, header); m != nil {
		t.Fatal(m)
	}
}

func TestSliceHeaderFloat32(t *testing.T) {
	slice := []float32{ 0.0, 1.0, 2.0, 3.0, 4.0, 5.0 }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	if m := MatchHeaders(FLOAT32, slice, header); m != nil {
		t.Fatal(m)
	}
}

func TestSliceHeaderFloat64(t *testing.T) {
	slice := []float64{ 0.0, 1.0, 2.0, 3.0, 4.0, 5.0 }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	if m := MatchHeaders(FLOAT64, slice, header); m != nil {
		t.Fatal(m)
	}
}

func TestSliceHeaderComplex64(t *testing.T) {
	slice := []complex64{ 0, 0, 0, 0, 0, 0 }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	if m := MatchHeaders(COMPLEX64, slice, header); m != nil {
		t.Fatal(m)
	}
}

func TestSliceHeaderComplex128(t *testing.T) {
	slice := []complex128{ 0, 0, 0, 0, 0, 0 }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	if m := MatchHeaders(COMPLEX128, slice, header); m != nil {
		t.Fatal(m)
	}
}

func TestScale(t *testing.T) {
	var h 	*reflect.SliceHeader

	if Scale(h, 0, 0) != nil {
		t.Fatal("Scaling a nil *SliceHeader should return a nil *SliceHeader")
	}

	b := []byte{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	h, _, _ = SliceHeader(b)
	hs := Scale(h, 1, 2)

	FailOnBadBufferSize(t,
		hs.Len == h.Len / 2,
		hs.Cap == h.Cap / 2,
	)

	i32 := []int32{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	h, _, _ = SliceHeader(i32)
	hs = Scale(h, INT32.Size(), BYTE.Size())
	switch 	is, bs := int(INT32.Size()), int(BYTE.Size()); {
	case hs.Len != h.Len * is / bs:
		t.Fatalf("new slice length should be %v not %v", h.Len * is / bs, hs.Len)
	case hs.Cap != h.Cap * is / bs:
		t.Fatalf("new slice capacity should be %v not %v", h.Cap * is / bs, hs.Cap)
	}
}

func TestResliceNil(t *testing.T) {
	var h 	*reflect.SliceHeader

	hs := &reflect.SliceHeader{}
	h, _, _ = SliceHeader(Reslice(h, reflect.SliceOf(INT32)))

	FailOnBadBufferSize(t,
		h.Len == hs.Len,
		hs.Cap == hs.Cap,
	)
}

func TestReslice(t *testing.T) {
	b := []byte{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15 }
	s := Reslice(b, reflect.SliceOf(INT32)).([]int32)
	is := int(INT32.Size())
	FailOnBadBufferSize(t,
		len(s) == len(b) / is,
		cap(s) == cap(b) / is,
	)
}

func TestPointerSlice(t *testing.T) {
	b := []byte{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15 }
	p := PointerSlice(b)
	ps := int(POINTER.Size())
	FailOnBadBufferSize(t,
		len(p) == len(b) / ps,
		len(p) == cap(b) / ps,
	)
}

func TestUintptrSlice(t *testing.T) {
	b := []byte{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15 }
	p := UintptrSlice(b)
	us := int(UINTPTR.Size())
	FailOnBadBufferSize(t,
		len(p) == len(b) / us,
		len(p) == cap(b) / us,
	)
}

func TestInterfaceSlice(t *testing.T) {
	t.Fatal("broken by replacement of BasicType with reflect.Type")
	b := []byte{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15 }
	p := InterfaceSlice(b)
	is := int(INTERFACE.Size())
	FailOnBadBufferSize(t,
		len(p) == len(b) / is,
		len(p) == cap(b) / is,
	)
}

func TestBoolSlice(t *testing.T) {
	b := []byte{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15 }
	p := BoolSlice(b)
	bs := int(BOOLEAN.Size())
	FailOnBadBufferSize(t,
		len(p) == len(b) / bs,
		len(p) == cap(b) / bs,
	)
}

func TestIntSlice(t *testing.T) {
	b := []byte{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15 }
	p := IntSlice(b)
	is := int(INT.Size())
	FailOnBadBufferSize(t,
		len(p) == len(b) / is,
		len(p) == cap(b) / is,
	)
}

func TestInt8Slice(t *testing.T) {
	b := []byte{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15 }
	p := Int8Slice(b)
	is := int(INT8.Size())
	FailOnBadBufferSize(t,
		len(p) == len(b) / is,
		len(p) == cap(b) / is,
	)
}

func TestInt16Slice(t *testing.T) {
	b := []byte{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15 }
	p := Int16Slice(b)
	is := int(INT16.Size())
	FailOnBadBufferSize(t,
		len(p) == len(b) / is,
		len(p) == cap(b) / is,
	)
}

func TestInt32Slice(t *testing.T) {
	b := []byte{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15 }
	p := Int32Slice(b)
	is := int(INT32.Size())
	FailOnBadBufferSize(t,
		len(p) == len(b) / is,
		len(p) == cap(b) / is,
	)
}

func TestInt64Slice(t *testing.T) {
	b := []byte{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15 }
	p := Int64Slice(b)
	is := int(INT64.Size())
	FailOnBadBufferSize(t,
		len(p) == len(b) / is,
		len(p) == cap(b) / is,
	)
}

func TestUintSlice(t *testing.T) {
	b := []byte{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15 }
	p := UintSlice(b)
	us := int(UINT.Size())
	FailOnBadBufferSize(t,
		len(p) == len(b) / us,
		len(p) == cap(b) / us,
	)
}

func TestUint8Slice(t *testing.T) {
	b := []byte{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15 }
	p := Uint8Slice(b)
	us := int(UINT8.Size())
	FailOnBadBufferSize(t,
		len(p) == len(b) / us,
		len(p) == cap(b) / us,
	)
}

func TestUint16Slice(t *testing.T) {
	b := []byte{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15 }
	p := Uint16Slice(b)
	us := int(UINT16.Size())
	FailOnBadBufferSize(t,
		len(p) == len(b) / us,
		len(p) == cap(b) / us,
	)
}

func TestUint32Slice(t *testing.T) {
	b := []byte{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15 }
	p := Uint32Slice(b)
	us := int(UINT32.Size())
	FailOnBadBufferSize(t,
		len(p) == len(b) / us,
		len(p) == cap(b) / us,
	)
}

func TestUint64Slice(t *testing.T) {
	b := []byte{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15 }
	p := Uint64Slice(b)
	us := int(UINT64.Size())
	FailOnBadBufferSize(t,
		len(p) == len(b) / us,
		len(p) == cap(b) / us,
	)
}

func TestFloat32Slice(t *testing.T) {
	b := []byte{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15 }
	p := Float32Slice(b)
	fs := int(FLOAT32.Size())
	FailOnBadBufferSize(t,
		len(p) == len(b) / fs,
		len(p) == cap(b) / fs,
	)
}

func TestFloat64Slice(t *testing.T) {
	b := []byte{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15 }
	p := Float64Slice(b)
	fs := int(FLOAT64.Size())
	FailOnBadBufferSize(t,
		len(p) == len(b) / fs,
		len(p) == cap(b) / fs,
	)
}

func TestComplex64Slice(t *testing.T) {
	b := []byte{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15 }
	p := Complex64Slice(b)
	cs := int(COMPLEX64.Size())
	FailOnBadBufferSize(t,
		len(p) == len(b) / cs,
		len(p) == cap(b) / cs,
	)
}

func TestComplex128Slice(t *testing.T) {
	b := []byte{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15 }
	p := Complex128Slice(b)
	cs := int(COMPLEX128.Size())
	FailOnBadBufferSize(t,
		len(p) == len(b) / cs,
		len(p) == cap(b) / cs,
	)
}