package widgets

// Widget is an element in the node tree that can be paint.
type Widget interface {
	Rendable

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
