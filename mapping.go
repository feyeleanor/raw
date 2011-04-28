package raw

type Mapping interface {
	Container
	New() Mapping
	At(key interface{}) interface{}
	Store(key, value interface{})
	Keys() Sequence
}

func Merge(d, s Mapping) {
	Each(s.Keys(), func(k interface{}) {
		d.Store(k, s.At(k))
	})
}

func Delete(m Mapping, k interface{}) {
	m.Store(k, MakeBlank(m))
}