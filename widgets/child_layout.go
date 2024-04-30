package widgets

import (
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/geometry"
)

// ChildLayout define position and size of a child widget.
type ChildLayout struct {
	Widget Widget
	// Bounds relative to layout origin.
	Bounds geometry.Rectangle
}

// Reset resets child layout.
func (cl *ChildLayout) Reset() {
	cl.Widget = nil
	cl.Bounds = geometry.Rectangle{}
}

// ChildrenLayout contains position and size of children.
type ChildrenLayout struct {
	Origin  geometry.Vec2D
	layouts []ChildLayout
}

// Size returns number of ChildLayout stored.
func (cl *ChildrenLayout) Size() int {
	return len(cl.layouts)
}

// Get returns ChildLayout relative to parent origin.
func (cl *ChildrenLayout) Get(i int) ChildLayout {
	childLayout := cl.layouts[i]
	childLayout.Bounds = childLayout.Bounds.MoveBy(cl.Origin)

	return childLayout
}

// Append appends the given ChildLayout.
func (cl *ChildrenLayout) Append(child ChildLayout) {
	cl.layouts = append(cl.layouts, child)
}

// Reset resets stored child layouts.
func (cl *ChildrenLayout) Reset() {
	cl.Origin = geometry.Vec2D{}
	cl.layouts = cl.layouts[:0]
}

// Draw draws all children on the given surface.
func (cl *ChildrenLayout) Draw(srf draw.Surface) {
	for i := 0; i < cl.Size(); i++ {
		childLayout := cl.Get(i)
		childLayout.Widget.Draw(draw.NewSubSurface(srf, childLayout.Bounds))
	}
}
