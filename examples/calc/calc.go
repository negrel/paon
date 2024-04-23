package main

import (
	"github.com/negrel/paon/colors"
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
	"github.com/negrel/paon/widgets"
)

type Calc struct {
	widgets.BaseLayout
	model *Model
}

func NewCalc() *Calc {
	calc := &Calc{}
	calc.BaseLayout = widgets.NewBaseLayout(calc)

	screen := NewScreen()
	m := NewModel(screen)

	err := calc.AppendChild(NewPadding(screen, 1, 2, 0))
	if err != nil {
		panic(err)
	}
	err = calc.AppendChild(NewPadding(NewKeyPad(m), 0, 2, 1))
	if err != nil {
		panic(err)
	}

	return calc
}

// Layout implements layout.Layout.
func (c *Calc) Layout(co layout.Constraint) geometry.Size {
	c.BaseLayout.ChildrenLayout.Reset()
	screen := c.FirstChild()

	// Constraint screen to use at most a third of available space.
	screenCo := co.ForceSize(co.MaxSize.WithHeight(co.MaxSize.Height / 3))
	screenSize := screenCo.ApplyOnSize(screen.Layout(screenCo))

	// Store screen layout.
	c.BaseLayout.ChildrenLayout.Append(widgets.ChildLayout{
		Widget: screen,
		Bounds: geometry.Rectangle{
			Origin:   geometry.Vec2D{},
			RectSize: screenSize,
		},
	})

	// Constraint keypad to use at most all remaining space.
	keyPad := c.LastChild()
	keyPadCo := co.ForceSize(co.MaxSize.WithHeight(co.MaxSize.Height - co.MaxSize.Height/3))
	// Store keyPad layout.
	c.BaseLayout.ChildrenLayout.Append(widgets.ChildLayout{
		Widget: keyPad,
		Bounds: geometry.Rectangle{
			Origin:   geometry.Vec2D{X: 0, Y: screenSize.Height},
			RectSize: keyPadCo.ApplyOnSize(keyPad.Layout(keyPadCo)),
		},
	})

	return co.MaxSize
}

func (c *Calc) Draw(srf draw.Surface) {
	for x := 0; x < srf.Size().Width; x++ {
		for y := 0; y < srf.Size().Height; y++ {
			srf.Set(geometry.Vec2D{X: x, Y: y}, draw.Cell{
				Style: draw.CellStyle{
					Background: colors.ColorFromHex(0x1e1e1e),
				},
			})
		}
	}

	c.ChildrenLayout.Draw(srf)
}
