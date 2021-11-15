package backend

import (
	"github.com/negrel/paon/pdk/draw"
	pdkevents "github.com/negrel/paon/pdk/events"
)

// Terminal define a generic terminal in raw mode used for rendering and event loops.
type Terminal interface {
	draw.Surface

	// Clear clears the entire console surface and replace cell by empty black one.
	Clear()

	// Flush flushes the cell buffer.
	Flush()

	// Start initializes the console for use. This starts the event loop and rendering.
	Start(chan<- pdkevents.Event) error

	// Stop deinitializes the console.
	Stop()
}
