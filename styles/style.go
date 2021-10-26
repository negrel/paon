package styles

import (
	"github.com/negrel/paon/pdk/events"
	"github.com/negrel/paon/styles/property"
)

// Styled is a generic interface for object that have a Style.
type Styled interface {
	Style() Style
}

// Style is a set of property.Property object.
// Property change can be observed by listening to EventPropertyChange events.
type Style interface {
	events.Target

	// Set sets the given property.
	Set(property.Property)

	// Get gets a property.
	Get(property.ID) property.Property

	// Del deletes a property.
	Del(property.ID)
}

var _ Style = &style{}

type style struct {
	events.Target

	props       []property.Property
	customProps map[property.ID]property.Property
}

var noOpTarget = events.NewNoOpTarget()

// New returns a new Style object configured with the given options.
func New(target events.Target) Style {
	if target == nil {
		target = noOpTarget
	}

	return newStyle(target)
}

func newStyle(target events.Target) *style {
	style := &style{
		Target:      target,
		props:       make([]property.Property, property.LastID()+1),
		customProps: make(map[property.ID]property.Property, 8),
	}

	return style
}

// Del implements the Style interface.
func (s *style) Del(id property.ID) {
	if !property.IsCustomPropID(id) {
		s.props[uint32(id)] = nil
	} else {
		delete(s.customProps, id)
	}
}

// Set implements the Style interface.
func (s *style) Set(prop property.Property) {
	if !property.IsCustomPropID(prop.ID()) {
		s.props[uint32(prop.ID())] = prop
	} else {
		s.customProps[prop.ID()] = prop
	}
}

// Get implements the Style interface.
func (s style) Get(id property.ID) property.Property {
	if !property.IsCustomPropID(id) {
		return s.props[uint32(id)]
	}

	return s.customProps[id]
}

// WeightedStyle extends the Style interface with a Weight method.
type WeightedStyle interface {
	Style

	// Weight returns the weight of this style.
	Weight() int
}

var _ WeightedStyle = Weighted{}

// Weighted is a simple wrapper around Style that implements
// the WeightedStyle interface.
type Weighted struct {
	Style
	weight int
}

// NewWeighted returns a new WeightedStyle object.
func NewWeighted(style Style, weight int) Weighted {
	return Weighted{Style: style, weight: weight}
}

// Weight implements the WeightedStyle interface.
func (ws Weighted) Weight() int {
	return ws.weight
}
