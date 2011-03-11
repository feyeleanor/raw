package raw

import "reflect"
import "testing"
import "unsafe"

func TestSliceHeaderByte(t *testing.T) {
	slice := []byte{ 0, 1, 2, 3, 4, 5 }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	h, s, al := SliceHeader(slice)

	if s != SIZE[BYTE] { t.Fatalf("slice element size should be %v not %v", SIZE[BYTE], s) }
	if al != ALIGNMENT[BYTE] { t.Fatalf("slice element alignment should be %v not %v", ALIGNMENT[BYTE], al) }
	if h.Data != header.Data { t.Fatalf("slice headers should point to the same memory: %v - %v", h.Data, header.Data) }
	if h.Len != header.Len { t.Fatalf("slice header lengths should be the same: %v - %v", h.Len, header.Len) }
	if h.Cap != header.Cap { t.Fatalf("slice header capacities should be the same: %v - %v", h.Cap, header.Cap) }
}

func TestSliceHeaderPointer(t *testing.T) {
	a, b, c, d, e, f := 0, 1, 2, 3, 4, 5
	slice := []unsafe.Pointer{ unsafe.Pointer(&a), unsafe.Pointer(&b), unsafe.Pointer(&c), unsafe.Pointer(&d), unsafe.Pointer(&e), unsafe.Pointer(&f) }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	h, s, al := SliceHeader(slice)

	if s != SIZE[POINTER] { t.Fatalf("slice element size should be %v not %v", SIZE[POINTER], s) }
	if al != ALIGNMENT[POINTER] { t.Fatalf("slice element alignment should be %v not %v", ALIGNMENT[POINTER], al) }
	if h.Data != header.Data { t.Fatalf("slice headers should point to the same memory: %v - %v", h.Data, header.Data) }
	if h.Len != header.Len { t.Fatalf("slice header lengths should be the same: %v - %v", h.Len, header.Len) }
	if h.Cap != header.Cap { t.Fatalf("slice header capacities should be the same: %v - %v", h.Cap, header.Cap) }
}

func TestSliceHeaderUintptr(t *testing.T) {
	a, b, c, d, e, f := 0, 1, 2, 3, 4, 5
	slice := []uintptr{ uintptr(unsafe.Pointer(&a)), uintptr(unsafe.Pointer(&b)), uintptr(unsafe.Pointer(&c)), uintptr(unsafe.Pointer(&d)), uintptr(unsafe.Pointer(&e)), uintptr(unsafe.Pointer(&f)) }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	h, s, al := SliceHeader(slice)

	if s != SIZE[UINTPTR] { t.Fatalf("slice element size should be %v not %v", SIZE[UINTPTR], s) }
	if al != ALIGNMENT[UINTPTR] { t.Fatalf("slice element alignment should be %v not %v", ALIGNMENT[UINTPTR], al) }
	if h.Data != header.Data { t.Fatalf("slice headers should point to the same memory: %v - %v", h.Data, header.Data) }
	if h.Len != header.Len { t.Fatalf("slice header lengths should be the same: %v - %v", h.Len, header.Len) }
	if h.Cap != header.Cap { t.Fatalf("slice header capacities should be the same: %v - %v", h.Cap, header.Cap) }
}

func TestSliceHeaderInterface(t *testing.T) {
	slice := []interface{}{ 0, 1, 2, 3, 4, 5 }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	h, s, al := SliceHeader(slice)

	if s != SIZE[INTERFACE] { t.Fatalf("slice element size should be %v not %v", SIZE[INTERFACE], s) }
	if al != ALIGNMENT[INTERFACE] { t.Fatalf("slice element alignment should be %v not %v", ALIGNMENT[INTERFACE], al) }
	if h.Data != header.Data { t.Fatalf("slice headers should point to the same memory: %v - %v", h.Data, header.Data) }
	if h.Len != header.Len { t.Fatalf("slice header lengths should be the same: %v - %v", h.Len, header.Len) }
	if h.Cap != header.Cap { t.Fatalf("slice header capacities should be the same: %v - %v", h.Cap, header.Cap) }
}

func TestSliceHeaderBool(t *testing.T) {
	slice := []bool{ true, false, true, false, false, true }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	h, s, al := SliceHeader(slice)

	if s != SIZE[BOOLEAN] { t.Fatalf("slice element size should be %v not %v", SIZE[BOOLEAN], s) }
	if al != ALIGNMENT[BOOLEAN] { t.Fatalf("slice element alignment should be %v not %v", ALIGNMENT[BOOLEAN], al) }
	if h.Data != header.Data { t.Fatalf("slice headers should point to the same memory: %v - %v", h.Data, header.Data) }
	if h.Len != header.Len { t.Fatalf("slice header lengths should be the same: %v - %v", h.Len, header.Len) }
	if h.Cap != header.Cap { t.Fatalf("slice header capacities should be the same: %v - %v", h.Cap, header.Cap) }
}

