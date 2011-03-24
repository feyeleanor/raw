package raw

import "reflect"
import "testing"

func TestMapMakeMap(t *testing.T) {
	SHOULD_MATCH := "Map elements m[%v] and b[%v] should match but are %v and %v"
	b := map[int]int{ 0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9 }
	m := MakeMap(b)

	switch {
	case m == nil:				
	case m.Len() != len(b):		t.Fatalf("Slice length should be %v not %v", len(b), m.Len())
	case m.At(0) != b[0]:		t.Fatalf(SHOULD_MATCH, 0, 0, m.At(0), b[0])
	case m.At(1) != b[1]:		t.Fatalf(SHOULD_MATCH, 1, 1, m.At(1), b[1])
	case m.At(2) != b[2]:		t.Fatalf(SHOULD_MATCH, 2, 2, m.At(2), b[2])
	case m.At(3) != b[3]:		t.Fatalf(SHOULD_MATCH, 3, 3, m.At(3), b[3])
	case m.At(4) != b[4]:		t.Fatalf(SHOULD_MATCH, 4, 4, m.At(4), b[4])
	case m.At(5) != b[5]:		t.Fatalf(SHOULD_MATCH, 5, 5, m.At(5), b[5])
	case m.At(6) != b[6]:		t.Fatalf(SHOULD_MATCH, 6, 6, m.At(6), b[6])
	case m.At(7) != b[7]:		t.Fatalf(SHOULD_MATCH, 7, 7, m.At(7), b[7])
	case m.At(8) != b[8]:		t.Fatalf(SHOULD_MATCH, 8, 8, m.At(8), b[8])
	case m.At(9) != b[9]:		t.Fatalf(SHOULD_MATCH, 9, 9, m.At(9), b[9])
	}
}

func TestMapClone(t *testing.T) {
	SHOULD_MATCH := "Map elements m[%v] and c[%v] should match but are %v and %v"
	b := map[int]int{ 0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9 }
	m := MakeMap(b)
	c := m.Clone()

	switch {
	case c.Len() != m.Len():	t.Fatalf("Map length should be %v not %v", m.Len(), c.Len())
	case c.At(0) != m.At(0):	t.Fatalf(SHOULD_MATCH, 0, 0, m.At(0), c.At(0))
	case c.At(1) != m.At(1):	t.Fatalf(SHOULD_MATCH, 1, 1, m.At(1), c.At(1))
	case c.At(2) != m.At(2):	t.Fatalf(SHOULD_MATCH, 2, 2, m.At(2), c.At(2))
	case c.At(3) != m.At(3):	t.Fatalf(SHOULD_MATCH, 3, 3, m.At(3), c.At(3))
	case c.At(4) != m.At(4):	t.Fatalf(SHOULD_MATCH, 4, 4, m.At(4), c.At(4))
	case c.At(5) != m.At(5):	t.Fatalf(SHOULD_MATCH, 5, 5, m.At(5), c.At(5))
	case c.At(6) != m.At(6):	t.Fatalf(SHOULD_MATCH, 6, 6, m.At(6), c.At(6))
	case c.At(7) != m.At(7):	t.Fatalf(SHOULD_MATCH, 7, 7, m.At(7), c.At(7))
	case c.At(8) != m.At(8):	t.Fatalf(SHOULD_MATCH, 8, 8, m.At(8), c.At(8))
	case c.At(9) != m.At(9):	t.Fatalf(SHOULD_MATCH, 9, 9, m.At(9), c.At(9))
	}
}

func TestMapKeyType(t *testing.T) {
	b := map[int]int{ 0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9 }
	m := MakeMap(b)
	e := reflect.Typeof(int(0))

	switch {
	case e != m.KeyType():		t.Fatalf("Map claims element type %v when should be %v", m.KeyType(), e)
	}
}

func TestMapElementType(t *testing.T) {
	b := map[int]int{ 0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9 }
	m := MakeMap(b)
	e := reflect.Typeof(int(0))

	switch {
	case e != m.ElementType():	t.Fatalf("Map claims element type %v when should be %v", m.ElementType(), e)
	}
}

func TestMapCopy(t *testing.T) {
	b := map[int]int{ 0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9 }
	m := MakeMap(b)
	m.Copy(1, 3)

	switch {
	case b[1] != b[3]:			t.Fatalf("Elements b[1] and b[3] should match but are %v and %v", b[1], b[3])
	case b[3] != 3:				t.Fatalf("Element b[3] should be %v but is %v", 3, b[3])
	case m.At(1) != m.At(3):	t.Fatalf("Map elements m[1] and m[3] should match but are %v and %v", m.At(1), m.At(3))
	}
}

func TestMapSwap(t *testing.T) {
	HAS_VALUE := "b[%v] should be %v rather than %v"
	b := map[int]int{ 0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9 }
	m := MakeMap(b)
	m.Swap(1, 3)

	switch {
	case b[0] != 0:				t.Fatalf(HAS_VALUE, 0, 0, b[0])
	case b[1] != 3:				t.Fatalf(HAS_VALUE, 1, 3, b[1])
	case b[2] != 2:				t.Fatalf(HAS_VALUE, 2, 2, b[2])
	case b[3] != 1:				t.Fatalf(HAS_VALUE, 3, 1, b[3])
	case b[4] != 4:				t.Fatalf(HAS_VALUE, 4, 4, b[4])
	}
}


