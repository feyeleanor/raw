package raw

import "reflect"
import "testing"

func initChannelTest() (b chan int, c *Channel) {
	b = make(chan int, 16)
	for _, v := range []int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 } {
		b <- v
	}
	c = MakeChannel(b)
	return
}

func TestChannelMakeChannel(t *testing.T) {
	SHOULD_RECEIVE := "Should receive %v but received %v"

	b, c := initChannelTest()
	switch {
	case c == nil:					t.Fatal("MakeChannel returned a nil value")
	case c.Len() != len(b):			t.Fatalf("Channel length should be %v not %v", len(b), c.Len())
	case c.Cap() != cap(b):			t.Fatalf("Channel capacity should be %v not %v", cap(b), c.Cap())
	}

	for i := 0; i < 10; i++ {
		if v, _ := c.Recv(); v.Interface() != i {
			t.Fatalf(SHOULD_RECEIVE, i, v.Interface())
		}
	}

	if v, _ := c.TryRecv(); v != nil {
		t.Fatalf(SHOULD_RECEIVE, nil, v.Interface())
	}

	if v, _ := reflect.NewValue(b).(*reflect.ChanValue).TryRecv(); v != nil {
		t.Fatalf(SHOULD_RECEIVE, nil, v.Interface())
	}
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

func TestChannelWhile(t *testing.T) {
	b, c := initChannelTest()
	count := 0
	c.While(func(i interface{}) bool {
		count++
		return i.(int) < 6
	})
	switch {
	case count != 7:				t.Fatalf("Count should be %v not %v", 7, count)
	case c.Len() != 3:				t.Fatalf("Channel length should be %v not %v", 3, c.Len())
	case c.Len() != len(b):			t.Fatalf("Channel length should be %v not %v", len(b), c.Len())
	}
}

func TestChannelUntil(t *testing.T) {
	b, c := initChannelTest()
	count := 0
	c.Until(func(i interface{}) bool {
		count++
		return i.(int) == 6
	})
	switch {
	case count != 7:				t.Fatalf("Count should be %v not %v", 7, count)
	case c.Len() != 3:				t.Fatalf("Channel length should be %v not %v", 3, c.Len())
	case c.Len() != len(b):			t.Fatalf("Channel length should be %v not %v", len(b), c.Len())
	}
}