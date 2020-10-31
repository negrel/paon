//go:generate mockgen -source ./widget.go -destination mock/widget.go Widget

package widgets

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/negrel/paon/internal/events"
	"github.com/negrel/paon/internal/render"
)

// Widget define any TUI components in the widget tree.
type Widget interface {
	fmt.Stringer
	render.Renderable
	events.EventTarget

	// Return the name of the widget.
	Name() string
	// Return the Unique Universal IDentifier of the widget.
	UUID() uuid.UUID

	// Parent return the parent layout.
	Parent() Layout
	adoptedBy(parent Layout)

	isAttached() bool
	Window() Widget
}
