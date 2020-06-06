package ui

// Layout define any widgets that can containt child
// widgets.
type Layout interface {
	// Children returns all the child widgets
	Children() []Widget

	// childAt returns the child widget at the given slot.
	ChildAt(slot uint) Widget
}
