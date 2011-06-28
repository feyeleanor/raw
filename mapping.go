package raw

type Mapping interface {
	Len() int
	At(key interface{}) interface{}
	Store(key, value interface{})
	Keys() interface{}
}

func Merge(d, s Mapping) {
	Each(s.Keys(), func(k interface{}) {
		d.Store(k, s.At(k))
	})
}