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

func (s *Slice) New(length, capacity int) Sequence {
	return &Slice{ reflect.MakeSlice(s.Type(), length, capacity) }
}

func (s *Slice) Blit(destination, source, count int) {
	reflect.Copy(s.Slice(destination, destination + count), s.Slice(source, source + count))
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

func (s *Slice) Overwrite(offset int, source interface{}) {
	switch source := source.(type) {
	case *Slice:		s.Overwrite(offset, *source)
	case Slice:			if offset == 0 {
							reflect.Copy(s.Value, source.Value)
						} else {
							reflect.Copy(s.Slice(offset, s.Len()), source.Value)
						}
	default:			switch v := reflect.ValueOf(source); v.Kind() {
						case reflect.Slice:		s.Overwrite(offset, NewSlice(source))
						default:				s.Store(offset, v.Interface())
						}
	}
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

func (s *Slice) Prepend(i interface{}) {
	switch v := i.(type) {
	case *Slice:		s.Prepend(*v)
	case Slice:			n := s.New(s.Len() + v.Len(), s.Len() + v.Len()).(*Slice)
						n.Overwrite(0, v)
						n.Overwrite(v.Len(), s)
						s.setValue(n.Value)
	default:			switch v := reflect.ValueOf(i); v.Kind() {
						case reflect.Slice:
							n := s.New(s.Len() + v.Len(), v.Len() + s.Len()).(*Slice)
							n.Overwrite(0, NewSlice(i))
							n.Overwrite(v.Len(), s)
							s.setValue(n.Value)
						default:
							n := s.New(s.Len() + 1, s.Len() + 1).(*Slice)
							n.Overwrite(0, i)
							n.Overwrite(1, s)
							s.setValue(n.Value)
						}
	}
}

func (s *Slice) At(i int) interface{} {
	return s.Index(i).Interface()
}

func (s *Slice) Store(i int, value interface{}) {
	s.Index(i).Set(reflect.ValueOf(value))
}

func (s *Slice) Repeat(count int) *Slice {
	length := s.Len() * count
	capacity := s.Cap()
	if capacity < length {
		capacity = length
	}
	destination := s.New(length, capacity).(*Slice)
	for start, end := 0, s.Len(); count > 0; count-- {
		reflect.Copy(destination.Slice(start, end), s.Value)
		start = end
		end += s.Len()
	}
	return destination
}

func (s *Slice) Section(start, end int) Sequence {
	return &Slice{ s.Value.Slice(start, end) }
}

func (s *Slice) Reallocate(capacity int) {
	length := s.Len()
	if length > capacity {
		length = capacity
	}
	x := s.New(length, capacity).(*Slice)
	reflect.Copy(x.Value, s.Value)
	s.Value = x.Value
}

func (s *Slice) Feed(c chan<- interface{}, f func(x interface{}) interface{}) {
	go func() {
		for i, l := 0, s.Len(); i < l; i++ {
			c <- f(s.Index(i).Interface())
		}
		close(c)
	}()
}

func (s *Slice) Pipe(f func(x interface{}) interface{}) (c chan interface{}) {
	c = make(chan interface{}, StandardChannelBuffer)
	s.Feed(c, f)
	return
}