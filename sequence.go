package raw

type Sequence interface {
	Container
	New(length, capacity int) Sequence
	Cap() int
	SetLen(l int)
	At(i int) interface{}
	Set(i int, x interface{})
	Section(start, end int) Sequence
}

type Deque interface {
	Append(i interface{})
	Prepend(i interface{})
}

func NewSequence(i interface{}) Sequence {
	return NewContainer(i).(Sequence)
}

func First(s Sequence, i int) Sequence {
	return s.Section(0, i)
}

func PopFirst(s Sequence) (i interface{}, x Sequence) {
	return s.At(0), s.Section(1, s.Len())
}

func Last(s Sequence, i int) Sequence {
	l := s.Len()
	return s.Section(l - i, l)
}

func PopLast(s Sequence) (i interface{}, x Sequence) {
	l := s.Len() - 1
	return s.At(l), s.Section(0, l)
}

func Clear(s Sequence, start, end int) {
	blank := MakeBlank(s)
	if end > s.Len() {
		end = s.Len()
	}
	end++
	for ; start < end; start++ {
		s.Set(start, blank)
	}
}

func CopyElements(s Sequence, destination, source, count int) {
	switch {
	case count == 0:
	case destination > source:
		count--
		destination = destination + count
		for end := source + count; end >= source; end-- {
			s.Set(destination, s.At(end))
			destination--
		}
	case destination < source:
		for end := source + count; source < end; source++ {
			s.Set(destination, s.At(source))
			destination++
		}
	}	
}