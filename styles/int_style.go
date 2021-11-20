package styles

import (
	"unsafe"

	"github.com/negrel/paon/pdk/id"
	"github.com/negrel/paon/pdk/id/store"
	"github.com/negrel/paon/styles/property"
)

// IntStyle define objects containing property.Int style properties.
type IntStyle interface {
	Int(property.IntID) *property.Int
	SetInt(property.IntID, *property.Int)
}

var _ IntStyle = intStyle{}

type intStyle struct {
	ints store.PtrSlice
}

// NewIntStyle returns a new IntStyle instance.
func NewIntStyle() IntStyle {
	return newIntStyle()
}

func newIntStyle() intStyle {
	return intStyle{
		ints: store.NewPtrSlice(int(property.IntIDCount() + 1)),
	}
}

func (is intStyle) Int(i property.IntID) *property.Int {
	return (*property.Int)(is.ints.Get(id.ID(i)))
}

func (is intStyle) SetInt(i property.IntID, p *property.Int) {
	is.ints.Set(id.ID(i), unsafe.Pointer(p))
}
