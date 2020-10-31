package tree

import (
	"github.com/negrel/debuggo/pkg/assert"
)

type ChildNode interface {
	Node

	Remove() Node
	Before(...Node) error
	After(...Node) error
	ReplaceWith(...Node) error
}

type childNode struct {
	Node
}

func makeChildNode(node Node) ChildNode {
	if cn, ok := node.(ChildNode); ok {
		return cn
	}
	assert.NotNil(node.Parent(), "to make a child node, node must have a non-nil parent")

	return childNode{
		Node: node,
	}
}

func (c childNode) leaveSiblings() {
	if prev := c.Previous(); prev != nil {
		prev.setNext(c.Next())
		c.setPrevious(nil)
	}
	if next := c.Next(); next != nil {
		next.setPrevious(c.Previous())
		c.setNext(nil)
	}
}

func (c childNode) leaveParent() {
	if parent := c.Parent(); parent != nil {
		if firstChild := parent.FirstChild(); firstChild != nil {
			if firstChild.isSame(c.Node) {
				parent.setFirstChild(nil)
			}
		}
		if lastChild := parent.LastChild(); lastChild != nil {
			if lastChild.isSame(c.Node) {
				parent.setLastChild(nil)
			}
		}
		c.setParent(nil)
	}
}

func (c childNode) Remove() Node {
	c.leaveSiblings()
	c.leaveParent()

	c.setOwner(nil)

	return c.Node
}

func (c childNode) Before(nodes ...Node) (err error) {
	parent := c.Parent()
	for _, node := range nodes {
		_, err = parent.InsertBefore(node, c.Node)
		if err != nil {
			return
		}
	}

	return nil
}

func (c childNode) After(nodes ...Node) (err error) {
	return makeChildNode(c.Next()).Before(nodes...)
}

func (c childNode) ReplaceWith(nodes ...Node) (err error) {
	err = c.Before(nodes...)
	c.Remove()

	return err
}
