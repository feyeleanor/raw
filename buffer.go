package raw

type Buffer interface {
	Reallocate(capacity int)
	Cap() int
	New(capacity int) Buffer
}

func Reallocate(b Buffer, capacity int) {
	switch {
	case capacity < 0:			b.Reallocate(0)
	case capacity != b.Cap():	b.Reallocate(capacity)
	}
}

func Extend(b Buffer, count int) {
	Reallocate(b, b.Cap() + count)
}

func Shrink(b Buffer, count int) {
	Reallocate(b, b.Cap() - count)
}

func DoubleCapacity(b Buffer) {
	Reallocate(b, b.Cap() * 2)
}

func HalveCapacity(b Buffer) {
	Reallocate(b, b.Cap() / 2)
}