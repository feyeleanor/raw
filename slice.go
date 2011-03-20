package raw

import "reflect"

var StandardSlack float32 = 1.1

type Slice struct {
				*reflect.SliceValue
	Slack		float32					"how much spare capacity to allow on resize attempts"
}

func MakeSlice(i interface{}) (s *Slice) {
	switch v := reflect.NewValue(i).(type) {
	case *reflect.SliceValue:		s = &Slice{ v, StandardSlack }
	case nil:						panic(i)
	case *reflect.InterfaceValue:	s = MakeSlice(v.Elem())
	case *reflect.PtrValue:			s = MakeSlice(v.Elem())
	}
	return
}

func (s *Slice) ElementType() reflect.Type {
	return s.Type().(*reflect.SliceType).Elem()
}

func (s *Slice) Copy(destination, source int) {
	s.Elem(destination).SetValue(s.Elem(source))
}

func (s *Slice) Swap(left, right int) {
	x := s.Elem(left)
	y := s.Elem(right)
	temp := reflect.NewValue(x.Interface())
	x.SetValue(y)
	y.SetValue(temp)
}

func (s *Slice) Clear(start, end int) {
	if end > s.Len() {
		end = s.Len()
	}
	blank := reflect.NewValue(s.ElementType())
	for ; start < end; start++ {
		s.Elem(start).SetValue(blank)
	}
}

func (s *Slice) Repeat(count int) *Slice {
	length := s.Len() * count
	capacity := s.Cap()
	if capacity < length {
		capacity = length
	}
	destination := reflect.MakeSlice(s.Type().(*reflect.SliceType), length, capacity)
	for a, l := 0, s.Len(); count > 1; count-- {
		reflect.Copy(destination.Slice(a, l), s.SliceValue)
	}
	return &Slice{ destination, StandardSlack }
}

func (s *Slice) Clone() *Slice {
	destination := reflect.MakeSlice(s.Type().(*reflect.SliceType), s.Len(), s.Cap())
	reflect.Copy(destination, s.SliceValue)
	return &Slice{ destination, StandardSlack }
}

func (s *Slice) Count(f func(x interface{}) bool) (c int) {
	for i := s.Len() - 1; i > -1; i-- {
		if f(s.Elem(i).Interface()) {
			c++
		}
	}
	return
}

func (s *Slice) Any(f func(x interface{}) bool) bool {
	for i := s.Len() - 1; i > -1; i-- {
		if f(s.Elem(i).Interface()) {
			return true
		}
	}
	return false
}

func (s *Slice) All(f func(x interface{}) bool) bool {
	for i := s.Len() - 1; i > -1; i-- {
		if !f(s.Elem(i).Interface()) {
			return false
		}
	}
	return true
}

func (s *Slice) None(f func(x interface{}) bool) bool {
	for i := s.Len() - 1; i > -1; i-- {
		if f(s.Elem(i).Interface()) {
			return false
		}
	}
	return true
}

func (s *Slice) One(f func(x interface{}) bool) bool {
	c := 0
	for i := s.Len() - 1; i > -1; i-- {
		switch {
		case c > 1:							return false
		case f(s.Elem(i).Interface()):		c++
		}
	}
	return c == 1
}

func (s *Slice) At(i int) interface{} {
	return s.Elem(i).Interface()
}

func (s *Slice) Set(i int, value interface{}) {
	s.Elem(i).SetValue(reflect.NewValue(value))
}

func (s *Slice) Collect(f func(x interface{}) interface{}) *Slice {
	destination := &Slice{ reflect.MakeSlice(s.Type().(*reflect.SliceType), s.Len(), s.Cap()), StandardSlack }
	for i := s.Len() - 1; i > 0; i-- {
		destination.Set(i, f(s.At(i)))
	}
	return destination
}

func (s *Slice) Inject(seed interface{}, f func(memo, x interface{}) interface{}) interface{} {
	end := s.Len()
	for i := 0; i < end; i++ {
		seed = f(seed, s.At(i))
	}
	return seed
}

func (s *Slice) Combine(o *Slice, f func(x, y interface{}) interface{}) *Slice {
	l := s.Len()
	if s.Len() > o.Len() {
		l = o.Len()
	}
	destination := &Slice{ reflect.MakeSlice(s.Type().(*reflect.SliceType), l, l), StandardSlack }
	for i := 0; i < l; i++ {
		destination.Set(i, f(s.At(i), o.At(i)))
	}
	return destination
}

func (s *Slice) Cycle(count int, f func(i int, x interface{})) interface{} {
	j := 0
	l := s.Len()
	switch count {
	case 0:		for {
					for i := 0; i < l; i++ {
						f(j, s.At(i))
					}
					j++
				}
	default:	for k := 0; j < count; j++ {
					for i := 0; i < l; i++ {
						f(k, s.At(i))
					}
					k++
				}
	}
	return j
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

func (s *Slice) Extend(count int) {
	s.Resize(s.Cap() + count)
}

func (s *Slice) Shrink(count int) {
	s.Resize(s.Cap() - count)
}

func (s *Slice) DoubleCapacity() {
	s.Resize(s.Cap() * 2)
}

func (s *Slice) HalveCapacity() {
	s.Resize(s.Cap() / 2)
}

/*
func (b IntBuffer) Feed(c chan<- int, f func(i, x int) int) {
	d := b.Clone()
	go func() {
		for i, v := range d { c <- f(i, v) }
		close(c)
	}()
}

func (b IntBuffer) Pipe(f func(i, x int) int) <-chan int {
	c := make(chan int)
	b.Feed(c, f)
	return c
}
*/