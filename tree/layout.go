package tree

import (
	e "github.com/negrel/ginger/errors"
)

var _ Node = &Layout{}

// Layout are container for widgets.
type Layout struct {
	*BaseNode

	childNodes  []Node
	layoutStyle Style
}

/*****************************************************
 ********************* INTERFACE *********************
 *****************************************************/
// ANCHOR Interface

/*****************************************************
 ***************** GETTERS & SETTERS *****************
 *****************************************************/
// ANCHOR Getters & setter

// FirstChild return the first child node.
func (l *Layout) FirstChild() Node {
	if len := len(l.childNodes); len > 0 {
		return l.childNodes[0]
	}

	return nil
}

// Item return the child node at the given index.
func (l *Layout) Item(index int) Node {
	return l.childNodes[index]
}

// LastChild return the last child node.
func (l *Layout) LastChild() Node {
	if len := len(l.childNodes); len > 0 {
		return l.childNodes[len-1]
	}

	return nil
}

/*****************************************************
 ********************** METHODS **********************
 *****************************************************/
// ANCHOR Methods

// AppendChild add a child as last child
func (l *Layout) AppendChild(node Node) error {
	return l.InsertBefore(node, nil)
}

// Contains return wether or not it contain the
// given node.
func (l *Layout) Contains(node Node) bool {
	// Checking first class node child
	for _, child := range l.childNodes {
		if child == node {
			return true
		}
	}

	for _, child := range l.childNodes {
		if lchild, isLayout := child.(Layout); isLayout {
			if lchild.Contains(node) {
				return true
			}
		}
	}

	return false
}

// IndexOf return the index of the given node if
// it is a children, otherwise it return -1.
func (l *Layout) IndexOf(node Node) int {
	for i, v := range l.childNodes {
		// Comparing pointer
		if v == node {
			return i
		}
	}

	return -1
}

// InsertBefore insert the given node before the given
// reference child.
func (l *Layout) InsertBefore(node, referenceChild Node) error {
	// Checking that node to insert is not the parent
	// of the layout itself
	if lnode, isLayout := node.(Layout); isLayout {
		if lnode.Contains(l) {
			return e.HierarchyRequestError("The new child contains the parent.")
		}
	}

	// Checking that node to insert is not the layout itself
	if l == node {
		return e.HierarchyRequestError("The new child contains the parent.")
	}

	// Inserting
	var index int = l.IndexOf(referenceChild)

	if index == -1 {
		index = len(l.childNodes)
	}

	var before []Node = append(l.childNodes[:index], node)
	var after []Node = l.childNodes[index:]

	l.childNodes = append(before, after...)

	// Changing node parent and removing from parent the child
	bnode, ok := node.(*BaseNode)
	if ok {
		bnode.parent = l
	} else {
		return e.TypeError("The given Node is not a node.")
	}

	return nil
}

// RemoveChild remove the given child node
func (l *Layout) RemoveChild(node Node) {
	var index int = l.IndexOf(node)

	// If found
	if index != -1 {
		child := l.childNodes[index].(*BaseNode)

		child.parent = nil
	}
}
