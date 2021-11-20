package store

import "github.com/negrel/paon/pdk/id"

// Iface define a generic interface for interface{} store.
type Iface interface {
	Set(id.ID, interface{})
	Get(id.ID) interface{}
	Swap(id.ID, interface{}) interface{}
}

var _ Iface = IfaceSlice{}

// IfaceSlice is an implementation Iface store.
type IfaceSlice []interface{}

// NewIfaceSlice returns a new IfaceSlice store of the given size
func NewIfaceSlice(size int) IfaceSlice {
	return IfaceSlice(make([]interface{}, size))
}

// Set assigns the given value to the given id in the store.
func (ps IfaceSlice) Set(id id.ID, value interface{}) {
	ps[id] = value
}

// Get returns the value associated to this id.
func (ps IfaceSlice) Get(id id.ID) interface{} {
	return ps[id]
}

// Swap swaps the store value of the given id with the given new one. The old
// value is returned.
func (ps IfaceSlice) Swap(id id.ID, newValue interface{}) interface{} {
	old := ps.Get(id)
	ps.Set(id, newValue)

	return old
}

var _ Iface = IfaceMap{}

// IfaceMap is an implementation of Iface store that use an internal map.
type IfaceMap map[id.ID]interface{}

// NewIfaceMap returns a new IfaceMap store.
func NewIfaceMap() IfaceMap {
	return make(IfaceMap)
}

// Set assigns the given value to the given id in the store.
func (pm IfaceMap) Set(id id.ID, value interface{}) {
	pm[id] = value
}

// Get returns the value associated to this id.
func (pm IfaceMap) Get(id id.ID) interface{} {
	return pm[id]
}

// Swap swaps the store value of the given id with the given new one. The old
// value is returned.
func (pm IfaceMap) Swap(id id.ID, newValue interface{}) interface{} {
	old := pm.Get(id)
	pm.Set(id, newValue)

	return old
}
