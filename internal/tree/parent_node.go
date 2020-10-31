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

var _ ParentNode = &containerNode{}

type containerNode struct {
	Node

	firstChild ChildNode
	lastChild  ChildNode
}

func newParentNode(nodeType NodeType) ParentNode {
	assert.NotEqual(nodeType, TextNode, "TextNode is not a valid ParentNode type")

	return &containerNode{
		Node: &node{
			nodeType:    nodeType,
			isContainer: true,
		},
	}
}

func (cn *containerNode) isSame(other Node) bool {
	return cn == other
}

func (cn *containerNode) FirstChild() ChildNode {
	return cn.firstChild
}

func (cn *containerNode) setFirstChild(child ChildNode) {
	cn.firstChild = child
}

func (cn *containerNode) LastChild() ChildNode {
	return cn.lastChild
}

func (cn *containerNode) setLastChild(child ChildNode) {
	cn.lastChild = child
}

func (cn *containerNode) InsertBefore(reference, newChild Node) (ChildNode, error) {
	assert.NotNil(newChild, "child must be non-nil to be inserted")

	// InsertBefore(nil, node) is equivalent to AppendChild(node)
	if reference == nil {
		log.Infoln("can't insert before a nil reference, appending the child")
		return cn.AppendChild(newChild)
	}

	if !cn.isSame(reference.Parent()) {
		return nil, errors.New("can't insert node, the reference node is not a child of this node")
	}

	// Some check to keep a valid node tree
	err := cn.ensurePreInsertionValidity(newChild)
	if err != nil {
		return nil, err
	}

	// newChild and reference are the same
	if reference == newChild {
		log.Infoln("can't insert a node before itself, reference is now node next sibling")
		reference = newChild.Next()
		if reference == nil {
			log.Infoln("can't insert before a nil reference, appending the child")
			return cn.AppendChild(newChild)
		}
	}

	cn.prepareNewChild(newChild)
	child := makeChildNode(newChild)
	cn.performInsertBefore(reference, child)

	return makeChildNode(newChild), nil
}

func (cn *containerNode) prepareNewChild(newChild Node) {
	makeChildNode(newChild).Remove()
}

func (cn *containerNode) performInsertBefore(reference Node, newChild ChildNode) {
	prev := reference.Previous()
	reference.setPrevious(newChild)

	if prev != nil {
		prev.setNext(newChild)
		newChild.setPrevious(prev)
	} else {
		cn.setFirstChild(newChild)
	}

	newChild.setNext(reference)
	cn.adopt(newChild)
}

func (cn *containerNode) adopt(child Node) {
	child.setParent(cn)
	child.setOwner(cn.Owner())
}

func (cn *containerNode) AppendChild(newChild Node) (ChildNode, error) {
	assert.NotNil(newChild, "child must be non-nil to be appended")

	err := cn.ensurePreInsertionValidity(newChild)
	if err != nil {
		return nil, err
	}

	cn.prepareNewChild(newChild)
	child := makeChildNode(newChild)

	if cn.lastChild != nil {
		cn.lastChild.setNext(child)
		child.setPrevious(cn.lastChild)
	} else {
		cn.setFirstChild(child)
	}

	cn.setLastChild(child)
	cn.adopt(child)

	return child, nil
}

func (cn *containerNode) ensurePreInsertionValidity(newChild Node) error {
	assert.NotNil(newChild, "child must be non-nil to be validate for insertion")

	childType := newChild.Type()
	if childType == DocumentNode {
		return errors.New("document node can't be inserted")
	}
	if !newChild.isParentNode() || childType == TextNode {
		return nil
	}

	if newChild.contains(cn) {
		return errors.New("the new child contains the parent")
	}

	return nil
}
