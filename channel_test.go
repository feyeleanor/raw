package raw

import "reflect"
import "testing"

func initChannelTest() (b chan int, c *Channel) {
	b = make(chan int, 16)
	go func() {
		for _, v := range []int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 } {
			b <- v
		}
		close(b)
	}()
	c = NewChannel(b)
	return
}

func TestChannelElementType(t *testing.T) {
	_, c := initChannelTest()
	e := reflect.Typeof(int(0))
	switch {
	case e != c.ElementType():		t.Fatalf("Channel claims element type %v when should be %v", c.ElementType(), e)
	}
}

func TestChannelDirection(t *testing.T) {
	b, c := initChannelTest()
	d := reflect.NewValue(b).(*reflect.ChanValue).Type().(*reflect.ChanType).Dir()
	switch {
	case d != c.Direction():		t.Fatalf("Channel claims element type %v when should be %v", c.Direction(), d)
	case !c.Receiver():				t.Fatalf("Channel should report as a receiver not %v", c.Direction())
	case !c.Bidirectional():		t.Fatalf("Channel should report as bidirectional not %v", c.Direction())
	case !c.Sender():				t.Fatalf("Channel should report as a sender not %v", c.Direction())
	}
}

func TestChannelEach(t *testing.T) {
	_, c := initChannelTest()
	sum := 0
	count := c.Each(func(i interface{}) {
		sum += i.(int)
	})
	switch {
	case count != 10:				t.Fatalf("Item count should be 10 and not %v", count)
	case sum != 45:					t.Fatalf("Sum should be 45 and not %v", sum)
	}
}

func TestChannelFirst(t *testing.T) {
	b, c := initChannelTest()
	s := []int{}
	c.First(5, func(i interface{}) {
		s = append(s, i.(int))
	})
	switch {
	case len(s) != 5:				t.Fatalf("Slice length should be %v not %v", 5, len(s))
	case c.Len() != len(b):			t.Fatalf("Channel length should be %v not %v", len(b), c.Len())
	}
}

/*
func TestChannelFeed(t *testing.T) {
	_, c := initChannelTest()
	o := make(chan interface{})
	i := 0
	c.Feed(o, func(x interface{}) (r interface{}) {
		r = i * x.(int)
		i++
		return
	})
	n := []int{}
	MakeChannel(o).First(10, func(x interface{}) {
		n = append(n, x.(int))
	})
	close(o)
	switch {
	case n[0] != 0:					t.Fatalf("%v: expected %v but got %v", 0, 0, n[0])
	case n[1] != 1:					t.Fatalf("%v: expected %v but got %v", 1, 1, n[1])
	case n[2] != 4:					t.Fatalf("%v: expected %v but got %v", 2, 4, n[2])
	case n[3] != 9:					t.Fatalf("%v: expected %v but got %v", 3, 9, n[3])
	case n[4] != 16:				t.Fatalf("%v: expected %v but got %v", 4, 16, n[4])
	case n[5] != 25:				t.Fatalf("%v: expected %v but got %v", 5, 25, n[5])
	case n[6] != 36:				t.Fatalf("%v: expected %v but got %v", 6, 36, n[6])
	case n[7] != 49:				t.Fatalf("%v: expected %v but got %v", 7, 49, n[7])
	case n[8] != 64:				t.Fatalf("%v: expected %v but got %v", 8, 64, n[8])
	case n[9] != 81:				t.Fatalf("%v: expected %v but got %v", 9, 81, n[9])
	}
}

func TestChannelPipe(t *testing.T) {
	_, c := initChannelTest()
	i := 0
	o := c.Pipe(func(x interface{}) (r interface{}) {
		r = i * x.(int)
		i++
		return 
	})
	n := []int{}
	for x := range o {
		n = append(n, x.(int))
	}
	switch {
	case n[0] != 0:					t.Fatalf("%v: expected %v but got %v", 0, 0, n[0])
	case n[1] != 1:					t.Fatalf("%v: expected %v but got %v", 1, 1, n[1])
	case n[2] != 4:					t.Fatalf("%v: expected %v but got %v", 2, 4, n[2])
	case n[3] != 9:					t.Fatalf("%v: expected %v but got %v", 3, 9, n[3])
	case n[4] != 16:				t.Fatalf("%v: expected %v but got %v", 4, 16, n[4])
	case n[5] != 25:				t.Fatalf("%v: expected %v but got %v", 5, 25, n[5])
	case n[6] != 36:				t.Fatalf("%v: expected %v but got %v", 6, 36, n[6])
	case n[7] != 49:				t.Fatalf("%v: expected %v but got %v", 7, 49, n[7])
	case n[8] != 64:				t.Fatalf("%v: expected %v but got %v", 8, 64, n[8])
	case n[9] != 81:				t.Fatalf("%v: expected %v but got %v", 9, 81, n[9])
	}
}
*/