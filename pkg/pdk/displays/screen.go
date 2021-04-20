package displays

import (
	"context"

	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pkg/pdk/draw"
	"github.com/negrel/paon/pkg/pdk/events"
	"github.com/negrel/paon/pkg/pdk/render"
)

// Screen define a screen.
type Screen interface {
	events.Target
	render.Surface

	// Start starts the screen so it can produce events, display things and more.
	Start(ctx context.Context) error

	// Stop clears the screen and stop listening to events.
	Stop()

	// Canvas returns a draw.Canvas object.
	Canvas() draw.Canvas

	// Bounds returns the current Screen bounds.
	Bounds() geometry.Rectangle
}
