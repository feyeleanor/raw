package raw

type Sequence interface {
	Container
	New(length, capacity int) Sequence
	Cap() int
	SetLen(l int)
	At(i int) interface{}
	Store(i int, x interface{})
	Section(start, end int) Sequence
}

type Blitter interface {
	Blit(destination, source, count int)
}

type Deque interface {
	Append(i interface{})
	Prepend(i interface{})
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
	switch {
	case end < start:		return
	case end > s.Len():		end = s.Len()
	}
	blank := MakeBlank(s)
	for ; start <= end; start++ {
		s.Store(start, blank)
	}
}

func CopyElements(c Container, destination, source, count int) {
	if count == 0 {
		return
	}
	switch c := c.(type) {
	case Blitter:			c.Blit(destination, source, count)
	case Sequence:			switch {
							case count == 0:
							case destination > source:
								count--
								destination = destination + count
								for end := source + count; end >= source; end-- {
									c.Store(destination, c.At(end))
									destination--
								}
							case destination < source:
								for end := source + count; source < end; source++ {
									c.Store(destination, c.At(source))
									destination++
								}
							}
	}
}