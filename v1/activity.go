package ginger

import (
	"image"

	"github.com/negrel/ginger/v1/style"
	"github.com/negrel/ginger/v1/widget"
)

// Activity is a screen of the application
type Activity struct {
	screen       *Screen
	paintChannel chan style.Frame
	Root         widget.Widget
}

// Start the activity
func (ac *Activity) Start() {
	scr := *ac.screen
	x, y := scr.Size()

	f := ac.Root.Draw(widget.Constraint{
		R: image.Rectangle{
			Min: image.Point{},
			Max: image.Point{
				X: x,
				Y: y,
			},
		},
	})

	ac.paintChannel <- *f
}
