package raw

type Enumerable interface {
	Len() int
	Each(func(x interface{})) int
}

func Count(e Enumerable, f func(x interface{}) bool) (c int) {
	e.Each(func(x interface{}) {
		if f(x) {
			c++
		}
	})
	return
}

//	While processes values from an Enumerable whilst a condition is true or until the end of the Enumerable is reached
//	Returns the count of items which pass the test
func While(e Enumerable, f func(x interface{}) bool) (n int) {
	defer func() {
		if x := recover(); x != nil {
			panic(x)
		}
	}()
	e.Each(func(x interface{}) {
		if f(x) {
			n++
		} else {
			panic(nil)
		}
	})
	return
}

//	Until processes values from an Enumerable until a condition is true or until the end of the Enumerable is reached
//	Returns the count of items which fail the test
func Until(e Enumerable, f func(x interface{}) bool) (n int) {
	defer func() {
		if x := recover(); x != nil {
			panic(x)
		}
	}()
	e.Each(func(x interface{}) {
		if !f(x) {
			n++
		} else {
			panic(nil)
		}
	})
	return
}

func Any(e Enumerable, f func(x interface{}) bool) bool {
	if c := Until(e, f); c < e.Len() {
		return true
	}
	return false
}

func All(e Enumerable, f func(x interface{}) bool) bool {
	if c := While(e, f); c == e.Len() {
		return true
	}
	return false
}

func None(e Enumerable, f func(x interface{}) bool) bool {
	if c := Until(e, f); c == e.Len() {
		return true
	}
	return false
}

func One(e Enumerable, f func(x interface{}) bool) (b bool) {
	c := 0
	defer func() {
		if x := recover(); x != nil {
			panic(x)
		}
		b = c == 1
	}()
	e.Each(func(x interface{}) {
		switch {
		case c > 1:				panic(nil)
		case f(x):				c++
		}
	})
	return
}

func Most(e Enumerable, f func(x interface{}) bool) (b bool) {
	threshold := e.Len() / 2
	c := 0
	defer func() {
		if x := recover(); x != nil {
			panic(x)
		}
		b = c > threshold
	}()
	e.Each(func(x interface{}) {
		switch {
		case c == threshold:	panic(nil)
		case f(x):				c++
		}
	})
	return
}