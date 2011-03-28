package raw

import "reflect"

type Map struct {
	*reflect.MapValue
}

// Creates a Slice from a given object, raising a runtime panic if the object cannot be represented as a *reflect.MapValue.
func MakeMap(i interface{}) (m *Map) {
	switch v := reflect.NewValue(i).(type) {
	case *reflect.MapValue:			m = &Map{ v }
	case *reflect.InterfaceValue:	m = MakeMap(v.Elem())
	case *reflect.PtrValue:			m = MakeMap(v.Elem())
	default:						panic(i)
	}
	return
}

// Create an independent duplicate of the Map, copy all contents to the new assigned memory
func (m *Map) Clone() *Map {
	destination := &Map{ reflect.MakeMap(m.Type().(*reflect.MapType)) }
	for _, k := range m.Keys() {
		destination.SetElem(k, m.Elem(k))
	}
	return destination
}

// Returns the runtime type of the keys referencing values in the Map.
func (m *Map) KeyType() reflect.Type {
	return m.Type().(*reflect.MapType).Key()
}

// Returns the runtime type of the elements contained within the Map.
func (m *Map) ElementType() reflect.Type {
	return m.Type().(*reflect.MapType).Elem()
}

func (m *Map) At(k interface{}) interface{} {
	return m.Elem(reflect.NewValue(k)).Interface()
}

func (m *Map) Set(k, value interface{}) {
	m.SetElem(reflect.NewValue(k), reflect.NewValue(value))
}

// Copies a value from one location in the Map to another.
func (m *Map) Copy(destination, source interface{}) {
	m.SetElem(reflect.NewValue(destination), reflect.NewValue(source))
}

func (m *Map) Clear(i interface{}) {
	m.Elem(reflect.NewValue(i)).SetValue(reflect.MakeZero(m.ElementType()))
}


func (m *Map) Each(f func(v interface{})) int {
	keys := m.Keys()
	for _, k := range keys {
		f(m.Elem(k).Interface())
	}
	return len(keys)
}

//	Create a new Map with identical keys to the existing Map but with values transformed according to a function.
func (m *Map) Collect(f func(x interface{}) interface{}) *Map {
	destination := &Map{ reflect.MakeMap(m.Type().(*reflect.MapType)) }
	for _, k := range m.Keys() {
		destination.SetElem(k, reflect.NewValue(f(m.Elem(k).Interface())))
	}
	return destination
}

//	Reduce the values contained in the Map to a single value.
//	This is inherently unstable as Go makes no guarantees about the order in which map keys will be enumerable.
func (m *Map) Inject(seed interface{}, f func(memo, x interface{}) interface{}) interface{} {
	for _, k := range m.Keys() {
		seed = f(seed, m.Elem(k).Interface())
	}
	return seed
}

//	Create a new Map whose keys are the union of two existing Maps with their values combined according to a function.
func (m *Map) Combine(o *Map, f func(x, y interface{}) interface{}) *Map {
	destination := &Map{ reflect.MakeMap(m.Type().(*reflect.MapType)) }
	for _, k := range m.Keys() {
		destination.SetElem(k, reflect.NewValue(f(m.Elem(k).Interface(), o.Elem(k).Interface())))
	}
	for _, k := range o.Keys() {
		if destination.Elem(k) == nil {
			destination.SetElem(k, reflect.NewValue(f(m.Elem(k).Interface(), o.Elem(k).Interface())))
		}
	}
	return destination
}

func (m *Map) Cycle(count int, f func(v interface{})) (limit int) {
	switch count {
	case 0:		for {
					for _, k := range m.Keys() {
						f(m.Elem(k).Interface())
					}
					limit++
				}
	default:	for ; count > 0; count-- {
					for _, k := range m.Keys() {
						f(m.Elem(k).Interface())
					}
					limit++
				}
	}
	return
}

func (m *Map) Feed(c chan<- interface{}, f func(k, v interface{}) interface{}) {
	go func() {
		for _, k := range m.Keys() {
			c <- f(k.Interface(), m.Elem(k).Interface())
		}
		close(c)
	}()
}

func (m *Map) Pipe(f func(k, v interface{}) interface{}) <-chan interface{} {
	c := make(chan interface{}, StandardChannelBuffer)
	m.Clone().Feed(c, f)
	return c
}

/*
func (m *Map) Tee(c chan<- interface{}, f func(k, v interface{}) interface{}) <-chan interface{} {
	t := make(chan interface{}, StandardChannelBuffer)
	go func() {
		for _, k := range m.Keys() {
			x := f(k.Interface(), m.Elem(k).Interface())
			c <- x
			t <- x
		}
		close(t)
	}()
	return t
}
*/