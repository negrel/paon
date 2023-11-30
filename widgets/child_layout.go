package widgets

import (
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/tree"
)

// ChildLayout define position and size of a child Node.
type ChildLayout struct {
	// Node associated to rectangle.
	Node *tree.Node
	// Bounds relative to layout origin.
	Bounds geometry.Rectangle
}

// ChildrenLayout contains position and size of children Node.
type ChildrenLayout struct {
	origin  geometry.Vec2D
	layouts []ChildLayout
}

// Size returns number of ChildLayout stored.
func (cl *ChildrenLayout) Size() int {
	return len(cl.layouts)
}

// Get returns ChildLayout relative to parent origin.
func (cl *ChildrenLayout) Get(i int) ChildLayout {
	childLayout := cl.layouts[i]
	childLayout.Bounds = childLayout.Bounds.MoveBy(cl.origin)

	return childLayout
}

// Append appends the given ChildLayout.
func (cl *ChildrenLayout) Append(child ChildLayout) {
	cl.layouts = append(cl.layouts, child)
}

func (cl *ChildrenLayout) reset() {
	cl.origin = geometry.Vec2D{}
	cl.layouts = cl.layouts[:0]
}
