package raw

import "reflect"

var StandardChannelBuffer	int		= 16

type Slice struct {
	reflect.Value
}

func newSliceFromValue(v reflect.Value) (s *Slice) {
	switch v.Kind() {
	case reflect.Slice:						if !v.CanAddr() {
												x := reflect.New(v.Type()).Elem()
												x.Set(v)
												v = x
											}
											s = &Slice{ v }
	case reflect.Ptr, reflect.Interface:	s = newSliceFromValue(v.Elem())
	}
	return
}

func NewSlice(i interface{}) (s *Slice) {
	return newSliceFromValue(reflect.ValueOf(i))
}

func (s *Slice) setValue(v reflect.Value) {
	if !s.CanAddr() {
		x := reflect.New(s.Type()).Elem()
		x.Set(s.Value)
		s.Value = x
	}
	s.Value = v
}

func (s *Slice) Set(i interface{}) {
	s.setValue(reflect.ValueOf(i))
}

func (s *Slice) Append(i interface{}) {
	switch v := i.(type) {
	case *Slice:		s.Append(*v)
	case Slice:			s.setValue(reflect.AppendSlice(s.Value, v.Value))
	default:			switch v := reflect.ValueOf(i); v.Kind() {
						case reflect.Slice:		s.setValue(reflect.AppendSlice(s.Value, v))
						default:				s.setValue(reflect.Append(s.Value, v))
						}
	}
}

func (s *Slice) At(i int) interface{} {
	return s.Index(i).Interface()
}

func (s *Slice) Store(i int, value interface{}) {
	s.Index(i).Set(reflect.ValueOf(value))
}

func (s *Slice) Repeat(count int) (destination *Slice) {
	length := s.Len() * count
	capacity := s.Cap()
	if capacity < length {
		capacity = length
	}
	destination = &Slice{ Value: reflect.MakeSlice(s.Type(), length, capacity) }
	for start, end := 0, s.Len(); count > 0; count-- {
		reflect.Copy(destination.Slice(start, end), s.Value)
		start = end
		end += s.Len()
	}
	return
}