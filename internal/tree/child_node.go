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

func (cn childNode) Remove() Node {
	err := cn.Parent().RemoveChild(cn)
	if err != nil {
		panic(err)
	}

	return cn.Node
}

func (cn childNode) Before(nodes ...Node) (err error) {
	parent := cn.Parent()
	for _, node := range nodes {
		_, err = parent.InsertBefore(node, cn.Node)
		if err != nil {
			return
		}
	}

	return nil
}

func (cn childNode) After(nodes ...Node) (err error) {
	return makeChildNode(cn.Next()).Before(nodes...)
}

func (cn childNode) ReplaceWith(nodes ...Node) (err error) {
	err = cn.Before(nodes...)
	cn.Remove()

	return err
}
