package raw

import "testing"

//func BenchmarkNewISlice(b *testing.B) {}

//func BenchmarkISliceNew(b *testing.B) {}

//func BenchmarkISliceBlit(b *testing.B) {}

//func BenchmarkISliceOverwrite(b *testing.B) {}

func BenchmarkISliceAppend1x1(b *testing.B) {
	v := ISlice{ 0: 0 }
	for i := 0; i < b.N; i++ {
		(&ISlice{ 0: 0 }).Append(v)
	}
}

func BenchmarkISliceAppend1x10(b *testing.B) {
	b.StopTimer()
		v := make(ISlice, 0, 10)
		for i := 0; i < 10; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		(&ISlice{ 0: 0 }).Append(v)
	}
}

func BenchmarkISliceAppend1x100(b *testing.B) {
	b.StopTimer()
		v := make(ISlice, 0, 100)
		for i := 0; i < 100; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		(&ISlice{ 0: 0 }).Append(v)
	}
}

func BenchmarkISliceAppend1x1000(b *testing.B) {
	b.StopTimer()
		v := make(ISlice, 0, 1000)
		for i := 0; i < 1000; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		(&ISlice{ 0: 0 }).Append(v)
	}
}

//func BenchmarkISlicePrepend(b *testing.B) {}

func BenchmarkISliceAt(b *testing.B) {
	v := ISlice{ 0 }
	for i := 0; i < b.N; i++ {
		_ = v.At(0)
	}
}

func BenchmarkISliceStore(b *testing.B) {
	v := ISlice{ 0: 0 }
	for i := 0; i < b.N; i++ {
		v.Store(0, 0)
	}
}

func BenchmarkISliceRepeat1x1(b *testing.B) {
	v := ISlice{ 0: 0 }
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(1)
	}
}

func BenchmarkISliceRepeat1x10(b *testing.B) {
	v := ISlice{ 0: 0 }
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(10)
	}
}

func BenchmarkISliceRepeat1x100(b *testing.B) {
	v := ISlice{ 0: 0 }
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(100)
	}
}

func BenchmarkISliceRepeat1x1000(b *testing.B) {
	v := ISlice{ 0: 0 }
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(1000)
	}
}

func BenchmarkISliceRepeat10x1(b *testing.B) {
	b.StopTimer()
		v := make(ISlice, 0, 10)
		for i := 0; i < 10; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(1)
	}
}

func BenchmarkISliceRepeat10x10(b *testing.B) {
	b.StopTimer()
		v := make(ISlice, 0, 10)
		for i := 0; i < 10; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(10)
	}
}

func BenchmarkISliceRepeat10x100(b *testing.B) {
	b.StopTimer()
		v := make(ISlice, 0, 10)
		for i := 0; i < 10; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(100)
	}
}

func BenchmarkISliceRepeat10x1000(b *testing.B) {
	b.StopTimer()
		v := make(ISlice, 0, 10)
		for i := 0; i < 10; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(1000)
	}
}

func BenchmarkISliceRepeat100x1(b *testing.B) {
	b.StopTimer()
		v := make(ISlice, 0, 100)
		for i := 0; i < 100; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(1)
	}
}

func BenchmarkISliceRepeat100x10(b *testing.B) {
	b.StopTimer()
		v := make(ISlice, 0, 100)
		for i := 0; i < 100; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(10)
	}
}

func BenchmarkISliceRepeat100x100(b *testing.B) {
	b.StopTimer()
		v := make(ISlice, 0, 100)
		for i := 0; i < 100; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(100)
	}
}

func BenchmarkISliceRepeat100x1000(b *testing.B) {
	b.StopTimer()
		v := make(ISlice, 0, 100)
		for i := 0; i < 100; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(1000)
	}
}

//func BenchmarkISliceSection(b *testing.B) {}

//func BenchmarkISliceReallocate(b *testing.B) {}

//func BenchmarkISliceFeed(b *testing.B) {}

//func BenchmarkISlicePipe(b *testing.B) {}