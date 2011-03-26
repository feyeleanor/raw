package raw

import "reflect"

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

//	While reads values from a channel into a function whilst a condition is true or until the channel is closed
func (c *Channel) While(f func(x interface{}) bool) {
	for {
		if v, open := c.Recv(); open {
			if f(v.Interface()) {
				continue
			}
		}
		break
	}
}

//	Until reads values from a channel into a function until a condition is met or the channel is closed
func (c *Channel) Until(f func(x interface{}) bool) {
	for {
		if v, open := c.Recv(); open {
			if !f(v.Interface()) {
				continue
			}
		}
		break
	}
}