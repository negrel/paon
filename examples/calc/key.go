package main

import (
	"fmt"

	"github.com/negrel/paon/colors"
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
	"github.com/negrel/paon/widgets"
)

type Key struct {
	widgets.BaseWidget
	key     string
	style   draw.CellStyle
	clicked bool
}

func NewKey(key string, style draw.CellStyle, m *Model) *Key {
	w := &Key{
		key:     key,
		style:   style,
		clicked: false,
	}
	w.BaseWidget = widgets.NewBaseWidget(w)
	w.AddEventListener(events.MousePressListener(func(ev events.Event, _ events.MouseEventData) {
		w.clicked = true
		w.NeedRender()
	}))
	w.AddEventListener(events.ClickListener(func(ev events.Event, data events.ClickEventData) {
		w.clicked = false
		w.NeedRender()

		if len(key) == 1 && key[0] >= '0' && key[0] <= '9' {
			m.AddDigit(key)
			return
		}

		switch key {
		case ".":
			m.AddDecimal()
		case "=":
			m.Compute()
		case "AC":
			m.Clear()
		default:
			m.AddOperator(key)
		}

	}))

	return w
}

// Layout implements layout.Layout.
func (k *Key) Layout(co layout.Constraint) geometry.Size {
	return co.MaxSize
}

// Draw implements draw.Drawer.
func (k *Key) Draw(srf draw.Surface) {
	key := []rune(fmt.Sprint(k.key))

	for y := 0; y < srf.Size().Height; y++ {
		isCenterY := y == srf.Size().Height/2

		for x := 0; x < srf.Size().Width; x++ {
			isCenterX := x+len(key)/2 == srf.Size().Width/2
			if isCenterX && isCenterY {
				for i, r := range key {
					srf.Set(geometry.Vec2D{X: x + i, Y: y}, draw.Cell{
						Style:   k.style,
						Content: r,
					})
				}

				x += len(key) - 1
			} else {
				srf.Set(geometry.Vec2D{X: x, Y: y}, draw.Cell{
					Style:   k.style,
					Content: ' ',
				})
			}
		}
	}

	borderTopColor := mapColor(k.style.Background, increaseBrightness)
	borderBottomColor := mapColor(k.style.Background, decreaseBrightness)
	if k.clicked {
		borderTopColor, borderBottomColor = borderBottomColor, borderTopColor
	}

	// Borders.
	for x := 0; x < srf.Size().Width; x++ {
		srf.Set(geometry.Vec2D{X: x, Y: 0}, draw.Cell{
			Style: draw.CellStyle{
				Background: k.style.Background,
				Foreground: borderTopColor,
			},
			Content: '▔',
		})
		srf.Set(geometry.Vec2D{X: x, Y: srf.Size().Height - 1}, draw.Cell{
			Style: draw.CellStyle{
				Background: k.style.Background,
				Foreground: borderBottomColor,
			},
			Content: '▁',
		})
	}
}

func mapColor(c colors.Color, fn func(uint8) uint8) colors.Color {
	r := fn(c.R())
	g := fn(c.G())
	b := fn(c.B())

	return colors.ColorFromRGB(r, g, b)
}

func increaseBrightness(channel uint8) uint8 {
	if channel+0x33 < channel {
		return 255
	}
	return channel + 0x33
}

func decreaseBrightness(channel uint8) uint8 {
	if channel-0x33 > channel {
		return 0
	}
	return channel - 0x33
}
