//go:generate mockgen -source ./widget.go -destination mock/widget.go Widget

package widgets

import (
	"github.com/google/uuid"
)

// Widget define any TUI components in the widget tree.
type Widget interface {
	Name() string
	UUID() uuid.UUID

	Parent() Layout
	setParent(parent Layout)

	String() string
}
