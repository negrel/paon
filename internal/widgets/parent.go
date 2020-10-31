package widgets

import (
	"github.com/negrel/debuggo/pkg/assert"
)

type Parent interface {
	Node

	isAncestorOf(child Node) bool

	FirstChild() Node
	LastChild() Node

	AppendChild(child Node) error
	InsertBefore(reference, child Node) error
	RemoveChild(child Node)
}

var _ Parent = &parent{}

type parent struct {
	*node

	firstChild Node
	lastChild  Node
}

func (p parent) isAncestorOf(child Node) bool {
	return child.isDescendantOf(p)
}

func (p parent) FirstChild() Node {
	return p.firstChild
}

func (p parent) LastChild() Node {
	return p.lastChild
}

func newParent(name string) *parent {
	return &parent{
		node: newNode(name),
	}
}

func (p parent) AppendChild(child Node) error {
	assert.NotNil(child, "child must be non-nil to be appended")

	// If child have a parent, remove it from the parent
	if parent := child.Parent(); parent != nil {

	}
}

func (p parent) InsertBefore(reference, child Node) error {
	return nil

}

func (p parent) RemoveChild(child Node) {

}
