package main

import (
	"github.com/negrel/paon/colors"
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
	"github.com/negrel/paon/widgets"
)

type Screen struct {
	widgets.BaseWidget
	result string
}

func NewScreen() *Screen {
	screen := &Screen{
		result: "0",
	}
	screen.BaseWidget = widgets.NewBaseWidget(screen)
	return screen
}

// Layout implements layout.Layout.
func (s *Screen) Layout(co layout.Constraint) geometry.Size {
	return co.MaxSize
}

func (s *Screen) SetText(str string) {
	s.result = str
	s.NeedRender()
}

// Draw implements draw.Drawer.
func (s *Screen) Draw(srf draw.Surface) {
	for x := 0; x < srf.Size().Width; x++ {
		for y := 0; y < srf.Size().Height; y++ {
			srf.Set(geometry.Vec2D{X: x, Y: y}, draw.Cell{
				Style: draw.CellStyle{
					Background: colors.ColorFromHex(0x3b689f),
				},
			})
		}
	}

	y := srf.Size().Height/2 - 1

	for i, r := range s.result {
		srf.Set(geometry.Vec2D{
			X: srf.Size().Width - len(s.result) + i - 1,
			Y: y,
		}, draw.Cell{
			Style: draw.CellStyle{
				Foreground: colors.ColorWhite,
				Background: colors.ColorFromHex(0x3b689f),
			},
			Content: r,
		})
	}
}
