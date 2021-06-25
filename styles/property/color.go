package property

import (
	"github.com/negrel/paon/styles/value"
)

var _ Property = Color{}

// Color define any property that contains a color value.
type Color struct {
	Prop
	value.Color
}

// NewColor returns a new Color property with the given value.
func NewColor(id ID, color value.Color) Color {
	return Color{
		Prop:  NewProp(id),
		Color: color,
	}
}
