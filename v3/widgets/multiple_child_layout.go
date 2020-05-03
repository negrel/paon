package widgets

// MultipleChildLayout are layout that contain
// multiple child.
type MultipleChildLayout interface {
	Widget

	// -- GETTERS & SETTERS --

	// Children of the layout.
	Children() []Widget

	// -- METHODS --

	// Append the given widget to the child list,
	// if the widget already have a parent, remove the child from
	// the parent and adopt it. Panic
	// if widget is a parent of the layout.
	AppendChild(Widget)

	// ChildAt return the child widget
	// at the given index or nil.
	ChildAt(int) Widget

	// Drop the given child widget.
	DropChild(Widget)

	// Insert the given child before the
	// the second before child.
	InsertBefore(Widget, Widget)
}

// MultipleChildCore define the core of every
// multiple child layouts.
type MultipleChildCore struct {
	*Core

	Children []Widget
}


