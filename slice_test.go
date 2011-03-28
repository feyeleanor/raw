package raw

import "reflect"
import "testing"

const(
	NO_TESTS				string = "Tests yet to be implemented"
	FURTHER_TESTS_NEEDED	string = "Add further test filters"
)

func initSliceTest() (b []int, s *Slice) {
	b = []int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	s = MakeSlice(b)
	return
}

func TestSliceMakeSlice(t *testing.T) {
	SHOULD_MATCH := "Slice elements s[%v] and b[%v] should match but are %v and %v"

	b, s := initSliceTest()
	switch {
	case s == nil:				t.Fatal("MakeSlice returned a nil value")
	case s.Len() != len(b):		t.Fatalf("Slice length should be %v not %v", len(b), s.Len())
	case s.Cap() != cap(b):		t.Fatalf("Slice capacity should be %v not %v", cap(b), s.Cap())
	case s.At(0) != b[0]:		t.Fatalf(SHOULD_MATCH, 0, 0, s.At(0), b[0])
	case s.At(1) != b[1]:		t.Fatalf(SHOULD_MATCH, 1, 1, s.At(1), b[1])
	case s.At(2) != b[2]:		t.Fatalf(SHOULD_MATCH, 2, 2, s.At(2), b[2])
	case s.At(3) != b[3]:		t.Fatalf(SHOULD_MATCH, 3, 3, s.At(3), b[3])
	case s.At(4) != b[4]:		t.Fatalf(SHOULD_MATCH, 4, 4, s.At(4), b[4])
	case s.At(5) != b[5]:		t.Fatalf(SHOULD_MATCH, 5, 5, s.At(5), b[5])
	case s.At(6) != b[6]:		t.Fatalf(SHOULD_MATCH, 6, 6, s.At(6), b[6])
	case s.At(7) != b[7]:		t.Fatalf(SHOULD_MATCH, 7, 7, s.At(7), b[7])
	case s.At(8) != b[8]:		t.Fatalf(SHOULD_MATCH, 8, 8, s.At(8), b[8])
	case s.At(9) != b[9]:		t.Fatalf(SHOULD_MATCH, 9, 9, s.At(9), b[9])
	}
}

