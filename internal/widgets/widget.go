//go:generate mockgen -source ./widget.go -destination mock/widget.go Widget

package widgets

import (
	"github.com/google/uuid"

	"github.com/negrel/paon/internal/style"
)

// Widget define any TUI components in the widget tree.
type Widget interface {
	Name() string
	UUID() uuid.UUID

	Parent() Layout
	adoptedBy(parent Layout)
	// Position in parent children collection
	slot() int

	Style() *style.Node
	String() string
}
