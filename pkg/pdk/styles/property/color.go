package property

import (
	"github.com/negrel/paon/pkg/pdk/styles/value"
)

var _ Property = Color{}

type Color struct {
	Prop
	Value value.Color
}

func MakeColor(id ID, value value.Color) Color {
	return Color{
		Prop:  MakeProp(id),
		Value: value,
	}
}
