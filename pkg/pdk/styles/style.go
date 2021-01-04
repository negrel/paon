package styles

import (
	"github.com/negrel/paon/internal/events"
	"github.com/negrel/paon/pkg/pdk/styles/property"
)

type Style interface {
	Set(property.Property)
	Get(property.ID) property.Property
	Del(property.ID)
}

var _ Style = style{}

type style struct {
	events.Target
	props map[property.ID]property.Property
}

func MakeStyle() Style {
	return style{
		props: make(map[property.ID]property.Property, 8),
	}
}

func (s style) Del(id property.ID) {
	delete(s.props, id)
}

func (s style) Set(prop property.Property) {
	if prop == nil {
		return
	}

	s.props[prop.ID()] = prop
}

func (s style) Get(id property.ID) property.Property {
	return s.props[id]
}
