package widgets

import "github.com/negrel/ginger/v2/render"

// Widget is an element in the node tree that can be paint.
type Widget interface {
	render.Rendable

	// ----------- GETTERS & SETTERS -----------

	// Name return the name of the widget.
	Name() string

	// Attached return wether or not this node is in a tree
	// whose root is attached to something
	Attached() bool

	// Attach the node to the given owner.
	Attach(Layout)

	// Detach the node from the owner
	Detach()

	// Owner return the root of this entire subtree
	// (nil if unattached).
	Owner() Layout

	// Parent return the parent node.
	Parent() Layout
	setParent(Layout)
}
