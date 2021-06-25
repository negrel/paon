package backend

import (
	"github.com/negrel/paon/pdk/draw"
	"github.com/negrel/paon/events"
)

// Console define a generic console used for rendering and event loops.
type Console interface {
	draw.Canvas

	// Clear clears the entire console surface and replace cell by empty black one.
	Clear()

	// Flush flushes the cell buffer.
	Flush()

	// Start initializes the console for use. This starts the event loop and rendering.
	Start(chan<- events.Event) error

	// Stop deinitializes the console.
	Stop()
}
