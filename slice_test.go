package raw

import "testing"

func TestSliceMakeSlice(t *testing.T) {
	t.Fatal("Implement Test")
}

func TestSliceSizing(t *testing.T) {
	b := []int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	s := MakeSlice(b)

	switch {
	case s.Len() != len(b):		t.Fatalf("Slice length should be %v not %v", len(b), s.Len())
	case s.Cap() != cap(b):		t.Fatalf("Slice capacity should be %v not %v", cap(b), s.Cap())
	}
}

func TestSliceCopy(t *testing.T) {
	b := []int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	s := MakeSlice(b)
	s.Copy(1, 3)

	switch {
	case b[1] != b[3]:			t.Fatalf("Elements b[1] and b[3] should match but are %v and %v", b[1], b[3])
	}
}

func TestSliceSwap(t *testing.T) {
	b := []int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	s := MakeSlice(b)
	s.Swap(1, 3)

	switch {
	case b[1] != 3:				t.Fatalf("b[1] should be 3 rather than %v", b[1])
	case b[3] != 1:				t.Fatalf("b[3] should be 1 rather than %v", b[3])
	}
}

func TestSliceClear(t *testing.T) {
	t.Fatal("Implement Test")
}

func TestSliceRepeat(t *testing.T) {
	b := []int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	s := MakeSlice(b)
	c := 10
	s = s.Repeat(c)

	switch {
	case s.Len() != len(b) * c:	t.Fatalf("Slice length should be %v not %v", len(b) * c, s.Len())
	case s.Cap() != cap(b) * c:	t.Fatalf("Slice capacity should be %v not %v", cap(b) * c, s.Cap())
	default:					t.Fatal("Add further test filters")
	}
}

func TestSliceClone(t *testing.T) {
	b := []int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	s := MakeSlice(b)
	clone := s.Clone()

	switch {
	case clone.Len() != len(b):	t.Fatalf("Slice length should be %v not %v", len(b), clone.Len())
	case clone.Cap() != cap(b):	t.Fatalf("Slice capacity should be %v not %v", cap(b), clone.Cap())
	default:					t.Fatal("Add further test filters")
	}
}

func TestSliceCount(t *testing.T) {
	b := []int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	s := MakeSlice(b)

	if c := s.Count(func(i interface{}) bool { return i.(int) > 4 }); c != 5 {
		t.Fatalf("Item count should be 5 and not %v", c)
	}

	if c := s.Count(func(i interface{}) bool { return i.(int) < 5 }); c != 5 {
		t.Fatalf("Item count should be 5 and not %v", c)
	}
}

func TestSliceAny(t *testing.T) {
	b := []int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	s := MakeSlice(b)
	if !s.Any(func(i interface{}) bool { return i.(int) > 4 }) {
		t.Fatal("Should have values greater than 4")
	}

	if !s.Any(func(i interface{}) bool { return i.(int) < 5 }) {
		t.Fatal("Should have values less than 5")
	}
}

func TestSliceAll(t *testing.T) {
	b := []int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	s := MakeSlice(b)
	if !s.All(func(i interface{}) bool { return i.(int) < 11 }) {
		t.Fatal("All values should be below 11")
	}

	if !s.All(func(i interface{}) bool { return i.(int) > -1 }) {
		t.Fatal("All values should be above -1")
	}
}

func TestSliceNone(t *testing.T) {
	b := []int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	s := MakeSlice(b)
	if !s.None(func(i interface{}) bool { return i.(int) < 0 }) {
		t.Fatal("No values should be below 0")
	}

	if !s.None(func(i interface{}) bool { return i.(int) > 9 }) {
		t.Fatal("No values should be above 9")
	}
}

func TestSliceOne(t *testing.T) {
	b := []int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	s := MakeSlice(b)
	r := s.One(func(i interface{}) bool {
		return i.(int) == 0
	})
	if !r { t.Fatal("") }
}

func TestSliceAt(t *testing.T) {
	t.Fatal("Implement Test")
}

func TestSliceSet(t *testing.T) {
	t.Fatal("Implement Test")
}

