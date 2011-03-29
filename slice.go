package raw

import "reflect"

var StandardSlack			float32	= 1.1
var StandardChannelBuffer	int		= 16

type Slice struct {
				*reflect.SliceValue
	Slack		float32					"how much spare capacity to allow on resize attempts"
}

// Create an independent duplicate of the Slice, copy all contents to the new assigned memory
func (s *Slice) Clone() Sequence {
	destination := reflect.MakeSlice(s.Type().(*reflect.SliceType), s.Len(), s.Cap())
	reflect.Copy(destination, s.SliceValue)
	return &Slice{ destination, StandardSlack }
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

// Copies a value from one location in the Slice to another.
func (s *Slice) Copy(destination, source int) {
	s.Elem(destination).SetValue(s.Elem(source))
}

// Copies a subslice from one location in the Slice to another.
func (s *Slice) CopySlice(destination, source, count int) {
	reflect.Copy(s.Slice(destination, destination + count), s.Slice(source, source + count))
//	s.Elem(destination).SetValue(s.Elem(source))
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
	return &Slice{ destination, StandardSlack }
}

//	First reads a specified number of values into a function, terminating if the end of Slice is reached
func (s *Slice) First(i int, f func(x interface{})) {
	for c := 0; c < i && c < s.Len(); c++ {
		f(s.Elem(c).Interface())
	}
}

//	First reads a specified number of values into a function, starting at the end of slice and terminating if the start of Slice is reached
func (s *Slice) Last(i int, f func(x interface{})) {
	for e := s.Len() - 1; i > 0 && e > 0; i-- {
		f(s.Elem(e).Interface())
		e--
	}
}

func (s *Slice) Resize(capacity int) {
	length := s.Len()
	switch {
	case capacity > int(float32(s.Cap()) * s.Slack):
		fallthrough
	case capacity < int(float32(s.Cap()) / s.Slack):
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