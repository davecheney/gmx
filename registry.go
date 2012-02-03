package gmx

import (
	"sync"
)

type Registry struct {
	config  Config
	mu      sync.Mutex
	entries map[string]func() interface{}
}

func newRegistry(c Config) *Registry {
	return &Registry{
		config:  c,
		entries: make(map[string]func() interface{}),
	}
}

func (r *Registry) register(name string, getter func() interface{}) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.entries[name] = getter
	log.Printf("register: %s, %#v", name, getter)
}

var nilfunc = func() interface{} { return nil }

func (r *Registry) getter(name string) func() interface{} {
	r.mu.Lock()
	defer r.mu.Unlock()
	f, ok := r.entries[name]
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
