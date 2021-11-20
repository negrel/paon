package property

import (
	"github.com/negrel/paon/pdk/id"
	"github.com/negrel/paon/pdk/id/store"
)

// IntID define a unique ID accross all Int properties.
type IntID id.ID

var (
	intRegistry = id.Registry{}
	intMap      = store.NewStringMap()
)

// NewIntID returns a new unique int property ID.
func NewIntID(name string) IntID {
	id := intRegistry.New()
	intMap.Set(id, name)

	return IntID(id)
}

// String implements the fmt.Stringer interface.
func (ii IntID) String() string {
	return intMap.Get(id.ID(ii))
}

// IntIDCount returns the number of IntID generated.
func IntIDCount() uint32 {
	return uint32(intRegistry.Last())
}

// Int define a read-only integer value.
type Int int

// NewInt returns a new Int with the given value.
func NewInt(value int) Int {
	return Int(value)
}

// Value returns the int value.
func (i *Int) Value() int {
	return int(*i)
}

var (
	_IDZIndex = NewIntID("z-index")
)

// ZIndex returns the IntID property of the "z-index" property.
func ZIndex() IntID {
	return _IDZIndex
}
