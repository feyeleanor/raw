package raw

import "reflect"
import "unsafe"

func SliceHeader(i interface{}) (Header *reflect.SliceHeader, ElementSize, ElementAlignment int) {
	value := reflect.NewValue(i)
	switch value := value.(type) {
	case nil:						panic(i)
	case *reflect.InterfaceValue:	Header, ElementSize, ElementAlignment = SliceHeader(value.Elem())

	case *reflect.SliceValue:		Header = (*reflect.SliceHeader)(unsafe.Pointer(value.UnsafeAddr()))
									ElementType := value.Type().(*reflect.SliceType).Elem()
									ElementSize = int(ElementType.Size())
									ElementAlignment = ElementType.Align()

	case *reflect.PtrValue:			Header, ElementSize, ElementAlignment = SliceHeader(value.Elem())
	}
	return
}

func Scale(oldHeader *reflect.SliceHeader, oldElementSize, newElementSize int) (h *reflect.SliceHeader) {
	if oldHeader != nil {
		s := float64(oldElementSize) / float64(newElementSize)
		h = &reflect.SliceHeader{ Data: oldHeader.Data }
		h.Len = int(float64(oldHeader.Len) * s)
		h.Cap = int(float64(oldHeader.Cap) * s)
	}
	return
}

func Reslice(slice, sliceType interface{}, elementSize int) interface{} {
	b := ByteSlice(slice)
	h := Scale(&reflect.SliceHeader{ uintptr(DataAddress(b)), len(b), cap(b) }, 1, elementSize)
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

func Uint64Slice(i interface{}) []uint64 {
	switch i := i.(type) {
	case []int64:
		return *(*[]uint64)(unsafe.Pointer(&i))
	case []float64:
		return *(*[]uint64)(unsafe.Pointer(&i))
	}
	return Reslice(i, SLICE_TYPE[UINT32], SIZE[UINT32]).([]uint64)
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

func Complex64Slice(i interface{}) []complex64 {
	switch i := i.(type) {
	case []int64:
		return *(*[]complex64)(unsafe.Pointer(&i))
	case []uint64:
		return *(*[]complex64)(unsafe.Pointer(&i))
	case []float64:
		return *(*[]complex64)(unsafe.Pointer(&i))
	}
	return Reslice(i, SLICE_TYPE[COMPLEX64], SIZE[COMPLEX64]).([]complex64)
}

func Complex128Slice(i interface{}) []complex128 {
	return Reslice(i, SLICE_TYPE[COMPLEX128], SIZE[COMPLEX128]).([]complex128)
}