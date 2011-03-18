package raw

import "reflect"

type Slice reflect.SliceValue

func MakeSlice(i interface{}) (s *Slice) {
	switch v := reflect.NewValue(i).(type) {
	case *reflect.SliceValue:		s = (*Slice)(v)
	case nil:						panic(i)
	case *reflect.InterfaceValue:	s = MakeSlice(v.Elem())
	case *reflect.PtrValue:			s = MakeSlice(v.Elem())
	}
	return
}

func (s *Slice) Len() int {
	return (*reflect.SliceValue)(s).Len()
}

func (s *Slice) Cap() int {
	return (*reflect.SliceValue)(s).Len()
}

func (s *Slice) Type() reflect.Type {
	return (*reflect.SliceValue)(s).Type()
}

func (s *Slice) ElementType() reflect.Type {
	return s.Type().(*reflect.SliceType).Elem()
}

func (s *Slice) Copy(destination, source int) {
	v := (*reflect.SliceValue)(s)
	v.Elem(destination).SetValue(v.Elem(source))
}

func (s *Slice) Swap(left, right int) {
	x := (*reflect.SliceValue)(s).Elem(left)
	y := (*reflect.SliceValue)(s).Elem(right)
	temp := reflect.NewValue(x.Interface())
	x.SetValue(y)
	y.SetValue(temp)
}

func (s *Slice) Clear(start, end int) {
	if end > s.Len() {
		end = s.Len()
	}
	blank := reflect.NewValue(s.ElementType())
	v := (*reflect.SliceValue)(s)
	for ; start < end; start++ {
		v.Elem(start).SetValue(blank)
	}
}

func (s *Slice) Replicate(count int) *Slice {
	destination := reflect.MakeSlice(s.Type().(*reflect.SliceType), s.Len(), s.Cap())
	source := (*reflect.SliceValue)(s)
	for ; count > 1; count-- {
		destination = reflect.AppendSlice(destination, source)
	}
	return (*Slice)(destination)
}

func (s *Slice) Clone() *Slice {
	destination := reflect.MakeSlice(s.Type().(*reflect.SliceType), 0, 0)
	return (*Slice)(reflect.AppendSlice(destination, (*reflect.SliceValue)(s)))
}

func (s *Slice) Count(f func(x interface{}) bool) (c int) {
	v := (*reflect.SliceValue)(s)
	for i := s.Len() - 1; i > -1; i-- {
		if f(v.Elem(i).Interface()) {
			c++
		}
	}
	return
}

func (s *Slice) Any(f func(x interface{}) bool) bool {
	v := (*reflect.SliceValue)(s)
	for i := s.Len() - 1; i > -1; i-- {
		if f(v.Elem(i).Interface()) {
			return true
		}
	}
	return false
}

func (s *Slice) All(f func(x interface{}) bool) bool {
	v := (*reflect.SliceValue)(s)
	for i := s.Len() - 1; i > -1; i-- {
		if !f(v.Elem(i).Interface()) {
			return false
		}
	}
	return true
}

func (s *Slice) None(f func(x interface{}) bool) bool {
	v := (*reflect.SliceValue)(s)
	for i := s.Len() - 1; i > -1; i-- {
		if f(v.Elem(i).Interface()) {
			return false
		}
	}
	return true
}

func (s *Slice) One(f func(x interface{}) bool) bool {
	v := (*reflect.SliceValue)(s)
	c := 0
	for i := s.Len() - 1; i > -1; i-- {
		switch {
		case c > 1:							return false
		case f(v.Elem(i).Interface()):		c++
		}
	}
	return c == 1
}

/*
func (b *IntBuffer) Resize(length int) {
	if length > cap(*b) {
		x := *b
		*b = make(IntBuffer, length)
		copy(*b, x)
	} else {
		*b = (*b)[:length]
	}
}

func (b *IntBuffer) Extend(count int) {
	b.Resize(len(*b) + count)
}

func (b *IntBuffer) Shrink(count int) {
	b.Resize(len(*b) - count)
}

func (b IntBuffer) Collect(f func(x int) int) IntBuffer {
	n := make(IntBuffer, len(b))
	for i, x := range b {
		n[i] = f(x)
	}
	return n
}

func (b IntBuffer) Inject(seed int, f func(memo, x int) int) int {
	for _, x := range b {
		seed = f(seed, x)
	}
	return seed
}

func (b IntBuffer) Cycle(count int, f func(i, x int)) (j int) {
	switch count {
	case 0:		for {
					for _, x := range b {
						f(j, x)
					}
					j++
				}
	default:	for k := 0; j < count; j++ {
					for _, x := range b {
						f(k, x)
					}
					k++
				}
	}
	return
}

func (b IntBuffer) Combine(o IntBuffer, f func(x, y int) int) IntBuffer {
	if len(b) != len(o) {
		panic(b)
	}
	n := make(IntBuffer, len(b))
	for i, x := range b {
		n[i] = f(x, o[i])
	}
	return n
}


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