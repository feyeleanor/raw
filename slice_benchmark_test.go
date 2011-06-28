package raw

import "testing"

//func BenchmarkNewSlice(b *testing.B) {}

func BenchmarkSliceAppend1x1(b *testing.B) {
	v := NewSlice([]int{ 0: 0 })
	for i := 0; i < b.N; i++ {
		NewSlice([]int{ 0: 0 }).Append(v)
	}
}

func BenchmarkSliceAppend1x10(b *testing.B) {
	b.StopTimer()
		v := NewSlice(make([]int, 0, 10))
		for i := 0; i < 10; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		NewSlice([]int{ 0: 0 }).Append(v)
	}
}

func BenchmarkSliceAppend1x100(b *testing.B) {
	b.StopTimer()
		v := NewSlice(make([]int, 0, 100))
		for i := 0; i < 100; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		NewSlice([]int{ 0: 0 }).Append(v)
	}
}

func BenchmarkSliceAppend1x1000(b *testing.B) {
	b.StopTimer()
		v := NewSlice(make([]int, 0, 1000))
		for i := 0; i < 1000; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		NewSlice([]int{ 0: 0 }).Append(v)
	}
}

func BenchmarkSliceAt(b *testing.B) {
	v := NewSlice([]int{ 0 })
	for i := 0; i < b.N; i++ {
		_ = v.At(0)
	}
}

func BenchmarkSliceStore(b *testing.B) {
	v := NewSlice([]int{ 0: 0 })
	for i := 0; i < b.N; i++ {
		v.Store(0, 0)
	}
}

func BenchmarkSliceRepeat1x1(b *testing.B) {
	v := NewSlice([]int{ 0: 0 })
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(1)
	}
}

func BenchmarkSliceRepeat1x10(b *testing.B) {
	v := NewSlice([]int{ 0: 0 })
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(10)
	}
}

func BenchmarkSliceRepeat1x100(b *testing.B) {
	v := NewSlice([]int{ 0: 0 })
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(100)
	}
}

func BenchmarkSliceRepeat1x1000(b *testing.B) {
	v := NewSlice([]int{ 0: 0 })
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(1000)
	}
}

func BenchmarkSliceRepeat10x1(b *testing.B) {
	b.StopTimer()
		v := NewSlice(make([]int, 0, 10))
		for i := 0; i < 10; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(1)
	}
}

func BenchmarkSliceRepeat10x10(b *testing.B) {
	b.StopTimer()
		v := NewSlice(make([]int, 0, 10))
		for i := 0; i < 10; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(10)
	}
}

func BenchmarkSliceRepeat10x100(b *testing.B) {
	b.StopTimer()
		v := NewSlice(make([]int, 0, 10))
		for i := 0; i < 10; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(100)
	}
}

func BenchmarkSliceRepeat10x1000(b *testing.B) {
	b.StopTimer()
		v := NewSlice(make([]int, 0, 10))
		for i := 0; i < 10; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(1000)
	}
}

func BenchmarkSliceRepeat100x1(b *testing.B) {
	b.StopTimer()
		v := NewSlice(make([]int, 0, 100))
		for i := 0; i < 100; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(1)
	}
}

func BenchmarkSliceRepeat100x10(b *testing.B) {
	b.StopTimer()
		v := NewSlice(make([]int, 0, 100))
		for i := 0; i < 100; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(10)
	}
}

func BenchmarkSliceRepeat100x100(b *testing.B) {
	b.StopTimer()
		v := NewSlice(make([]int, 0, 100))
		for i := 0; i < 100; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(100)
	}
}

func BenchmarkSliceRepeat100x1000(b *testing.B) {
	b.StopTimer()
		v := NewSlice(make([]int, 0, 100))
		for i := 0; i < 100; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(1000)
	}
}

//func BenchmarkSliceFeed(b *testing.B) {}

//func BenchmarkSlicePipe(b *testing.B) {}