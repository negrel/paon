package tree

import (
	"errors"
	"fmt"

	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/debuggo/pkg/log"
)

// ParentNode define a Node that can have children Node.
type ParentNode interface {
	Node

	isAncestorOf(child Node) bool

	FirstChild() Node
	LastChild() Node

	AppendChild(newChild Node) error
	InsertBefore(reference, newChild Node) error
	RemoveChild(child Node) error
}

var _ ParentNode = &parentNode{}

type parentNode struct {
	*node

	firstChild Node
	lastChild  Node
}

// NewParent returns a ParentNode Node with zero child.
func NewParent(name string) ParentNode {
	return newParent(name)
}

func newParent(name string) *parentNode {
	return &parentNode{
		node: newNode(name),
	}
}

func (pn *parentNode) isAncestorOf(child Node) bool {
	return child.isDescendantOf(pn)
}

func (pn *parentNode) FirstChild() Node {
	return pn.firstChild
}

func (pn *parentNode) LastChild() Node {
	return pn.lastChild
}

func (pn *parentNode) adopt(child Node) {
	child.setParent(pn)
	child.setRoot(pn.root)
	if pn.root != nil {
		pn.root.register(child)
	}
}

func (pn *parentNode) AppendChild(newChild Node) (err error) {
	assert.NotNil(newChild, "child must be non-nil to be appended")
	log.Debugln("appending", newChild, "in", pn)

	if err = pn.ensurePreInsertionValidity(newChild); err != nil {
		return fmt.Errorf("can't append child, %v", err)
	}
	pn.prepareChildForInsertion(newChild)

	if pn.lastChild != nil {
		pn.lastChild.setNext(newChild)
		newChild.setPrevious(pn.lastChild)
	} else {
		pn.firstChild = newChild
	}

	pn.lastChild = newChild
	pn.adopt(newChild)

	return nil
}

func (pn *parentNode) ensurePreInsertionValidity(child Node) error {
	// check if child is not a parentNode of pn
	if parentNode, isParent := child.(ParentNode); isParent {
		if parentNode.isAncestorOf(pn) {
			return errors.New("child contains the parentNode")
		}
	}

	return nil
}

func (pn *parentNode) prepareChildForInsertion(newChild Node) {
	if parent := newChild.Parent(); parent != nil {
		err := parent.RemoveChild(newChild)
		assert.Nil(err, err)
	}
	assert.Nil(newChild.Root())
	assert.Nil(newChild.Parent())
	assert.Nil(newChild.Previous())
	assert.Nil(newChild.Next())
}

func (pn *parentNode) InsertBefore(reference, newChild Node) error {
	assert.NotNil(newChild, "child must be non-nil to be appended")
	log.Debugln("inserting", newChild, "before", reference, "in", pn)

	// InsertBefore(nil, node) is equal to AppendChild(node)
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
			return pn.AppendChild(newChild)
		}
	}

	pn.prepareChildForInsertion(newChild)

	if previous := reference.Previous(); previous != nil {
		previous.setNext(newChild)
		newChild.setPrevious(previous)
	} else {
		pn.firstChild = newChild
	}
	newChild.setNext(reference)
	reference.setPrevious(newChild)

	pn.adopt(newChild)

	return nil
}

func (pn *parentNode) RemoveChild(child Node) error {
	assert.NotNil(child, "child must be non-nil to be removed")
	log.Debugln("removing", child, "from", "pn")

	// if not a child of pn
	if !pn.IsSame(child.Parent()) {
		return errors.New("can't remove child, the node is not a child of this node")
	}

	// Removing siblings link
	next := child.Next()
	prev := child.Previous()
	if next != nil {
		child.setNext(nil)
		next.setPrevious(prev)
	} else {
		pn.lastChild = prev
	}

	if prev != nil {
		child.setPrevious(nil)
		prev.setNext(next)
	} else {
		pn.firstChild = next
	}
	// Removing parentNode & root link
	child.setParent(nil)
	child.setRoot(nil)
	if pn.isConnected() {
		pn.root.unregister(child)
	}

	return nil
}
