package raw

import "testing"

func TestCopySequence(t *testing.T) {
	SHOULD_MATCH := "Slice elements s[%v] and c[%v] should match but are %v and %v"

	_, s := initSliceTest()
	c := CopySequence(s)
	c.Set(0, 1000)
	switch {
	case c.Len() != s.Len():	t.Fatalf("Slice length should be %v not %v", s.Len(), c.Len())
	case c.Cap() != s.Cap():	t.Fatalf("Slice capacity should be %v not %v", s.Cap(), c.Cap())
	case s.At(0) != 0:			t.Fatalf("Slice element s[%v] should be %v and not", 0, 0, s.At(0))
	case c.At(0) != 1000:		t.Fatalf("Slice element c[%v] should be %v and not", 0, 1000, c.At(0))
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
}

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
	c := CopySequence(s)
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

	c = CopySequence(s)
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