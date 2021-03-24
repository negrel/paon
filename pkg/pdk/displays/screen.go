package displays

import (
	"github.com/negrel/paon/pkg/pdk/draw"
	"github.com/negrel/paon/pkg/pdk/events"
	"github.com/negrel/paon/pkg/pdk/render"
)

// Screen define a screen.
type Screen interface {
	events.Target
	render.Surface

	// Start starts the screen so it can produce events, display things and more.
	Start() error

	// Canvas returns a draw.Canvas object.
	Canvas() draw.Canvas

	// Stop stops the screen and makes it unusable.
	Stop()
}