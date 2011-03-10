package raw

import "C"
import "reflect"
import "unsafe"

func SliceHeader(i interface{}) (Header *reflect.SliceHeader, ElementSize, ElementAlignment int) {
	value := reflect.NewValue(i)
	switch value := value.(type) {
	case *reflect.PtrValue:
		Header, ElementSize, ElementAlignment = SliceHeader(value.Elem())

	case *reflect.InterfaceValue:
		Header, ElementSize, ElementAlignment = SliceHeader(value.Elem())

	case *reflect.SliceValue:
		Header = (*reflect.SliceHeader)(unsafe.Pointer(value.UnsafeAddr()))
		ElementType := value.Type().(*reflect.SliceType).Elem()
		ElementSize = int(ElementType.Size())
		ElementAlignment = ElementType.Align()
	}
	return
}

func Scale(h *reflect.SliceHeader, oldElementSize, newElementSize int) *reflect.SliceHeader {
	s := float64(oldElementSize) / float64(newElementSize)
	return &reflect.SliceHeader{ Data: h.Data, Len: int(float64(h.Len) * s), Cap: int(float64(h.Cap) * s) }
}

func Reslice(slice, sliceType interface{}, elementSize int) interface{} {
	b := ByteSlice(slice)
	h := Scale(&reflect.SliceHeader{ uintptr(unsafe.Pointer(&(b[0]))), len(b), cap(b) }, 1, elementSize)
	return unsafe.Unreflect(sliceType, unsafe.Pointer(h))
}

func PointerSlice(i interface{}) []unsafe.Pointer {
	return Reslice(i, SLICE_TYPE[POINTER], SIZE[POINTER]).([]unsafe.Pointer)
}

func UintptrSlice(i interface{}) []uintptr {
	return Reslice(i, SLICE_TYPE[UINTPTR], SIZE[UINTPTR]).([]uintptr)
}

func InterfaceSlice(i interface{}) []interface{} {
	return Reslice(i, SLICE_TYPE[INTERFACE], SIZE[INTERFACE]).([]interface{})
}

func BoolSlice(i interface{}) []bool {
	return Reslice(i, SLICE_TYPE[BOOLEAN], SIZE[BOOLEAN]).([]bool)
}

func IntSlice(i interface{}) []int {
	if i, ok := i.([]uint); ok {
		return *(*[]int)(unsafe.Pointer(&i))
	}
	return Reslice(i, SLICE_TYPE[INT], SIZE[INT]).([]int)
}

func Int8Slice(i interface{}) []int8 {
	if i, ok := i.([]uint8); ok {
		return *(*[]int8)(unsafe.Pointer(&i))
	}
	return Reslice(i, SLICE_TYPE[INT8], SIZE[INT8]).([]int8)
}

func Int16Slice(i interface{}) []int16 {
	if i, ok := i.([]uint16); ok {
		return *(*[]int16)(unsafe.Pointer(&i))
	}
	return Reslice(i, SLICE_TYPE[INT16], SIZE[INT16]).([]int16)
}

func Int32Slice(i interface{}) []int32 {
	switch i := i.(type) {
	case []uint32:
		return *(*[]int32)(unsafe.Pointer(&i))
	case []float32:
		return *(*[]int32)(unsafe.Pointer(&i))
	}
	return Reslice(i, SLICE_TYPE[INT32], SIZE[INT32]).([]int32)
}

func UintSlice(i interface{}) []uint {
	if i, ok := i.([]int); ok {
		return *(*[]uint)(unsafe.Pointer(&i))
	}
	return Reslice(i, SLICE_TYPE[UINT], SIZE[UINT]).([]uint)
}

func Uint8Slice(i interface{}) []uint8 {
	if i, ok := i.([]int8); ok {
		return *(*[]uint8)(unsafe.Pointer(&i))
	}
	return Reslice(i, SLICE_TYPE[UINT8], SIZE[UINT8]).([]uint8)
}

func Uint16Slice(i interface{}) []uint16 {
	if i, ok := i.([]int16); ok {
		return *(*[]uint16)(unsafe.Pointer(&i))
	}
	return Reslice(i, SLICE_TYPE[UINT16], SIZE[UINT16]).([]uint16)
}

func Uint32Slice(i interface{}) []uint32 {
	switch i := i.(type) {
	case []int32:
		return *(*[]uint32)(unsafe.Pointer(&i))
	case []float32:
		return *(*[]uint32)(unsafe.Pointer(&i))
	}
	return Reslice(i, SLICE_TYPE[UINT32], SIZE[UINT32]).([]uint32)
}

func Float32Slice(i interface{}) []float32 {
	switch i := i.(type) {
	case []int32:
		return *(*[]float32)(unsafe.Pointer(&i))
	case []uint32:
		return *(*[]float32)(unsafe.Pointer(&i))
	}
	return Reslice(i, SLICE_TYPE[FLOAT32], SIZE[FLOAT32]).([]float32)
}

func Float64Slice(i interface{}) []float64 {
	switch i := i.(type) {
	case []int64:
		return *(*[]float64)(unsafe.Pointer(&i))
	case []uint64:
		return *(*[]float64)(unsafe.Pointer(&i))
	case []complex64:
		return *(*[]float64)(unsafe.Pointer(&i))
	}
	return Reslice(i, SLICE_TYPE[FLOAT64], SIZE[FLOAT64]).([]float64)
}

/*
func intsToFloats(i []int) []float64 { return *(*[]float64)(unsafe.Pointer(&i)) }
func floatsToInts(f []float64) []int { return *(*[]int)(unsafe.Pointer(&f)) }
*/