package layout

import "github.com/negrel/ginger/v1/widget"

// Layout is a widget that lay out is childrens widgets.
type Layout interface {
	widget.Widget

	// AppendChild method append the given child and
	// trigger a reflow.
	AppendChild(widget.Widget) error

	// Reflow force a repaint of all the child
	Reflow()
}
