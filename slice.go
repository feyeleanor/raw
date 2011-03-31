package raw

import "reflect"

var StandardChannelBuffer	int		= 16

type Slice struct {
	*reflect.SliceValue
}

func NewSlice(i interface{}) *Slice {
	return NewContainer(i).(*Slice)
}

// Create an independent duplicate of the Slice, copy all contents to the new assigned memory
func (s *Slice) Clone() Sequence {
	destination := reflect.MakeSlice(s.Type().(*reflect.SliceType), s.Len(), s.Cap())
	reflect.Copy(destination, s.SliceValue)
	return &Slice{ destination }
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
	destination := reflect.MakeSlice(s.Type().(*reflect.SliceType), length, capacity)
	for start, end := 0, s.Len(); count > 0; count-- {
		reflect.Copy(destination.Slice(start, end), s.SliceValue)
		start = end
		end += s.Len()
	}
	return &Slice{ destination }
}

func (s *Slice) Section(start, end int) Sequence {
	return &Slice{ s.SliceValue.Slice(start, end) }
}

func (s *Slice) Resize(capacity int) {
	length := s.Len()
	switch {
	case capacity != s.Cap():
		if capacity < 0 {
			capacity = 0
		}
		if length > capacity {
			length = capacity
		}
		x := reflect.MakeSlice(s.Type().(*reflect.SliceType), length, capacity)
		reflect.Copy(x, s.SliceValue)
		s.SetValue(x)
	}
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