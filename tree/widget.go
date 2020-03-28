package tree

// Widget are finished component.
type Widget struct {
	Node
	Drawable

	*BaseNode

	computedStyle  Style
	inheritedStyle Style
	classStyle     Style
	ownStyle       Style
}
