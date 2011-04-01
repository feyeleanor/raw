package raw

import "reflect"

type Mapping interface {
	Container
	New() Mapping
	KeyType() reflect.Type
	At(key interface{}) interface{}
	Set(key, value interface{})
	Keys() Sequence
	Clone() Mapping
}

func NewMapping(i interface{}) Mapping {
	return NewContainer(i).(Mapping)
}