func TestSliceCollect(t *testing.T) {
	b := []int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	s := MakeSlice(b)
	r := s.Collect(func(i interface{}) interface{} {
		return i.(int) * 2
	})
	switch {
	case r == nil:				t.Fatal("Collect() returned a nil value")
	case r.Cap() != cap(b):		t.Fatalf("capacity should be %v but is %v", cap(b), r.Cap())
	case r.Len() != len(b):		t.Fatalf("capacity should be %v but is %v", len(b), r.Len())
	case r.At(0) != 0:			t.Fatalf("r[0] == %v", r.At(0))
	case r.At(1) != 2:			t.Fatalf("r[1] == %v", r.At(1))
	case r.At(2) != 4:			t.Fatalf("r[2] == %v", r.At(2))
	case r.At(3) != 6:			t.Fatalf("r[3] == %v", r.At(3))
	case r.At(4) != 8:			t.Fatalf("r[4] == %v", r.At(4))
	case r.At(5) != 10:			t.Fatalf("r[5] == %v", r.At(5))
	case r.At(6) != 12:			t.Fatalf("r[6] == %v", r.At(6))
	case r.At(7) != 14:			t.Fatalf("r[7] == %v", r.At(7))
	case r.At(8) != 16:			t.Fatalf("r[8] == %v", r.At(8))
	case r.At(9) != 18:			t.Fatalf("r[9] == %v", r.At(9))
	}
	t.Fatal("Add further test filters")
}

func TestSliceInject(t *testing.T) {
	b := []int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	s := MakeSlice(b)
	r := s.Inject(0, func(memo, x interface{}) interface{} {
		return memo.(int) + x.(int)
	})
	switch {
	case r == nil:				t.Fatal("Inject() returned a nil value")
	case r != 45:				t.Fatalf("r should be 45 but is %v", r)
	}
	t.Fatal("Add further test filters")
}

func TestSliceCombine(t *testing.T) {
	b := []int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	s := MakeSlice(b)
	r := s.Combine(s, func(x, y interface{}) interface{} {
		return x.(int) * y.(int)
	})
	switch {
	case r == nil:				t.Fatal("Combine() returned a nil value")
	case r.Cap() != cap(b):		t.Fatalf("capacity should be %v but is %v", cap(b), r.Cap())
	case r.Len() != len(b):		t.Fatalf("capacity should be %v but is %v", len(b), r.Len())
	case r.At(0) != 0:			t.Fatalf("r[0] == %v", r.At(0))
	case r.At(1) != 1:			t.Fatalf("r[1] == %v", r.At(1))
	case r.At(2) != 4:			t.Fatalf("r[2] == %v", r.At(2))
	case r.At(3) != 9:			t.Fatalf("r[3] == %v", r.At(3))
	case r.At(4) != 16:			t.Fatalf("r[4] == %v", r.At(4))
	case r.At(5) != 25:			t.Fatalf("r[5] == %v", r.At(5))
	case r.At(6) != 36:			t.Fatalf("r[6] == %v", r.At(6))
	case r.At(7) != 49:			t.Fatalf("r[7] == %v", r.At(7))
	case r.At(8) != 64:			t.Fatalf("r[8] == %v", r.At(8))
	case r.At(9) != 81:			t.Fatalf("r[9] == %v", r.At(9))
	}
	t.Fatal("Add further test filters")
}

func TestSliceCycle(t *testing.T) {
	b := []int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	s := MakeSlice(b)
	r := s.Cycle(5, func(i int, x interface{}) {})
	switch {
	case r == nil:				t.Fatal("Cycle() returned a nil value")
	case r != 5:				t.Fatalf("r should be 5 but is %v", r)
	}
	t.Fatal("Add further test filters")
}

