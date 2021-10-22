package id

import "sync/atomic"

// Registry define an identifier registry.
// All id generated from a registry are unique.
type Registry struct {
	counter uint32
}

// New returns a new unique ID.
func (r *Registry) New() ID {
	id := atomic.AddUint32(&r.counter, 1)
	return ID(id)
}

// ID define a unique number.
type ID uint32

var globalReg = Registry{
	counter: 0,
}

// New returns a new unique ID from the global Registry.
func New() ID {
	return globalReg.New()
}
