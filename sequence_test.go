package raw

import "testing"

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