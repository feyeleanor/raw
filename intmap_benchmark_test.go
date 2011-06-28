package raw

import "testing"

//func BenchmarkNewIntMap(b *testing.B) {}

func BenchmarkIntMapAt(b *testing.B) {
	v := IntMap{ 0: 0 }
	for i := 0; i < b.N; i++ {
		_ = v.At(0)
	}
}

func BenchmarkIntMapStore(b *testing.B) {
	v := IntMap{ 0: 0 }
	for i := 0; i < b.N; i++ {
		v.Store(0, 0)
	}
}

func BenchmarkIntMapCopyElementManually(b *testing.B) {
	v := IntMap{ 0: 0, 1: 1 }
	for i := 0; i < b.N; i++ {
		v.Store(1, v.At(0))
	}
}

func BenchmarkIntMapCopyElement(b *testing.B) {
	v := IntMap{ 0: 0, 1: 1 }
	for i := 0; i < b.N; i++ {
		v.CopyElement(1, 0)
	}
}

func BenchmarkIntMapKeys1(b *testing.B) {
	v := IntMap{ 0: 0 }
	for i := 0; i < b.N; i++ {
		_ = v.Keys()
	}
}

func BenchmarkIntMapKeys10(b *testing.B) {
	b.StopTimer()
		v := make(IntMap)
		for i := 0; i < 10; i++ {
			v.Store(i, i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Keys()
	}
}

func BenchmarkIntMapKeys100(b *testing.B) {
	b.StopTimer()
		v := make(IntMap)
		for i := 0; i < 100; i++ {
			v.Store(i, i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Keys()
	}
}

func BenchmarkIntMapKeys1000(b *testing.B) {
	b.StopTimer()
		v := make(IntMap)
		for i := 0; i < 1000; i++ {
			v.Store(i, i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Keys()
	}
}

func BenchmarkIntMapEach1(b *testing.B) {
	b.StopTimer()
		v := IntMap{ 0: 0 }
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Each(func(x interface{}) {})
	}
}

func BenchmarkIntMapEach10(b *testing.B) {
	b.StopTimer()
		v := make(IntMap)
		for i := 0; i < 10; i++ {
			v.Store(i, i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Each(func(x interface{}) {})
	}
}

func BenchmarkIntMapEach100(b *testing.B) {
	b.StopTimer()
		v := make(IntMap)
		for i := 0; i < 100; i++ {
			v.Store(i, i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Each(func(x interface{}) {})
	}
}

func BenchmarkIntMapEach1000(b *testing.B) {
	b.StopTimer()
		v := make(IntMap)
		for i := 0; i < 1000; i++ {
			v.Store(i, i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Each(func(x interface{}) {})
	}
}

func BenchmarkIntMapCollect1(b *testing.B) {
	v := IntMap{ 0: 0 }
	for i := 0; i < b.N; i++ {
		_ = v.Collect(func(x interface{}) interface{} { return 0 })
	}
}

func BenchmarkIntMapCollect10(b *testing.B) {
	b.StopTimer()
		v := make(IntMap)
		for i := 0; i < 10; i++ {
			v.Store(i, i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Collect(func(x interface{}) interface{} { return 0 })
	}
}

func BenchmarkIntMapCollect100(b *testing.B) {
	b.StopTimer()
		v := make(IntMap)
		for i := 0; i < 100; i++ {
			v.Store(i, i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Collect(func(x interface{}) interface{} { return 0 })
	}
}

func BenchmarkIntMapCollect1000(b *testing.B) {
	b.StopTimer()
		v := make(IntMap)
		for i := 0; i < 1000; i++ {
			v.Store(i, i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Collect(func(x interface{}) interface{} { return 0 })
	}
}

func BenchmarkIntMapReduce1(b *testing.B) {
	v := IntMap{ 0: 0 }
	for i := 0; i < b.N; i++ {
		_ = v.Reduce(0, func(memo, x interface{}) interface{} { return nil })
	}
}

func BenchmarkIntMapReduce10(b *testing.B) {
	b.StopTimer()
		v := make(IntMap)
		for i := 0; i < 10; i++ {
			v.Store(i, i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Reduce(0, func(memo, x interface{}) interface{} { return nil })
	}
}

func BenchmarkIntMapReduce100(b *testing.B) {
	b.StopTimer()
		v := make(IntMap)
		for i := 0; i < 100; i++ {
			v.Store(i, i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Reduce(0, func(memo, x interface{}) interface{} { return nil })
	}
}

func BenchmarkIntMapReduce1000(b *testing.B) {
	b.StopTimer()
		v := make(IntMap)
		for i := 0; i < 1000; i++ {
			v.Store(i, i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Reduce(0, func(memo, x interface{}) interface{} { return nil })
	}
}

//func BenchmarkIntMapFeed(b *testing.B) {}

//func BenchmarkIntMapPipe(b *testing.B) {}