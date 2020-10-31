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

	AppendChild(newChild Node) (ChildNode, error)
	InsertBefore(newChild Node, reference Node) (ChildNode, error)
}

var _ ContainerNode = &containerNode{}

type containerNode struct {
	Node

	firstChild ChildNode
	lastChild  ChildNode
}

func makeContainerNode() ContainerNode {
	return &containerNode{
		Node: &node{
			nType:       DocumentNode,
			isContainer: true,
		},
	}
}

func (c *containerNode) FirstChild() ChildNode {
	return c.firstChild
}

func (c *containerNode) setFirstChild(child ChildNode) {
	c.firstChild = child
}

func (c *containerNode) LastChild() ChildNode {
	return c.lastChild
}

func (c *containerNode) setLastChild(child ChildNode) {
	c.lastChild = child
}

func (c *containerNode) InsertBefore(newChild, reference Node) (ChildNode, error) {
	assert.NotNil(newChild, "child must be non-nil to be inserted")

	// InsertBefore(node, nil) is equivalent to AppendChild(node)
	if reference == nil {
		return c.AppendChild(newChild)
	}

	// Some check to keep a valid node tree
	err := c.ensurePreInsertionValidity(newChild)
	if err != nil {
		return nil, err
	}

	// newChild and reference are the same
	if reference == newChild {
		reference = newChild.Next()
		if reference == nil {
			return c.AppendChild(newChild)
		}
	}

	c.prepareNewChild(newChild)
	child := makeChildNode(newChild)
	c.performInsertBefore(child, reference)

	return makeChildNode(newChild), nil
}

func (c *containerNode) prepareNewChild(newChild Node) {
	makeChildNode(newChild).Remove()
}

func (c *containerNode) performInsertBefore(newChild ChildNode, reference Node) {
	prev := reference.Previous()
	reference.setPrevious(newChild)

	if prev != nil {
		prev.setNext(newChild)
		newChild.setPrevious(prev)
	} else {
		c.setFirstChild(newChild)
	}

	newChild.setNext(reference)
	c.adopt(newChild)
}

func (c *containerNode) adopt(child Node) {
	child.setParent(c)
	child.setOwner(c.Owner())
}

func (c *containerNode) AppendChild(newChild Node) (ChildNode, error) {
	assert.NotNil(newChild, "child must be non-nil to be appended")

	err := c.ensurePreInsertionValidity(newChild)
	if err != nil {
		return nil, err
	}

	c.prepareNewChild(newChild)
	child := makeChildNode(newChild)

	if c.firstChild == nil {
		c.setFirstChild(child)
	}
	if c.lastChild != nil {
		c.lastChild.setNext(child)
		child.setPrevious(c.lastChild)
	}

	c.setLastChild(child)
	c.adopt(child)

	return child, nil
}

func (c *containerNode) ensurePreInsertionValidity(newChild Node) error {
	assert.NotNil(newChild, "nil ")

	childType := newChild.Type()
	if childType == DocumentNode {
		return errors.New("document node can't be inserted")
	}
	if !newChild.isContainerNode() || childType == TextNode {
		return nil
	}

	if newChild.contains(c) {
		return errors.New("the new child contains the parent")
	}

	return nil
}
