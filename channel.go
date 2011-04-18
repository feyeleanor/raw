package raw

import "reflect"

func WaitFor(f func()) {
	done := make(chan bool)
	go func() {
		f()
		done <- true
	}()
	<-done
}

type Channel struct {
	reflect.Value
}

func NewChannel(i interface{}) (c *Channel) {
	if v := reflect.NewValue(i); v.Kind() == reflect.Chan {
		c = &Channel{ v }
	} else {
		c = NewChannel(v.Elem())
	}
	return
}

// Returns the runtime type of the elements travelling along the Channel.
func (c *Channel) ElementType() reflect.Type {
	return c.Type().Elem()
}

func (c *Channel) Direction() reflect.ChanDir {
	return c.Type().ChanDir()
}

func (c *Channel) Receiver() bool {
	return (c.Direction() & reflect.RecvDir) == reflect.RecvDir
}

func (c *Channel) Sender() bool {
	return (c.Direction() & reflect.SendDir) == reflect.SendDir
}

func (c *Channel) Bidirectional() bool {
	return c.Direction() == reflect.BothDir
}

func (c *Channel) Recv() (x interface{}, open bool) {
	v, open := c.Value.Recv()
	x = v.Interface()
	return
}

func (c *Channel) TryRecv() (x interface{}, open bool) {
	v, open := c.Value.TryRecv()
	x = v.Interface()
	return
}

func (c *Channel) Send(x interface{}) {
	c.Value.Send(reflect.NewValue(x))
}

func (c *Channel) TrySend(x interface{}) {
	c.Value.TrySend(reflect.NewValue(x))
}

func (c *Channel) Each(f func(x interface{})) (n int) {
	for {
		if v, open := c.Recv(); open {
			f(v)
			n++
		} else {
			return
		}
	}
	return
}

//	First reads a specified number of values into a function, terminating if the Channel is closed
func (c *Channel) First(i int, f func(x interface{})) {
	for ; i > 0; i-- {
		v, open := c.Recv()
		f(v)
		if !open {
			break
		}
	}
}

func (c *Channel) Feed(o chan<- interface{}, f func(x interface{}) interface{}) {
	for {
		if v, open := c.Recv(); open {
			o <- f(v)
		} else {
			return
		}
	}
}

func (c *Channel) Pipe(f func(x interface{}) interface{}) (o chan interface{}) {
	o = make(chan interface{}, StandardChannelBuffer)
	c.Feed(o, f)
	return
}