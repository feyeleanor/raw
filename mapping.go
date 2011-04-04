package raw

import "reflect"

type Mapping interface {
	Container
	New() Mapping
	KeyType() reflect.Type
	At(key interface{}) interface{}
	Set(key, value interface{})
	Keys() Sequence
}

func NewMapping(i interface{}) Mapping {
	return NewContainer(i).(Mapping)
}

func Merge(d, s Mapping) {
	if d.KeyType() == s.KeyType() && d.ElementType() == s.ElementType() {
		Each(s.Keys(), func(k interface{}) {
			d.Set(k, s.At(k))
		})
	}
}