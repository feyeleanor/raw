package raw

func Each(container interface{}, f func(x interface{})) (i int) {
	defer Catch()
	switch c := container.(type) {
	case Sequence:
		for i = 0; i < c.Len(); i++ {
			f(c.At(i))
		}
	}
	return
}