func TestSliceClone(t *testing.T) {
	SHOULD_MATCH := "Slice elements s[%v] and c[%v] should match but are %v and %v"

	_, s := initSliceTest()
	c := s.Clone()
	switch {
	case c.Len() != s.Len():	t.Fatalf("Slice length should be %v not %v", s.Len(), c.Len())
	case c.Cap() != s.Cap():	t.Fatalf("Slice capacity should be %v not %v", s.Cap(), c.Cap())
	case c.At(0) != s.At(0):	t.Fatalf(SHOULD_MATCH, 0, 0, s.At(0), c.At(0))
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
	case s.Cap() == cap(b):		t.Fatalf("Slice capacity should be greater than %v", cap(b))
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

func TestSliceElementType(t *testing.T) {
	_, s := initSliceTest()
	e := reflect.Typeof(int(0))
	switch {
	case e != s.ElementType():	t.Fatalf("Slice claims element type %v when should be %v", s.ElementType(), e)
	}
}

func TestSliceCopy(t *testing.T) {
	b, s := initSliceTest()
	s.Copy(1, 3)
	switch {
	case b[1] != b[3]:			t.Fatalf("Elements b[1] and b[3] should match but are %v and %v", b[1], b[3])
	case b[3] != 3:				t.Fatalf("Element b[3] should be %v but is %v", 3, b[3])
	case s.At(1) != s.At(3):	t.Fatalf("Slice elements s[1] and s[3] should match but are %v and %v", s.At(1), s.At(3))
	}
}

func TestSliceCopySlice(t *testing.T) {
	SHOULD_MATCH := "Slice elements c[%V] and s[%V] should match but are %v and %v"
	SHOULD_NOT_MATCH := "Slice elements s[%v] and s[%v] should not match but are both %v"

	b, s := initSliceTest()
	c := s.Clone()
	c.CopySlice(1, 3, 5)
	switch {
	case b[1] == b[3]:			t.Fatalf("Elements b[1] and b[3] should not match but are %v and %v", b[1], b[3])
	case s.At(1) == s.At(3):	t.Fatalf(SHOULD_NOT_MATCH, 1, 3, s.At(1))
	case s.At(2) == s.At(4):	t.Fatalf(SHOULD_NOT_MATCH, 2, 4, s.At(2))
	case s.At(3) == s.At(5):	t.Fatalf(SHOULD_NOT_MATCH, 3, 5, s.At(3))
	case s.At(4) == s.At(6):	t.Fatalf(SHOULD_NOT_MATCH, 4, 6, s.At(4))
	case s.At(5) == s.At(7):	t.Fatalf(SHOULD_NOT_MATCH, 5, 7, s.At(5))
	case c.At(1) != s.At(3):	t.Fatalf(SHOULD_MATCH, 1, 3, c.At(1), s.At(3))
	case c.At(2) != s.At(4):	t.Fatalf(SHOULD_MATCH, 2, 4, c.At(2), s.At(4))
	case c.At(3) != s.At(5):	t.Fatalf(SHOULD_MATCH, 3, 5, c.At(3), s.At(5))
	case c.At(4) != s.At(6):	t.Fatalf(SHOULD_MATCH, 4, 6, c.At(4), s.At(6))
	case c.At(5) != s.At(7):	t.Fatalf(SHOULD_MATCH, 5, 7, c.At(5), s.At(7))
	case c.At(6) != s.At(6):	t.Fatalf(SHOULD_MATCH, 6, 6, c.At(6), s.At(6))
	}
}

func TestSliceSwap(t *testing.T) {
	HAS_VALUE := "b[%v] should be %v rather than %v"

	b, s := initSliceTest()
	s.Swap(1, 3)
	switch {
	case b[0] != 0:				t.Fatalf(HAS_VALUE, 0, 0, b[0])
	case b[1] != 3:				t.Fatalf(HAS_VALUE, 1, 3, b[1])
	case b[2] != 2:				t.Fatalf(HAS_VALUE, 2, 2, b[2])
	case b[3] != 1:				t.Fatalf(HAS_VALUE, 3, 1, b[3])
	case b[4] != 4:				t.Fatalf(HAS_VALUE, 4, 4, b[4])
	}
}

func TestSliceClear(t *testing.T) {
	HAS_VALUE := "b[%v] should be %v rather than %v"

	b, s := initSliceTest()
	s.Clear(1, 3)
	switch {
	case b[0] != 0:				t.Fatalf(HAS_VALUE, 0, 0, b[0])
	case b[1] != 0:				t.Fatalf(HAS_VALUE, 1, 0, b[1])
	case b[2] != 0:				t.Fatalf(HAS_VALUE, 2, 0, b[2])
	case b[3] != 0:				t.Fatalf(HAS_VALUE, 3, 0, b[3])
	case b[4] != 4:				t.Fatalf(HAS_VALUE, 4, 4, b[4])
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

func TestSliceEach(t *testing.T) {
	_, s := initSliceTest()
	c := 0
	n := s.Each(func(i interface{}) {
		c += i.(int)
	})
	switch {
	case n != 10:				t.Fatalf("Count should be 10 and not %v", n)
	case c != 45:				t.Fatalf("Sum should be 45 and not %v", c)
	}
}

func TestSliceFirst(t *testing.T) {
	SHOULD_MATCH := "Slice elements s[%v] and r[%v] should match but are %v and %v"

	_, s := initSliceTest()
	r := []int{}
	s.First(5, func(i interface{}) {
		r = append(r, i.(int))
	})
	switch {
	case len(r) != 5:			t.Fatalf("Slice length should be %v not %v", 5, len(r))
	case s.At(0) != r[0]:		t.Fatalf(SHOULD_MATCH, 0, 0, s.At(0), r[0])
	case s.At(1) != r[1]:		t.Fatalf(SHOULD_MATCH, 1, 1, s.At(1), r[1])
	case s.At(2) != r[2]:		t.Fatalf(SHOULD_MATCH, 2, 2, s.At(2), r[2])
	case s.At(3) != r[3]:		t.Fatalf(SHOULD_MATCH, 3, 3, s.At(3), r[3])
	case s.At(4) != r[4]:		t.Fatalf(SHOULD_MATCH, 4, 4, s.At(4), r[4])
	}
}

func TestSliceLast(t *testing.T) {
	SHOULD_MATCH := "Slice elements s[%v] and r[%v] should match but are %v and %v"

	_, s := initSliceTest()
	r := []int{}
	s.Last(5, func(i interface{}) {
		r = append(r, i.(int))
	})
	switch {
	case len(r) != 5:			t.Fatalf("Slice length should be %v not %v", 5, len(r))
	case s.At(9) != r[0]:		t.Fatalf(SHOULD_MATCH, 9, 0, s.At(9), r[0])
	case s.At(8) != r[1]:		t.Fatalf(SHOULD_MATCH, 8, 1, s.At(8), r[1])
	case s.At(7) != r[2]:		t.Fatalf(SHOULD_MATCH, 7, 2, s.At(7), r[2])
	case s.At(6) != r[3]:		t.Fatalf(SHOULD_MATCH, 6, 3, s.At(6), r[3])
	case s.At(5) != r[4]:		t.Fatalf(SHOULD_MATCH, 5, 4, s.At(5), r[4])
	}
}

func TestSliceCollect(t *testing.T) {
	INCORRECT_VALUE := "r[%v] == %v"

	b, s := initSliceTest()
	r := s.Collect(func(i interface{}) interface{} {
		return i.(int) * 2
	})
	switch {
	case r == nil:				t.Fatal("Collect() returned a nil value")
	case r.Cap() != cap(b):		t.Fatalf("capacity should be %v but is %v", cap(b), r.Cap())
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

func TestSliceInject(t *testing.T) {
	_, s := initSliceTest()
	r := s.Inject(0, func(memo, x interface{}) interface{} {
		return memo.(int) + x.(int)
	})
	switch {
	case r == nil:				t.Fatal("Inject() returned a nil value")
	case r != 45:				t.Fatalf("r should be 45 but is %v", r)
	}
}

func TestSliceCombine(t *testing.T) {
	INCORRECT_VALUE := "r[%v] == %v"

	b, s := initSliceTest()
	r := s.Combine(s, func(x, y interface{}) interface{} {
		return x.(int) * y.(int)
	})
	switch {
	case r == nil:				t.Fatal("Combine() returned a nil value")
	case r.Cap() != cap(b):		t.Fatalf("capacity should be %v but is %v", cap(b), r.Cap())
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

func TestSliceCycle(t *testing.T) {
	_, s := initSliceTest()
	r := s.Cycle(5, func(x interface{}) {})
	switch {
	case r == nil:				t.Fatal("Cycle() returned a nil value")
	case r != 5:				t.Fatalf("cycle count should be 5 but is %v", r)
	}
}

func TestSliceResize(t *testing.T) {
	b, s := initSliceTest()
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
	b, s := initSliceTest()
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
	b, s := initSliceTest()
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
	b, s := initSliceTest()
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
	b, s := initSliceTest()
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
	b, s := initSliceTest()
	c := make(chan interface{})
	i := 0
	s.Feed(c, func(x interface{}) (r interface{}) {
		r = i * x.(int)
		i++
		return
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
	b, s := initSliceTest()
	i := 0
	c := s.Pipe(func(x interface{}) (r interface{}) {
		r = i * x.(int)
		i++
		return 
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

func TestSliceTee(t *testing.T) {
	t.Fatal(NO_TESTS)
}