package raw

import "testing"

func TestEach(t *testing.T) {
	_, s := initSliceTest()
	c := 0
	n := Each(s, func(i interface{}) {
		c += i.(int)
	})
	switch {
	case n != 10:				t.Fatalf("Count should be 10 and not %v", n)
	case c != 45:				t.Fatalf("Sum should be 45 and not %v", c)
	}
}