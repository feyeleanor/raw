package raw

type Buffer interface {
	Resize(int)
	Len() int
	Cap() int
}

func Extend(b Buffer, count int) {
	b.Resize(b.Cap() + count)
}

func Shrink(b Buffer, count int) {
	b.Resize(b.Cap() - count)
}

func DoubleCapacity(b Buffer) {
	b.Resize(b.Cap() * 2)
}

func HalveCapacity(b Buffer) {
	b.Resize(b.Cap() / 2)
}