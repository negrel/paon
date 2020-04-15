package node

// Leaf is a end branch node in the node tree.
type Leaf interface {
	// ----------- GETTERS & SETTERS -----------

	// Attached return wether or not this node is in a tree
	// whose root is attached to something
	Attached() bool

	// Attach the node to the given owner.
	Attach(Branch)

	// Detach the node from the owner
	Detach()

	// Owner return the root of this entire subtree
	// (nil if unattached).
	Owner() Branch

	// Parent return the parent node.
	Parent() Branch
	setParent(Branch)
}
