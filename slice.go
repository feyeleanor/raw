package raw

import "reflect"

var StandardChannelBuffer	int		= 16

type Slice struct {
	reflect.Value
}

func NewSlice(i interface{}) *Slice {
	return NewContainer(i).(*Slice)
}

func (s *Slice) New(length, capacity int) Sequence {
	return &Slice{ reflect.MakeSlice(s.Type(), length, capacity) }
}

func (s *Slice) Blit(destination, source, count int) {
	reflect.Copy(s.Slice(destination, destination + count), s.Slice(source, source + count))
}

func (s *Slice) Overwrite(offset int, source interface{}) {
	switch source := source.(type) {
	case *Slice:		s.Overwrite(offset, *source)
	case Slice:			if offset == 0 {
							reflect.Copy(s.Value, source.Value)
						} else {
							reflect.Copy(s.Value.Slice(offset, s.Len()), source.Value)
						}
	default:			switch v := reflect.NewValue(source); v.Kind() {
						case reflect.Slice:		s.Overwrite(offset, NewSlice(source))
						default:				s.Set(offset, v.Interface())
						}
	}
}

func (s *Slice) Append(i interface{}) {
	switch v := i.(type) {
	case *Slice:		s.Value.Set(reflect.AppendSlice(s.Value, v.Value))
	case Slice:			s.Value.Set(reflect.AppendSlice(s.Value, v.Value))
	default:			switch v := reflect.NewValue(i); v.Kind() {
						case reflect.Slice:		s.Value.Set(reflect.AppendSlice(s.Value, v))
						default:				s.Value.Set(reflect.Append(s.Value, v))
						}
	}
}

func (s *Slice) Prepend(i interface{}) {
	switch v := i.(type) {
	case *Slice:		s.Prepend(*v)
	case Slice:			n := s.New(s.Len() + v.Len(), s.Len() + v.Len()).(*Slice)
						n.Overwrite(0, v)
						n.Overwrite(v.Len(), s)
						s.Value.Set(n.Value)
	default:			switch v := reflect.NewValue(i); v.Kind() {
						case reflect.Slice:
							n := s.New(s.Len() + v.Len(), v.Len() + s.Len()).(*Slice)
							n.Overwrite(0, NewSlice(i))
							n.Overwrite(v.Len(), s)
							s.Value.Set(n.Value)
						default:
							n := s.New(s.Len() + 1, s.Len() + 1).(*Slice)
							n.Overwrite(0, i)
							n.Overwrite(1, s)
							s.Value.Set(n.Value)
						}
	}
}

// Returns the runtime type of the elements contained within the Slice.
func (s *Slice) ElementType() reflect.Type {
	return s.Type().Elem()
}

func (s *Slice) At(i int) interface{} {
	return s.Index(i).Interface()
}

func (s *Slice) Set(i int, value interface{}) {
	s.Index(i).Set(reflect.NewValue(value))
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