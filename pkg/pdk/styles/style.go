package styles

import (
	"github.com/negrel/paon/internal/tree"
	"github.com/negrel/paon/pkg/pdk/events"
	"github.com/negrel/paon/pkg/pdk/id"
	"github.com/negrel/paon/pkg/pdk/styles/property"
)

// Stylable define object that can have a Style.
type Stylable interface {
	id.Identifiable
	tree.Node

	ParentStyle() Style
	Style() Style
}

// Style is a set of property.Property object.
type Style interface {
	id.Identifiable
	events.Target

	Set(property.Property)
	Get(property.ID) property.Property
	Del(property.ID)
}

var _ Style = style{}

type style struct {
	events.Target

	id    id.ID
	props map[property.ID]property.Property
}

func MakeStyle() Style {
	return style{
		id:     id.Make(),
		Target: events.MakeTarget(),
		props:  make(map[property.ID]property.Property, 8),
	}
}

// ID implements the identifiable interface.
func (s style) ID() id.ID {
	return s.id
}

// Del implements the Style interface.
func (s style) Del(id property.ID) {
	delete(s.props, id)
}

// Set implements the Style interface.
func (s style) Set(prop property.Property) {
	old := s.props[prop.ID()]
	s.props[prop.ID()] = prop

	s.DispatchEvent(makeEventSetProperty(old, prop))
}

// Get implements the Style interface.
func (s style) Get(id property.ID) property.Property {
	return s.props[id]
}
