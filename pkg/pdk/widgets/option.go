package widgets

import (
	"github.com/negrel/paon/internal/events"
)

type Option func(widget Widget)
type LayoutOption func(layout Layout)

func Listener(eventType events.EventType, listener events.Listener) Option {
	return func(widget Widget) {
		widget.AddEventListener(eventType, &listener)
	}
}
