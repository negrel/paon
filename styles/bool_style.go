package styles

import "github.com/negrel/paon/styles/property"

// BoolStyle define objects containing property.Bool style properties.
type BoolStyle interface {
	Bool(property.BoolID) *property.Bool
	SetBool(property.BoolID, *property.Bool)
}

var _ BoolStyle = boolStyle{}

type boolStyle struct {
	bools []*property.Bool
}

// NewBoolStyle returns a new BoolStyle instance.
func NewBoolStyle() BoolStyle {
	return newBoolStyle()
}

func newBoolStyle() boolStyle {
	return boolStyle{
		bools: make([]*property.Bool, property.BoolIDCount()+1),
	}
}

func (bs boolStyle) Bool(id property.BoolID) *property.Bool {
	return bs.bools[uint32(id)]
}

func (bs boolStyle) SetBool(id property.BoolID, c *property.Bool) {
	bs.bools[uint32(id)] = c
}
