package widgets

import (
	"errors"
	"fmt"

	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/debuggo/pkg/log"
	"github.com/negrel/paon/pkg/pdk/id"
)

// Layout is a Widget that can contain child Widget.
type Layout interface {
	Widget

	// Returns the first child Widget of this Layout.
	FirstChild() Widget
	// Returns the last child Widget of this Layout.
	LastChild() Widget

	// Adds the given Widget at the end of the child list.
	// An error is returned if the given Widget is an ancestor
	// of this Layout.
	// This function panic if a nil value is given.
	AppendChild(child Widget) error

	// Inserts the given Widget before the reference Widget in the child list.
	// If the reference is nil the child is appended.
	// An error is returned if the given Widget is an ancestor
	// of this Layout.
	// This function panic if a nil child value is given.
	InsertBefore(reference, child Widget) error

	// Removes the given Widget of the child list of this Layout.
	// An error is returned if the Widget is not a direct child of
	// this Layout.
	RemoveChild(child Widget) error

	// IsAncestorOf return true if the given Layout is a child of this Layout.
	IsAncestorOf(child Widget) bool
}

var _ Layout = &layout{}

type layout struct {
	*widget

	firstChild Widget
	lastChild  Widget
}

func NewLayout(name string, opts ...Option) Layout {
	return newLayout(name, opts...)
}

func newLayout(name string, opts ...Option) *layout {
	l := &layout{
		widget: newWidget(name, opts...),
	}

	return l
}

// ID implements the id.Identifiable interface.
func (l *layout) ID() id.ID {
	return l.widget.ID()
}

// LastChild implements the Layout interface.
func (l *layout) FirstChild() Widget {
	return l.firstChild
}

// LastChild implements the Layout interface.
func (l *layout) LastChild() Widget {
	return l.lastChild
}

func (l *layout) AppendChild(newChild Widget) (err error) {
	assert.NotNil(newChild, "child must be non-nil to be appended")

	if err = l.ensurePreInsertionValidity(newChild); err != nil {
		return fmt.Errorf("can't append child, %v", err)
	}
	l.appendChild(newChild)

	return nil
}

func (l *layout) appendChild(newChild Widget) {
	l.prepareChildForInsertion(newChild)

	if l.lastChild != nil {
		l.lastChild.setNext(newChild)
		newChild.setPrevious(l.lastChild)
	} else {
		l.firstChild = newChild
	}

	l.lastChild = newChild

	newChild.setParent(l)
}

func (l *layout) ensurePreInsertionValidity(child Widget) error {
	// check if child is not a parentWidget of pn
	if layout, isLayout := child.(Layout); isLayout {
		if layout.IsAncestorOf(l) {
			return errors.New("child contains the parentWidget")
		}
	}

	return nil
}

func (l *layout) prepareChildForInsertion(newChild Widget) {
	if parent := newChild.Parent(); parent != nil {
		err := parent.RemoveChild(newChild)
		assert.Nil(err)
	}
	assert.Nil(newChild.Root())
	assert.Nil(newChild.Parent())
	assert.Nil(newChild.Previous())
	assert.Nil(newChild.Next())
}

func (l *layout) InsertBefore(reference, newChild Widget) error {
	assert.NotNil(newChild, "child must be non-nil to be appended")

	// InsertBefore(nil, node) is equal to AppendChild(node)
	if reference == nil {
		return l.AppendChild(newChild)
	}
	if referenceIsNotChild := !l.IsSame(reference.Parent()); referenceIsNotChild {
		return errors.New("can't insert child, the given reference is not a child of this node")
	}

	if err := l.ensurePreInsertionValidity(newChild); err != nil {
		return fmt.Errorf("can't insert child, %v", err)
	}

	// newChild and reference are the same
	if reference == newChild {
		log.Debugln("can't insert child before itself, reference is now child next sibling")
		reference = newChild.Next()
		if reference == nil {
			log.Debugln("can't insert before a nil reference, appending the child")
			l.appendChild(newChild)
			return nil
		}
	}

	l.insertBefore(reference, newChild)
	return nil
}

func (l *layout) insertBefore(reference, newChild Widget) {
	l.prepareChildForInsertion(newChild)

	if previous := reference.Previous(); previous != nil {
		previous.setNext(newChild)
		newChild.setPrevious(previous)
	} else {
		l.firstChild = newChild
	}
	newChild.setNext(reference)
	reference.setPrevious(newChild)

	newChild.setParent(l)
}

func (l *layout) RemoveChild(child Widget) error {
	assert.NotNil(child, "child must be non-nil to be removed")

	// if not a child of pn
	if !l.IsSame(child.Parent()) {
		return errors.New("can't remove child, the node is not a child of this node")
	}

	// Removing siblings link
	next := child.Next()
	prev := child.Previous()
	if next != nil {
		child.setNext(nil)
		next.setPrevious(prev)
	} else {
		l.lastChild = prev
	}

	if prev != nil {
		child.setPrevious(nil)
		prev.setNext(next)
	} else {
		l.firstChild = next
	}
	// Removing parentWidget & root link
	child.setParent(nil)

	return nil
}

func (l *layout) IsAncestorOf(widget Widget) bool {
	if widget == nil {
		return false
	}

	return widget.IsDescendantOf(l)
}
