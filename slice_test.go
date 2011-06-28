package raw

import "testing"

const(
	NO_TESTS				string = "Tests yet to be implemented"
	FURTHER_TESTS_NEEDED	string = "Add further test filters"
)

func initSliceTest() (b []int, s *Slice) {
	b = []int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	s = NewSlice(b)
	return
}

func TestNewSlice(t *testing.T) {
	SHOULD_CONTAIN := "%v[%v] should contain %v but contains %v"

	s := NewSlice([]int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 })
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
}

func TestSliceAppend(t *testing.T) {
	SHOULD_MATCH := "Slice elements s[%v] and s[%v] should match but are %v and %v"
	HAS_VALUE := "s[%v] should be %v rather than %v"

	t.Log("Append b[] to s")
	b, s := initSliceTest()
	s.Append(b)
	switch {
	case s.Len() != len(b) * 2:	t.Fatalf("Slice length should be %v not %v", len(b) * 2, s.Len())
	case s.Cap() != cap(b) * 2:	t.Fatalf("Slice capacity should be %v not %v", cap(b) * 2, s.Cap())
	case s.At(0) != s.At(10):	t.Fatalf(SHOULD_MATCH, 0, 10, s.At(0), s.At(10))
	case s.At(1) != s.At(11):	t.Fatalf(SHOULD_MATCH, 1, 11, s.At(1), s.At(11))
	case s.At(2) != s.At(12):	t.Fatalf(SHOULD_MATCH, 2, 12, s.At(2), s.At(12))
	case s.At(3) != s.At(13):	t.Fatalf(SHOULD_MATCH, 3, 13, s.At(3), s.At(13))
	case s.At(4) != s.At(14):	t.Fatalf(SHOULD_MATCH, 4, 14, s.At(4), s.At(14))
	case s.At(5) != s.At(15):	t.Fatalf(SHOULD_MATCH, 5, 15, s.At(5), s.At(15))
	case s.At(6) != s.At(16):	t.Fatalf(SHOULD_MATCH, 6, 16, s.At(6), s.At(16))
	case s.At(7) != s.At(17):	t.Fatalf(SHOULD_MATCH, 7, 17, s.At(7), s.At(17))
	case s.At(8) != s.At(18):	t.Fatalf(SHOULD_MATCH, 8, 18, s.At(8), s.At(18))
	case s.At(9) != s.At(19):	t.Fatalf(SHOULD_MATCH, 9, 19, s.At(9), s.At(19))
	}

	t.Log("Append s to s")
	b, s = initSliceTest()
	s.Append(s)
	switch {
	case s.Len() != len(b) * 2:	t.Fatalf("Slice length should be %v not %v", len(b) * 2, s.Len())
	case s.Cap() != cap(b) * 2:	t.Fatalf("Slice capacity should be %v not %v", cap(b) * 2, s.Cap())
	case s.At(0) != s.At(10):	t.Fatalf(SHOULD_MATCH, 0, 10, s.At(0), s.At(10))
	case s.At(1) != s.At(11):	t.Fatalf(SHOULD_MATCH, 1, 11, s.At(1), s.At(11))
	case s.At(2) != s.At(12):	t.Fatalf(SHOULD_MATCH, 2, 12, s.At(2), s.At(12))
	case s.At(3) != s.At(13):	t.Fatalf(SHOULD_MATCH, 3, 13, s.At(3), s.At(13))
	case s.At(4) != s.At(14):	t.Fatalf(SHOULD_MATCH, 4, 14, s.At(4), s.At(14))
	case s.At(5) != s.At(15):	t.Fatalf(SHOULD_MATCH, 5, 15, s.At(5), s.At(15))
	case s.At(6) != s.At(16):	t.Fatalf(SHOULD_MATCH, 6, 16, s.At(6), s.At(16))
	case s.At(7) != s.At(17):	t.Fatalf(SHOULD_MATCH, 7, 17, s.At(7), s.At(17))
	case s.At(8) != s.At(18):	t.Fatalf(SHOULD_MATCH, 8, 18, s.At(8), s.At(18))
	case s.At(9) != s.At(19):	t.Fatalf(SHOULD_MATCH, 9, 19, s.At(9), s.At(19))
	}

	t.Log("Append *s to s")
	b, s = initSliceTest()
	s.Append(*s)
	switch {
	case s.Len() != len(b) * 2:	t.Fatalf("Slice length should be %v not %v", len(b) * 2, s.Len())
	case s.Cap() != cap(b) * 2:	t.Fatalf("Slice capacity should be %v not %v", cap(b) * 2, s.Cap())
	case s.At(0) != s.At(10):	t.Fatalf(SHOULD_MATCH, 0, 10, s.At(0), s.At(10))
	case s.At(1) != s.At(11):	t.Fatalf(SHOULD_MATCH, 1, 11, s.At(1), s.At(11))
	case s.At(2) != s.At(12):	t.Fatalf(SHOULD_MATCH, 2, 12, s.At(2), s.At(12))
	case s.At(3) != s.At(13):	t.Fatalf(SHOULD_MATCH, 3, 13, s.At(3), s.At(13))
	case s.At(4) != s.At(14):	t.Fatalf(SHOULD_MATCH, 4, 14, s.At(4), s.At(14))
	case s.At(5) != s.At(15):	t.Fatalf(SHOULD_MATCH, 5, 15, s.At(5), s.At(15))
	case s.At(6) != s.At(16):	t.Fatalf(SHOULD_MATCH, 6, 16, s.At(6), s.At(16))
	case s.At(7) != s.At(17):	t.Fatalf(SHOULD_MATCH, 7, 17, s.At(7), s.At(17))
	case s.At(8) != s.At(18):	t.Fatalf(SHOULD_MATCH, 8, 18, s.At(8), s.At(18))
	case s.At(9) != s.At(19):	t.Fatalf(SHOULD_MATCH, 9, 19, s.At(9), s.At(19))
	}

	t.Log("Append 100 to s")
	b, s = initSliceTest()
	s.Append(100)
	switch {
	case s.Cap() <= cap(b):		t.Fatalf("Slice capacity should be greater than %v", cap(b))
	case s.Len() != len(b) + 1:	t.Fatalf("Slice length should be %v not %v", len(b) + 1, s.Len())
	case s.At(0) != 0:			t.Fatalf(HAS_VALUE, 0, 0, s.At(0))
	case s.At(1) != 1:			t.Fatalf(HAS_VALUE, 1, 1, s.At(1))
	case s.At(2) != 2:			t.Fatalf(HAS_VALUE, 2, 2, s.At(2))
	case s.At(3) != 3:			t.Fatalf(HAS_VALUE, 3, 3, s.At(3))
	case s.At(4) != 4:			t.Fatalf(HAS_VALUE, 4, 4, s.At(4))
	case s.At(5) != 5:			t.Fatalf(HAS_VALUE, 5, 5, s.At(5))
	case s.At(6) != 6:			t.Fatalf(HAS_VALUE, 6, 6, s.At(6))
	case s.At(7) != 7:			t.Fatalf(HAS_VALUE, 7, 7, s.At(7))
	case s.At(8) != 8:			t.Fatalf(HAS_VALUE, 8, 8, s.At(8))
	case s.At(9) != 9:			t.Fatalf(HAS_VALUE, 9, 9, s.At(9))
	case s.At(10) != 100:		t.Fatalf(HAS_VALUE, 10, 100, s.At(10))
	}
}

