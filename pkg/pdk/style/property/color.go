package property

import (
	"github.com/negrel/paon/pkg/pdk/style/value"
)

var _ Property = Color{}

type Color struct {
	Prop
	Value value.Color
}

func MakeColorProp(id ID, value value.Color) Color {
	return Color{
		Prop:  MakeProp(id),
		Value: value,
	}
}
