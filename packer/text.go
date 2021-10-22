package packer

import (
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/pdk/layout"
)

var _ layout.BoxPacker = Text{}

type Text struct {
	src   *string
	lines []string
}

func NewText(data *string) Text {
	return Text{
		src:   data,
		lines: make([]string, 0),
	}
}

// Pack implements the layout.Packer interface.
func (t Text) Pack(co layout.Constraint) layout.BoxModel {
	size := geometry.NewSize(len(*t.src), 1)

	return layout.NewBox(co.ApplyOnSize(size))
}
