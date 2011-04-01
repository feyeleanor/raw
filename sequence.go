package raw

type Sequence interface {
	Container
	Buffer
	SetLen(l int)
	At(i int) interface{}
	Set(i int, x interface{})
	Section(start, end int) Sequence
	Copy(source Sequence)
}

type Queue interface {
	Container
	Buffer
	Push(interface{})
	Pull() interface{}
}

type Stack interface {
	Container
	Buffer
	Push(interface{})
	Pop(interface{})
}

func NewSequence(i interface{}) Sequence {
	return NewContainer(i).(Sequence)
}

func CopySequence(s Sequence) (n Sequence) {
	n = s.New(s.Cap()).(Sequence)
	n.SetLen(s.Len())
	n.Copy(s)
	return
}

func First(s Sequence, i int) Sequence {
	return s.Section(0, i)
}

func Last(s Sequence, i int) Sequence {
	length := s.Len()
	return s.Section(length - i, length)
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