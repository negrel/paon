package tree

// Node ...
type Node interface {
	// NextSibling return the next sibling element.
	NextSibling() Node

	// Parent return the parent node.
	Parent() *Layout

	// PreviousSibling return the previous sibling element.
	PreviousSibling() Node
}

var _ Node = &BaseNode{}

// BaseNode are minimal node.
type BaseNode struct {
	Node

	parent *Layout
}

/*****************************************************
 ********************* INTERFACE *********************
 *****************************************************/
// ANCHOR Interface

// NODE

// Parent return the widget parent node.
func (w *BaseNode) Parent() *Layout {
	return w.parent
}

// NextSibling return the next sibling node.
func (w *BaseNode) NextSibling() Node {
	var index int = w.parent.IndexOf(w)

	if index == -1 {
		return nil
	}

	return w.parent.Item(index + 1)
}

// PreviousSibling return the previous sibling node.
func (w *BaseNode) PreviousSibling() Node {
	var index int = w.parent.IndexOf(w)

	if index == -1 {
		return nil
	}

	return w.parent.Item(index - 1)
}
