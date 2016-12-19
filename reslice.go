package raw

import (
	"reflect"
	"unsafe"
)

func sliceHeaderFromValue(v reflect.Value) (s *reflect.SliceHeader) {
	switch v.Kind() {
	case reflect.Slice:						if !v.CanAddr() {
												x := reflect.New(v.Type()).Elem()
												x.Set(v)
												v = x
											}
											s = (*reflect.SliceHeader)(unsafe.Pointer(v.UnsafeAddr()))
	case reflect.Ptr, reflect.Interface:	s = sliceHeaderFromValue(v.Elem())
	}
	return
}

func SliceHeader(i interface{}) (Header *reflect.SliceHeader, ElementSize, ElementAlignment int) {
	value := reflect.ValueOf(i)
	if Header = sliceHeaderFromValue(value); Header != nil {
		ElementType := value.Type().Elem()
		ElementSize = int(ElementType.Size())
		ElementAlignment = int(ElementType.Align())
	} else {
		panic(i)
	}
	return
}

func Scale(oldHeader *reflect.SliceHeader, oldElementSize, newElementSize uintptr) (h *reflect.SliceHeader) {
	if oldHeader != nil {
		s := float64(oldElementSize) / float64(newElementSize)
		h = &reflect.SliceHeader{ Data: oldHeader.Data }
		h.Len = int(float64(oldHeader.Len) * s)
		h.Cap = int(float64(oldHeader.Cap) * s)
	}
	return
}

func Reslice(slice interface{}, sliceType reflect.Type) interface{} {
	b := ByteSlice(slice)
	h := Scale(&reflect.SliceHeader{ uintptr(DataAddress(b)), len(b), cap(b) }, 1, sliceType.Elem().Size())
	return reflect.NewAt(sliceType, unsafe.Pointer(h)).Elem().Interface()
}

func PointerSlice(i interface{}) []unsafe.Pointer {
	return Reslice(i, reflect.SliceOf(POINTER)).([]unsafe.Pointer)
}

func UintptrSlice(i interface{}) []uintptr {
	return Reslice(i, reflect.SliceOf(UINTPTR)).([]uintptr)
}

func InterfaceSlice(i interface{}) []interface{} {
	return Reslice(i, reflect.SliceOf(INTERFACE)).([]interface{})
}

func BoolSlice(i interface{}) []bool {
	return Reslice(i, reflect.SliceOf(BOOLEAN)).([]bool)
}

func IntSlice(i interface{}) []int {
	if i, ok := i.([]uint); ok {
		return *(*[]int)(unsafe.Pointer(&i))
	}
	return Reslice(i, reflect.SliceOf(INT)).([]int)
}

func Int8Slice(i interface{}) []int8 {
	if i, ok := i.([]uint8); ok {
		return *(*[]int8)(unsafe.Pointer(&i))
	}
	return Reslice(i, reflect.SliceOf(INT8)).([]int8)
}

func Int16Slice(i interface{}) []int16 {
	if i, ok := i.([]uint16); ok {
		return *(*[]int16)(unsafe.Pointer(&i))
	}
	return Reslice(i, reflect.SliceOf(INT16)).([]int16)
}

func Int32Slice(i interface{}) []int32 {
	switch i := i.(type) {
	case []uint32:
		return *(*[]int32)(unsafe.Pointer(&i))
	case []float32:
		return *(*[]int32)(unsafe.Pointer(&i))
	}
	return Reslice(i, reflect.SliceOf(INT32)).([]int32)
}

func Int64Slice(i interface{}) []int64 {
	switch i := i.(type) {
	case []uint64:
		return *(*[]int64)(unsafe.Pointer(&i))
	case []float64:
		return *(*[]int64)(unsafe.Pointer(&i))
	}
	return Reslice(i, reflect.SliceOf(INT64)).([]int64)
}

func UintSlice(i interface{}) []uint {
	if i, ok := i.([]int); ok {
		return *(*[]uint)(unsafe.Pointer(&i))
	}
	return Reslice(i, reflect.SliceOf(UINT)).([]uint)
}

func Uint8Slice(i interface{}) []uint8 {
	if i, ok := i.([]int8); ok {
		return *(*[]uint8)(unsafe.Pointer(&i))
	}
	return Reslice(i, reflect.SliceOf(UINT8)).([]uint8)
}

func Uint16Slice(i interface{}) []uint16 {
	if i, ok := i.([]int16); ok {
		return *(*[]uint16)(unsafe.Pointer(&i))
	}
	return Reslice(i, reflect.SliceOf(UINT16)).([]uint16)
}

func Uint32Slice(i interface{}) []uint32 {
	switch i := i.(type) {
	case []int32:
		return *(*[]uint32)(unsafe.Pointer(&i))
	case []float32:
		return *(*[]uint32)(unsafe.Pointer(&i))
	}
	return Reslice(i, reflect.SliceOf(UINT32)).([]uint32)
}

func Uint64Slice(i interface{}) []uint64 {
	switch i := i.(type) {
	case []int64:
		return *(*[]uint64)(unsafe.Pointer(&i))
	case []float64:
		return *(*[]uint64)(unsafe.Pointer(&i))
	}
	return Reslice(i, reflect.SliceOf(UINT64)).([]uint64)
}

func Float32Slice(i interface{}) []float32 {
	switch i := i.(type) {
	case []int32:
		return *(*[]float32)(unsafe.Pointer(&i))
	case []uint32:
		return *(*[]float32)(unsafe.Pointer(&i))
	}
	return Reslice(i, reflect.SliceOf(FLOAT32)).([]float32)
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
	return Reslice(i, reflect.SliceOf(FLOAT64)).([]float64)
}

func Complex64Slice(i interface{}) []complex64 {
	switch i := i.(type) {
	case []int64:
		return *(*[]complex64)(unsafe.Pointer(&i))
	case []uint64:
		return *(*[]complex64)(unsafe.Pointer(&i))
	case []float64:
		return *(*[]complex64)(unsafe.Pointer(&i))
	}
	return Reslice(i, reflect.SliceOf(COMPLEX64)).([]complex64)
}

func Complex128Slice(i interface{}) []complex128 {
	return Reslice(i, reflect.SliceOf(COMPLEX128)).([]complex128)
}