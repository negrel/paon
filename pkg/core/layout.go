package core

import (
	"fmt"

	"github.com/negrel/debugo"
)

// Layout define any widgets that can contain child
// widgets.
type Layout interface {
	Widget

	/*
	 * --- GETTERS & SETTERS ---
	 */

	// Children of the layout.
	Children() []Widget

	/*
	 * --- METHODS ---
	 */

	// Append the given widget to the child list,
	// if the widget already have a parent, remove the child from
	// the parent and adopt it. Panic
	// if widget is a parent of the layout.
	AppendChild(Widget)

	// ChildAt return the child widget
	// at the given index or nil.
	ChildAt(int) (Widget, bool)

	// Drop the given child widget.
	DropChild(Widget)

	// Insert the given child before the
	// the second given child. If the reference
	// widget is nil, the child is appended.
	InsertBefore(Widget, Widget)
}

var _ Layout = &LayoutCore{}

// LayoutCore is the core of every layout.
// It implements the common method for the Widget
// interface.
type LayoutCore struct {
	*WidgetCore

	children []Widget
}

// NewLayoutCore returns a new layout core.
func NewLayoutCore(name string) *LayoutCore {
	return &LayoutCore{
		WidgetCore: NewWidgetCore(name),
		children:   []Widget{},
	}
}

/*****************************************************
 ***************** Getters & Setters *****************
 *****************************************************/
// SECTION Getters & Setters

// Children returns the children widgets of the layout.
func (lc *LayoutCore) Children() []Widget {
	return lc.children
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// SECTION Methods

// AppendChild try to append the given child and return an
// error otherwise. If the child is somewhere else
func (lc *LayoutCore) AppendChild(child Widget) {
	// check that the child is not an ancestor of the layout
	debugo.Assert(func() (bool, string) {
		var node Widget = lc
		for node != nil {
			if node == child {
				return false, fmt.Sprintf("%v - %v can't append the child, the given child is an ancestor of the layout", lc.uid, lc.name)
			}

			node = node.Parent()
		}

		return true, ""
	}())

	// drop child parent if necessary
	if parent := child.Parent(); parent != nil {
		parent.DropChild(child)
	}

	lc.children = append(lc.children, child)
	child.adoptedBy(lc, len(lc.children)-1)
}

// ChildAt returns the child widget at the given slot.
func (lc *LayoutCore) ChildAt(slot int) (Widget, bool) {
	if slot < 0 || slot >= len(lc.children) {
		return nil, false
	}

	return lc.children[slot], true
}

// DropChild drop the given widget if it's direct
// child of the layout.
func (lc *LayoutCore) DropChild(child Widget) {
	slot := child.slot()

	// checking that the given widget is a child of the layout
	debugo.Assert(func() (bool, string) {
		if lc.children[slot] != child {
			return false, fmt.Sprintf("%v - %v can't drop the given widget %v because it's not a child", lc.uid, lc.name, child.Name())
		}

		return true, ""
	}())

	// removing parent/child link
	lc.children = append(lc.children[:slot-1], lc.children[slot+1:]...)
	child.abandoned()
}

// InsertBefore insert the given child before the second
// given child.
func (lc *LayoutCore) InsertBefore(child, reference Widget) {
	if reference == nil {
		lc.AppendChild(child)
	}

	before := reference.slot()
	lc.children = append(
		append(lc.children[:before-1], child),
		lc.children[before+1:]...,
	)
}