func TestSliceHeaderUint(t *testing.T) {
	slice := []uint{ 0, 1, 2, 3, 4, 5 }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	h, s, al := SliceHeader(slice)

	if s != SIZE[UINT] { t.Fatalf("slice element size should be %v not %v", SIZE[UINT], s) }
	if al != ALIGNMENT[UINT] { t.Fatalf("slice element alignment should be %v not %v", ALIGNMENT[UINT], al) }
	if h.Data != header.Data { t.Fatalf("slice headers should point to the same memory: %v - %v", h.Data, header.Data) }
	if h.Len != header.Len { t.Fatalf("slice header lengths should be the same: %v - %v", h.Len, header.Len) }
	if h.Cap != header.Cap { t.Fatalf("slice header capacities should be the same: %v - %v", h.Cap, header.Cap) }
}

func TestSliceHeaderUint8(t *testing.T) {
	slice := []uint8{ 0, 1, 2, 3, 4, 5 }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	h, s, al := SliceHeader(slice)

	if s != SIZE[UINT8] { t.Fatalf("slice element size should be %v not %v", SIZE[UINT8], s) }
	if al != ALIGNMENT[UINT8] { t.Fatalf("slice element alignment should be %v not %v", ALIGNMENT[UINT8], al) }
	if h.Data != header.Data { t.Fatalf("slice headers should point to the same memory: %v - %v", h.Data, header.Data) }
	if h.Len != header.Len { t.Fatalf("slice header lengths should be the same: %v - %v", h.Len, header.Len) }
	if h.Cap != header.Cap { t.Fatalf("slice header capacities should be the same: %v - %v", h.Cap, header.Cap) }
}

func TestSliceHeaderUint16(t *testing.T) {
	slice := []uint16{ 0, 1, 2, 3, 4, 5 }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	h, s, al := SliceHeader(slice)

	if s != SIZE[UINT16] { t.Fatalf("slice element size should be %v not %v", SIZE[UINT16], s) }
	if al != ALIGNMENT[UINT16] { t.Fatalf("slice element alignment should be %v not %v", ALIGNMENT[UINT16], al) }
	if h.Data != header.Data { t.Fatalf("slice headers should point to the same memory: %v - %v", h.Data, header.Data) }
	if h.Len != header.Len { t.Fatalf("slice header lengths should be the same: %v - %v", h.Len, header.Len) }
	if h.Cap != header.Cap { t.Fatalf("slice header capacities should be the same: %v - %v", h.Cap, header.Cap) }
}

func TestSliceHeaderUint32(t *testing.T) {
	slice := []uint32{ 0, 1, 2, 3, 4, 5 }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	h, s, al := SliceHeader(slice)

	if s != SIZE[UINT32] { t.Fatalf("slice element size should be %v not %v", SIZE[UINT32], s) }
	if al != ALIGNMENT[UINT32] { t.Fatalf("slice element alignment should be %v not %v", ALIGNMENT[UINT32], al) }
	if h.Data != header.Data { t.Fatalf("slice headers should point to the same memory: %v - %v", h.Data, header.Data) }
	if h.Len != header.Len { t.Fatalf("slice header lengths should be the same: %v - %v", h.Len, header.Len) }
	if h.Cap != header.Cap { t.Fatalf("slice header capacities should be the same: %v - %v", h.Cap, header.Cap) }
}


func TestSliceHeaderInt(t *testing.T) {
	slice := []int{ 0, 1, 2, 3, 4, 5 }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	h, s, al := SliceHeader(slice)

	if s != SIZE[INT] { t.Fatalf("slice element size should be %v not %v", SIZE[INT], s) }
	if al != ALIGNMENT[INT] { t.Fatalf("slice element alignment should be %v not %v", ALIGNMENT[INT], al) }
	if h.Data != header.Data { t.Fatalf("slice headers should point to the same memory: %v - %v", h.Data, header.Data) }
	if h.Len != header.Len { t.Fatalf("slice header lengths should be the same: %v - %v", h.Len, header.Len) }
	if h.Cap != header.Cap { t.Fatalf("slice header capacities should be the same: %v - %v", h.Cap, header.Cap) }
}

func TestSliceHeaderInt8(t *testing.T) {
	slice := []int8{ 0, 1, 2, 3, 4, 5 }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	h, s, al := SliceHeader(slice)

	if s != SIZE[INT8] { t.Fatalf("slice element size should be %v not %v", SIZE[INT8], s) }
	if al != ALIGNMENT[INT8] { t.Fatalf("slice element alignment should be %v not %v", ALIGNMENT[INT8], al) }
	if h.Data != header.Data { t.Fatalf("slice headers should point to the same memory: %v - %v", h.Data, header.Data) }
	if h.Len != header.Len { t.Fatalf("slice header lengths should be the same: %v - %v", h.Len, header.Len) }
	if h.Cap != header.Cap { t.Fatalf("slice header capacities should be the same: %v - %v", h.Cap, header.Cap) }
}

