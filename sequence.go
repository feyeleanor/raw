package raw

import "reflect"

type Sequence interface {
	Container
	At(i int) interface{}
	Set(i int, x interface{})
	Section(start, end int) Sequence
	Clone() Sequence
}

type Queue interface {
	Container
	Push(interface{})
	Pull() interface{}
}

type Stack interface {
	Container
	Push(interface{})
	Pop(interface{})
}

func NewSequence(i interface{}) Sequence {
	return NewContainer(i).(Sequence)
}

func First(s Sequence, i int) Sequence {
	return s.Section(0, i)
}

func Last(s Sequence, i int) Sequence {
	length := s.Len()
	return s.Section(length - i, length)
}

func Clear(s Sequence, start, end int) {
	blank := MakeBlank(s)
	if end > s.Len() {
		end = s.Len()
	}
	end++
	for ; start < end; start++ {
		s.Set(start, blank)
	}
}

func CopyElements(s Sequence, destination, source, count int) {
	switch {
	case count == 0:
	case destination > source:
		count--
		destination = destination + count
		for end := source + count; end >= source; end-- {
			s.Set(destination, s.At(end))
			destination--
		}
	case destination < source:
		for end := source + count; source < end; source++ {
			s.Set(destination, s.At(source))
			destination++
		}
	}	
}

func Cycle(s Sequence, count int, f func(x interface{})) interface{} {
	j := 0
	l := s.Len()
	switch count {
	case 0:		for {
					for i := 0; i < l; i++ {
						f(s.At(i))
					}
					j++
				}
	default:	for k := 0; j < count; j++ {
					for i := 0; i < l; i++ {
						f(s.At(i))
					}
					k++
				}
	}
	return j
}

func Combine(left, right Sequence, f func(x, y interface{}) interface{}) (s Sequence) {
	if t := left.Type(); t == right.Type() {
		switch l, r := left.Len(), right.Len(); {
		case l == r:
			s = &Slice{ reflect.MakeSlice(t.(*reflect.SliceType), l, l) }
			for i := 0; i < l; i++ {
				s.Set(i, f(left.At(i), right.At(i)))
			}

		case l > r:
			s = &Slice{ reflect.MakeSlice(t.(*reflect.SliceType), l, l) }
			for i := 0; i < r; i++ {
				s.Set(i, f(left.At(i), right.At(i)))
			}
			v := reflect.MakeZero(t)
			for i := r; i < l; i++ {
				s.Set(i, f(left.At(i), v))
			}

		case l < r:
			s = &Slice{ reflect.MakeSlice(t.(*reflect.SliceType), r, r) }
			for i := 0; i < l; i++ {
				s.Set(i, f(left.At(i), right.At(i)))
			}
			v := reflect.MakeZero(t)
			for i := l; i < r; i++ {
				s.Set(i, f(v, right.At(i)))
			}
		}
	}
	return
}