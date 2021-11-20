package store

import (
	"unsafe"

	"github.com/negrel/paon/pdk/id"
)

// Ptr define a generic interface for unsafe.Pointer store.
type Ptr interface {
	Set(id.ID, unsafe.Pointer)
	Get(id.ID) unsafe.Pointer
	Swap(id.ID, unsafe.Pointer) unsafe.Pointer
}

var _ Ptr = PtrSlice{}

// PtrSlice is an implementation Ptr store.
type PtrSlice []unsafe.Pointer

// NewPtrSlice returns a new PtrSlice store of the given size
func NewPtrSlice(size int) PtrSlice {
	return PtrSlice(make([]unsafe.Pointer, size))
}

// Set assigns the given value to the given id in the store.
func (ps PtrSlice) Set(id id.ID, value unsafe.Pointer) {
	ps[id] = value
}

// Get returns the value associated to this id.
func (ps PtrSlice) Get(id id.ID) unsafe.Pointer {
	return ps[id]
}

// Swap swaps the store value of the given id with the given new one. The old
// value is returned.
func (ps PtrSlice) Swap(id id.ID, newValue unsafe.Pointer) unsafe.Pointer {
	old := ps.Get(id)
	ps.Set(id, newValue)

	return old
}

var _ Ptr = PtrMap{}

// PtrMap is an implementation of Ptr store that use an internal map.
type PtrMap map[id.ID]unsafe.Pointer

// NewPtrMap returns a new PtrMap store.
func NewPtrMap() PtrMap {
	return make(PtrMap)
}

// Set assigns the given value to the given id in the store.
func (pm PtrMap) Set(id id.ID, value unsafe.Pointer) {
	pm[id] = value
}

// Get returns the value associated to this id.
func (pm PtrMap) Get(id id.ID) unsafe.Pointer {
	return pm[id]
}

// Swap swaps the store value of the given id with the given new one. The old
// value is returned.
func (pm PtrMap) Swap(id id.ID, newValue unsafe.Pointer) unsafe.Pointer {
	old := pm.Get(id)
	pm.Set(id, newValue)

	return old
}