func TestSliceHeaderInt16(t *testing.T) {
	slice := []int16{ 0, 1, 2, 3, 4, 5 }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	h, s, al := SliceHeader(slice)

	if s != SIZE[INT16] { t.Fatalf("slice element size should be %v not %v", SIZE[INT16], s) }
	if al != ALIGNMENT[INT16] { t.Fatalf("slice element alignment should be %v not %v", ALIGNMENT[INT16], al) }
	if h.Data != header.Data { t.Fatalf("slice headers should point to the same memory: %v - %v", h.Data, header.Data) }
	if h.Len != header.Len { t.Fatalf("slice header lengths should be the same: %v - %v", h.Len, header.Len) }
	if h.Cap != header.Cap { t.Fatalf("slice header capacities should be the same: %v - %v", h.Cap, header.Cap) }
}

func TestSliceHeaderInt32(t *testing.T) {
	slice := []int32{ 0, 1, 2, 3, 4, 5 }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	h, s, al := SliceHeader(slice)

	if s != SIZE[INT32] { t.Fatalf("slice element size should be %v not %v", SIZE[INT32], s) }
	if al != ALIGNMENT[INT32] { t.Fatalf("slice element alignment should be %v not %v", ALIGNMENT[INT32], al) }
	if h.Data != header.Data { t.Fatalf("slice headers should point to the same memory: %v - %v", h.Data, header.Data) }
	if h.Len != header.Len { t.Fatalf("slice header lengths should be the same: %v - %v", h.Len, header.Len) }
	if h.Cap != header.Cap { t.Fatalf("slice header capacities should be the same: %v - %v", h.Cap, header.Cap) }
}

func TestSliceHeaderFloat32(t *testing.T) {
	slice := []float32{ 0.0, 1.0, 2.0, 3.0, 4.0, 5.0 }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	h, s, al := SliceHeader(slice)

	if s != SIZE[FLOAT32] { t.Fatalf("slice element size should be %v not %v", SIZE[FLOAT32], s) }
	if al != ALIGNMENT[FLOAT32] { t.Fatalf("slice element alignment should be %v not %v", ALIGNMENT[FLOAT32], al) }
	if h.Data != header.Data { t.Fatalf("slice headers should point to the same memory: %v - %v", h.Data, header.Data) }
	if h.Len != header.Len { t.Fatalf("slice header lengths should be the same: %v - %v", h.Len, header.Len) }
	if h.Cap != header.Cap { t.Fatalf("slice header capacities should be the same: %v - %v", h.Cap, header.Cap) }
}

func TestSliceHeaderFloat64(t *testing.T) {
	slice := []float64{ 0.0, 1.0, 2.0, 3.0, 4.0, 5.0 }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	h, s, al := SliceHeader(slice)

	if s != SIZE[FLOAT64] { t.Fatalf("slice element size should be %v not %v", SIZE[FLOAT64], s) }
	if al != ALIGNMENT[FLOAT64] { t.Fatalf("slice element alignment should be %v not %v", ALIGNMENT[FLOAT64], al) }
	if h.Data != header.Data { t.Fatalf("slice headers should point to the same memory: %v - %v", h.Data, header.Data) }
	if h.Len != header.Len { t.Fatalf("slice header lengths should be the same: %v - %v", h.Len, header.Len) }
	if h.Cap != header.Cap { t.Fatalf("slice header capacities should be the same: %v - %v", h.Cap, header.Cap) }
}

func TestScale(t *testing.T) {
	t.Fatal("implement test")
}

func TestReslice(t *testing.T) {
	t.Fatal("implement test")
}

func TestPointerSlice(t *testing.T) {
	t.Fatal("implement test")
}

func TestUintptrSlice(t *testing.T) {
	t.Fatal("implement test")
}

func TestInterfaceSlice(t *testing.T) {
	t.Fatal("implement test")
}

func TestBoolSlice(t *testing.T) {
	t.Fatal("implement test")
}

func TestIntSlice(t *testing.T) {
	t.Fatal("implement test")
}

func TestInt8Slice(t *testing.T) {
	t.Fatal("implement test")
}

func TestInt16Slice(t *testing.T) {
	t.Fatal("implement test")
}

func TestInt32Slice(t *testing.T) {
	t.Fatal("implement test")
}

func TestInt64Slice(t *testing.T) {
	t.Fatal("implement test")
}

func TestUintSlice(t *testing.T) {
	t.Fatal("implement test")
}

func TestUint8Slice(t *testing.T) {
	t.Fatal("implement test")
}

func TestUint16Slice(t *testing.T) {
	t.Fatal("implement test")
}

func TestUint32Slice(t *testing.T) {
	t.Fatal("implement test")
}

func TestUint64Slice(t *testing.T) {
	t.Fatal("implement test")
}

func TestFloat32Slice(t *testing.T) {
	t.Fatal("implement test")
}

func TestFloat64Slice(t *testing.T) {
	t.Fatal("implement test")
}

func TestComplex64Slice(t *testing.T) {
	t.Fatal("implement test")
}

func TestComplex128Slice(t *testing.T) {
	t.Fatal("implement test")
}