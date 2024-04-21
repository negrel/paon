package widgets

import (
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/render"
	"github.com/negrel/paon/tree"
)

// Widget is a generic interface that define any component part of the
// widget/element tree. Any types that implement the Widget interface can be
// added to the widget tree.
type Widget interface {
	events.Target
	render.Renderable
	tree.Node[Widget]
}
