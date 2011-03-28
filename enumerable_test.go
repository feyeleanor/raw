package raw

import "testing"

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

func TestMost(t *testing.T) {
	LOGIC_FAILURE := "%v should return %v for %v detected"
	_, s := initSliceTest()
	if Most(s, func(i interface{}) bool { return i.(int) > 0 }) {
		t.Fatalf(LOGIC_FAILURE, "Slice", true, "most values")
	}

	if Most(s, func(i interface{}) bool { return i.(int) == 0 }) {
		t.Fatalf(LOGIC_FAILURE, "Slice", false, "single value")
	}

	if Most(s, func(i interface{}) bool { return i.(int) < 0 }) {
		t.Fatalf(LOGIC_FAILURE, "Slice", false, "no values")
	}

	_, m := initMapTest()
	if Most(m, func(i interface{}) bool { return i.(int) > 0 }) {
		t.Fatalf(LOGIC_FAILURE, "Map", true, "most values")
	}

	if Most(m, func(i interface{}) bool { return i.(int) == 0 }) {
		t.Fatalf(LOGIC_FAILURE, "Map", false, "single value")
	}

	if Most(m, func(i interface{}) bool { return i.(int) < 0 }) {
		t.Fatalf(LOGIC_FAILURE, "Map", false, "no values")
	}
}