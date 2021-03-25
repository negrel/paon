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

	// IsAncestorOf return true if the given Node is a child of this.
	IsAncestorOf(child Node) bool

	// Return the first child Node of this.
	FirstChildNode() Node

	// Return the last child Node of this.
	LastChildNode() Node

	// Append the given child to the list of child Node. An error is returned
	// if the given child is an ancestor of this Node.
	AppendChildNode(newChild Node) error

	// Insert the given child before the given reference child Node. If the
	// reference is nil, the child is appended. An error is returned
	// if the given child is an ancestor of this Node or if the reference
	// is not a direct child of this.
	InsertBeforeNode(reference, newChild Node) error

	// Remove the given direct child Node of this. Return an error otherwise.
	RemoveChildNode(child Node) error
}

var _ ParentNode = &parentNode{}

type parentNode struct {
	*node

	ptr2Composite ParentNode
	firstChild    Node
	lastChild     Node
}

// NewParent returns a ParentNode Node with no child.
func NewParent() ParentNode {
	return newParent()
}

func newParent() *parentNode {
	p := &parentNode{
		node: newNode(),
	}
	p.ptr2Composite = p

	return p
}

// NewCompositeParent returns a new ParentNode object with no child. Unlike NewParent,
// this function returns a ParentNode object that can be used in composite struct. The ptr argument
// must be a pointer to the composite struct. Therefore, inserted Node will receives a pointer to the
// composite struct as a parent.
func NewCompositeParent(ptr ParentNode) ParentNode {
	return newCompositeParent(ptr)
}

func newCompositeParent(ptr ParentNode) *parentNode {
	return &parentNode{
		node:          newNode(),
		ptr2Composite: ptr,
	}
}

func (pn *parentNode) IsAncestorOf(child Node) bool {
	if child == nil {
		return false
	}

	return child.IsDescendantOf(pn)
}

func (pn *parentNode) FirstChildNode() Node {
	return pn.firstChild
}

func (pn *parentNode) LastChildNode() Node {
	return pn.lastChild
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
	pn.prepareChildForInsertion(newChild)

	if pn.lastChild != nil {
		pn.lastChild.setNextNode(newChild)
		newChild.setPreviousNode(pn.lastChild)
	} else {
		pn.firstChild = newChild
	}

	pn.lastChild = newChild

	newChild.setParentNode(pn.ptr2Composite)
}

func (pn *parentNode) ensurePreInsertionValidity(child Node) error {
	// check if child is not a parentNode of pn
	if parentNode, isParent := child.(ParentNode); isParent {
		if parentNode.IsAncestorOf(pn) {
			return errors.New("child contains the parentNode")
		}
	}

	return nil
}

func (pn *parentNode) prepareChildForInsertion(newChild Node) {
	if parent := newChild.ParentNode(); parent != nil {
		err := parent.RemoveChildNode(newChild)
		assert.Nil(err)
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
	pn.prepareChildForInsertion(newChild)

	if previous := reference.PreviousNode(); previous != nil {
		previous.setNextNode(newChild)
		newChild.setPreviousNode(previous)
	} else {
		pn.firstChild = newChild
	}
	newChild.setNextNode(reference)
	reference.setPreviousNode(newChild)

	newChild.setParentNode(pn.ptr2Composite)
}

func (pn *parentNode) RemoveChildNode(child Node) error {
	assert.NotNil(child, "child must be non-nil to be removed")

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

	return nil
}
