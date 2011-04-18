package raw

import "reflect"

type Mapping interface {
	Container
	New() Mapping
	KeyType() reflect.Type
	At(key interface{}) interface{}
	Store(key, value interface{})
	Keys() Sequence
}

func Merge(d, s Mapping) {
	if d.KeyType() == s.KeyType() && d.ElementType() == s.ElementType() {
		Each(s.Keys(), func(k interface{}) {
			d.Store(k, s.At(k))
		})
	}
}