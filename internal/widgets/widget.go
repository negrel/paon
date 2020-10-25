//go:generate mockgen -source ./widget.go -destination mock/widget.go Widget

package widgets

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/negrel/paon/internal/render"
	"github.com/negrel/paon/internal/style"
)

// Widget define any TUI components in the widget tree.
type Widget interface {
	fmt.Stringer
	render.Renderable

	// Return the name of the widget.
	Name() string
	// Return the Unique Universal IDentifier of the widget.
	UUID() uuid.UUID

	// Parent return the parent layout.
	Parent() Layout
	adoptedBy(parent Layout)
	// Position in parent children collection
	slot() int

	// Return the style of the widget
	Style() *style.Node
}