func TestSliceResize(t *testing.T) {
	b := []int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	s := MakeSlice(b)
	switch s.Resize(20); {
	case b == nil:				t.Fatal("Resize() created a nil value for original slice")
	case s == nil:				t.Fatal("Resize() created a nil value for Slice")
	case cap(b) != 10:			t.Fatalf("original slice capacity should be 10 but is %v", cap(b))
	case len(b) != 10:			t.Fatalf("original slice length should be 10 but is %v", len(b))
	case s.Cap() != 20:			t.Fatalf("Slice capacity should be 20 but is %v", s.Cap())
	case s.Len() != 10:			t.Fatalf("Slice length should be 10 but is %v", s.Len())
	}

	switch s.Resize(5); {
	case b == nil:				t.Fatal("Resize() created a nil value for original slice")
	case s == nil:				t.Fatal("Resize() created a nil value for Slice")
	case cap(b) != 10:			t.Fatalf("original slice capacity should be 10 but is %v", cap(b))
	case len(b) != 10:			t.Fatalf("original slice length should be 10 but is %v", len(b))
	case s.Cap() != 5:			t.Fatalf("Slice capacity should be 5 but is %v", s.Cap())
	case s.Len() != 5:			t.Fatalf("Slice length should be 5 but is %v", s.Len())
	}
}

func TestSliceExtend(t *testing.T) {
	b := []int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	s := MakeSlice(b)
	switch s.Extend(10); {
		case b == nil:				t.Fatal("Resize() created a nil value for original slice")
		case s == nil:				t.Fatal("Resize() created a nil value for Slice")
		case cap(b) != 10:			t.Fatalf("original slice capacity should be 10 but is %v", cap(b))
		case len(b) != 10:			t.Fatalf("original slice length should be 10 but is %v", len(b))
		case s.Cap() != 20:			t.Fatalf("Slice capacity should be 20 but is %v", s.Cap())
		case s.Len() != 10:			t.Fatalf("Slice length should be 10 but is %v", s.Len())
	}

	switch s.Extend(-15); {
	case b == nil:				t.Fatal("Resize() created a nil value for original slice")
	case s == nil:				t.Fatal("Resize() created a nil value for Slice")
	case cap(b) != 10:			t.Fatalf("original slice capacity should be 10 but is %v", cap(b))
	case len(b) != 10:			t.Fatalf("original slice length should be 10 but is %v", len(b))
	case s.Cap() != 5:			t.Fatalf("Slice capacity should be 5 but is %v", s.Cap())
	case s.Len() != 5:			t.Fatalf("Slice length should be 5 but is %v", s.Len())
	}
}

func TestSliceShrink(t *testing.T) {
	b := []int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	s := MakeSlice(b)
	switch s.Shrink(-10); {
		case b == nil:				t.Fatal("Resize() created a nil value for original slice")
		case s == nil:				t.Fatal("Resize() created a nil value for Slice")
		case cap(b) != 10:			t.Fatalf("original slice capacity should be 10 but is %v", cap(b))
		case len(b) != 10:			t.Fatalf("original slice length should be 10 but is %v", len(b))
		case s.Cap() != 20:			t.Fatalf("Slice capacity should be 20 but is %v", s.Cap())
		case s.Len() != 10:			t.Fatalf("Slice length should be 10 but is %v", s.Len())
	}

	switch s.Shrink(15); {
	case b == nil:				t.Fatal("Resize() created a nil value for original slice")
	case s == nil:				t.Fatal("Resize() created a nil value for Slice")
	case cap(b) != 10:			t.Fatalf("original slice capacity should be 10 but is %v", cap(b))
	case len(b) != 10:			t.Fatalf("original slice length should be 10 but is %v", len(b))
	case s.Cap() != 5:			t.Fatalf("Slice capacity should be 5 but is %v", s.Cap())
	case s.Len() != 5:			t.Fatalf("Slice length should be 5 but is %v", s.Len())
	}
}

func TestSliceDoubleCapacity(t *testing.T) {
	b := []int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	s := MakeSlice(b)
	switch s.DoubleCapacity(); {
		case b == nil:				t.Fatal("Resize() created a nil value for original slice")
		case s == nil:				t.Fatal("Resize() created a nil value for Slice")
		case cap(b) != 10:			t.Fatalf("original slice capacity should be 10 but is %v", cap(b))
		case len(b) != 10:			t.Fatalf("original slice length should be 10 but is %v", len(b))
		case s.Cap() != 20:			t.Fatalf("Slice capacity should be 20 but is %v", s.Cap())
		case s.Len() != 10:			t.Fatalf("Slice length should be 10 but is %v", s.Len())
	}
}

