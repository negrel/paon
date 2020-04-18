package widgets

// Layout are widgets that lay out is children.
type Layout interface {
	Widget

	// ----------- METHODS -----------

	// AdoptChild method mark the given node as child of
	// this node.
	AdoptChild(Widget)

	// DropChild method disconnect the given child from
	// this node.
	DropChild(Widget)
}
