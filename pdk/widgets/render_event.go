package widgets

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/pdk/events"
	"github.com/negrel/paon/pdk/render"
)

var needRenderEventType = events.NewType("need-render")

// NeedRenderEventType returns the events.Type of need render events.
func NeedRenderEventType() events.Type {
	return needRenderEventType
}

// NeedRenderListener returns an events.Listener that will call the given handler
// on need render events.
func NeedRenderListener(handler func(NeedRenderEvent)) *events.Listener {
	return &events.Listener{
		Type: needRenderEventType,
		Handle: func(event events.Event) {
			assert.IsType(event, NewNeedRenderEvent(nil))
			handler(event.(NeedRenderEvent))
		},
	}
}

var _ events.Event = NeedRenderEvent{}

// NeedRenderEvent is triggered when a Widget need to be redrawn.
type NeedRenderEvent struct {
	events.Event
	Renderable render.Renderable
}

// NewNeedRenderEvent returns a new NeedRenderEvent.
func NewNeedRenderEvent(renderable render.Renderable) NeedRenderEvent {
	return NeedRenderEvent{
		Event:      events.NewEvent(needRenderEventType),
		Renderable: renderable,
	}
}
