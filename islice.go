package raw

type ISlice []interface{}

func (i ISlice) New(length, capacity int) Sequence {
	return make(ISlice, length, capacity)
}

func (i ISlice) Len() int {
	return len(i)
}

func (i ISlice) Cap() int {
	return cap(i)
}

func (i ISlice) Blit(destination, source, count int) {
	copy(i[destination:], i[source:source + count])
}

func (i ISlice) Overwrite(offset int, source interface{}) {
	switch source := source.(type) {
	case *ISlice:		i.Overwrite(offset, *source)
	case ISlice:		copy(i[offset:], source)
	default:			i[offset] = source
	}
}

func (i *ISlice) Append(v interface{}) {
	switch v := v.(type) {
	case *ISlice:			i.Append(*v)
	case *[]interface{}:	i.Append(*v)
	case ISlice:			*i = append(*i, v...)
	case []interface{}:		*i = append(*i, v...)
	default:				*i = append(*i, v)
	}
}

func (i *ISlice) Prepend(v interface{}) {
	switch v := v.(type) {
	case *ISlice:			i.Prepend(*v)
	case *[]interface{}:	i.Prepend(*v)
	case ISlice:			n := make([]interface{}, len(*i) + len(v), len(*i) + len(v))
							copy(n, v)
							copy(n[len(v):], *i)
							*i = n
	case []interface{}:		n := make([]interface{}, len(*i) + len(v), len(*i) + len(v))
							copy(n, v)
							copy(n[len(v):], *i)
							*i = n
	default:				n := make([]interface{}, len(*i) + 1, len(*i) + 1)
							n[0] = v
							copy(n[1:], *i)
							*i = n
	}
}

func (i ISlice) At(index int) interface{} {
	return i[index]
}

func (i ISlice) Store(index int, value interface{}) {
	i[index] = value
}

func (i ISlice) Repeat(count int) ISlice {
	length := len(i) * count
	capacity := cap(i)
	if capacity < length {
		capacity = length
	}
	destination := make([]interface{}, length, capacity)
	for start, end := 0, len(i); count > 0; count-- {
		copy(destination[start:end], i)
		start = end
		end += len(i)
	}
	return destination
}

func (i ISlice) Section(start, end int) Sequence {
	return i[start:end]
}

func (i *ISlice) Reallocate(capacity int) {
	length := len(*i)
	if length > capacity {
		length = capacity
	}
	x := make([]interface{}, length, capacity)
	copy(x, *i)
	*i = x
}

func (i ISlice) Feed(c chan<- interface{}, f func(x interface{}) interface{}) {
	go func() {
		for index, l := 0, len(i); index < l; index++ {
			c <- f(i[index])
		}
		close(c)
	}()
}

func (i ISlice) Pipe(f func(x interface{}) interface{}) (c chan interface{}) {
	c = make(chan interface{}, StandardChannelBuffer)
	i.Feed(c, f)
	return
}