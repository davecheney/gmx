package gmx

import (
	"sync"
)

type Registry struct {
	mu      sync.Mutex
	entries map[string]func() interface{}
}

func newRegistry() *Registry {
	return &Registry{
		entries: make(map[string]func() interface{}),
	}
}

func (r *Registry) register(key string, f func() interface{}) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.entries[key] = f
}

var nilfunc = func() interface{} { return nil }

func (r *Registry) value(key string) func() interface{} {
	r.mu.Lock()
	defer r.mu.Unlock()
	f, ok := r.entries[key]
	if !ok {
		return nilfunc
	}
	return f
}

func (r *Registry) keys() (k []string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for e := range r.entries {
		k = append(k, e)
	}
	return
}
