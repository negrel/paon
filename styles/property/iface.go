package property

import (
	"github.com/negrel/paon/pdk/id"
	"github.com/negrel/paon/pdk/id/store"
)

// IfaceID define a unique ID accross all Iface properties.
type IfaceID id.ID

var (
	ifaceRegistry = id.Registry{}
	ifaceMap      = store.NewStringMap()
)

// NewIfaceID returns a new unique interface property ID.
func NewIfaceID(name string) IfaceID {
	id := ifaceRegistry.New()
	ifaceMap.Set(id, name)

	return IfaceID(id)
}

// String implements the fmt.Stringer interface.
func (ii IfaceID) String() string {
	return ifaceMap.Get(id.ID(ii))
}

// IfaceIDCount returns the number of IfaceID generated.
func IfaceIDCount() uint32 {
	return uint32(ifaceRegistry.Last())
}

var (
	_IDBorderCharSet = NewIfaceID("border-charset")
)

// BorderCharSet returns the IfaceID of the "border-charset" property.
func BorderCharSet() IfaceID {
	return _IDBorderCharSet
}
