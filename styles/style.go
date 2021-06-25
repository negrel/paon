package styles

import (
	"sync"

	"github.com/negrel/paon/events"
	"github.com/negrel/paon/pdk/id"
	"github.com/negrel/paon/styles/property"
)

// Stylable define object that can have a Style.
type Stylable interface {
	id.Identifiable

	ParentStyle() Style
	Style() Style
}

// Style is a set of property.Property object.
// Property change can be observed by listening to EventPropertyChange events.
type Style interface {
	events.Target

	Weight() int

	Set(property.Property)
	Get(property.ID) property.Property
	Del(property.ID)
}

var _ Style = &style{}

type style struct {
	*sync.RWMutex
	events.Target

	weight int
	props  map[property.ID]property.Property
}

// NewStyle returns a new Style object configured with the given options.
func NewStyle(options ...Option) Style {
	style := &style{
		RWMutex: &sync.RWMutex{},
		Target:  events.NewTarget(),
		props:   make(map[property.ID]property.Property, 8),
	}

	for _, option := range options {
		option(style)
	}

	if style.Target == nil {
		style.Target = events.NewTarget()
	}

	return style
}

// Weight implements the Style interface.
func (s *style) Weight() int {
	s.RLock()
	defer s.RUnlock()

	return s.weight
}

// Del implements the Style interface.
func (s *style) Del(id property.ID) {
	s.Lock()
	defer s.Unlock()

	delete(s.props, id)
}

// Set implements the Style interface.
func (s *style) Set(prop property.Property) {
	s.Lock()
	defer s.Unlock()

	old := s.props[prop.ID()]
	s.props[prop.ID()] = prop

	s.DispatchEvent(newEventPropertyChange(old, prop))
}

// Get implements the Style interface.
func (s *style) Get(id property.ID) property.Property {
	s.RLock()
	defer s.RUnlock()

	return s.props[id]
}
