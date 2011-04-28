package raw

import "testing"

func initMapTest() (b map[int] int, m *Map) {
	b = map[int]int{ 0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9 }
	m = NewMap(b)
	return
}

func TestNewMap(t *testing.T) {
	SHOULD_CONTAIN := "%v[%v] should contain %v but contains %v"

	m := NewMap(map[int]int{ 0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9 })
	switch {
	case m == nil:				t.Fatal("Make map returned a nil value")
	case m.Len() != 10:			t.Fatalf("Map length should be %v not %v", 10, m.Len())
	case m.At(0) != 0:			t.Fatalf(SHOULD_CONTAIN, "map", 0, m.At(0))
	case m.At(1) != 1:			t.Fatalf(SHOULD_CONTAIN, "map", 1, m.At(1))
	case m.At(2) != 2:			t.Fatalf(SHOULD_CONTAIN, "map", 2, m.At(2))
	case m.At(3) != 3:			t.Fatalf(SHOULD_CONTAIN, "map", 3, m.At(3))
	case m.At(4) != 4:			t.Fatalf(SHOULD_CONTAIN, "map", 4, m.At(4))
	case m.At(5) != 5:			t.Fatalf(SHOULD_CONTAIN, "map", 5, m.At(5))
	case m.At(6) != 6:			t.Fatalf(SHOULD_CONTAIN, "map", 6, m.At(6))
	case m.At(7) != 7:			t.Fatalf(SHOULD_CONTAIN, "map", 7, m.At(7))
	case m.At(8) != 8:			t.Fatalf(SHOULD_CONTAIN, "map", 8, m.At(8))
	case m.At(9) != 9:			t.Fatalf(SHOULD_CONTAIN, "map", 9, m.At(9))
	}
}

func TestMapNew(t *testing.T) {
	t.Log(NO_TESTS)
}

func TestMapCopyElement(t *testing.T) {
	b, m := initMapTest()
	m.CopyElement(1, 3)
	switch {
	case b[1] != b[3]:			t.Fatalf("Elements b[1] and b[3] should match but are %v and %v", b[1], b[3])
	case b[3] != 3:				t.Fatalf("Element b[3] should be %v but is %v", 3, b[3])
	case m.At(1) != m.At(3):	t.Fatalf("Map elements m[1] and m[3] should match but are %v and %v", m.At(1), m.At(3))
	}
}

func TestMapEach(t *testing.T) {
	_, m := initMapTest()
	c := 0
	m.Each(func(v interface{}) {
		c += v.(int)
	})
	if c != 45 {
		t.Fatalf("Sum should be 45 and not %v", c)
	}
}

func TestMapCollect(t *testing.T) {
	INCORRECT_VALUE := "r[%v] == %v"

	b, m := initMapTest()
	r := m.Collect(func(i interface{}) interface{} {
		return i.(int) * 2
	})
	switch {
	case r == nil:				t.Fatal("Collect() returned a nil value")
	case r.Len() != len(b):		t.Fatalf("capacity should be %v but is %v", len(b), r.Len())
	case m.At(0) != 0:			t.Fatalf(INCORRECT_VALUE, 0, m.At(0))
	case m.At(1) != 1:			t.Fatalf(INCORRECT_VALUE, 1, m.At(1))
	case m.At(2) != 2:			t.Fatalf(INCORRECT_VALUE, 2, m.At(2))
	case m.At(3) != 3:			t.Fatalf(INCORRECT_VALUE, 3, m.At(3))
	case m.At(4) != 4:			t.Fatalf(INCORRECT_VALUE, 4, m.At(4))
	case m.At(5) != 5:			t.Fatalf(INCORRECT_VALUE, 5, m.At(5))
	case m.At(6) != 6:			t.Fatalf(INCORRECT_VALUE, 6, m.At(6))
	case m.At(7) != 7:			t.Fatalf(INCORRECT_VALUE, 7, m.At(7))
	case m.At(8) != 8:			t.Fatalf(INCORRECT_VALUE, 8, m.At(8))
	case m.At(9) != 9:			t.Fatalf(INCORRECT_VALUE, 9, m.At(9))
	case r.At(0) != 0:			t.Fatalf(INCORRECT_VALUE, 0, r.At(0))
	case r.At(1) != 2:			t.Fatalf(INCORRECT_VALUE, 1, r.At(1))
	case r.At(2) != 4:			t.Fatalf(INCORRECT_VALUE, 2, r.At(2))
	case r.At(3) != 6:			t.Fatalf(INCORRECT_VALUE, 3, r.At(3))
	case r.At(4) != 8:			t.Fatalf(INCORRECT_VALUE, 4, r.At(4))
	case r.At(5) != 10:			t.Fatalf(INCORRECT_VALUE, 5, r.At(5))
	case r.At(6) != 12:			t.Fatalf(INCORRECT_VALUE, 6, r.At(6))
	case r.At(7) != 14:			t.Fatalf(INCORRECT_VALUE, 7, r.At(7))
	case r.At(8) != 16:			t.Fatalf(INCORRECT_VALUE, 8, r.At(8))
	case r.At(9) != 18:			t.Fatalf(INCORRECT_VALUE, 9, r.At(9))
	}
}

func TestMapReduce(t *testing.T) {
	_, m := initMapTest()
	r := m.Reduce(0, func(seed, v interface{}) interface{} {
		return seed.(int) + v.(int)
	})
	switch {
	case r == nil:				t.Fatal("Reduce() returned a nil value")
	case r != 45:				t.Fatalf("r should be 45 but is %v", r)
	}
}

func TestMapFeed(t *testing.T) {
	b, m := initMapTest()
	c := make(chan interface{})
	m.Feed(c, func(k, v interface{}) (r interface{}) {
		return k.(int) * v.(int)
	})
	n := []int{}
	for x := range c {
		n = append(n, x.(int))
	}
	switch {
	case len(n) != len(b):			t.Fatalf("Length of slice should be the same length of map but is %v", len(n))
	}
}

func TestMapPipe(t *testing.T) {
	b, m := initMapTest()
	c := m.Pipe(func(k, v interface{}) (r interface{}) {
		return k.(int) * v.(int)
	})
	n := []int{}
	for x := range c {
		n = append(n, x.(int))
	}
	switch {
	case len(n) != len(b):			t.Fatalf("Length of slice should be the same length of map but is %v", len(n))
	}
}