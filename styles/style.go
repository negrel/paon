package styles

import (
	"github.com/negrel/paon/pdk/events"
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
// This function panic if the target is nil.
// The returned Style will support properties ID (Color, Int, Iface...)
// created before the Style itself.
func New(target events.Target) Style {
	if target == nil {
		panic("style events.Target must no be nil")
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
