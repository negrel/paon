package styles

import (
	"github.com/negrel/paon/styles/property"
)

// IfaceStyle define objects containing property.Iface style properties.
type IfaceStyle interface {
	Iface(property.IfaceID) interface{}
	SetIface(property.IfaceID, interface{})
}

var _ IfaceStyle = ifaceStyle{}

type ifaceStyle struct {
	ifaces []interface{}
}

// NewIfaceStyle returns a new IfaceStyle instance.
func NewIfaceStyle() IfaceStyle {
	return newIfaceStyle()
}

func newIfaceStyle() ifaceStyle {
	return ifaceStyle{
		ifaces: make([]interface{}, property.IfaceIDCount()+1),
	}
}

func (is ifaceStyle) Iface(id property.IfaceID) interface{} {
	return is.ifaces[uint32(id)]
}

func (is ifaceStyle) SetIface(id property.IfaceID, i interface{}) {
	is.ifaces[uint32(id)] = i
}
