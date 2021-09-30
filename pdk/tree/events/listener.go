package events

import "github.com/negrel/paon/pdk/events"

// Listener define a wrapper around events.Listener for tree events.
type Listener struct {
	events.Listener

	bubbles bool
}

// NewListener returns a new Listener that wraps the given events.Listener.
// If bubbles is true, the listener will be called only if a bubble phase occur.
func NewListener(listener events.Listener, bubbles bool) Listener {
	return Listener{
		Listener: listener,
		bubbles:  bubbles,
	}
}
