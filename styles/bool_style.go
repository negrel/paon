package styles

import (
	"unsafe"

	"github.com/negrel/paon/pdk/id"
	"github.com/negrel/paon/pdk/id/store"
	"github.com/negrel/paon/styles/property"
)

// BoolStyle define objects containing property.Bool style properties.
type BoolStyle interface {
	Bool(property.BoolID) *property.Bool
	SetBool(property.BoolID, *property.Bool)
}

var _ BoolStyle = boolStyle{}

type boolStyle struct {
	store.PtrSlice
}

// NewBoolStyle returns a new BoolStyle instance.
func NewBoolStyle() BoolStyle {
	return newBoolStyle()
}

func newBoolStyle() boolStyle {
	return boolStyle{
		store.NewPtrSlice(int(property.BoolIDCount() + 1)),
	}
}

func (bs boolStyle) Bool(i property.BoolID) *property.Bool {
	return (*property.Bool)(bs.Get(id.ID(i)))
}

func (bs boolStyle) SetBool(i property.BoolID, c *property.Bool) {
	bs.PtrSlice.Set(id.ID(i), unsafe.Pointer(c))
}
