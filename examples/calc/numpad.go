package main

import (
	"github.com/negrel/paon/colors"
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
	"github.com/negrel/paon/widgets"
)

type NumPad struct {
	widgets.BaseLayout
}

func NewNumPad(m *Model) *NumPad {
	numPad := &NumPad{}
	numPad.BaseLayout = widgets.NewBaseLayout(numPad)

	keyStyle := draw.CellStyle{
		Foreground: colors.ColorWhite,
		Background: colors.ColorFromHex(0x24292f),
	}

	// First row
	numPad.AppendChild(NewPadding(NewKey("7", keyStyle, m), 1, 2, 0, 0))
	numPad.AppendChild(NewPadding(NewKey("8", keyStyle, m), 1, 2, 0, 0))
	numPad.AppendChild(NewPadding(NewKey("9", keyStyle, m), 1, 2, 0, 0))

	// Second row
	numPad.AppendChild(NewPadding(NewKey("4", keyStyle, m), 1, 2, 0, 0))
	numPad.AppendChild(NewPadding(NewKey("5", keyStyle, m), 1, 2, 0, 0))
	numPad.AppendChild(NewPadding(NewKey("6", keyStyle, m), 1, 2, 0, 0))

	// Third row
	numPad.AppendChild(NewPadding(NewKey("1", keyStyle, m), 1, 2, 0, 0))
	numPad.AppendChild(NewPadding(NewKey("2", keyStyle, m), 1, 2, 0, 0))
	numPad.AppendChild(NewPadding(NewKey("3", keyStyle, m), 1, 2, 0, 0))

	// Last row
	numPad.AppendChild(NewPadding(NewKey("0", keyStyle, m), 1, 2, 0, 0))
	numPad.AppendChild(NewPadding(NewKey(".", keyStyle, m), 1, 2, 0, 0))

	return numPad
}

// Layout implements layout.Layout.
func (np *NumPad) Layout(co layout.Constraint) geometry.Size {
	np.BaseLayout.ChildrenLayout.Reset()
	child := np.FirstChild()

	keyCo := co.ForceSize(co.MaxSize.
		WithHeight(co.MaxSize.Height / 4).
		WithWidth(co.MaxSize.Width / 3))

	for row := 0; row < 4; row++ {
		y := row * keyCo.MaxSize.Height

		for col := 0; col < 3; col++ {
			x := col * keyCo.MaxSize.Width

			if row == 3 && col == 2 {
				break
			}

			childSize := child.Layout(keyCo)

			np.BaseLayout.ChildrenLayout.Append(widgets.ChildLayout{
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
