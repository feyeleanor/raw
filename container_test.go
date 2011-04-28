package raw

import "testing"

func TestCopy(t *testing.T) {
	SHOULD_MATCH := "Slice elements s[%v] and copy[%v] should match but are %v and %v"

	_, s := initSliceTest()
	c := Copy(s).(Sequence)
	c.Store(0, 1000)
	switch {
	case c.Len() != s.Len():	t.Fatalf("Slice length should be %v not %v", s.Len(), c.Len())
	case c.Cap() != s.Cap():	t.Fatalf("Slice capacity should be %v not %v", s.Cap(), c.Cap())
	case s.At(0) != 0:			t.Fatalf("Slice element s[%v] should be %v and not", 0, 0, s.At(0))
	case c.At(0) != 1000:		t.Fatalf("Slice element copy[%v] should be %v and not", 0, 1000, c.At(0))
	case c.At(1) != s.At(1):	t.Fatalf(SHOULD_MATCH, 1, 1, s.At(1), c.At(1))
	case c.At(2) != s.At(2):	t.Fatalf(SHOULD_MATCH, 2, 2, s.At(2), c.At(2))
	case c.At(3) != s.At(3):	t.Fatalf(SHOULD_MATCH, 3, 3, s.At(3), c.At(3))
	case c.At(4) != s.At(4):	t.Fatalf(SHOULD_MATCH, 4, 4, s.At(4), c.At(4))
	case c.At(5) != s.At(5):	t.Fatalf(SHOULD_MATCH, 5, 5, s.At(5), c.At(5))
	case c.At(6) != s.At(6):	t.Fatalf(SHOULD_MATCH, 6, 6, s.At(6), c.At(6))
	case c.At(7) != s.At(7):	t.Fatalf(SHOULD_MATCH, 7, 7, s.At(7), c.At(7))
	case c.At(8) != s.At(8):	t.Fatalf(SHOULD_MATCH, 8, 8, s.At(8), c.At(8))
	case c.At(9) != s.At(9):	t.Fatalf(SHOULD_MATCH, 9, 9, s.At(9), c.At(9))
	}

	_, m := initMapTest()
	n := Copy(m).(Mapping)
	n.Store(0, 1000)
	switch {
	case n.Len() != s.Len():	t.Fatalf("Slice length should be %v not %v", s.Len(), c.Len())
	case s.At(0) != 0:			t.Fatalf("Slice element s[%v] should be %v and not", 0, 0, s.At(0))
	case n.At(0) != 1000:		t.Fatalf("Slice element copy[%v] should be %v and not", 0, 1000, c.At(0))
	case n.At(1) != s.At(1):	t.Fatalf(SHOULD_MATCH, 1, 1, s.At(1), c.At(1))
	case n.At(2) != s.At(2):	t.Fatalf(SHOULD_MATCH, 2, 2, s.At(2), c.At(2))
	case n.At(3) != s.At(3):	t.Fatalf(SHOULD_MATCH, 3, 3, s.At(3), c.At(3))
	case n.At(4) != s.At(4):	t.Fatalf(SHOULD_MATCH, 4, 4, s.At(4), c.At(4))
	case n.At(5) != s.At(5):	t.Fatalf(SHOULD_MATCH, 5, 5, s.At(5), c.At(5))
	case n.At(6) != s.At(6):	t.Fatalf(SHOULD_MATCH, 6, 6, s.At(6), c.At(6))
	case n.At(7) != s.At(7):	t.Fatalf(SHOULD_MATCH, 7, 7, s.At(7), c.At(7))
	case n.At(8) != s.At(8):	t.Fatalf(SHOULD_MATCH, 8, 8, s.At(8), c.At(8))
	case n.At(9) != s.At(9):	t.Fatalf(SHOULD_MATCH, 9, 9, s.At(9), c.At(9))
	}
}

