package raw

import "reflect"

type Map struct {
	reflect.Value
}

func NewMap(i interface{}) (m *Map) {
	if v := reflect.ValueOf(i); v.Kind() == reflect.Map {
		m = &Map{ v }
	} else {
 		m = NewMap(v.Elem())
	}
	return
}

func (m *Map) At(k interface{}) (v interface{}) {
	switch k := k.(type) {
	case reflect.Value:
		if x := m.MapIndex(k); x.IsValid() {
			v = x.Interface()
		}
	default:
		v = m.At(reflect.ValueOf(k))
	}
	return 
}

func (m *Map) Store(k, value interface{}) {
	v := reflect.ValueOf(value)
	switch k := k.(type) {
	case reflect.Value:		m.SetMapIndex(k, v)
	default:				m.SetMapIndex(reflect.ValueOf(k), v)
	}
}

// Copies a value from one location in the Map to another.
func (m *Map) CopyElement(destination, source interface{}) {
	m.SetMapIndex(reflect.ValueOf(destination), reflect.ValueOf(source))
}

func (m *Map) Keys() interface{} {
	return NewSlice(m.MapKeys())
}

func (m *Map) Each(f func(v interface{})) (count int) {
	keys := m.MapKeys()
	count = len(keys)
	for _, k := range keys {
		f(m.MapIndex(k).Interface())
	}
	return
}

//	Create a new Map with identical keys to the existing Map but with values transformed according to a function.
func (m Map) Collect(f func(x interface{}) interface{}) (r Map) {
	r.Value = reflect.MakeMap(m.Type())
	for _, k := range m.MapKeys() {
		r.SetMapIndex(k, reflect.ValueOf(f(m.MapIndex(k).Interface())))
	}
	return
}

//	Reduce the values contained in the Map to a single value.
//	This is inherently unstable as Go makes no guarantees about the order in which map keys will be enumerable.
func (m *Map) Reduce(seed interface{}, f func(memo, x interface{}) interface{}) interface{} {
	keys := m.MapKeys()
	for _, k := range keys {
		seed = f(seed, m.MapIndex(k).Interface())
	}
	return seed
}

func (m *Map) Feed(c chan<- interface{}, f func(k, v interface{}) interface{}) {
	go func() {
		keys := m.MapKeys()
		for _, k := range keys {
			c <- f(k.Interface(), m.MapIndex(k).Interface())
		}
		close(c)
	}()
}

func (m *Map) Pipe(f func(k, v interface{}) interface{}) <-chan interface{} {
	c := make(chan interface{}, StandardChannelBuffer)
	m.Feed(c, f)
	return c
}