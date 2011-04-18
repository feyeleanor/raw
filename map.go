package raw

import "reflect"

type Map struct {
	reflect.Value
}

func NewMap(i interface{}) (m *Map) {
	if v := reflect.NewValue(i); v.Kind() == reflect.Map {
		m = &Map{ v }
	} else {
 		m = NewMap(v.Elem())
	}
	return
}

func (m *Map) New() Mapping {
	return &Map{ reflect.MakeMap(m.Type()) }
}

// Returns the runtime type of the keys referencing values in the Map.
func (m *Map) KeyType() reflect.Type {
	return m.Type().Key()
}

// Returns the runtime type of the elements contained within the Map.
func (m *Map) ElementType() reflect.Type {
	return m.Type().Elem()
}

func (m *Map) At(k interface{}) (v interface{}) {
	switch k := k.(type) {
	case reflect.Value:
		if x := m.MapIndex(k); x.IsValid() {
			v = x.Interface()
		}
	default:
		v = m.At(reflect.NewValue(k))
	}
	return 
}

func (m *Map) Store(k, value interface{}) {
	switch k := k.(type) {
	case reflect.Value:		m.SetMapIndex(k, reflect.NewValue(value))
	default:				m.SetMapIndex(reflect.NewValue(k), reflect.NewValue(value))
	}
}

// Copies a value from one location in the Map to another.
func (m *Map) CopyElement(destination, source interface{}) {
	m.SetMapIndex(reflect.NewValue(destination), reflect.NewValue(source))
}

func (m *Map) Clear(i interface{}) {
	m.SetMapIndex(reflect.NewValue(i), reflect.Zero(m.ElementType()))
}

func (m *Map) Keys() Sequence {
	return NewSlice(m.Value.MapKeys())
}

func (m *Map) Each(f func(v interface{})) int {
	keys := m.Keys()
	Each(keys, func(k interface{}) {
		f(m.At(k))
	})
	return keys.Len()
}

//	Create a new Map with identical keys to the existing Map but with values transformed according to a function.
func (m *Map) Collect(f func(x interface{}) interface{}) (r *Map) {
	r = m.New().(*Map)
	Each(m.Keys(), func(k interface{}) {
		r.Store(k, f(m.At(k)))
	})
	return
}

//	Reduce the values contained in the Map to a single value.
//	This is inherently unstable as Go makes no guarantees about the order in which map keys will be enumerable.
func (m *Map) Inject(seed interface{}, f func(memo, x interface{}) interface{}) interface{} {
	Each(m.Keys(), func(k interface{}) {
		seed = f(seed, m.At(k))
	})
	return seed
}

func (m *Map) Feed(c chan<- interface{}, f func(k, v interface{}) interface{}) {
	go func() {
		Each(m.Keys(), func(k interface{}) {
			switch k := k.(type) {
			case reflect.Value:
				c <- f(k.Interface(), m.At(k))
			default:
				c <- f(k, m.At(k))
			}
		})
		close(c)
	}()
}

func (m *Map) Pipe(f func(k, v interface{}) interface{}) <-chan interface{} {
	c := make(chan interface{}, StandardChannelBuffer)
	m.Feed(c, f)
	return c
}