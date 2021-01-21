package property

import (
	"github.com/negrel/paon/pkg/pdk/styles/value"
)

var _ Property = Color{}

type Color struct {
	Prop
	value.Color
}

func MakeColor(id ID, color value.Color) Color {
	return Color{
		Prop:  MakeProp(id),
		Color: color,
	}
}
