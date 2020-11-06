package paon

import (
	"github.com/negrel/debuggo/pkg/assert"

	"github.com/negrel/paon/internal/events"
	"github.com/negrel/paon/internal/render"
	"github.com/negrel/paon/internal/widgets"
)

var _ widgets.Widget = &root{}

type root struct {
	widgets.Widget
	events.Target

	children widgets.Widget
}

func newRoot(child widgets.Widget) *root {
	assert.NotNil(child, "child must be non-nil")

	return &root{
		Widget:   widgets.NewWidget("root", widgets.Opt(renderRoot(child), 0)),
		children: child,
	}
}

func renderRoot(child widgets.Widget) func(render.Surface) {
	return func(buffer render.Surface) {
		child.Render(buffer)
	}
}

// func (r *root) DispatchEvent(event events.Event) {
// 	r.Target.DispatchEvent(event)
// 	r.children.DispatchEvent(event)
// }
