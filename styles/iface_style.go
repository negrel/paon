package styles

import (
	"github.com/negrel/paon/pdk/id"
	"github.com/negrel/paon/pdk/id/store"
	"github.com/negrel/paon/styles/property"
)

// IfaceStyle define objects containing property.Iface style properties.
type IfaceStyle interface {
	Iface(property.IfaceID) interface{}
	SetIface(property.IfaceID, interface{})
}

var _ IfaceStyle = ifaceStyle{}

type ifaceStyle struct {
	ifaces store.IfaceSlice
}

// NewIfaceStyle returns a new IfaceStyle instance.
func NewIfaceStyle() IfaceStyle {
	return newIfaceStyle()
}

func newIfaceStyle() ifaceStyle {
	return ifaceStyle{
		ifaces: store.NewIfaceSlice(int(property.IfaceIDCount() + 1)),
	}
}

func (is ifaceStyle) Iface(i property.IfaceID) interface{} {
	return is.ifaces.Get(id.ID(i))
}

func (is ifaceStyle) SetIface(i property.IfaceID, iface interface{}) {
	is.ifaces.Set(id.ID(i), iface)
}
