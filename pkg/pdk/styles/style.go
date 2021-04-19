package styles

import (
	"sync"

	"github.com/negrel/paon/pkg/pdk/events"
	"github.com/negrel/paon/pkg/pdk/id"
	"github.com/negrel/paon/pkg/pdk/styles/property"
)

// Stylable define object that can have a Style.
type Stylable interface {
	id.Identifiable

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
	sync.RWMutex
	events.Target

	id    id.ID
	props map[property.ID]property.Property
}

func NewStyle() Style {
	return style{
		id:     id.Make(),
		Target: events.MakeTarget(),
		props:  make(map[property.ID]property.Property, 8),
	}
}

// ID implements the identifiable interface.
func (s style) ID() id.ID {
	s.RLock()
	defer s.RUnlock()

	return s.id
}

// Del implements the Style interface.
func (s style) Del(id property.ID) {
	s.Lock()
	defer s.Unlock()

	delete(s.props, id)
}

// Set implements the Style interface.
func (s style) Set(prop property.Property) {
	s.Lock()
	defer s.Unlock()

	old := s.props[prop.ID()]
	s.props[prop.ID()] = prop

	s.DispatchEvent(makeEventPropertyChange(old, prop))
}

// Get implements the Style interface.
func (s style) Get(id property.ID) property.Property {
	s.RLock()
	defer s.RUnlock()

	return s.props[id]
}
