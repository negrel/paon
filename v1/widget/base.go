package widget

import "github.com/negrel/ginger/v1/style"

var _ Widget = &Base{}

// Base is use as base for every widget
type Base struct {
	parent Widget
}

// AdoptedBy set the widget parent
func (b *Base) AdoptedBy(w Widget) {
	b.parent = w
}

// Draw the widget
func (b *Base) Draw(c Constraint) *style.Frame {
	return &style.Frame{}
}

// repaint trigger a repaint of this component
func (b *Base) reflow() {
}

// repaint trigger a repaint of this component
func (b *Base) repaint(f style.Frame) {
	b.parent.repaint(f)
}

// Parent return the widget parent
func (b *Base) Parent() Widget {
	return b.parent
}
