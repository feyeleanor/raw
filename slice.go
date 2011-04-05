package raw

import "reflect"

var StandardChannelBuffer	int		= 16

type Slice struct {
	*reflect.SliceValue
}

func NewSlice(i interface{}) *Slice {
	return NewContainer(i).(*Slice)
}

func (s *Slice) New(length, capacity int) Sequence {
	return &Slice{ reflect.MakeSlice(s.Type().(*reflect.SliceType), length, capacity) }
}

func (s *Slice) Blit(destination, source, count int) {
	reflect.Copy(s.Slice(destination, destination + count), s.Slice(source, source + count))
}

func (s *Slice) Overwrite(offset int, source interface{}) {
	switch source := source.(type) {
	case *Slice:		s.Overwrite(offset, *source)
	case Slice:			if offset == 0 {
							reflect.Copy(s.SliceValue, source.SliceValue)
						} else {
							reflect.Copy(s.SliceValue.Slice(offset, s.Len()), source.SliceValue)
						}
	default:			switch v := reflect.NewValue(source).(type) {
						case *reflect.SliceValue:		s.Overwrite(offset, NewSlice(source))
						default:						s.Set(offset, v.Interface())
						}
	}
}

func (s *Slice) Append(i interface{}) {
	switch v := i.(type) {
	case *Slice:		s.SetValue(reflect.AppendSlice(s.SliceValue, v.SliceValue))
	case Slice:			s.SetValue(reflect.AppendSlice(s.SliceValue, v.SliceValue))
	default:			switch v := reflect.NewValue(i).(type) {
						case *reflect.SliceValue:		s.SetValue(reflect.AppendSlice(s.SliceValue, v))
						default:						s.SetValue(reflect.Append(s.SliceValue, v))
						}
	}
}

func (s *Slice) Prepend(i interface{}) {
	switch v := i.(type) {
	case *Slice:		s.Prepend(*v)
	case Slice:			n := s.New(s.Len() + v.Len(), s.Len() + v.Len()).(*Slice)
						n.Overwrite(0, v)
						n.Overwrite(v.Len(), s)
						s.SetValue(n.SliceValue)
	default:			switch v := reflect.NewValue(i).(type) {
						case *reflect.SliceValue:
							n := s.New(s.Len() + v.Len(), v.Len() + s.Len()).(*Slice)
							n.Overwrite(0, NewSlice(i))
							n.Overwrite(v.Len(), s)
							s.SetValue(n.SliceValue)
						default:
							n := s.New(s.Len() + 1, s.Len() + 1).(*Slice)
							n.Overwrite(0, i)
							n.Overwrite(1, s)
							s.SetValue(n.SliceValue)
						}
	}
}

// Returns the runtime type of the elements contained within the Slice.
func (s *Slice) ElementType() reflect.Type {
	return s.Type().(*reflect.SliceType).Elem()
}

func (s *Slice) At(i int) interface{} {
	return s.Elem(i).Interface()
}

func (s *Slice) Set(i int, value interface{}) {
	s.Elem(i).SetValue(reflect.NewValue(value))
}

func (s *Slice) Repeat(count int) *Slice {
	length := s.Len() * count
	capacity := s.Cap()
	if capacity < length {
		capacity = length
	}
	destination := s.New(length, capacity).(*Slice)
	for start, end := 0, s.Len(); count > 0; count-- {
		reflect.Copy(destination.Slice(start, end), s.SliceValue)
		start = end
		end += s.Len()
	}
	return destination
}

func (s *Slice) Section(start, end int) Sequence {
	return &Slice{ s.SliceValue.Slice(start, end) }
}

func (s *Slice) Reallocate(capacity int) {
	length := s.Len()
	if length > capacity {
		length = capacity
	}
	x := s.New(length, capacity).(*Slice)
	reflect.Copy(x.SliceValue, s.SliceValue)
	s.SliceValue = x.SliceValue
}

func (s *Slice) Feed(c chan<- interface{}, f func(x interface{}) interface{}) {
	go func() {
		for i, l := 0, s.Len(); i < l; i++ {
			c <- f(s.Elem(i).Interface())
		}
		close(c)
	}()
}

func (s *Slice) Pipe(f func(x interface{}) interface{}) (c chan interface{}) {
	c = make(chan interface{}, StandardChannelBuffer)
	s.Feed(c, f)
	return
}