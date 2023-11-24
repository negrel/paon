package widgets

import (
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/events/resize"
	"github.com/negrel/paon/render"
)

var (
	NeedRenderEventType = events.NewType("need-render")
)

// NeedRenderEvent is emitted by Root widget when widgets tree needs to be rendered.
type NeedRenderEvent struct {
	events.Event
}

// NeedRenderEventListener returns an events.Listener that will call the given handler
// on resize events.
func NeedRenderEventListener(handler func(NeedRenderEvent)) (events.Type, events.Handler) {
	return NeedRenderEventType, events.HandlerFunc(func(ev events.Event) {
		handler(ev.(NeedRenderEvent))
	})
}

// Root define root of widgets tree.
type Root struct {
	Widget
	rootRenderable
}

// NewRoot returns a new Root widget that wraps the given Widget and its
// render.Renderable.
func NewRoot(w Widget) Root {
	root := Root{
		Widget: w,
		rootRenderable: rootRenderable{
			target:     w,
			Renderable: w.Node().Unwrap().(render.RenderableAccessor).Renderable(),
		},
	}

	// Place root inside node.
	root.Node().Swap(root)

	root.AddEventListener(resize.Listener(func(resize.Event) {
		root.rootRenderable.MarkDirty()
	}))

	return root
}

// Renderable implements render.RenderableAccessor.
func (r Root) Renderable() render.Renderable {
	return r.rootRenderable
}

type rootRenderable struct {
	target events.Target
	render.Renderable
}

func (rr rootRenderable) MarkDirty() {
	if !rr.Renderable.IsDirty() {
		rr.target.DispatchEvent(NeedRenderEvent{
			Event: events.NewEvent(NeedRenderEventType),
		})
	}
	rr.Renderable.MarkDirty()
}
