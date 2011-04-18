package raw

import "reflect"

type Container interface {
	Len() int
	Type() reflect.Type
	ElementType() reflect.Type
}

type Enumerable interface {
	Each(func(x interface{})) int
}

func Compatible(l, r Container) (b bool) {
	switch l := l.(type) {
	case Sequence:
		if r, ok := r.(Sequence); ok {
			b = l.ElementType() == r.ElementType()
		}

	case Mapping:
		if r, ok := r.(Mapping); ok {
			b = l.KeyType() == r.KeyType() && l.ElementType() == r.ElementType()
		}
	}
	return
}

func Copy(c Container) Container {
	switch c := c.(type) {
	case Sequence:
		n := c.New(c.Len(), c.Cap())
		for i := 0; i < c.Len(); i++ {
			n.Store(i, c.At(i))
		}
		return n
	case Mapping:
		n := c.New()
		Each(c.Keys(), func(k interface{}) {
			n.Store(k, c.At(k))
		})
		return n
	}
	return nil
}

func MakeBlank(c Container) interface{} {
	return reflect.Zero(c.ElementType()).Interface()
}

func SwapElements(c Container, left, right interface{}) {
	switch c := c.(type) {
	case Sequence:
		l := left.(int)
		r := right.(int)
		temp := c.At(l)
		c.Store(l, c.At(r))
		c.Store(r, temp)
	case Mapping:
		temp := c.At(left)
		c.Store(left, c.At(right))
		c.Store(right, temp)
	}
}

func Each(c Container, f func(x interface{})) (i int) {
	defer Catch()
	switch c := c.(type) {
	case Sequence:
		for i = 0; i < c.Len(); i++ {
			f(c.At(i))
		}
	case Enumerable:
		i = c.Each(f)
	}
	return
}

func Cycle(c Container, count int, f func(x interface{})) (i int) {
	switch {
	case count == 0:	for { Each(c, f) }
	default:			for ; i < count; i++ { Each(c, f) }
	}
	return
}

func Collect(c Container, f func(x interface{}) interface{}) (s Sequence) {
	s = &Slice{ reflect.MakeSlice(c.Type(), c.Len(), c.Len()) }
	i := 0
	Each(c, func(x interface{}) {
		s.Store(i, f(x))
		i++
	})
	return
}

func Inject(c Container, seed interface{}, f func(memo, x interface{}) interface{}) (r interface{}) {
	r = seed
	Each(c, func(x interface{}) {
		r = f(r, x)
	})
	return
}

//	While processes values from a Container type whilst a condition is true or until the end of the Container is reached
//	Returns the count of items which pass the test
func While(c Container, f func(x interface{}) bool) (i int) {
	Each(c, func(x interface{}) {
		if f(x) {
			i++
		} else {
			Throw()
		}
	})
	return
}

//	Until processes values from a Container type until a condition is true or until the end of the Container is reached
//	Returns the count of items which fail the test
func Until(c Container, f func(x interface{}) bool) (i int) {
	Each(c, func(x interface{}) {
		if f(x) {
			Throw()
		} else {
			i++
		}
	})
	return
}

func Count(c Container, f func(x interface{}) bool) (n int) {
	Each(c, func(x interface{}) {
		if f(x) { n++ }
	})
	return
}

func Any(c Container, f func(x interface{}) bool) bool {
	return Until(c, f) < c.Len()
}

func All(c Container, f func(x interface{}) bool) bool {
	return While(c, f) == c.Len()
}

func None(c Container, f func(x interface{}) bool) bool {
	return Until(c, f) == c.Len()
}

func One(c Container, f func(x interface{}) bool) (b bool) {
	Each(c, func(x interface{}) {
		if f(x) {
			if b {
				b = false
				Throw()
			} else {
				b = true
			}
		}
	})
	return
}

func Dense(c Container, proportion float32, f func(x interface{}) bool) (b bool) {
	threshold := int(float32(c.Len()) * proportion)
	i := 0
	Each(c, func(x interface{}) {
		switch {
		case i == threshold:
			b = true
			Throw()
		case f(x):
			i++
		}
	})
	return
}

func Most(c Container, f func(x interface{}) bool) bool {
	return Dense(c, 0.5, f)
}

func Combine(left, right Container, f func(x, y interface{}) interface{}) (c Container) {
	defer Catch()

	if !Compatible(left, right) { Throw() }
	switch left := left.(type) {
	case Sequence:
		right := right.(Sequence)
		l, r := left.Len(), right.Len()
		var s Sequence
		switch {
		case l == r:
			s = left.New(l, l)
			for i := 0; i < l; i++ {
				s.Store(i, f(left.At(i), right.At(i)))
			}

		case l > r:
			s = left.New(l, l)
			for i := 0; i < r; i++ {
				s.Store(i, f(left.At(i), right.At(i)))
			}
			for i := r; i < l; i++ {
				s.Store(i, f(left.At(i), s.At(i)))
			}

		case l < r:
			s = left.New(r, r)
			for i := 0; i < l; i++ {
				s.Store(i, f(left.At(i), right.At(i)))
			}
			for i := l; i < r; i++ {
				s.Store(i, f(s.At(i), right.At(i)))
			}
		}
		c = s

	case Mapping:
		if right, ok := right.(Mapping); ok {
			m := left.New()
			Each(left.Keys(), func(k interface{}) {
				m.Store(k, f(left.At(k), right.At(k)))
			})
			Each(right.Keys(), func(k interface{}) {
				if m.At(k) == nil {
					m.Store(k, f(left.At(k), right.At(k)))
				}
			})
			c = m
		}
	}
	return
}