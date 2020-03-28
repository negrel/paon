package tree

import (
	s "github.com/negrel/ginger/style"
)

// View represent a graphical component.
type View interface {
	// Parent return the parent of the view.
	Parent() ViewParent

	// Change view parent and remove it from is previous
	// parent
	AdoptedBy(ViewParent)

	// Abandoned remove the parent
	Abandoned()

	// PreviousSibling return the previous view sibling.
	PreviousSibling() View

	// NextSibling return the next viewsibling.
	NextSibling() View
}

type _view struct {
	View

	parent        ViewParent
	computedStyle s.Style
	ownStyle      s.Style
}

/*****************************************************
 ***************** GETTERS & SETTERS *****************
 *****************************************************/
// ANCHOR Getters & setter

func (v *_view) Parent() ViewParent {
	return v.parent
}

func (v *_view) AdoptedBy(vp ViewParent) {
	v.parent.RemoveChild(v)
	v.parent = vp
}

func (v *_view) Abandoned() {
	v.parent = nil
}

func (v *_view) PreviousSibling() View {
	index, err := v.parent.IndexOf(v)
	if err != nil {

	}

	return v.parent.Item(index - 1)
}

func (v *_view) NextSibling() View {
	index, err := v.parent.IndexOf(v)
	if err != nil {

	}

	return v.parent.Item(index + 1)
}

/*****************************************************
 ********************** METHODS **********************
 *****************************************************/
// ANCHOR Methods