func TestMapCount(t *testing.T) {
	b := map[int]int{ 0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9 }
	m := MakeMap(b)

	if c := m.Count(func(i interface{}) bool { return i.(int) > 4 }); c != 5 {
		t.Fatalf("Item count should be 5 and not %v", c)
	}

	if c := m.Count(func(i interface{}) bool { return i.(int) < 5 }); c != 5 {
		t.Fatalf("Item count should be 5 and not %v", c)
	}
}

func TestMapAny(t *testing.T) {
	b := map[int]int{ 0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9 }
	m := MakeMap(b)

	if !m.Any(func(i interface{}) bool { return i.(int) > 4 }) {
		t.Fatal("Should have values greater than 4")
	}

	if !m.Any(func(i interface{}) bool { return i.(int) < 5 }) {
		t.Fatal("Should have values less than 5")
	}
}

func TestMapAll(t *testing.T) {
	b := map[int]int{ 0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9 }
	m := MakeMap(b)

	if !m.All(func(i interface{}) bool { return i.(int) < 11 }) {
		t.Fatal("All values should be below 11")
	}

	if !m.All(func(i interface{}) bool { return i.(int) > -1 }) {
		t.Fatal("All values should be above -1")
	}
}

func TestMapNone(t *testing.T) {
	b := map[int]int{ 0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9 }
	m := MakeMap(b)

	if !m.None(func(i interface{}) bool { return i.(int) < 0 }) {
		t.Fatal("No values should be below 0")
	}

	if !m.None(func(i interface{}) bool { return i.(int) > 9 }) {
		t.Fatal("No values should be above 9")
	}
}

func TestMapOne(t *testing.T) {
	b := map[int]int{ 0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9 }
	m := MakeMap(b)

	if !m.One(func(i interface{}) bool { return i.(int) == 0 }) {
		t.Fatal("Should return true")
	}

	if m.One(func(i interface{}) bool { return i.(int) == -1 }) {
		t.Fatal("Should return false")
	}
}

func TestMapMany(t *testing.T) {
	LOGIC_FAILURE := "Should return %v for %v detected"
	b := map[int]int{ 0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9 }
	m := MakeMap(b)

	if !m.Many(func(i interface{}) bool { return i.(int) > 0 }) {
		t.Fatal(LOGIC_FAILURE, true, "many values")
	}

	if m.Many(func(i interface{}) bool { return i.(int) == 0 }) {
		t.Fatal(LOGIC_FAILURE, false, "single value")
	}

	if m.Many(func(i interface{}) bool { return i.(int) < 0 }) {
		t.Fatal(LOGIC_FAILURE, false, "no values")
	}
}

func TestMapCollect(t *testing.T) {
	INCORRECT_VALUE := "r[%v] == %v"
	b := map[int]int{ 0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9 }
	m := MakeMap(b)
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

func TestMapInject(t *testing.T) {
	b := map[int]int{ 0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9 }
	m := MakeMap(b)
	r := m.Inject(0, func(seed, v interface{}) interface{} {
		return seed.(int) + v.(int)
	})
	switch {
	case r == nil:				t.Fatal("Inject() returned a nil value")
	case r != 45:				t.Fatalf("r should be 45 but is %v", r)
	}
}

func TestMapCombine(t *testing.T) {
	INCORRECT_VALUE := "r[%v] == %v"
	b := map[int]int{ 0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9 }
	m := MakeMap(b)
	r := m.Combine(m, func(x, y interface{}) interface{} {
		return x.(int) * y.(int)
	})
	switch {
	case r == nil:				t.Fatal("Combine() returned a nil value")
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
	case r.At(1) != 1:			t.Fatalf(INCORRECT_VALUE, 1, r.At(1))
	case r.At(2) != 4:			t.Fatalf(INCORRECT_VALUE, 2, r.At(2))
	case r.At(3) != 9:			t.Fatalf(INCORRECT_VALUE, 3, r.At(3))
	case r.At(4) != 16:			t.Fatalf(INCORRECT_VALUE, 4, r.At(4))
	case r.At(5) != 25:			t.Fatalf(INCORRECT_VALUE, 5, r.At(5))
	case r.At(6) != 36:			t.Fatalf(INCORRECT_VALUE, 6, r.At(6))
	case r.At(7) != 49:			t.Fatalf(INCORRECT_VALUE, 7, r.At(7))
	case r.At(8) != 64:			t.Fatalf(INCORRECT_VALUE, 8, r.At(8))
	case r.At(9) != 81:			t.Fatalf(INCORRECT_VALUE, 9, r.At(9))
	}
}

func TestMapCycle(t *testing.T) {
	b := map[int]int{ 0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9 }
	m := MakeMap(b)
	r := m.Cycle(5, func(x interface{}) {})
	switch {
	case r != 5:				t.Fatalf("cycle count should be 5 but is %v", r)
	}
}

func TestMapFeed(t *testing.T) {
	b := map[int]int{ 0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9 }
	m := MakeMap(b)
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
	b := map[int]int{ 0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9 }
	m := MakeMap(b)
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

/*
func TestMapTee(t *testing.T) {
	t.Fatal(NO_TESTS)
}
*/