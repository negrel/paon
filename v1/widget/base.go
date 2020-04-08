package widget

import (
	"image"

	"github.com/negrel/ginger/v1/painting"
)

var _ Widget = &Base{}

// Base is use as base for every widget
type Base struct {
	active bool
	parent Widget
	cache  painting.Frame
}

// AdoptedBy set the widget parent
func (b *Base) AdoptedBy(w Widget) {
	b.parent = w
}

// Draw the widget
func (b *Base) Draw(image.Rectangle) *painting.Frame {
	return nil
}

// repaint trigger a repaint of this component
func (b *Base) reflow() {
}

// repaint trigger a repaint of this component
func (b *Base) repaint(f painting.Frame) {
	b.parent.repaint(f)
}

// Parent return the widget parent
func (b *Base) Parent() Widget {
	return b.parent
}
