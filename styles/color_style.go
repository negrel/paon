package styles

import (
	"github.com/negrel/paon/styles/property"
)

// ColorStyle define objects containing property.Color style properties.
type ColorStyle interface {
	Color(property.ColorID) *property.Color
	SetColor(property.ColorID, *property.Color)
}

var _ ColorStyle = colorStyle{}

type colorStyle struct {
	colors []*property.Color
}

// NewColorStyle returns a new ColorStyle instance.
func NewColorStyle() ColorStyle {
	return newColorStyle()
}

func newColorStyle() colorStyle {
	return colorStyle{
		colors: make([]*property.Color, property.ColorIDCount()+1),
	}
}

func (cs colorStyle) Color(id property.ColorID) *property.Color {
	return cs.colors[uint32(id)]
}

func (cs colorStyle) SetColor(id property.ColorID, c *property.Color) {
	cs.colors[uint32(id)] = c
}
