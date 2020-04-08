package ginger

import (
	"image"

	"github.com/negrel/ginger/v1/painting"
	"github.com/negrel/ginger/v1/widget"
)

// Activity is a screen of the application
type Activity struct {
	screen       *Screen
	paintChannel chan *painting.Frame
	Root         widget.Widget
}

// Start the activity
func (ac *Activity) Start() {
	scr := *ac.screen
	w, h := scr.Size()
	bounds := image.Rect(0, 0, w, h)
	frame := ac.Root.Draw(bounds)

	ac.paintChannel <- frame
}
