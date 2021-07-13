package styles

import (
	"sync"

	"github.com/negrel/paon/styles/property"
)

// Style is a set of property.Property object.
// Property change can be observed by listening to EventPropertyChange events.
type Style interface {
	// Set sets the given property.
	Set(property.Property)

	// Get gets a property.
	Get(property.ID) property.Property

	// Del deletes a property.
	Del(property.ID)
}

func idToIndex(id property.ID) int {
	return int(id - property.FirstID())
}

var _ Style = &style{}

type style struct {
	*sync.RWMutex

	props       []property.Property
	customProps map[property.ID]property.Property
}

// New returns a new Style object configured with the given options.
func New() Style {
	return newStyle()
}

func newStyle() *style {
	style := &style{
		RWMutex:     &sync.RWMutex{},
		props:       make([]property.Property, property.LastID()-property.FirstID()+1),
		customProps: make(map[property.ID]property.Property, 8),
	}

	return style
}

// Del implements the Style interface.
func (s *style) Del(id property.ID) {
	s.Lock()
	defer s.Unlock()

	if !property.IsCustomPropID(id) {
		s.props[idToIndex(id)] = nil
	} else {
		delete(s.customProps, id)
	}
}

// Set implements the Style interface.
func (s *style) Set(prop property.Property) {
	s.Lock()
	defer s.Unlock()

	if !property.IsCustomPropID(prop.ID()) {
		s.props[idToIndex(prop.ID())] = prop
	} else {
		s.customProps[prop.ID()] = prop
	}
}

// Get implements the Style interface.
func (s *style) Get(id property.ID) property.Property {
	s.RLock()
	defer s.RUnlock()

	if !property.IsCustomPropID(id) {
		return s.props[idToIndex(id)]
	}

	return s.customProps[id]
}

// WeightedStyle extends the Style interface with a Weight method.
type WeightedStyle interface {
	Style

	// Weight returns the weight of this style.
	Weight() int
}

var _ WeightedStyle = weightedStyle{}

type weightedStyle struct {
	Style
	weight int
}

// NewWeighted returns a new WeightedStyle object.
func NewWeighted(style Style, weight int) WeightedStyle {
	return weightedStyle{Style: style, weight: weight}
}

func (ws weightedStyle) Weight() int {
	return ws.weight
}
