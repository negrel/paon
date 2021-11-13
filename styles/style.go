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
	ColorStyle
	IntStyle
	IntUnitStyle
	IfaceStyle
}

var _ Style = &style{}

type style struct {
	events.Target
	colorStyle
	intStyle
	intUnitStyle
	ifaceStyle
}

// New returns a new Style object with the given target.
// A new target is created if tge given is nil.
// The returned Style will support properties ID (Color, Int, Iface...)
// created before the Style itself.
func New(target events.Target) Style {
	if target == nil {
		target = events.NewTarget()
	}

	return newStyle(target)
}

func newStyle(target events.Target) *style {
	style := &style{
		Target:       target,
		colorStyle:   newColorStyle(),
		intStyle:     newIntStyle(),
		intUnitStyle: newIntUnitStyle(),
		ifaceStyle:   newIfaceStyle(),
	}

	return style
}

func (s *style) SetColor(id property.ColorID, c *property.Color) {
	old := s.colorStyle.Color(id)
	if old != c {
		s.colorStyle.SetColor(id, c)
		s.DispatchEvent(NewColorChanged(id, old, c))
	}
}

func (s *style) SetIface(id property.IfaceID, i interface{}) {
	old := s.ifaceStyle.Iface(id)
	if old != i {
		s.ifaceStyle.SetIface(id, i)
		s.DispatchEvent(NewIfaceChanged(id, old, i))
	}
}

func (s *style) SetInt(id property.IntID, i *property.Int) {
	old := s.intStyle.Int(id)
	if old != i {
		s.intStyle.SetInt(id, i)
		s.DispatchEvent(NewIntChanged(id, old, i))
	}
}

func (s *style) SetIntUnit(id property.IntUnitID, i *property.IntUnit) {
	old := s.intUnitStyle.IntUnit(id)
	if old != i {
		s.intUnitStyle.SetIntUnit(id, i)
		s.DispatchEvent(NewIntUnitChanged(id, old, i))
	}
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
