package main

import (
	"github.com/negrel/paon/colors"
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
	"github.com/negrel/paon/widgets"
)

type OpPad struct {
	widgets.BaseLayout
}

func NewOpPad(m *Model) *OpPad {
	opPad := &OpPad{}
	opPad.BaseLayout = widgets.NewBaseLayout(opPad)

	keyStyle := draw.CellStyle{
		Foreground: colors.ColorBlack,
		Background: colors.ColorFromHex(0xfea62a),
	}

	opPad.AppendChild(NewPadding(NewKey("/", keyStyle, m), 1, 0, 0, 0))

	opPad.AppendChild(NewPadding(NewKey("*", keyStyle, m), 1, 0, 0, 0))

	opPad.AppendChild(NewPadding(NewKey("-", keyStyle, m), 1, 0, 0, 0))

	opPad.AppendChild(NewPadding(NewKey("+", keyStyle, m), 1, 0, 0, 0))

	return opPad
}

// Layout implements layout.Layout.
func (op *OpPad) Layout(co layout.Constraint) geometry.Size {
	op.BaseLayout.ChildrenLayout.Reset()

	child := op.FirstChild()

	keyCo := co.ForceSize(co.MaxSize.WithHeight(co.MaxSize.Height / 4))

	for row := 0; row < 4; row++ {
		y := row * keyCo.MaxSize.Height

		for col := 0; col < 1; col++ {
			x := col * keyCo.MaxSize.Width

			childSize := child.Layout(keyCo)

			op.BaseLayout.ChildrenLayout.Append(widgets.ChildLayout{
				Widget: child,
				Bounds: geometry.Rectangle{
					Origin:   geometry.Vec2D{X: x, Y: y},
					RectSize: childSize,
				},
			})

			child = child.NextSibling()
		}
	}

	return co.MaxSize
}
