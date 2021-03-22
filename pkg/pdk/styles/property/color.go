package property

import (
	"github.com/negrel/paon/pkg/pdk/styles/value"
)

var _ Property = Color{}

// Color define any property that contains a color value.
type Color struct {
	Prop
	value.Color
}

// MakeColor returns a new Color property with the given value.
func MakeColor(id ID, color value.Color) Color {
	return Color{
		Prop:  MakeProp(id),
		Color: color,
	}
}
