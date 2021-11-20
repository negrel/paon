package store

import "github.com/negrel/paon/pdk/id"

// String define a generic interface for string store.
type String interface {
	Set(id.ID, string)
	Get(id.ID) string
	Swap(id.ID, string) string
}

// StringSlice is an implementation of String store.
type StringSlice []string

// NewStringSlice returns a new string store of the given size.
func NewStringSlice(size int) StringSlice {
	return StringSlice(make([]string, size))
}

// Set assigns the given value to the given id in the store.
func (ss StringSlice) Set(id id.ID, value string) {
	ss[id] = value
}

// Get returns the value associated to this id.
func (ss StringSlice) Get(id id.ID) string {
	return ss[id]
}

// Swap swaps the store value of the given id with the given new one. The old
// value is returned.
func (ss StringSlice) Swap(id id.ID, newValue string) string {
	old := ss.Get(id)
	ss.Set(id, newValue)

	return old
}

// StringMap is an implementation of String store.
type StringMap map[id.ID]string

// NewStringMap returns a new string store that can store element of the given id.Registry.
func NewStringMap() StringMap {
	return make(StringMap)
}

// Set assigns the given value to the given id in the store.
func (ss StringMap) Set(id id.ID, value string) {
	ss[id] = value
}

// Get returns the value associated to this id.
func (ss StringMap) Get(id id.ID) string {
	return ss[id]
}

// Swap swaps the store value of the given id with the given new one. The old
// value is returned.
func (ss StringMap) Swap(id id.ID, newValue string) string {
	old := ss.Get(id)
	ss.Set(id, newValue)

	return old
}
