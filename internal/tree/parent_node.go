package tree

import (
	"errors"

	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/debuggo/pkg/log"
)

type ParentNode interface {
	Node

	FirstChild() ChildNode
	setFirstChild(ChildNode)
	LastChild() ChildNode
	setLastChild(ChildNode)

	AppendChild(Node) (ChildNode, error)
	InsertBefore(reference Node, newChild Node) (ChildNode, error)
	RemoveChild(Node) error
}

var _ ParentNode = &parentNode{}

type parentNode struct {
	Node

	firstChild ChildNode
	lastChild  ChildNode
}

func newParentNode(nodeType NodeType) *parentNode {
	assert.NotEqual(nodeType, TextNodeType, "TextNode is not a valid ParentNode type")

	return &parentNode{
		Node: &node{
			nodeType:    nodeType,
			isContainer: true,
		},
	}
}

func (pn *parentNode) isSame(other Node) bool {
	return pn == other
}

func (pn *parentNode) FirstChild() ChildNode {
	return pn.firstChild
}

func (pn *parentNode) setFirstChild(child ChildNode) {
	pn.firstChild = child
}

func (pn *parentNode) LastChild() ChildNode {
	return pn.lastChild
}

func (pn *parentNode) setLastChild(child ChildNode) {
	pn.lastChild = child
}

func (pn *parentNode) InsertBefore(reference, newChild Node) (ChildNode, error) {
	assert.NotNil(newChild, "child must be non-nil to be inserted")

	// InsertBefore(nil, node) is equivalent to AppendChild(node)
	if reference == nil {
		log.Infoln("can't insert before a nil reference, appending the child")
		return pn.AppendChild(newChild)
	}

	if !pn.isSame(reference.Parent()) {
		return nil, errors.New("can't insert node, the reference node is not a child of this node")
	}

	// Some check to keep a valid node tree
	err := pn.ensurePreInsertionValidity(newChild)
	if err != nil {
		return nil, err
	}

	// newChild and reference are the same
	if reference == newChild {
		log.Infoln("can't insert a node before itself, reference is now node next sibling")
		reference = newChild.Next()
		if reference == nil {
			log.Infoln("can't insert before a nil reference, appending the child")
			return pn.AppendChild(newChild)
		}
	}

	pn.prepareNewChild(newChild)
	child := makeChildNode(newChild)
	pn.performInsertBefore(reference, child)

	return makeChildNode(newChild), nil
}

func (pn *parentNode) prepareNewChild(newChild Node) {
	makeChildNode(newChild).Remove()
}

func (pn *parentNode) performInsertBefore(reference Node, newChild ChildNode) {
	prev := reference.Previous()
	reference.setPrevious(newChild)

	if prev != nil {
		prev.setNext(newChild)
		newChild.setPrevious(prev)
	} else {
		pn.setFirstChild(newChild)
	}

	newChild.setNext(reference)
	pn.adopt(newChild)
}

func (pn *parentNode) adopt(child Node) {
	child.setParent(pn)
	child.setOwner(pn.Owner())
}

func (pn *parentNode) AppendChild(newChild Node) (ChildNode, error) {
	assert.NotNil(newChild, "child must be non-nil to be appended")

	err := pn.ensurePreInsertionValidity(newChild)
	if err != nil {
		return nil, err
	}

	pn.prepareNewChild(newChild)
	child := makeChildNode(newChild)

	if pn.lastChild != nil {
		pn.lastChild.setNext(child)
		child.setPrevious(pn.lastChild)
	} else {
		pn.setFirstChild(child)
	}

	pn.setLastChild(child)
	pn.adopt(child)

	return child, nil
}

func (pn *parentNode) ensurePreInsertionValidity(newChild Node) error {
	assert.NotNil(newChild, "child must be non-nil to be validate for insertion")

	childType := newChild.Type()
	if childType == DocumentNodeType {
		return errors.New("document node can't be inserted")
	}
	if !newChild.isParentNode() || childType == TextNodeType {
		return nil
	}

	if newChild.contains(pn) {
		return errors.New("the new child contains the parent")
	}

	return nil
}

func (pn *parentNode) RemoveChild(child Node) error {
	assert.NotNil(child, "child must be non-nil to be removed")

	if !pn.isSame(child.Parent()) {
		return errors.New("can't remove child, the node is not a child of the this node")
	}

	// Handling siblings
	if prev := child.Previous(); prev != nil {
		prev.setNext(child.Next())
		child.setPrevious(nil)
	}
	if next := child.Next(); next != nil {
		next.setPrevious(child.Previous())
		child.setNext(nil)
	}

	// Handling parent
	if firstChild := pn.firstChild; firstChild != nil {
		if firstChild.isSame(child) {
			pn.setFirstChild(nil)
		}
	} else if lastChild := pn.lastChild; lastChild != nil {
		if lastChild.isSame(child) {
			pn.setLastChild(nil)
		}
	}
	child.setParent(nil)

	return nil
}
