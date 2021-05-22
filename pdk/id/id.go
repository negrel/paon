package id

import "sync/atomic"

// ID define a unique number.
type ID int32

var idCounter int32 = 0

// New returns a new ID.
func New() ID {
	id := atomic.AddInt32(&idCounter, 1)

	return ID(id)
}
