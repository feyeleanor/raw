package raw

import "reflect"

type Mapping interface {
	Container
	KeyType() reflect.Type
	At(key interface{}) interface{}
	Set(key, value interface{})
	Clone() Mapping
}