func TestSliceRepeat(t *testing.T) {
	SHOULD_MATCH := "Slice elements s[%v] and s[%v] should match but are %v and %v"

	b, s := initSliceTest()
	c := 3
	s = s.Repeat(c)
	switch {
	case s.Len() != len(b) * c:	t.Fatalf("Slice length should be %v not %v", len(b) * c, s.Len())
	case s.Cap() != cap(b) * c:	t.Fatalf("Slice capacity should be %v not %v", cap(b) * c, s.Cap())
	case s.At(0) != s.At(10):	t.Fatalf(SHOULD_MATCH, 0, 10, s.At(0), s.At(10))
	case s.At(1) != s.At(11):	t.Fatalf(SHOULD_MATCH, 1, 11, s.At(1), s.At(11))
	case s.At(9) != s.At(19):	t.Fatalf(SHOULD_MATCH, 9, 19, s.At(9), s.At(19))
	case s.At(0) != s.At(20):	t.Fatalf(SHOULD_MATCH, 0, 20, s.At(0), s.At(20))
	case s.At(1) != s.At(21):	t.Fatalf(SHOULD_MATCH, 1, 21, s.At(1), s.At(21))
	case s.At(9) != s.At(29):	t.Fatalf(SHOULD_MATCH, 9, 19, s.At(9), s.At(29))
	}
}