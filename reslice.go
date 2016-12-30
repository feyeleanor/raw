package raw

import (
	"reflect"
	"unsafe"
)

func sliceHeaderFromValue(v reflect.Value) (s *reflect.SliceHeader) {
	switch v.Kind() {
	case reflect.Slice:
		if !v.CanAddr() {
			x := reflect.New(v.Type()).Elem()
			x.Set(v)
			v = x
		}
		s = (*reflect.SliceHeader)(unsafe.Pointer(v.UnsafeAddr()))
	case reflect.Ptr, reflect.Interface:
		s = sliceHeaderFromValue(v.Elem())
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

func ResliceOf(t reflect.Type, i interface{}) interface{} {
	return Reslice(i, reflect.SliceOf(t))
}

func PointerSlice(i interface{}) []unsafe.Pointer {
	return ResliceOf(POINTER, i).([]unsafe.Pointer)
}

func UintptrSlice(i interface{}) []uintptr {
	return ResliceOf(UINTPTR, i).([]uintptr)
}

func InterfaceSlice(i interface{}) []interface{} {
	return ResliceOf(INTERFACE, i).([]interface{})
}

func BoolSlice(i interface{}) []bool {
	return ResliceOf(BOOLEAN, i).([]bool)
}

func IntSlice(i interface{}) []int {
	if i, ok := i.([]uint); ok {
		return *(*[]int)(unsafe.Pointer(&i))
	}
	return ResliceOf(INT, i).([]int)
}

func Int8Slice(i interface{}) []int8 {
	if i, ok := i.([]uint8); ok {
		return *(*[]int8)(unsafe.Pointer(&i))
	}
	return ResliceOf(INT8, i).([]int8)
}

func Int16Slice(i interface{}) []int16 {
	if i, ok := i.([]uint16); ok {
		return *(*[]int16)(unsafe.Pointer(&i))
	}
	return ResliceOf(INT16, i).([]int16)
}

func Int32Slice(i interface{}) (r []int32) {
	switch i := i.(type) {
	case []uint32:
		r = *(*[]int32)(unsafe.Pointer(&i))
	case []float32:
		r = *(*[]int32)(unsafe.Pointer(&i))
	default:
		r = ResliceOf(INT32, i).([]int32)
	}
	return
}

func Int64Slice(i interface{}) (r []int64) {
	switch i := i.(type) {
	case []uint64:
		r = *(*[]int64)(unsafe.Pointer(&i))
	case []float64:
		r = *(*[]int64)(unsafe.Pointer(&i))
	default:
		r = ResliceOf(INT64, i).([]int64)
	}
	return
}

func UintSlice(i interface{}) []uint {
	if i, ok := i.([]int); ok {
		return *(*[]uint)(unsafe.Pointer(&i))
	}
	return ResliceOf(UINT, i).([]uint)
}

func Uint8Slice(i interface{}) []uint8 {
	if i, ok := i.([]int8); ok {
		return *(*[]uint8)(unsafe.Pointer(&i))
	}
	return ResliceOf(UINT8, i).([]uint8)
}

func Uint16Slice(i interface{}) []uint16 {
	if i, ok := i.([]int16); ok {
		return *(*[]uint16)(unsafe.Pointer(&i))
	}
	return ResliceOf(UINT16, i).([]uint16)
}

func Uint32Slice(i interface{}) (r []uint32) {
	switch i := i.(type) {
	case []int32:
		r = *(*[]uint32)(unsafe.Pointer(&i))
	case []float32:
		r = *(*[]uint32)(unsafe.Pointer(&i))
	default:
		r = ResliceOf(UINT32, i).([]uint32)
	}
	return
}

func Uint64Slice(i interface{}) (r []uint64) {
	switch i := i.(type) {
	case []int64:
		r = *(*[]uint64)(unsafe.Pointer(&i))
	case []float64:
		r = *(*[]uint64)(unsafe.Pointer(&i))
	default:
		r = ResliceOf(UINT64, i).([]uint64)
	}
	return
}

func Float32Slice(i interface{}) (r []float32) {
	switch i := i.(type) {
	case []int32:
		r = *(*[]float32)(unsafe.Pointer(&i))
	case []uint32:
		r = *(*[]float32)(unsafe.Pointer(&i))
	default:
		r = ResliceOf(FLOAT32, i).([]float32)
	}
	return
}

func Float64Slice(i interface{}) (r []float64) {
	switch i := i.(type) {
	case []int64:
		r = *(*[]float64)(unsafe.Pointer(&i))
	case []uint64:
		r = *(*[]float64)(unsafe.Pointer(&i))
	case []complex64:
		r = *(*[]float64)(unsafe.Pointer(&i))
	default:
		r = ResliceOf(FLOAT64, i).([]float64)
	}
	return
}

func Complex64Slice(i interface{}) (r []complex64) {
	switch i := i.(type) {
	case []int64:
		r = *(*[]complex64)(unsafe.Pointer(&i))
	case []uint64:
		r = *(*[]complex64)(unsafe.Pointer(&i))
	case []float64:
		r = *(*[]complex64)(unsafe.Pointer(&i))
	default:
		r = ResliceOf(COMPLEX64, i).([]complex64)
	}
	return
}

func Complex128Slice(i interface{}) []complex128 {
	return ResliceOf(COMPLEX128, i).([]complex128)
}