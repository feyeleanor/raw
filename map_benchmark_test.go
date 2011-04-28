package raw

import "testing"

//func BenchmarkNewMap(b *testing.B) {}

//func BenchmarkMapNew(b *testing.B) {}

func BenchmarkMapAt(b *testing.B) {
	v := NewMap(map[int]int{ 0: 0 })
	for i := 0; i < b.N; i++ {
		_ = v.At(0)
	}
}

func BenchmarkMapStore(b *testing.B) {
	v := NewMap(map[int]int{ 0: 0 })
	for i := 0; i < b.N; i++ {
		v.Store(0, 0)
	}
}

func BenchmarkMapCopyElementManually(b *testing.B) {
	v := NewMap(map[int]int{ 0: 0, 1: 1 })
	for i := 0; i < b.N; i++ {
		v.Store(1, v.At(0))
	}
}

func BenchmarkMapCopyElement(b *testing.B) {
	v := NewMap(map[int]int{ 0: 0, 1: 1 })
	for i := 0; i < b.N; i++ {
		v.CopyElement(1, 0)
	}
}

func BenchmarkMapKeys1(b *testing.B) {
	v := NewMap(map[int]int{ 0: 0 })
	for i := 0; i < b.N; i++ {
		_ = v.Keys()
	}
}

func BenchmarkMapKeys10(b *testing.B) {
	b.StopTimer()
		v := NewMap(make(map[int]int))
		for i := 0; i < 10; i++ {
			v.Store(i, i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Keys()
	}
}

func BenchmarkMapKeys100(b *testing.B) {
	b.StopTimer()
		v := NewMap(make(map[int]int))
		for i := 0; i < 100; i++ {
			v.Store(i, i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Keys()
	}
}

func BenchmarkMapKeys1000(b *testing.B) {
	b.StopTimer()
		v := NewMap(make(map[int]int))
		for i := 0; i < 1000; i++ {
			v.Store(i, i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Keys()
	}
}

func BenchmarkMapEach1(b *testing.B) {
	b.StopTimer()
		v := NewMap(map[int]int{ 0: 0 })
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Each(func(x interface{}) {})
	}
}

func BenchmarkMapEach10(b *testing.B) {
	b.StopTimer()
		v := NewMap(make(map[int]int))
		for i := 0; i < 10; i++ {
			v.Store(i, i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Each(func(x interface{}) {})
	}
}

func BenchmarkMapEach100(b *testing.B) {
	b.StopTimer()
		v := NewMap(make(map[int]int))
		for i := 0; i < 100; i++ {
			v.Store(i, i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Each(func(x interface{}) {})
	}
}

func BenchmarkMapEach1000(b *testing.B) {
	b.StopTimer()
		v := NewMap(make(map[int]int))
		for i := 0; i < 1000; i++ {
			v.Store(i, i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Each(func(x interface{}) {})
	}
}

func BenchmarkMapCollect1(b *testing.B) {
	v := NewMap(map[int]int{ 0: 0 })
	for i := 0; i < b.N; i++ {
		_ = v.Collect(func(x interface{}) interface{} { return nil })
	}
}

func BenchmarkMapCollect10(b *testing.B) {
	b.StopTimer()
		v := NewMap(make(map[int]int))
		for i := 0; i < 10; i++ {
			v.Store(i, i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Collect(func(x interface{}) interface{} { return nil })
	}
}

func BenchmarkMapCollect100(b *testing.B) {
	b.StopTimer()
		v := NewMap(make(map[int]int))
		for i := 0; i < 100; i++ {
			v.Store(i, i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Collect(func(x interface{}) interface{} { return nil })
	}
}

func BenchmarkMapCollect1000(b *testing.B) {
	b.StopTimer()
		v := NewMap(make(map[int]int))
		for i := 0; i < 1000; i++ {
			v.Store(i, i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Collect(func(x interface{}) interface{} { return nil })
	}
}

func BenchmarkMapReduce1(b *testing.B) {
	v := NewMap(map[int]int{ 0: 0 })
	for i := 0; i < b.N; i++ {
		_ = v.Reduce(0, func(memo, x interface{}) interface{} { return nil })
	}
}

func BenchmarkMapReduce10(b *testing.B) {
	b.StopTimer()
		v := NewMap(make(map[int]int))
		for i := 0; i < 10; i++ {
			v.Store(i, i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Reduce(0, func(memo, x interface{}) interface{} { return nil })
	}
}

func BenchmarkMapReduce100(b *testing.B) {
	b.StopTimer()
		v := NewMap(make(map[int]int))
		for i := 0; i < 100; i++ {
			v.Store(i, i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Reduce(0, func(memo, x interface{}) interface{} { return nil })
	}
}

func BenchmarkMapReduce1000(b *testing.B) {
	b.StopTimer()
		v := NewMap(make(map[int]int))
		for i := 0; i < 1000; i++ {
			v.Store(i, i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Reduce(0, func(memo, x interface{}) interface{} { return nil })
	}
}

//func BenchmarkMapFeed(b *testing.B) {}

//func BenchmarkMapPipe(b *testing.B) {}