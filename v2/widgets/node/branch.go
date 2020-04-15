package node

// Branch are node that can contain leaf node.
type Branch interface {
	Leaf

	// ----------- METHODS -----------

	// AdoptChild method mark the given node as child of
	// this node.
	AdoptChild(Leaf)

	// DropChild method disconnect the given child from
	// this node.
	DropChild(Leaf)
}
