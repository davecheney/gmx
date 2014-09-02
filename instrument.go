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

type gauge struct {
	value int64
}

func (g *gauge) Inc() {
	atomic.AddInt64(&g.value, 1)
}

func (g *gauge) Dec() {
	atomic.AddInt64(&g.value, -1)
}

func NewGauge(name string) *gauge {
	g := new(gauge)
	Publish(name, func() interface{} {
		return g.value
	})
	return g
}
