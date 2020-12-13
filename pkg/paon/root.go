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
	r.Root.DispatchEvent(event)

	// Render the entire tree on screen resize
	if event.Type() == events.ResizeEventType {
		r.fullRender()
	}
}

func (r *root) fullRender() {
	screen.Clear()

	canvas := r.Render(screen.Size())
	engine.Draw(canvas)
}
