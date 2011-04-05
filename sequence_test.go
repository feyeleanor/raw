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

func TestQueue(t *testing.T) {
	SHOULD_MATCH := "Slice elements s[%v] and r[%v] should match but are %v and %v"
	HAS_VALUE := "Value should be %v but is %v"
	SHOULD_HAVE_LENGTH := "Slice length should be %v not %v"

	_, s := initSliceTest()
	x, r := PopFirst(s)
	switch {
	case x != 0:				t.Fatalf(HAS_VALUE, 0, x)
	case r.Len() != 9:			t.Fatalf(SHOULD_HAVE_LENGTH, 9, r.Len())
	case s.At(1) != r.At(0):	t.Fatalf(SHOULD_MATCH, 1, 0, s.At(1), r.At(0))
	case s.At(8) != r.At(7):	t.Fatalf(SHOULD_MATCH, 8, 7, s.At(8), r.At(7))
	}
	r.(Deque).Append(20)
	switch {
	case r.Len() != 10:			t.Fatalf(SHOULD_HAVE_LENGTH, 9, r.Len())
	case s.At(1) != r.At(0):	t.Fatalf(SHOULD_MATCH, 1, 0, s.At(1), r.At(0))
	case r.At(9) != 20:			t.Fatalf(HAS_VALUE, 20, r.At(9))
	}
	x, r = PopFirst(r)
	switch {
	case x != 1:				t.Fatalf(HAS_VALUE, 1, x)
	case r.Len() != 9:			t.Fatalf(SHOULD_HAVE_LENGTH, 8, r.Len())
	case s.At(2) != r.At(0):	t.Fatalf(SHOULD_MATCH, 2, 0, s.At(2), r.At(0))
	case r.At(8) != 20:			t.Fatalf(HAS_VALUE, 20, r.At(8))
	}
}

func TestStack(t *testing.T) {
	SHOULD_MATCH := "Slice elements s[%v] and r[%v] should match but are %v and %v"
	HAS_VALUE := "Value should be %v but is %v"
	SHOULD_HAVE_LENGTH := "Slice length should be %v not %v"

	_, s := initSliceTest()
	x, r := PopLast(s)
	switch {
	case x != 9:				t.Fatalf(HAS_VALUE, 9, x)
	case r.Len() != 9:			t.Fatalf(SHOULD_HAVE_LENGTH, 9, r.Len())
	case s.At(0) != r.At(0):	t.Fatalf(SHOULD_MATCH, 0, 0, s.At(0), r.At(0))
	case s.At(8) != r.At(8):	t.Fatalf(SHOULD_MATCH, 8, 8, s.At(8), r.At(8))
	}
	x, r = PopLast(r)
	switch {
	case x != 8:				t.Fatalf(HAS_VALUE, 8, x)
	case r.Len() != 8:			t.Fatalf(SHOULD_HAVE_LENGTH, 8, r.Len())
	case s.At(0) != r.At(0):	t.Fatalf(SHOULD_MATCH, 0, 0, s.At(0), r.At(0))
	case s.At(7) != r.At(7):	t.Fatalf(SHOULD_MATCH, 7, 7, s.At(7), r.At(7))
	}
	r.(Deque).Append(20)
	switch {
	case r.Len() != 9:			t.Fatalf(SHOULD_HAVE_LENGTH, 9, r.Len())
	case s.At(0) != r.At(0):	t.Fatalf(SHOULD_MATCH, 0, 0, s.At(0), r.At(0))
	case r.At(8) != 20:			t.Fatalf(HAS_VALUE, 20, r.At(8))
	}
	x, r = PopLast(r)
	switch {
	case x != 20:				t.Fatalf(HAS_VALUE, 20, x)
	case r.Len() != 8:			t.Fatalf(SHOULD_HAVE_LENGTH, 8, r.Len())
	case s.At(0) != r.At(0):	t.Fatalf(SHOULD_MATCH, 0, 0, s.At(0), r.At(0))
	case s.At(7) != r.At(7):	t.Fatalf(SHOULD_MATCH, 7, 7, s.At(7), r.At(7))
	}
}

func TestDeque(t *testing.T) {
	SHOULD_MATCH := "Slice elements s[%v] and r[%v] should match but are %v and %v"
	HAS_VALUE := "Value should be %v but is %v"
	SHOULD_HAVE_LENGTH := "Slice length should be %v not %v"

	_, s := initSliceTest()
	x, r := PopLast(s)
	switch {
	case x != 9:				t.Fatalf(HAS_VALUE, 9, x)
	case r.Len() != 9:			t.Fatalf(SHOULD_HAVE_LENGTH, 9, r.Len())
	case s.At(0) != r.At(0):	t.Fatalf(SHOULD_MATCH, 0, 0, s.At(0), r.At(0))
	case s.At(8) != r.At(8):	t.Fatalf(SHOULD_MATCH, 8, 8, s.At(8), r.At(8))
	}
	r.(Deque).Append(10)
	switch {
	case r.Len() != 10:			t.Fatalf(SHOULD_HAVE_LENGTH, 10, r.Len())
	case s.At(8) != r.At(8):	t.Fatalf(SHOULD_MATCH, 8, 8, s.At(8), r.At(8))
	case r.At(9) != 10:			t.Fatalf(HAS_VALUE, 10, r.At(9))
	}
	x, r = PopFirst(r)
	switch {
	case x != 0:				t.Fatalf(HAS_VALUE, 0, x)
	case r.Len() != 9:			t.Fatalf(SHOULD_HAVE_LENGTH, 9, r.Len())
	case s.At(1) != r.At(0):	t.Fatalf(SHOULD_MATCH, 1, 0, s.At(1), r.At(0))
	case s.At(8) != r.At(7):	t.Fatalf(SHOULD_MATCH, 8, 7, s.At(8), r.At(7))
	}
	r.(Deque).Prepend(20)
	switch {
	case r.Len() != 10:			t.Fatalf(SHOULD_HAVE_LENGTH, 10, r.Len())
	case r.At(0) != 20:			t.Fatalf(HAS_VALUE, 20, r.At(0))
	case s.At(1) != r.At(1):	t.Fatalf(SHOULD_MATCH, 1, 1, s.At(1), r.At(1))
	}
	x, r = PopFirst(r)
	switch {
	case x != 20:				t.Fatalf(HAS_VALUE, 20, x)
	case r.Len() != 9:			t.Fatalf(SHOULD_HAVE_LENGTH, 9, r.Len())
	case s.At(1) != r.At(0):	t.Fatalf(SHOULD_MATCH, 1, 0, s.At(1), r.At(0))
	case s.At(8) != r.At(7):	t.Fatalf(SHOULD_MATCH, 8, 7, s.At(8), r.At(7))
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
	c := Copy(s).(Sequence)
	CopyElements(c, 1, 3, 5)
	if _, ok := c.(Blitter); !ok {
		t.Fatalf("c should be a Blitter")
	}
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

	c = Copy(s).(Sequence)
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