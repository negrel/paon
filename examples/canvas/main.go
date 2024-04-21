package main

import (
	"context"

	"github.com/negrel/paon"
	"github.com/negrel/paon/colors"
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/events/mouse"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
	"github.com/negrel/paon/render"
	"github.com/negrel/paon/widgets"
	"github.com/negrel/paon/widgets/button"
	"github.com/negrel/paon/widgets/vbox"
)

func main() {
	app, err := paon.NewApp()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start application.
	err = app.Start(ctx, widgets.NewRoot(
		vbox.New(vbox.WithChildren(
			button.New("Click to exit", button.OnClick(func(_ mouse.ClickEvent) {
				cancel()
			})),
			NewCanvas()),
		)))
	if err != nil {
		panic(err)
	}
}

type Canvas struct {
	*widgets.PanicWidget
	canvasRenderable
}

func NewCanvas() *Canvas {
	c := &Canvas{}
	c.PanicWidget = widgets.NewPanicWidget(c)
	c.canvasRenderable = canvasRenderable{
		VoidRenderable: render.NewVoidRenderable(c),
	}

	mousePress := false
	c.AddEventListener(mouse.PressListener(func(event mouse.Event) {
		mousePress = true
	}))
	c.AddEventListener(mouse.UpListener(func(event mouse.Event) {
		mousePress = false
	}))
	c.AddEventListener(mouse.EventListener(func(event mouse.Event) {
		if mousePress {
			c.canvasRenderable.cells[(event.AbsPosition.Y-1)*c.canvasRenderable.rootSize.Width+event.RelPosition.X] = true
			c.canvasRenderable.MarkDirty()
		}
	}))

	return c
}

// Renderable implements render.Renderable.
func (c *Canvas) Renderable() render.Renderable {
	return &c.canvasRenderable
}

type canvasRenderable struct {
	render.VoidRenderable
	cells    []bool
	rootSize geometry.Size
}

// Layout implements layout.Layout.
func (cr *canvasRenderable) Layout(co layout.Constraint) geometry.Size {
	if co.RootSize.Width*co.RootSize.Height != len(cr.cells) {
		cr.cells = make([]bool, co.RootSize.Width*co.RootSize.Height)
		cr.rootSize = co.RootSize
	}

	return co.MaxSize
}

// Draw implements draw.Drawer.
func (cr *canvasRenderable) Draw(surface draw.Surface) {
	cr.VoidRenderable.Draw(surface)

	x := 0
	y := 0
	for _, cell := range cr.cells {
		if cell {
			surface.Set(geometry.Vec2D{X: x, Y: y}, draw.Cell{
				Style: draw.CellStyle{
					Background: colors.ColorWhite,
				},
				Content: 0,
			})
		}

		x++
		if x == surface.Size().Width {
			x = 0
			y++
		}
	}
}
