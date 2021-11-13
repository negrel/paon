package styles

import (
	"github.com/negrel/paon/styles/property"
)

// IntStyle define objects containing property.Int style properties.
type IntStyle interface {
	Int(property.IntID) *property.Int
	SetInt(property.IntID, *property.Int)
}

var _ IntStyle = intStyle{}

type intStyle struct {
	ints []*property.Int
}

// NewIntStyle returns a new IntStyle instance.
func NewIntStyle() IntStyle {
	return newIntStyle()
}

func newIntStyle() intStyle {
	return intStyle{
		ints: make([]*property.Int, property.IntIDCount()+1),
	}
}

func (is intStyle) Int(id property.IntID) *property.Int {
	return is.ints[uint32(id)]
}

func (is intStyle) SetInt(id property.IntID, p *property.Int) {
	is.ints[uint32(id)] = p
}
