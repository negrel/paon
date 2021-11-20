package styles

import (
	"unsafe"

	"github.com/negrel/paon/pdk/id"
	"github.com/negrel/paon/pdk/id/store"
	"github.com/negrel/paon/styles/property"
)

// ColorStyle define objects containing property.Color style properties.
type ColorStyle interface {
	Color(property.ColorID) *property.Color
	SetColor(property.ColorID, *property.Color)
}

var _ ColorStyle = colorStyle{}

type colorStyle struct {
	colors store.Ptr
}

// NewColorStyle returns a new ColorStyle instance.
func NewColorStyle() ColorStyle {
	return newColorStyle()
}

func newColorStyle() colorStyle {
	return colorStyle{
		colors: store.NewPtrSlice(int(property.ColorIDCount() + 1)),
	}
}

func (cs colorStyle) Color(i property.ColorID) *property.Color {
	return (*property.Color)(cs.colors.Get(id.ID(i)))
}

func (cs colorStyle) SetColor(i property.ColorID, c *property.Color) {
	cs.colors.Set(id.ID(i), unsafe.Pointer(c))
}
