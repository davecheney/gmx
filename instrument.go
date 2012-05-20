package gmx

import (
	"sync/atomic"
)

type counter struct {
	value uint64
}

func (c *counter) Inc() {
	atomic.AddUint64(&c.value, 1)
}

func NewCounter(name string) *counter {
	c := new(counter)
	Publish(name, func() interface{} {
		return c.value
	})
	return c	
}
	
