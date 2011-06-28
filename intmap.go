package raw

import "reflect"

type IntMap map[int]interface{}

func (i IntMap) Len() int {
	return len(i)
}

func (i IntMap) Type() reflect.Type {
	return reflect.TypeOf(i)
}

func (i IntMap) At(k interface{}) interface{} {
	return i[k.(int)]
}

func (i IntMap) Store(k, value interface{}) {
	i[k.(int)] = value.(int)
}

func (i IntMap) CopyElement(destination, source interface{}) {
	i[destination.(int)] = i[source.(int)]
}

func (i IntMap) Keys() interface{} {
	return NewSlice(NewMap(i).MapKeys())
}

func (i IntMap) Each(f func(v interface{})) (count int) {
	for _, v := range i {
		f(v)
		count++
	}
	return
}

func (i IntMap) Collect(f func(x interface{}) interface{}) (r IntMap) {
	r = make(IntMap)
	for k, v := range i {
		r[k] = f(v).(int)
	}
	return
}

//	Reduce the values contained in the Map to a single value.
//	This is inherently unstable as Go makes no guarantees about the order in which map keys will be enumerable.
func (i IntMap) Reduce(seed interface{}, f func(memo, x interface{}) interface{}) interface{} {
	for _, v := range i {
		seed = f(seed, v)
	}
	return seed
}

func (i IntMap) Feed(c chan<- interface{}, f func(k, v interface{}) interface{}) {
	go func() {
		for k, v := range i {
			c <- f(k, v)
		}
		close(c)
	}()
}

func (i IntMap) Pipe(f func(k, v interface{}) interface{}) <-chan interface{} {
	c := make(chan interface{}, StandardChannelBuffer)
	i.Feed(c, f)
	return c
}