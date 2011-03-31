package raw

import "testing"

func TestFirst(t *testing.T) {
	SHOULD_MATCH := "Slice elements s[%v] and r[%v] should match but are %v and %v"

	_, s := initSliceTest()
	r := First(s, 5)
	switch {
	case r.Len() != 5:			t.Fatalf("Slice length should be %v not %v", 5, r.Len())
	case s.At(0) != r.At(0):	t.Fatalf(SHOULD_MATCH, 0, 0, s.At(0), r.At(0))
	case s.At(1) != r.At(1):	t.Fatalf(SHOULD_MATCH, 1, 1, s.At(1), r.At(1))
	case s.At(2) != r.At(2):	t.Fatalf(SHOULD_MATCH, 2, 2, s.At(2), r.At(2))
	case s.At(3) != r.At(3):	t.Fatalf(SHOULD_MATCH, 3, 3, s.At(3), r.At(3))
	case s.At(4) != r.At(4):	t.Fatalf(SHOULD_MATCH, 4, 4, s.At(4), r.At(4))
	}
}

func TestLast(t *testing.T) {
	SHOULD_MATCH := "Slice elements s[%v] and r[%v] should match but are %v and %v"

	_, s := initSliceTest()
	r := Last(s, 5)
	switch {
	case r.Len() != 5:			t.Fatalf("Slice length should be %v not %v", 5, r.Len())
	case s.At(5) != r.At(0):	t.Fatalf(SHOULD_MATCH, 5, 0, s.At(5), r.At(0))
	case s.At(6) != r.At(1):	t.Fatalf(SHOULD_MATCH, 6, 1, s.At(6), r.At(1))
	case s.At(7) != r.At(2):	t.Fatalf(SHOULD_MATCH, 7, 2, s.At(7), r.At(2))
	case s.At(8) != r.At(3):	t.Fatalf(SHOULD_MATCH, 8, 3, s.At(8), r.At(3))
	case s.At(9) != r.At(4):	t.Fatalf(SHOULD_MATCH, 9, 4, s.At(9), r.At(4))
	}
}

func TestClear(t *testing.T) {
	HAS_VALUE := "b[%v] should be %v rather than %v"

	_, s := initSliceTest()
	Clear(s, 1, 3)
	switch {
	case s.At(0) != 0:			t.Fatalf(HAS_VALUE, 0, 0, s.At(0))
	case s.At(1) != 0:			t.Fatalf(HAS_VALUE, 1, 0, s.At(1))
	case s.At(2) != 0:			t.Fatalf(HAS_VALUE, 2, 0, s.At(2))
	case s.At(3) != 0:			t.Fatalf(HAS_VALUE, 3, 0, s.At(3))
	case s.At(4) != 4:			t.Fatalf(HAS_VALUE, 4, 4, s.At(4))
	}
}

func TestCopyElements(t *testing.T) {
	SHOULD_CONTAIN := "Test %v: Slice element c[%v] should contain %v but contains %v"

	_, s := initSliceTest()
	c := s.Clone().(*Slice)
	CopyElements(c, 1, 3, 5)
	switch {
	case c.At(0) != 0:			t.Fatalf(SHOULD_CONTAIN, 1, 0, 0, c.At(0))
	case c.At(1) != 3:			t.Fatalf(SHOULD_CONTAIN, 1, 1, 3, c.At(1))
	case c.At(2) != 4:			t.Fatalf(SHOULD_CONTAIN, 1, 2, 4, c.At(2))
	case c.At(3) != 5:			t.Fatalf(SHOULD_CONTAIN, 1, 3, 5, c.At(3))
	case c.At(4) != 6:			t.Fatalf(SHOULD_CONTAIN, 1, 4, 6, c.At(4))
	case c.At(5) != 7:			t.Fatalf(SHOULD_CONTAIN, 1, 5, 7, c.At(5))
	case c.At(6) != 6:			t.Fatalf(SHOULD_CONTAIN, 1, 6, 6, c.At(6))
	case c.At(7) != 7:			t.Fatalf(SHOULD_CONTAIN, 1, 7, 7, c.At(7))
	case c.At(8) != 8:			t.Fatalf(SHOULD_CONTAIN, 1, 8, 8, c.At(8))
	case c.At(9) != 9:			t.Fatalf(SHOULD_CONTAIN, 1, 9, 9, c.At(9))
	}

	c = s.Clone().(*Slice)
	CopyElements(c, 3, 1, 5)
	switch {
	case c.At(0) != 0:			t.Fatalf(SHOULD_CONTAIN, 2, 0, 0, c.At(0))
	case c.At(1) != 1:			t.Fatalf(SHOULD_CONTAIN, 2, 1, 1, c.At(1))
	case c.At(2) != 2:			t.Fatalf(SHOULD_CONTAIN, 2, 2, 2, c.At(2))
	case c.At(3) != 1:			t.Fatalf(SHOULD_CONTAIN, 2, 3, 1, c.At(3))
	case c.At(4) != 2:			t.Fatalf(SHOULD_CONTAIN, 2, 4, 2, c.At(4))
	case c.At(5) != 3:			t.Fatalf(SHOULD_CONTAIN, 2, 5, 3, c.At(5))
	case c.At(6) != 4:			t.Fatalf(SHOULD_CONTAIN, 2, 6, 4, c.At(6))
	case c.At(7) != 5:			t.Fatalf(SHOULD_CONTAIN, 2, 7, 5, c.At(7))
	case c.At(8) != 8:			t.Fatalf(SHOULD_CONTAIN, 2, 8, 8, c.At(8))
	case c.At(9) != 9:			t.Fatalf(SHOULD_CONTAIN, 2, 9, 9, c.At(9))
	}
}

func TestCycle(t *testing.T) {
	_, s := initSliceTest()
	r := Cycle(s, 5, func(x interface{}) {})
	switch {
	case r == nil:				t.Fatal("Cycle() returned a nil value")
	case r != 5:				t.Fatalf("cycle count should be 5 but is %v", r)
	}
}

func TestCombine(t *testing.T) {
	INCORRECT_VALUE := "r[%v] == %v"

	b, s := initSliceTest()
	r := Combine(s, s, func(x, y interface{}) interface{} {
		return x.(int) * y.(int)
	})
	switch {
	case r == nil:				t.Fatal("Combine() returned a nil value")
	case r.Len() != len(b):		t.Fatalf("capacity should be %v but is %v", len(b), r.Len())
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
}