package tree

import (
	"errors"

	"github.com/negrel/debuggo/pkg/assert"
)

type ContainerNode interface {
	Node

	FirstChild() ChildNode
	setFirstChild(ChildNode)
	LastChild() ChildNode
	setLastChild(ChildNode)

	AppendChild(Node) (ChildNode, error)
	InsertBefore(reference Node, newChild Node) (ChildNode, error)
}

var _ ContainerNode = &containerNode{}

type containerNode struct {
	Node

	firstChild ChildNode
	lastChild  ChildNode
}

func newContainerNode() ContainerNode {
	return &containerNode{
		Node: &node{
			nType:       DocumentNode,
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
		return cn.AppendChild(newChild)
	}

	if !cn.isSame(reference.Parent()) {
		return nil, errors.New("the node before the child must be inserted is not a child of this node")
	}

	// Some check to keep a valid node tree
	err := cn.ensurePreInsertionValidity(newChild)
	if err != nil {
		return nil, err
	}

	// newChild and reference are the same
	if reference == newChild {
		reference = newChild.Next()
		if reference == nil {
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
	assert.NotNil(newChild, "nil ")

	childType := newChild.Type()
	if childType == DocumentNode {
		return errors.New("document node can't be inserted")
	}
	if !newChild.isContainerNode() || childType == TextNode {
		return nil
	}

	if newChild.contains(cn) {
		return errors.New("the new child contains the parent")
	}

	return nil
}
