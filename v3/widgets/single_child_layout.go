package widgets

import (
	"fmt"

	"github.com/negrel/debugo"
)

// SingleChildLayout define common methods
// for layout that have a single child.
type SingleChildLayout interface {
	Widget

	// -- GETTERS & SETTERS --

	// Child return the first (or only) child
	// widget.
	Child() Widget

	// -- METHODS --

	// Adopt the given widget. Panic if the given
	// child already have a parent, or if the
	// given child is an ancestor of the widget.
	AdoptChild(Widget)

	// Drop the single child widget.
	DropSingleChild()
}

var _ MultipleChildLayout = &SingleChildCore{}

// SingleChildCore define the core of every
// single child layouts.
type SingleChildCore struct {
	*Core

	child Widget
}

/*****************************************************
 ***************** Getters & Setters *****************
 *****************************************************/
// ANCHOR Getters & Setters

// Child implements SingleChildLayout interface.
func (scc *SingleChildCore) Child() Widget {
	return scc.child
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// AdoptChild implements the SingleChildLayout interface.
func (scc *SingleChildCore) AdoptChild(child Widget) {
	// Checking child ready to be adopted
	debugo.AssertF(func() (bool, string) {
		if child == nil {
			msg := "can't adopt the child. (child is nil or child parent is not nil)\n"
			l1 := fmt.Sprintf(" ├─ Child %v: %+v\n", child.Name(), child)
			l2 := fmt.Sprintf(" └─ Child parent %v: %+v", child.Parent().Name(), child.Parent())

			return false, fmt.Sprint(msg, l1, l2)
		}

		return true, ""
	})

	var node Widget = scc
	for node != nil {

		debugo.AssertF(func() (bool, string) {
			if node == child {
				msg := "can't adopt child, child is an ancestor of the layout\n"
				l1 := fmt.Sprintf(" └─ Child %v: %+v", child.Name(), child)
				return false, fmt.Sprint(msg, l1)
			}

			return true, ""
		})

		node = node.Parent()
	}

	child.adoptedBy(scc, 0)
	scc.child = child
}

// DropSingleChild implements the SingleChildLayout interface.
func (scc *SingleChildCore) DropSingleChild() {
	scc.child.abandoned()
	scc.child = nil
}

/*****************************************************
 ********************* Interface *********************
 *****************************************************/
// ANCHOR Interface

// AppendChild implements MultipleChildLayout interface.
//
// SingleChildLayout.AdoptChild method is more appropriated
// for SingleChildLayout.
func (scc *SingleChildCore) AppendChild(child Widget) {
	if scc.child != nil {
		scc.AdoptChild(child)
	}
}

// ChildAt implements MultipleChildLayout interface.
//
// SingleChildLayout.Child method is more appropriated
// for SingleChildLayout.
func (scc *SingleChildCore) ChildAt(i int) Widget {
	if i == 0 {
		return scc.child
	}

	return nil
}

// Children implements MultipleChildLayout interface.
//
// SingleChildLayout.Child method is more appropriated
// for SingleChildLayout.
func (scc *SingleChildCore) Children() []Widget {
	return []Widget{scc.child}
}

// DropChild implements MultipleChildLayout interface.
//
// SingleChildLayout.DropSingleChild method is more
// appropriated for SingleChildLayout.
func (scc *SingleChildCore) DropChild(child Widget) {
	if child == scc.child {
		scc.DropSingleChild()
	}
}

// InsertBefore implements MultipleChildLayout interface.
//
// SingleChildLayout.AdoptChild method is more appropriated
// for SingleChildLayout.
func (scc *SingleChildCore) InsertBefore(child Widget, _ Widget) {
	if scc.child != nil {
		scc.AdoptChild(child)
	}
}
