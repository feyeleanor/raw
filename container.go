package raw

import "reflect"

type Container interface {
	Len() int
	Type() reflect.Type
	ElementType() reflect.Type
}

type Enumerable interface {
	Container
	Each(func(x interface{})) int
}

func NewContainer(i interface{}) (c Container) {
	switch v := reflect.NewValue(i).(type) {
	case *reflect.SliceValue:		c = &Slice{ SliceValue: v }
	case *reflect.MapValue:			c = &Map{ v }
	case *reflect.ChanValue:		c = &Channel{ v }
	case *reflect.InterfaceValue:	c = NewContainer(v.Elem())
	case *reflect.PtrValue:			c = NewContainer(v.Elem())
	default:						panic(i)
	}
	return
}

func MakeBlank(c Container) interface{} {
	return reflect.MakeZero(c.ElementType()).Interface()
}

func Swap(c interface{}, left, right interface{}) {
	switch c := c.(type) {
	case Sequence:
		l := left.(int)
		r := right.(int)
		temp := c.At(l)
		c.Set(l, c.At(r))
		c.Set(r, temp)
	case Mapping:
		temp := c.At(left)
		c.Set(left, c.At(right))
		c.Set(right, temp)
	}
}

func Each(c Container, f func(x interface{})) (i int) {
	switch c := c.(type) {
	case Sequence:
		for i = 0; i < c.Len(); i++ {
			f(c.At(i))
		}
	case Enumerable:
		c.Each(f)
	}
	return
}

func Collect(c Container, f func(x interface{}) interface{}) (s Sequence) {
	s = &Slice{ reflect.MakeSlice(c.Type().(*reflect.SliceType), c.Len(), c.Len()) }
	i := 0
	Each(c, func(x interface{}) {
		s.Set(i, f(x))
		i++
	})
	return
}

func Inject(c Container, seed interface{}, f func(memo, x interface{}) interface{}) interface{} {
	Each(c, func(x interface{}) {
		seed = f(seed, x)
	})
	return seed
}

//	While processes values from a Container type whilst a condition is true or until the end of the Container is reached
//	Returns the count of items which pass the test
func While(c Container, f func(x interface{}) bool) (i int) {
	switch c := c.(type) {
	case Sequence:
		for ; i < c.Len(); i++ {
			if !f(c.At(i)) {
				break
			}
		}
	case Enumerable:
		defer func() {
			if x := recover(); x != nil {
				panic(x)
			}
		}()
		c.Each(func(x interface{}) {
			if f(x) {
				i++
			} else {
				panic(nil)
			}
		})
	}
	return
}

//	Until processes values from a Container type until a condition is true or until the end of the Container is reached
//	Returns the count of items which fail the test
func Until(c Container, f func(x interface{}) bool) (i int) {
	switch c := c.(type) {
	case Sequence:
		for ; i < c.Len(); i++ {
			if f(c.At(i)) {
				break
			}
		}
	case Enumerable:
		defer func() {
			if x := recover(); x != nil {
				panic(x)
			}
		}()
		c.Each(func(x interface{}) {
			if !f(x) {
				i++
			} else {
				panic(nil)
			}
		})
	}
	return
}

func Count(c Container, f func(x interface{}) bool) (n int) {
	switch c := c.(type) {
	case Sequence:
		for i := 0; i < c.Len(); i++ {
			if f(c.At(i)) {
				n++
			}
		}
	case Enumerable:
		c.Each(func(x interface{}) {
			if f(x) {
				n++
			}
		})
	}
	return
}

func Any(c Container, f func(x interface{}) bool) bool {
	if i := Until(c, f); i < c.Len() {
		return true
	}
	return false
}

func All(c Container, f func(x interface{}) bool) bool {
	if i := While(c, f); i == c.Len() {
		return true
	}
	return false
}

func None(c Container, f func(x interface{}) bool) bool {
	if i := Until(c, f); i == c.Len() {
		return true
	}
	return false
}

func One(c Container, f func(x interface{}) bool) (b bool) {
	i := 0
	switch c := c.(type) {
	case Sequence:
		for j := 0; j < c.Len(); j++ {
			switch {
			case i > 1:				break
			case f(c.At(j)):		i++
			}
		}
	case Enumerable:
		i := 0
		defer func() {
			if x := recover(); x != nil {
				panic(x)
			}
			b = i == 1
		}()
		c.Each(func(x interface{}) {
			switch {
			case i > 1:				panic(nil)
			case f(x):				i++
			}
		})
	}
	return i == 1
}

func Dense(c Container, proportion float32, f func(x interface{}) bool) (b bool) {
	threshold := int(float32(c.Len()) * proportion)
	i := 0
	defer func() {
		if x := recover(); x != nil {
			panic(x)
		}
		b = i == threshold
	}()
	switch c := c.(type) {
	case Sequence:
		for j := 0; j < c.Len(); j++ {
			switch {
			case i == threshold:	break
			case f(c.At(j)):		i++
			}
		}
	case Enumerable:
		c.Each(func(x interface{}) {
			switch {
			case i == threshold:	panic(nil)
			case f(x):				i++
			}
		})
	}
	return
}

func Most(c Container, f func(x interface{}) bool) bool {
	return Dense(c, 0.5, f)
}