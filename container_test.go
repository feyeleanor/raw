package raw

import "testing"

func TestNewContainer(t *testing.T) {
	SHOULD_CONTAIN := "%v[%v] should contain %v but contains %v"

	s := NewContainer([]int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }).(Sequence)
	switch {
	case s == nil:				t.Fatal("Make slice returned a nil value")
	case s.Len() != 10:			t.Fatalf("Slice length should be %v not %v", 10, s.Len())
	case s.At(0) != 0:			t.Fatalf(SHOULD_CONTAIN, "slice", 0, s.At(0))
	case s.At(1) != 1:			t.Fatalf(SHOULD_CONTAIN, "slice", 1, s.At(1))
	case s.At(2) != 2:			t.Fatalf(SHOULD_CONTAIN, "slice", 2, s.At(2))
	case s.At(3) != 3:			t.Fatalf(SHOULD_CONTAIN, "slice", 3, s.At(3))
	case s.At(4) != 4:			t.Fatalf(SHOULD_CONTAIN, "slice", 4, s.At(4))
	case s.At(5) != 5:			t.Fatalf(SHOULD_CONTAIN, "slice", 5, s.At(5))
	case s.At(6) != 6:			t.Fatalf(SHOULD_CONTAIN, "slice", 6, s.At(6))
	case s.At(7) != 7:			t.Fatalf(SHOULD_CONTAIN, "slice", 7, s.At(7))
	case s.At(8) != 8:			t.Fatalf(SHOULD_CONTAIN, "slice", 8, s.At(8))
	case s.At(9) != 9:			t.Fatalf(SHOULD_CONTAIN, "slice", 9, s.At(9))
	}

	m := NewContainer(map[int]int{ 0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9 }).(Mapping)
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

	b := make(chan int, 16)
	go func() {
		for _, v := range []int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 } {
			b <- v
		}
		close(b)
	}()
	c := NewContainer(b).(*Channel)

	switch {
	case c == nil:					t.Fatal("MakeChannel returned a nil value")
	case c.Len() != len(b):			t.Fatalf("Channel length should be %v not %v", len(b), c.Len())
	case c.Cap() != cap(b):			t.Fatalf("Channel capacity should be %v not %v", cap(b), c.Cap())
	}

	for i := 0; i < 10; i++ {
		switch v, open := c.Recv(); {
		case !open:					t.Fatalf("%v: channel should be open", i)
		case v != i:				t.Fatalf("Should receive %v but received %v", i, v)
		}
	}

	if _, open := c.TryRecv(); open {
		t.Fatal("Channel should be closed")
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

func TestInject(t *testing.T) {
	_, s := initSliceTest()
	r := Inject(s, 0, func(memo, x interface{}) interface{} {
		return memo.(int) + x.(int)
	})
	switch {
	case r == nil:				t.Fatal("Inject() returned a nil value")
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