package raw

import "reflect"
import "testing"

const(
	NO_TESTS				string = "Tests yet to be implemented"
	FURTHER_TESTS_NEEDED	string = "Add further test filters"
)

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

func TestNewChannel(t *testing.T) {
	b := make(chan int, 16)
	go func() {
		for _, v := range []int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 } {
			b <- v
		}
		close(b)
	}()
	c := NewChannel(b)

	switch {
	case c == nil:					t.Fatal("MakeChannel returned a nil value")
	case c.Len() != len(b):			t.Fatalf("Channel length should be %v not %v", len(b), c.Len())
	case c.Cap() != cap(b):			t.Fatalf("Channel capacity should be %v not %v", cap(b), c.Cap())
	}

	for i := 0; i < 10; i++ {
		switch v, open := c.Recv(); {
		case !open:					t.Fatalf("%v: channel should be open", i)
		case v != i:				t.Fatalf("Should receive %v but received %v", i, v)
		}
	}

	if _, open := c.TryRecv(); open {
		t.Fatal("Channel should be closed")
	}
}

func TestChannelDirection(t *testing.T) {
	b, c := initChannelTest()
	d := reflect.ValueOf(b).Type().ChanDir()
	switch {
	case c.Direction() != d:		t.Fatalf("Channel claims element type %v when should be %v", c.Direction(), d)
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