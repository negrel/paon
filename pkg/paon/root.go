package paon

import (
	"github.com/negrel/paon/internal/events"
	"github.com/negrel/paon/pkg/widgets"
)

type root struct {
	*widgets.Root
}

func newRoot(child widgets.Widget) *root {
	return &root{
		Root: widgets.NewRoot(child),
	}
}

func (r *root) DispatchEvent(event events.Event) {
	canvas := r.Render(screen.Size())
	engine.Draw(canvas)

	r.Layout.DispatchEvent(event)
}
