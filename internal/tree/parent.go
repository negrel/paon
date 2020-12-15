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

	FirstChildNode() Node
	LastChildNode() Node

	AppendChildNode(newChild Node) error
	InsertBeforeNode(reference, newChild Node) error
	RemoveChildNode(child Node) error
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

func (pn *parentNode) FirstChildNode() Node {
	return pn.firstChild
}

func (pn *parentNode) LastChildNode() Node {
	return pn.lastChild
}

func (pn *parentNode) adopt(child Node) {
	child.setParentNode(pn)
	if pn.root != nil {
		pn.root.register(child)
	}
}

func (pn *parentNode) AppendChildNode(newChild Node) (err error) {
	assert.NotNil(newChild, "child must be non-nil to be appended")

	if err = pn.ensurePreInsertionValidity(newChild); err != nil {
		return fmt.Errorf("can't append child, %v", err)
	}
	pn.appendChildNode(newChild)

	return nil
}

func (pn *parentNode) appendChildNode(newChild Node) {
	log.Debugln("appending", newChild, "in", pn)

	pn.prepareChildForInsertion(newChild)

	if pn.lastChild != nil {
		pn.lastChild.setNextNode(newChild)
		newChild.setPreviousNode(pn.lastChild)
	} else {
		pn.firstChild = newChild
	}

	pn.lastChild = newChild
	pn.adopt(newChild)
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
	if parent := newChild.ParentNode(); parent != nil {
		err := parent.RemoveChildNode(newChild)
		assert.Nil(err, err)
	}
	assert.Nil(newChild.RootNode())
	assert.Nil(newChild.ParentNode())
	assert.Nil(newChild.PreviousNode())
	assert.Nil(newChild.NextNode())
}

func (pn *parentNode) InsertBeforeNode(reference, newChild Node) error {
	assert.NotNil(newChild, "child must be non-nil to be appended")

	// InsertBeforeNode(nil, node) is equal to AppendChildNode(node)
	if reference == nil {
		return pn.AppendChildNode(newChild)
	}
	if referenceIsNotChild := !pn.IsSame(reference.ParentNode()); referenceIsNotChild {
		return errors.New("can't insert child, the given reference is not a child of this node")
	}

	if err := pn.ensurePreInsertionValidity(newChild); err != nil {
		return fmt.Errorf("can't insert child, %v", err)
	}

	// newChild and reference are the same
	if reference == newChild {
		log.Debugln("can't insert child before itself, reference is now child next sibling")
		reference = newChild.NextNode()
		if reference == nil {
			log.Debugln("can't insert before a nil reference, appending the child")
			pn.appendChildNode(newChild)
			return nil
		}
	}

	pn.insertBeforeNode(reference, newChild)
	return nil
}

func (pn *parentNode) insertBeforeNode(reference, newChild Node) {
	log.Debugln("inserting", newChild, "before", reference, "in", pn)
	pn.prepareChildForInsertion(newChild)

	if previous := reference.PreviousNode(); previous != nil {
		previous.setNextNode(newChild)
		newChild.setPreviousNode(previous)
	} else {
		pn.firstChild = newChild
	}
	newChild.setNextNode(reference)
	reference.setPreviousNode(newChild)

	pn.adopt(newChild)
}

func (pn *parentNode) RemoveChildNode(child Node) error {
	assert.NotNil(child, "child must be non-nil to be removed")
	log.Debugln("removing", child, "from", "pn")

	// if not a child of pn
	if !pn.IsSame(child.ParentNode()) {
		return errors.New("can't remove child, the node is not a child of this node")
	}

	// Removing siblings link
	next := child.NextNode()
	prev := child.PreviousNode()
	if next != nil {
		child.setNextNode(nil)
		next.setPreviousNode(prev)
	} else {
		pn.lastChild = prev
	}

	if prev != nil {
		child.setPreviousNode(nil)
		prev.setNextNode(next)
	} else {
		pn.firstChild = next
	}
	// Removing parentNode & root link
	child.setParentNode(nil)
	if pn.isConnected() {
		pn.root.unregister(child)
	}

	return nil
}
