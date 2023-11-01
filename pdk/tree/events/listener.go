package events

import "github.com/negrel/paon/pdk/events"

// Listener define a wrapper around events.Listener for tree events.
type Listener struct {
	events.Handler

	bubbles bool
}

// NewListener returns a new Listener that wraps the given events.Listener.
// If bubbles is true, the listener will be called only if a bubble phase occur.
func NewListener(listener events.Handler, bubbles bool) Listener {
	return Listener{
		Handler: listener,
		bubbles: bubbles,
	}
}
