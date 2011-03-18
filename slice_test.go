package raw

import "testing"

func TestMakeSlice(t *testing.T) {
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

func TestClear(t *testing.T) {
	t.Fatal("Implement Test")
}

func TestReplicate(t *testing.T) {
	b := []int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	s := MakeSlice(b)
	c := 10
	s = s.Replicate(c)

	switch {
	case s.Len() != len(b) * c:	t.Fatalf("Slice length should be %v not %v", len(b) * c, s.Len())
	case s.Cap() != cap(b) * c:	t.Fatalf("Slice capacity should be %v not %v", cap(b) * c, s.Cap())
	default:					t.Fatal("Add further test filters")
	}
}

func TestClone(t *testing.T) {
	b := []int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	s := MakeSlice(b)
	clone := s.Clone()

	switch {
	case clone.Len() != len(b):	t.Fatalf("Slice length should be %v not %v", len(b), clone.Len())
	case clone.Cap() != cap(b):	t.Fatalf("Slice capacity should be %v not %v", cap(b), clone.Cap())
	default:					t.Fatal("Add further test filters")
	}
}

func TestCount(t *testing.T) {
	b := []int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	s := MakeSlice(b)

	if c := s.Count(func(i interface{}) bool { return i.(int) > 4 }); c != 5 {
		t.Fatalf("Item count should be 5 and not %v", c)
	}

	if c := s.Count(func(i interface{}) bool { return i.(int) < 5 }); c != 5 {
		t.Fatalf("Item count should be 5 and not %v", c)
	}
}

func TestAny(t *testing.T) {
	b := []int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	s := MakeSlice(b)
	if !s.Any(func(i interface{}) bool { return i.(int) > 4 }) {
		t.Fatal("Should have values greater than 4")
	}

	if !s.Any(func(i interface{}) bool { return i.(int) < 5 }) {
		t.Fatal("Should have values less than 5")
	}
}

func TestAll(t *testing.T) {
	b := []int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	s := MakeSlice(b)
	if !s.All(func(i interface{}) bool { return i.(int) < 11 }) {
		t.Fatal("All values should be below 11")
	}

	if !s.All(func(i interface{}) bool { return i.(int) > -1 }) {
		t.Fatal("All values should be above -1")
	}
}

func TestNone(t *testing.T) {
	b := []int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	s := MakeSlice(b)
	if !s.None(func(i interface{}) bool { return i.(int) < 0 }) {
		t.Fatal("No values should be below 0")
	}

	if !s.None(func(i interface{}) bool { return i.(int) > 9 }) {
		t.Fatal("No values should be above 9")
	}
}

func TestOne(t *testing.T) {
	b := []int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	s := MakeSlice(b)
	r := s.One(func(i interface{}) bool {
		return i.(int) == 0
	})
	if !r { t.Fatal("") }
}