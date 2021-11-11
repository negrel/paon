package id

import "sync/atomic"

// ID define a unique number.
type ID uint32

// Registry define an identifier registry.
// All id generated from the same registry are unique.
type Registry struct {
	counter uint32
}

// New returns a new unique ID.
func (r *Registry) New() ID {
	id := atomic.AddUint32(&r.counter, 1)
	return ID(id)
}

// Last returns the most recent ID generated.
func (r *Registry) Last() ID {
	id := atomic.LoadUint32(&r.counter)
	return ID(id)
}

var globalReg = Registry{
	counter: 0,
}

// New returns a new unique ID from the global Registry.
func New() ID {
	return globalReg.New()
}
