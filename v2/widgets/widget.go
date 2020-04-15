package widgets

import (
	"github.com/negrel/ginger/v2/rendering"
	"github.com/negrel/ginger/v2/widgets/node"
)

// Widget is an element in the node tree that can be paint.
type Widget interface {
	node.Leaf
	rendering.Drawable
}