func TestSwapElements(t *testing.T) {
	HAS_VALUE := "%v[%v] should be %v rather than %v"

	_, s := initSliceTest()
	SwapElements(s, 1, 3)
	switch {
	case s.At(0) != 0:			t.Fatalf(HAS_VALUE, "slice", 0, 0, s.At(0))
	case s.At(1) != 3:			t.Fatalf(HAS_VALUE, "slice", 1, 3, s.At(1))
	case s.At(2) != 2:			t.Fatalf(HAS_VALUE, "slice", 2, 2, s.At(2))
	case s.At(3) != 1:			t.Fatalf(HAS_VALUE, "slice", 3, 1, s.At(3))
	case s.At(4) != 4:			t.Fatalf(HAS_VALUE, "slice", 4, 4, s.At(4))
	}

	_, m := initMapTest()
	SwapElements(m, 1, 3)
	switch {
	case m.At(0) != 0:			t.Fatalf(HAS_VALUE, "map", 0, 0, m.At(0))
	case m.At(1) != 3:			t.Fatalf(HAS_VALUE, "map", 1, 3, m.At(1))
	case m.At(2) != 2:			t.Fatalf(HAS_VALUE, "map", 2, 2, m.At(2))
	case m.At(3) != 1:			t.Fatalf(HAS_VALUE, "map", 3, 1, m.At(3))
	case m.At(4) != 4:			t.Fatalf(HAS_VALUE, "map", 4, 4, m.At(4))
	}
}

func TestEach(t *testing.T) {
	_, s := initSliceTest()
	c := 0
	n := Each(s, func(i interface{}) {
		c += i.(int)
	})
	switch {
	case n != 10:				t.Fatalf("Count should be 10 and not %v", n)
	case c != 45:				t.Fatalf("Sum should be 45 and not %v", c)
	}
}

func TestCycle(t *testing.T) {
	COUNT_INCORRECT := "%v cycle count should be %v but is %v"
	_, s := initSliceTest()
	r := Cycle(s, 5, func(x interface{}) {})
	switch {
	case r != 5:				t.Fatalf(COUNT_INCORRECT, "Slice", 5, r)
	}

	_, m := initMapTest()
	r = Cycle(m, 5, func(x interface{}) {})
	switch {
	case r != 5:				t.Fatalf(COUNT_INCORRECT, "Map", 5, r)
	}
}

