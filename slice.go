package raw

import "reflect"

var StandardChannelBuffer	int		= 16

type Slice struct {
	*reflect.SliceValue
}

func NewSlice(i interface{}) *Slice {
	return NewContainer(i).(*Slice)
}

func (s *Slice) New(capacity int) Buffer {
	return &Slice{ reflect.MakeSlice(s.Type().(*reflect.SliceType), 0, capacity) }
}

func (s *Slice) Copy(source Sequence) {
	switch source := source.(type) {
	case *Slice:		reflect.Copy(s.SliceValue, source.SliceValue)
	default:			s.Copy(NewSlice(source))
	}
}

// Append a value to the existing Slice
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
	destination := s.New(capacity).(*Slice)
	destination.SetLen(length)
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
	x := s.New(capacity).(*Slice)
	x.SetLen(length)
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

func (s *Slice) Tee(c chan<- interface{}, f func(x interface{}) interface{}) (t chan interface{}) {
	t = make(chan interface{}, StandardChannelBuffer)
	go func() {
		for i, l := 0, s.Len(); i < l; i++ {
			x := f(s.Elem(i).Interface())
			t <- x
			c <- x
		}
	}()
	return
}