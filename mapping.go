package raw

type Mapping interface {
	At(key interface{}) interface{}
	Set(key, value interface{})
}