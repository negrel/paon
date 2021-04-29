package tree

import (
	"errors"
	"fmt"

	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/debuggo/pkg/log"
)

type parentNode struct {
	*node

	firstChild Node
	lastChild  Node
}

// NewNode returns a parent Node with no child.
func NewNode(data interface{}) Node {
	return newNode(data)
}

func newNode(data interface{}) *parentNode {
	p := &parentNode{
		node: newLeafNode(data),
	}

	return p
}

func (pn *parentNode) IsAncestorOf(child Node) bool {
	if child == nil {
		return false
	}

	return child.IsDescendantOf(pn)
}

func (pn *parentNode) FirstChild() Node {
	return pn.firstChild
}

func (pn *parentNode) LastChild() Node {
	return pn.lastChild
}

func (pn *parentNode) AppendChild(newChild Node) (err error) {
	assert.NotNil(newChild, "child must be non-nil to be appended")

	if err = pn.ensurePreInsertionValidity(newChild); err != nil {
		return fmt.Errorf("can't append child, %v", err)
	}
	pn.appendChild(newChild)

	return nil
}

func (pn *parentNode) appendChild(newChild Node) {
	pn.prepareChildForInsertion(newChild)

	if pn.lastChild != nil {
		pn.lastChild.SetNext(newChild)
		newChild.SetPrevious(pn.lastChild)
	} else {
		pn.firstChild = newChild
	}

	pn.lastChild = newChild

	newChild.SetParent(pn)
}

func (pn *parentNode) ensurePreInsertionValidity(child Node) error {
	if child == nil {
		return errors.New("child is nil")
	}

	// check if child is not a parentNode of pn
	if child.IsAncestorOf(pn) {
		return errors.New("child contains the parentNode")
	}

	return nil
}

func (pn *parentNode) prepareChildForInsertion(newChild Node) {
	if parent := newChild.Parent(); parent != nil {
		err := parent.RemoveChild(newChild)
		assert.Nil(err)
	}
	assert.Nil(newChild.Root())
	assert.Nil(newChild.Parent())
	assert.Nil(newChild.Previous())
	assert.Nil(newChild.Next())
}

func (pn *parentNode) InsertBefore(reference, newChild Node) error {
	assert.NotNil(newChild, "child must be non-nil to be appended")

	// InsertBeforeNode(nil, node) is equal to AppendChildNode(node)
	if reference == nil {
		return pn.AppendChild(newChild)
	}
	if referenceIsNotChild := !pn.IsSame(reference.Parent()); referenceIsNotChild {
		return errors.New("can't insert child, the given reference is not a child of this node")
	}

	if err := pn.ensurePreInsertionValidity(newChild); err != nil {
		return fmt.Errorf("can't insert child, %v", err)
	}

	// newChild and reference are the same
	if reference == newChild {
		log.Debugln("can't insert child before itself, reference is now child next sibling")
		reference = newChild.Next()
		if reference == nil {
			log.Debugln("can't insert before a nil reference, appending the child")
			pn.appendChild(newChild)
			return nil
		}
	}

	pn.insertBefore(reference, newChild)
	return nil
}

func (pn *parentNode) insertBefore(reference, newChild Node) {
	pn.prepareChildForInsertion(newChild)

	if previous := reference.Previous(); previous != nil {
		previous.SetNext(newChild)
		newChild.SetPrevious(previous)
	} else {
		pn.firstChild = newChild
	}
	newChild.SetNext(reference)
	reference.SetPrevious(newChild)

	newChild.SetParent(pn)
}

func (pn *parentNode) RemoveChild(child Node) error {
	assert.NotNil(child, "child must be non-nil to be removed")

	// if not a child of pn
	if !pn.IsSame(child.Parent()) {
		return errors.New("can't remove child, the node is not a child of this node")
	}

	// Removing siblings link
	next := child.Next()
	prev := child.Previous()
	if next != nil {
		child.SetNext(nil)
		next.SetPrevious(prev)
	} else {
		pn.lastChild = prev
	}

	if prev != nil {
		child.SetPrevious(nil)
		prev.SetNext(next)
	} else {
		pn.firstChild = next
	}
	// Removing parentNode & root link
	child.SetParent(nil)

	return nil
}
