package widgets

import "github.com/gdamore/tcell"

// EventHandler are widget that can handle
// tcell events.
type EventHandler interface {
	// Handle the given event and return true
	// if the handler have consumed the event.
	Handle(tcell.Event) bool
}
