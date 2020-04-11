package widget

import "github.com/gdamore/tcell"

// EventListener are widget that can handle an event
type EventListener interface {
	// Handle the event if the handler has consumed the event,
	// it should return true. False otherwise.
	HandleEvent(tcell.Event) bool
}
