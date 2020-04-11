package widget

// Layout is a widget that lay out is childrens widgets.
type Layout interface {
	Widget

	// Abandon the given widget. If the given child is not a direct child
	// of the layout, an error is returned.
	Abandon(Widget) error

	// AppendChild method append the given child and
	// trigger a reflow.
	AppendChild(Widget) error

	// IndexOf return the index position of the given child, or an error
	// if the given child is not a direct children of the layout.
	// For example, in a row, the index N corresponds to the N th
	// child from the left. For a column, the 0 index correspond to the
	// top element.
	IndexOf(Widget) (int, error)

	// Reflow force a repaint of all the child
	Reflow()
}
