package paon

import (
	"github.com/negrel/debuggo/pkg/assert"

	"github.com/negrel/paon/internal/render"
	"github.com/negrel/paon/internal/utils"
	"github.com/negrel/paon/internal/widgets"
)

var _ widgets.Widget = &root{}

type root struct {
	*widgets.Node
	children widgets.Widget
}

func newRoot(child widgets.Widget) *root {
	assert.NotNil(child, "child must be non-nil")

	return &root{
		Node:     widgets.NewNodeWidget("paon_root"),
		children: child,
	}
}

func (r *root) Render(rect utils.Rectangle) render.Patch {
	return r.children.Render(rect)
}
