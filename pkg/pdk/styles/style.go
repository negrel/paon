package styles

import (
	"github.com/negrel/paon/pkg/pdk/events"
	"github.com/negrel/paon/pkg/pdk/styles/property"
)

// Stylised define object that are stylised and have a Style object.
type Stylised interface {
	Style() Style
}

// Style is a set of property.Property object.
type Style interface {
	events.Target

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
		Target: events.MakeTarget(),
		props:  make(map[property.ID]property.Property, 8),
	}
}

func (s style) Del(id property.ID) {
	delete(s.props, id)
}

func (s style) Set(prop property.Property) {
	old := s.props[prop.ID()]
	s.props[prop.ID()] = prop

	s.DispatchEvent(makeEventSetProperty(old, prop))
}

func (s style) Get(id property.ID) property.Property {
	return s.props[id]
}
