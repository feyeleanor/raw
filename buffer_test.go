package raw

import "testing"

func TestExtend(t *testing.T) {
	b, s := initSliceTest()
	Extend(s, 10)
	switch {
		case b == nil:				t.Fatal("Resize() created a nil value for original slice")
		case s == nil:				t.Fatal("Resize() created a nil value for Slice")
		case cap(b) != 10:			t.Fatalf("original slice capacity should be 10 but is %v", cap(b))
		case len(b) != 10:			t.Fatalf("original slice length should be 10 but is %v", len(b))
		case s.Cap() != 20:			t.Fatalf("Slice capacity should be 20 but is %v", s.Cap())
		case s.Len() != 10:			t.Fatalf("Slice length should be 10 but is %v", s.Len())
	}

	Extend(s, -15)
	switch {
	case b == nil:				t.Fatal("Resize() created a nil value for original slice")
	case s == nil:				t.Fatal("Resize() created a nil value for Slice")
	case cap(b) != 10:			t.Fatalf("original slice capacity should be 10 but is %v", cap(b))
	case len(b) != 10:			t.Fatalf("original slice length should be 10 but is %v", len(b))
	case s.Cap() != 5:			t.Fatalf("Slice capacity should be 5 but is %v", s.Cap())
	case s.Len() != 5:			t.Fatalf("Slice length should be 5 but is %v", s.Len())
	}
}

func TestShrink(t *testing.T) {
	b, s := initSliceTest()
	Shrink(s, -10)
	switch {
		case b == nil:				t.Fatal("Resize() created a nil value for original slice")
		case s == nil:				t.Fatal("Resize() created a nil value for Slice")
		case cap(b) != 10:			t.Fatalf("original slice capacity should be 10 but is %v", cap(b))
		case len(b) != 10:			t.Fatalf("original slice length should be 10 but is %v", len(b))
		case s.Cap() != 20:			t.Fatalf("Slice capacity should be 20 but is %v", s.Cap())
		case s.Len() != 10:			t.Fatalf("Slice length should be 10 but is %v", s.Len())
	}

	Shrink(s, 15)
	switch {
	case b == nil:				t.Fatal("Resize() created a nil value for original slice")
	case s == nil:				t.Fatal("Resize() created a nil value for Slice")
	case cap(b) != 10:			t.Fatalf("original slice capacity should be 10 but is %v", cap(b))
	case len(b) != 10:			t.Fatalf("original slice length should be 10 but is %v", len(b))
	case s.Cap() != 5:			t.Fatalf("Slice capacity should be 5 but is %v", s.Cap())
	case s.Len() != 5:			t.Fatalf("Slice length should be 5 but is %v", s.Len())
	}
}

func TestDoubleCapacity(t *testing.T) {
	b, s := initSliceTest()
	DoubleCapacity(s)
	switch {
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
	HalveCapacity(s)
	switch {
		case b == nil:				t.Fatal("Resize() created a nil value for original slice")
		case s == nil:				t.Fatal("Resize() created a nil value for Slice")
		case cap(b) != 10:			t.Fatalf("original slice capacity should be 10 but is %v", cap(b))
		case len(b) != 10:			t.Fatalf("original slice length should be 10 but is %v", len(b))
		case s.Cap() != 5:			t.Fatalf("Slice capacity should be 5 but is %v", s.Cap())
		case s.Len() != 5:			t.Fatalf("Slice length should be 5 but is %v", s.Len())
	}
}
