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

// Creates a Channel from a given object, raising a runtime panic if the object cannot be represented as a *reflect.ChanValue.
func MakeChannel(i interface{}) (c *Channel) {
	switch v := reflect.NewValue(i).(type) {
	case *reflect.ChanValue:		c = &Channel{ v }
	case *reflect.InterfaceValue:	c = MakeChannel(v.Elem())
	case *reflect.PtrValue:			c = MakeChannel(v.Elem())
	default:						panic(i)
	}
	return
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

func (c *Channel) Each(f func(x interface{})) (n int) {
	for {
		if v, open := c.Recv(); open {
			f(v.Interface())
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
		f(v.Interface())
		if !open {
			break
		}
	}
}

func (c *Channel) Feed(o chan<- interface{}, f func(x interface{}) interface{}) {
//	go func() {
		for {
			if v, open := c.Recv(); open {
				o <- f(v.Interface())
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
				x := f(v.Interface())
				t <- x
				o <- x
			}
		}
	}()
	return
}