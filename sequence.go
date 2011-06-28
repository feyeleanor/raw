package raw

type Sequence interface {
	Len() int
	Cap() int
	At(i int) interface{}
	Store(i int, x interface{})
}