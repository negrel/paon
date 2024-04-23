package main

import (
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
	"github.com/negrel/paon/minmax"
	"github.com/negrel/paon/widgets"
)

type Padding struct {
	widgets.BaseLayout
	Left, Right, Top, Bottom int
}

func NewPadding(w widgets.Widget, pad int, extraPads ...int) *Padding {
	p := &Padding{}
	p.BaseLayout = widgets.NewBaseLayout(p)
	p.AppendChild(w)

	switch len(extraPads) {
	case 3:
		p.Top = pad
		p.Right = extraPads[0]
		p.Bottom = extraPads[1]
		p.Left = extraPads[2]

	case 2:
		p.Top = pad
		p.Left = extraPads[0]
		p.Right = extraPads[0]
		p.Bottom = extraPads[1]

	case 1:
		p.Top = pad
		p.Bottom = pad
		p.Left = extraPads[0]
		p.Right = extraPads[0]

	case 0:
		p.Top = pad
		p.Right = pad
		p.Bottom = pad
		p.Left = pad

	default:
		panic("too much padding arguments")
	}

	return p
}

func (p *Padding) Layout(co layout.Constraint) geometry.Size {
	p.ChildrenLayout.Reset()

	child := p.FirstChild()
	if child == nil {
		return geometry.Size{Width: p.Left + p.Right, Height: p.Top + p.Bottom}
	}

	co.MaxSize.Width -= p.Left + p.Right
	co.MaxSize.Width = minmax.Max(co.MaxSize.Width, 0)

	co.MaxSize.Height -= p.Top + p.Bottom
	co.MaxSize.Height = minmax.Max(co.MaxSize.Height, 0)

	size := child.Layout(co)
	p.ChildrenLayout.Append(widgets.ChildLayout{
		Widget: child,
		Bounds: geometry.Rectangle{
			Origin:   geometry.Vec2D{X: p.Left, Y: p.Top},
			RectSize: size,
		},
	})

	size.Width += p.Left + p.Right
	size.Height += p.Top + p.Bottom

	return size
}