func TestCollect(t *testing.T) {
	INCORRECT_VALUE := "r[%v] == %v"

	b, s := initSliceTest()
	r := Collect(s, func(i interface{}) interface{} {
		return i.(int) * 2
	})
	switch {
	case r == nil:				t.Fatal("Collect() returned a nil value")
	case r.Len() != len(b):		t.Fatalf("slice length should be %v but is %v", len(b), r.Len())
	case s.At(0) != 0:			t.Fatalf(INCORRECT_VALUE, 0, s.At(0))
	case s.At(1) != 1:			t.Fatalf(INCORRECT_VALUE, 1, s.At(1))
	case s.At(2) != 2:			t.Fatalf(INCORRECT_VALUE, 2, s.At(2))
	case s.At(3) != 3:			t.Fatalf(INCORRECT_VALUE, 3, s.At(3))
	case s.At(4) != 4:			t.Fatalf(INCORRECT_VALUE, 4, s.At(4))
	case s.At(5) != 5:			t.Fatalf(INCORRECT_VALUE, 5, s.At(5))
	case s.At(6) != 6:			t.Fatalf(INCORRECT_VALUE, 6, s.At(6))
	case s.At(7) != 7:			t.Fatalf(INCORRECT_VALUE, 7, s.At(7))
	case s.At(8) != 8:			t.Fatalf(INCORRECT_VALUE, 8, s.At(8))
	case s.At(9) != 9:			t.Fatalf(INCORRECT_VALUE, 9, s.At(9))
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

func TestReduce(t *testing.T) {
	_, s := initSliceTest()
	r := Reduce(s, 0, func(memo, x interface{}) interface{} {
		return memo.(int) + x.(int)
	})
	switch {
	case r == nil:				t.Fatal("Reduce() returned a nil value")
	case r != 45:				t.Fatalf("r should be 45 but is %v", r)
	}
}

func TestWhile(t *testing.T) {
	_, s := initSliceTest()
	count := While(s, func(i interface{}) bool {
		return i.(int) < 6
	})
	switch {
	case count != 6:				t.Fatalf("Slice count should be %v not %v", 6, count)
	}

	b, c := initChannelTest()
	count = While(c, func(i interface{}) bool {
		return i.(int) < 6
	})
	switch {
	case count != 6:				t.Fatalf("Channel count should be %v not %v", 6, count)
	case c.Len() != 3:				t.Fatalf("Channel length should be %v not %v", 3, c.Len())
	case c.Len() != len(b):			t.Fatalf("Channel length should be %v not %v", len(b), c.Len())
	}
}

func TestUntil(t *testing.T) {
	_, s := initSliceTest()
	count := Until(s, func(i interface{}) bool {
		return i.(int) == 6
	})
	switch {
	case count != 6:				t.Fatalf("Slice count should be %v not %v", 6, count)
	}

	b, c := initChannelTest()
	count = Until(c, func(i interface{}) bool {
		return i.(int) == 6
	})
	switch {
	case count != 6:				t.Fatalf("Channel count should be %v not %v", 6, count)
	case c.Len() != 3:				t.Fatalf("Channel length should be %v not %v", 3, c.Len())
	case c.Len() != len(b):			t.Fatalf("Channel length should be %v not %v", len(b), c.Len())
	}
}

func TestCount(t *testing.T) {
	_, s := initSliceTest()
	if n := Count(s, func(i interface{}) bool { return i.(int) > 5 }); n != 4 {
		t.Fatalf("Slice count should be 4 and not %v", n)
	}

	if n := Count(s, func(i interface{}) bool { return i.(int) < 5 }); n != 5 {
		t.Fatalf("Slice count should be 5 and not %v", n)
	}

	_, m := initMapTest()
	if n := Count(m, func(i interface{}) bool { return i.(int) > 5 }); n != 4 {
		t.Fatalf("Map count should be 4 and not %v", n)
	}

	if n := Count(m, func(i interface{}) bool { return i.(int) < 5 }); n != 5 {
		t.Fatalf("Map count should be 5 and not %v", n)
	}

	_, c := initChannelTest()
	if n := Count(c, func(i interface{}) bool { return i.(int) > 5 }); n != 4 {
		t.Fatalf("Channel count should be 4 and not %v", n)
	}

	_, c = initChannelTest()
	if n := Count(c, func(i interface{}) bool { return i.(int) < 5 }); n != 5 {
		t.Fatalf("Channel count should be 5 and not %v", n)
	}
}

func TestAny(t *testing.T) {
	_, s := initSliceTest()
	if !Any(s, func(i interface{}) bool { return i.(int) > 4 }) {
		t.Fatal("Slice should have values greater than 4")
	}

	if !Any(s, func(i interface{}) bool { return i.(int) < 5 }) {
		t.Fatal("Slice should have values less than 5")
	}

	if Any(s, func(i interface{}) bool { return i.(int) < 0 }) {
		t.Fatal("Slice should not have values less than 0")
	}

	if Any(s, func(i interface{}) bool { return i.(int) > 10 }) {
		t.Fatal("Slice should not have values greater than 10")
	}

	_, m := initMapTest()
	if !Any(m, func(i interface{}) bool { return i.(int) > 4 }) {
		t.Fatal("Map should have values greater than 4")
	}

	if !Any(m, func(i interface{}) bool { return i.(int) < 5 }) {
		t.Fatal("Map should have values less than 5")
	}

	if Any(m, func(i interface{}) bool { return i.(int) < 0 }) {
		t.Fatal("Map should not have values less than 0")
	}

	if Any(m, func(i interface{}) bool { return i.(int) > 10 }) {
		t.Fatal("Map should not have values greater than 10")
	}
}

func TestAll(t *testing.T) {
	_, s := initSliceTest()
	if !All(s, func(i interface{}) bool { return i.(int) < 11 }) {
		t.Fatal("Slice values should all be less than 11")
	}

	if !All(s, func(i interface{}) bool { return i.(int) > -1 }) {
		t.Fatal("Slice values should all be greater than -1")
	}

	if All(s, func(i interface{}) bool { return i.(int) < 0 }) {
		t.Fatal("Slice should not have values less than 0")
	}

	if All(s, func(i interface{}) bool { return i.(int) > 10 }) {
		t.Fatal("Slice should not have values greater than 10")
	}

	_, m := initMapTest()
	if !All(m, func(i interface{}) bool { return i.(int) < 11 }) {
		t.Fatal("Map values should all be less than 11")
	}

	if !All(m, func(i interface{}) bool { return i.(int) > -1 }) {
		t.Fatal("Map values should all be greater than -1")
	}

	if All(m, func(i interface{}) bool { return i.(int) < 0 }) {
		t.Fatal("Map should not have values less than 0")
	}

	if All(m, func(i interface{}) bool { return i.(int) > 10 }) {
		t.Fatal("Map should not have values greater than 10")
	}
}

func TestNone(t *testing.T) {
	_, s := initSliceTest()
	if None(s, func(i interface{}) bool { return i.(int) > 0 }) {
		t.Fatal("There should be slice values greater than 0")
	}

	if None(s, func(i interface{}) bool { return i.(int) < 10 }) {
		t.Fatal("There should be slice values less than 10")
	}

	if !None(s, func(i interface{}) bool { return i.(int) < 0 }) {
		t.Fatal("There should not be slice values less than 0")
	}

	if !None(s, func(i interface{}) bool { return i.(int) > 10 }) {
		t.Fatal("There should not be slice values greater than 10")
	}

	_, m := initMapTest()
	if None(m, func(i interface{}) bool { return i.(int) > 0 }) {
		t.Fatal("There should be map values greater than 0")
	}

	if None(m, func(i interface{}) bool { return i.(int) < 10 }) {
		t.Fatal("There should be map values less than 10")
	}

	if !None(m, func(i interface{}) bool { return i.(int) < 0 }) {
		t.Fatal("There should not be map values less than 0")
	}

	if !None(m, func(i interface{}) bool { return i.(int) > 10 }) {
		t.Fatal("There should not be map values greater than 10")
	}
}

func TestOne(t *testing.T) {
	_, s := initSliceTest()
	if !One(s, func(i interface{}) bool { return i.(int) == 0 }) {
		t.Fatal("Slice should return true")
	}

	if One(s, func(i interface{}) bool { return i.(int) == -1 }) {
		t.Fatal("Slice should return false")
	}

	_, m := initMapTest()
	if !One(m, func(i interface{}) bool { return i.(int) == 0 }) {
		t.Fatal("Map should return true")
	}

	if One(m, func(i interface{}) bool { return i.(int) == -1 }) {
		t.Fatal("Map should return false")
	}
}

func TestDense(t *testing.T) {
	LOGIC_FAILURE := "%v should return %v for %v detected"
	_, s := initSliceTest()
	if !Dense(s, 0.5, func(i interface{}) bool { return i.(int) > 0 }) {
		t.Fatalf(LOGIC_FAILURE, "Slice", true, "most values")
	}

	if Dense(s, 0.5, func(i interface{}) bool { return i.(int) == 0 }) {
		t.Fatalf(LOGIC_FAILURE, "Slice", false, "single value")
	}

	if Dense(s, 0.5, func(i interface{}) bool { return i.(int) < 0 }) {
		t.Fatalf(LOGIC_FAILURE, "Slice", false, "no values")
	}

	_, m := initMapTest()
	if !Dense(m, 0.5, func(i interface{}) bool { return i.(int) > 0 }) {
		t.Fatalf(LOGIC_FAILURE, "Map", true, "most values")
	}

	if Dense(m, 0.5, func(i interface{}) bool { return i.(int) == 0 }) {
		t.Fatalf(LOGIC_FAILURE, "Map", false, "single value")
	}

	if Dense(m, 0.5, func(i interface{}) bool { return i.(int) < 0 }) {
		t.Fatalf(LOGIC_FAILURE, "Map", false, "no values")
	}
}

func TestMost(t *testing.T) {
	LOGIC_FAILURE := "%v should return %v for %v detected"
	_, s := initSliceTest()
	if !Most(s, func(i interface{}) bool { return i.(int) > 0 }) {
		t.Fatalf(LOGIC_FAILURE, "Slice", true, "most values")
	}

	if Most(s, func(i interface{}) bool { return i.(int) == 0 }) {
		t.Fatalf(LOGIC_FAILURE, "Slice", false, "single value")
	}

	if Most(s, func(i interface{}) bool { return i.(int) < 0 }) {
		t.Fatalf(LOGIC_FAILURE, "Slice", false, "no values")
	}

	_, m := initMapTest()
	if !Most(m, func(i interface{}) bool { return i.(int) > 0 }) {
		t.Fatalf(LOGIC_FAILURE, "Map", true, "most values")
	}

	if Most(m, func(i interface{}) bool { return i.(int) == 0 }) {
		t.Fatalf(LOGIC_FAILURE, "Map", false, "single value")
	}

	if Most(m, func(i interface{}) bool { return i.(int) < 0 }) {
		t.Fatalf(LOGIC_FAILURE, "Map", false, "no values")
	}
}

func TestCombine(t *testing.T) {
	INCORRECT_VALUE := "r[%v] == %v"

	_, s := initSliceTest()
	r := Combine(s, s, func(x, y interface{}) interface{} {
		return x.(int) * y.(int)
	}).(Sequence)
	switch {
	case r == nil:				t.Fatal("Combine() slice returned a nil value")
	case s.At(0) != 0:			t.Fatalf(INCORRECT_VALUE, 0, s.At(0))
	case s.At(1) != 1:			t.Fatalf(INCORRECT_VALUE, 1, s.At(1))
	case s.At(2) != 2:			t.Fatalf(INCORRECT_VALUE, 2, s.At(2))
	case s.At(3) != 3:			t.Fatalf(INCORRECT_VALUE, 3, s.At(3))
	case s.At(4) != 4:			t.Fatalf(INCORRECT_VALUE, 4, s.At(4))
	case s.At(5) != 5:			t.Fatalf(INCORRECT_VALUE, 5, s.At(5))
	case s.At(6) != 6:			t.Fatalf(INCORRECT_VALUE, 6, s.At(6))
	case s.At(7) != 7:			t.Fatalf(INCORRECT_VALUE, 7, s.At(7))
	case s.At(8) != 8:			t.Fatalf(INCORRECT_VALUE, 8, s.At(8))
	case s.At(9) != 9:			t.Fatalf(INCORRECT_VALUE, 9, s.At(9))
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

	_, m := initMapTest()
	x := Combine(m, m, func(x, y interface{}) interface{} {
		return x.(int) * y.(int)
	}).(Mapping)
	switch {
	case r == nil:				t.Fatal("Combine() map returned a nil value")
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
	case x.At(0) != 0:			t.Fatalf(INCORRECT_VALUE, 0, x.At(0))
	case x.At(1) != 1:			t.Fatalf(INCORRECT_VALUE, 1, x.At(1))
	case x.At(2) != 4:			t.Fatalf(INCORRECT_VALUE, 2, x.At(2))
	case x.At(3) != 9:			t.Fatalf(INCORRECT_VALUE, 3, x.At(3))
	case x.At(4) != 16:			t.Fatalf(INCORRECT_VALUE, 4, x.At(4))
	case x.At(5) != 25:			t.Fatalf(INCORRECT_VALUE, 5, x.At(5))
	case x.At(6) != 36:			t.Fatalf(INCORRECT_VALUE, 6, x.At(6))
	case x.At(7) != 49:			t.Fatalf(INCORRECT_VALUE, 7, x.At(7))
	case x.At(8) != 64:			t.Fatalf(INCORRECT_VALUE, 8, x.At(8))
	case x.At(9) != 81:			t.Fatalf(INCORRECT_VALUE, 9, x.At(9))
	}
}