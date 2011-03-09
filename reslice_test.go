package raw

import "reflect"
import "testing"
import "unsafe"

func TestSliceHeaderByte(t *testing.T) {
	slice := []byte{ 0, 1, 2, 3, 4, 5 }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	s, h := SliceHeader(slice)

	if s != _SIZE_BYTE { t.Fatalf("slice element size should be %v not %v", _SIZE_BYTE, s) }
	if h.Data != header.Data { t.Fatalf("slice headers should point to the same memory: %v - %v", h.Data, header.Data) }
	if h.Len != header.Len { t.Fatalf("slice header lengths should be the same: %v - %v", h.Len, header.Len) }
	if h.Cap != header.Cap { t.Fatalf("slice header capacities should be the same: %v - %v", h.Cap, header.Cap) }
}

func TestSliceHeaderUint(t *testing.T) {
	slice := []uint{ 0, 1, 2, 3, 4, 5 }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	s, h := SliceHeader(slice)

	if s != _SIZE_UINT { t.Fatalf("slice element size should be %v not %v", _SIZE_UINT, s) }
	if h.Data != header.Data { t.Fatalf("slice headers should point to the same memory: %v - %v", h.Data, header.Data) }
	if h.Len != header.Len { t.Fatalf("slice header lengths should be the same: %v - %v", h.Len, header.Len) }
	if h.Cap != header.Cap { t.Fatalf("slice header capacities should be the same: %v - %v", h.Cap, header.Cap) }
}

func TestSliceHeaderUint8(t *testing.T) {
	slice := []uint8{ 0, 1, 2, 3, 4, 5 }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	s, h := SliceHeader(slice)

	if s != _SIZE_UINT8 { t.Fatalf("slice element size should be %v not %v", _SIZE_UINT8, s) }
	if h.Data != header.Data { t.Fatalf("slice headers should point to the same memory: %v - %v", h.Data, header.Data) }
	if h.Len != header.Len { t.Fatalf("slice header lengths should be the same: %v - %v", h.Len, header.Len) }
	if h.Cap != header.Cap { t.Fatalf("slice header capacities should be the same: %v - %v", h.Cap, header.Cap) }
}

func TestSliceHeaderUint16(t *testing.T) {
	slice := []uint16{ 0, 1, 2, 3, 4, 5 }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	s, h := SliceHeader(slice)

	if s != _SIZE_UINT16 { t.Fatalf("slice element size should be %v not %v", _SIZE_UINT16, s) }
	if h.Data != header.Data { t.Fatalf("slice headers should point to the same memory: %v - %v", h.Data, header.Data) }
	if h.Len != header.Len { t.Fatalf("slice header lengths should be the same: %v - %v", h.Len, header.Len) }
	if h.Cap != header.Cap { t.Fatalf("slice header capacities should be the same: %v - %v", h.Cap, header.Cap) }
}

func TestSliceHeaderUint32(t *testing.T) {
	slice := []uint32{ 0, 1, 2, 3, 4, 5 }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	s, h := SliceHeader(slice)

	if s != _SIZE_UINT32 { t.Fatalf("slice element size should be %v not %v", _SIZE_UINT32, s) }
	if h.Data != header.Data { t.Fatalf("slice headers should point to the same memory: %v - %v", h.Data, header.Data) }
	if h.Len != header.Len { t.Fatalf("slice header lengths should be the same: %v - %v", h.Len, header.Len) }
	if h.Cap != header.Cap { t.Fatalf("slice header capacities should be the same: %v - %v", h.Cap, header.Cap) }
}


func TestSliceHeaderInt(t *testing.T) {
	slice := []int{ 0, 1, 2, 3, 4, 5 }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	s, h := SliceHeader(slice)

	if s != _SIZE_INT { t.Fatalf("slice element size should be %v not %v", _SIZE_INT, s) }
	if h.Data != header.Data { t.Fatalf("slice headers should point to the same memory: %v - %v", h.Data, header.Data) }
	if h.Len != header.Len { t.Fatalf("slice header lengths should be the same: %v - %v", h.Len, header.Len) }
	if h.Cap != header.Cap { t.Fatalf("slice header capacities should be the same: %v - %v", h.Cap, header.Cap) }
}

func TestSliceHeaderInt8(t *testing.T) {
	slice := []int8{ 0, 1, 2, 3, 4, 5 }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	s, h := SliceHeader(slice)

	if s != _SIZE_INT8 { t.Fatalf("slice element size should be %v not %v", _SIZE_INT8, s) }
	if h.Data != header.Data { t.Fatalf("slice headers should point to the same memory: %v - %v", h.Data, header.Data) }
	if h.Len != header.Len { t.Fatalf("slice header lengths should be the same: %v - %v", h.Len, header.Len) }
	if h.Cap != header.Cap { t.Fatalf("slice header capacities should be the same: %v - %v", h.Cap, header.Cap) }
}

func TestSliceHeaderInt16(t *testing.T) {
	slice := []int16{ 0, 1, 2, 3, 4, 5 }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	s, h := SliceHeader(slice)

	if s != _SIZE_INT16 { t.Fatalf("slice element size should be %v not %v", _SIZE_INT16, s) }
	if h.Data != header.Data { t.Fatalf("slice headers should point to the same memory: %v - %v", h.Data, header.Data) }
	if h.Len != header.Len { t.Fatalf("slice header lengths should be the same: %v - %v", h.Len, header.Len) }
	if h.Cap != header.Cap { t.Fatalf("slice header capacities should be the same: %v - %v", h.Cap, header.Cap) }
}

func TestSliceHeaderInt32(t *testing.T) {
	slice := []int32{ 0, 1, 2, 3, 4, 5 }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	s, h := SliceHeader(slice)

	if s != _SIZE_INT32 { t.Fatalf("slice element size should be %v not %v", _SIZE_INT32, s) }
	if h.Data != header.Data { t.Fatalf("slice headers should point to the same memory: %v - %v", h.Data, header.Data) }
	if h.Len != header.Len { t.Fatalf("slice header lengths should be the same: %v - %v", h.Len, header.Len) }
	if h.Cap != header.Cap { t.Fatalf("slice header capacities should be the same: %v - %v", h.Cap, header.Cap) }
}

func TestSliceHeaderFloat32(t *testing.T) {
	slice := []float32{ 0.0, 1.0, 2.0, 3.0, 4.0, 5.0 }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	s, h := SliceHeader(slice)

	if s != _SIZE_FLOAT32 { t.Fatalf("slice element size should be %v not %v", _SIZE_FLOAT32, s) }
	if h.Data != header.Data { t.Fatalf("slice headers should point to the same memory: %v - %v", h.Data, header.Data) }
	if h.Len != header.Len { t.Fatalf("slice header lengths should be the same: %v - %v", h.Len, header.Len) }
	if h.Cap != header.Cap { t.Fatalf("slice header capacities should be the same: %v - %v", h.Cap, header.Cap) }
}

func TestSliceHeaderFloat64(t *testing.T) {
	slice := []float64{ 0.0, 1.0, 2.0, 3.0, 4.0, 5.0 }
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	s, h := SliceHeader(slice)

	if s != _SIZE_FLOAT64 { t.Fatalf("slice element size should be %v not %v", _SIZE_FLOAT64, s) }
	if h.Data != header.Data { t.Fatalf("slice headers should point to the same memory: %v - %v", h.Data, header.Data) }
	if h.Len != header.Len { t.Fatalf("slice header lengths should be the same: %v - %v", h.Len, header.Len) }
	if h.Cap != header.Cap { t.Fatalf("slice header capacities should be the same: %v - %v", h.Cap, header.Cap) }
}