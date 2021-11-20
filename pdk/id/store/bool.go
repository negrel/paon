package store

import "github.com/negrel/paon/pdk/id"

// Bool define a generic interface for bool store.
type Bool interface {
	Set(id.ID, bool)
	Get(id.ID) bool
	Swap(id.ID, bool) bool
}

var _ Bool = BoolSlice{}

// BoolSlice is an implementation of Bool store.
type BoolSlice []bool

// NewBoolSlice returns a new BoolSlice store of the given size
func NewBoolSlice(size int) BoolSlice {
	return BoolSlice(make([]bool, size))
}

// Set assigns the given value to the given id in the store.
func (bs BoolSlice) Set(id id.ID, value bool) {
	bs[id] = value
}

// Get returns the value associated to this id.
func (bs BoolSlice) Get(id id.ID) bool {
	return bs[id]
}

// Swap swaps the store value of the given id with the given new one. The old
// value is returned.
func (bs BoolSlice) Swap(id id.ID, newValue bool) bool {
	old := bs.Get(id)
	bs.Set(id, newValue)

	return old
}

var _ Bool = BoolMap{}

// BoolMap is an implementation of Bool store.
type BoolMap map[id.ID]bool

// NewBoolMap returns a new BoolMap store.
func NewBoolMap() BoolMap {
	return make(BoolMap)
}

// Set assigns the given value to the given id in the store.
func (bm BoolMap) Set(id id.ID, value bool) {
	bm[id] = value
}

// Get returns the value associated to this id.
func (bm BoolMap) Get(id id.ID) bool {
	return bm[id]
}

// Swap swaps the store value of the given id with the given new one. The old
// value is returned.
func (bm BoolMap) Swap(id id.ID, newValue bool) bool {
	old := bm.Get(id)
	bm.Set(id, newValue)

	return old
}