func TestSliceHalveCapacity(t *testing.T) {
	b := []int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	s := MakeSlice(b)
	switch s.HalveCapacity(); {
		case b == nil:				t.Fatal("Resize() created a nil value for original slice")
		case s == nil:				t.Fatal("Resize() created a nil value for Slice")
		case cap(b) != 10:			t.Fatalf("original slice capacity should be 10 but is %v", cap(b))
		case len(b) != 10:			t.Fatalf("original slice length should be 10 but is %v", len(b))
		case s.Cap() != 5:			t.Fatalf("Slice capacity should be 5 but is %v", s.Cap())
		case s.Len() != 5:			t.Fatalf("Slice length should be 5 but is %v", s.Len())
	}
}

func TestSliceFeed(t *testing.T) {
	b := []int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	s := MakeSlice(b)
	c := make(chan interface{})
	s.Feed(c, func(i int, x interface{}) interface{} {
		return i * x.(int)
	})
	n := []int{}
	for x := range c {
		n = append(n, x.(int))
	}
	switch {
	case n[0] != b[0] * b[0]:		t.Fatalf("%v: expected %v but got %v", 0, b[0] * b[0], n[0])
	case n[1] != b[1] * b[1]:		t.Fatalf("%v: expected %v but got %v", 1, b[1] * b[1], n[1])
	case n[2] != b[2] * b[2]:		t.Fatalf("%v: expected %v but got %v", 2, b[2] * b[2], n[2])
	case n[3] != b[3] * b[3]:		t.Fatalf("%v: expected %v but got %v", 3, b[3] * b[3], n[3])
	case n[4] != b[4] * b[4]:		t.Fatalf("%v: expected %v but got %v", 4, b[4] * b[4], n[4])
	case n[5] != b[5] * b[5]:		t.Fatalf("%v: expected %v but got %v", 5, b[5] * b[5], n[5])
	case n[6] != b[6] * b[6]:		t.Fatalf("%v: expected %v but got %v", 6, b[6] * b[6], n[6])
	case n[7] != b[7] * b[7]:		t.Fatalf("%v: expected %v but got %v", 7, b[7] * b[7], n[7])
	case n[8] != b[8] * b[8]:		t.Fatalf("%v: expected %v but got %v", 8, b[8] * b[8], n[8])
	case n[9] != b[9] * b[9]:		t.Fatalf("%v: expected %v but got %v", 9, b[9] * b[9], n[9])
	}
}

func TestSlicePipe(t *testing.T) {
	b := []int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	s := MakeSlice(b)
	c := s.Pipe(func(i int, x interface{}) interface{} {
		return i * x.(int)
	})
	n := []int{}
	for x := range c {
		n = append(n, x.(int))
	}
	switch {
	case n[0] != b[0] * b[0]:		t.Fatalf("%v: expected %v but got %v", 0, b[0] * b[0], n[0])
	case n[1] != b[1] * b[1]:		t.Fatalf("%v: expected %v but got %v", 1, b[1] * b[1], n[1])
	case n[2] != b[2] * b[2]:		t.Fatalf("%v: expected %v but got %v", 2, b[2] * b[2], n[2])
	case n[3] != b[3] * b[3]:		t.Fatalf("%v: expected %v but got %v", 3, b[3] * b[3], n[3])
	case n[4] != b[4] * b[4]:		t.Fatalf("%v: expected %v but got %v", 4, b[4] * b[4], n[4])
	case n[5] != b[5] * b[5]:		t.Fatalf("%v: expected %v but got %v", 5, b[5] * b[5], n[5])
	case n[6] != b[6] * b[6]:		t.Fatalf("%v: expected %v but got %v", 6, b[6] * b[6], n[6])
	case n[7] != b[7] * b[7]:		t.Fatalf("%v: expected %v but got %v", 7, b[7] * b[7], n[7])
	case n[8] != b[8] * b[8]:		t.Fatalf("%v: expected %v but got %v", 8, b[8] * b[8], n[8])
	case n[9] != b[9] * b[9]:		t.Fatalf("%v: expected %v but got %v", 9, b[9] * b[9], n[9])
	}
}