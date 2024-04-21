package widgets

import "github.com/negrel/paon/events"

// Event define an event within a widget tree.
type Event interface {
	events.Event

	// Target returns targeted widget by the event.
	Target() Widget
}
