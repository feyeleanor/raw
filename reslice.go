package raw

import "C"
import "reflect"
import "unsafe"

func SliceHeader(i interface{}) (ElementSize int, Header *reflect.SliceHeader) {
	value := reflect.NewValue(i)
	switch value := value.(type) {
	case *reflect.PtrValue:
		ElementSize, Header = SliceHeader(value.Elem())

	case *reflect.InterfaceValue:
		ElementSize, Header = SliceHeader(value.Elem())

	case *reflect.SliceValue:
		Header = (*reflect.SliceHeader)(unsafe.Pointer(value.UnsafeAddr()))
		SliceType := value.Type().(*reflect.SliceType)
		ElementSize = int(SliceType.Elem().Size())
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

func PtrSlice(i interface{}) []uintptr {
	return Reslice(i, _SLICE_PTR, _SIZE_PTR).([]uintptr)
}

func InterfaceSlice(i interface{}) []interface{} {
	return Reslice(i, _SLICE_INTERFACE, _SIZE_INTERFACE).([]interface{})
}

func IntSlice(i interface{}) []int {
	return Reslice(i, _SLICE_INT, _SIZE_INT).([]int)
}

func Int8Slice(i interface{}) []int8 {
	return Reslice(i, _SLICE_INT8, _SIZE_INT8).([]int8)
}

func Int16Slice(i interface{}) []int16 {
	return Reslice(i, _SLICE_INT16, _SIZE_INT16).([]int16)
}

func Int32Slice(i interface{}) []int32 {
	return Reslice(i, _SLICE_INT32, _SIZE_INT32).([]int32)
}

func UintSlice(i interface{}) []uint {
	return Reslice(i, _SLICE_UINT, _SIZE_UINT).([]uint)
}

func Uint8Slice(i interface{}) []uint8 {
	return Reslice(i, _SLICE_UINT8, _SIZE_UINT8).([]uint8)
}

func Uint16Slice(i interface{}) []uint16 {
	return Reslice(i, _SLICE_UINT16, _SIZE_UINT16).([]uint16)
}

func Uint32Slice(i interface{}) []uint32 {
	return Reslice(i, _SLICE_UINT32, _SIZE_UINT32).([]uint32)
}

func Float32Slice(i interface{}) []float32 {
	return Reslice(i, _SLICE_FLOAT32, _SIZE_FLOAT32).([]float32)
}

func Float64Slice(i interface{}) []float64 {
	return Reslice(i, _SLICE_FLOAT64, _SIZE_FLOAT64).([]float64)
}