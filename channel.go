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
	*reflect.ChanValue
}

func NewChannel(i interface{}) *Channel {
	return NewContainer(i).(*Channel)
}

// Returns the runtime type of the elements travelling along the Channel.
func (c *Channel) ElementType() reflect.Type {
	return c.Type().(*reflect.ChanType).Elem()
}

func (c *Channel) Direction() reflect.ChanDir {
	return c.Type().(*reflect.ChanType).Dir()
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
	v, open := c.ChanValue.Recv()
	x = v.Interface()
	return
}

func (c *Channel) TryRecv() (x interface{}, open bool) {
	v, open := c.ChanValue.TryRecv()
	x = v.Interface()
	return
}

func (c *Channel) Send(x interface{}) {
	c.ChanValue.Send(reflect.NewValue(x))
}

func (c *Channel) TrySend(x interface{}) {
	c.ChanValue.TrySend(reflect.NewValue(x))
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
//	go func() {
		for {
			if v, open := c.Recv(); open {
				o <- f(v)
			} else {
				return
			}
		}
//	}()
}

func (c *Channel) Pipe(f func(x interface{}) interface{}) (o chan interface{}) {
	o = make(chan interface{}, StandardChannelBuffer)
	c.Feed(o, f)
	return
}

func (c *Channel) Tee(o chan<- interface{}, f func(x interface{}) interface{}) (t chan interface{}) {
	t = make(chan interface{}, StandardChannelBuffer)
	go func() {
		for  {
			if v, open := c.Recv(); open {
				x := f(v)
				t <- x
				o <- x
			}
		}
	}()
	return
}