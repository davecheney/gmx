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
	
type guage struct {
	value int64
}

func (g *guage) Inc() {
	atomic.AddInt64(&g.value, 1)
}

func (g *guage) Dec() {
	atomic.AddInt64(&g.value, -1)
}

func NewGuage(name string) *guage {
	g := new(guage)
	Publish(name, func() interface{} {
		return g.value
	})
	return g	
}
	
