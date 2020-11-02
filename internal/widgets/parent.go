package widgets

import (
	"errors"
	"fmt"

	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/debuggo/pkg/log"
)

type Parent interface {
	Node

	isAncestorOf(child Node) bool

	FirstChild() Node
	LastChild() Node

	AppendChild(newChild Node) error
	InsertBefore(reference, newChild Node) error
	RemoveChild(child Node) error
}

var _ Parent = &parent{}

type parent struct {
	*node

	firstChild Node
	lastChild  Node
}

func (p *parent) isAncestorOf(child Node) bool {
	return child.isDescendantOf(p)
}

func (p *parent) FirstChild() Node {
	return p.firstChild
}

func (p *parent) LastChild() Node {
	return p.lastChild
}

func newParent(name string) *parent {
	return &parent{
		node: newNode(name),
	}
}

func (p *parent) adopt(child Node) {
	child.setParent(p)
	child.setRoot(p.root)
	if p.root != nil {
		p.root.register(child)
	}
}

func (p *parent) AppendChild(newChild Node) (err error) {
	assert.NotNil(newChild, "child must be non-nil to be appended")
	log.Debugln("appending", newChild, "in", p)

	if err = p.ensurePreInsertionValidity(newChild); err != nil {
		return fmt.Errorf("can't append child, %v", err)
	}
	p.prepareChildForInsertion(newChild)

	if p.lastChild != nil {
		p.lastChild.setNext(newChild)
		newChild.setPrevious(p.lastChild)
	} else {
		p.firstChild = newChild
	}

	p.lastChild = newChild
	p.adopt(newChild)

	return nil
}

func (p *parent) ensurePreInsertionValidity(child Node) error {
	// check if child is not a parent of p
	if parentNode, isParent := child.(Parent); isParent {
		if parentNode.isAncestorOf(p) {
			return errors.New("child contains the parent")
		}
	}

	return nil
}

func (p *parent) prepareChildForInsertion(newChild Node) {
	if parent := newChild.Parent(); parent != nil {
		err := parent.RemoveChild(newChild)
		assert.Nil(err, err)
	}
	assert.Nil(newChild.Root())
	assert.Nil(newChild.Parent())
	assert.Nil(newChild.Previous())
	assert.Nil(newChild.Next())
}

func (p *parent) InsertBefore(reference, newChild Node) error {
	assert.NotNil(newChild, "child must be non-nil to be appended")
	log.Debugln("inserting", newChild, "before", reference, "in", p)

	// InsertBefore(nil, node) is equal to AppendChild(node)
	if reference == nil {
		return p.AppendChild(newChild)
	}
	if referenceIsNotChild := !p.isSame(reference.Parent()); referenceIsNotChild {
		return errors.New("can't insert child, the given reference is not a child of this node")
	}

	if err := p.ensurePreInsertionValidity(newChild); err != nil {
		return fmt.Errorf("can't insert child, %v", err)
	}

	// newChild and reference are the same
	if reference == newChild {
		log.Infoln("can't insert child before itself, reference is now child next sibling")
		reference = newChild.Next()
		if reference == nil {
			log.Infoln("can't insert before a nil reference, appending the child")
			return p.AppendChild(newChild)
		}
	}

	p.prepareChildForInsertion(newChild)

	if previous := reference.Previous(); previous != nil {
		previous.setNext(newChild)
		newChild.setPrevious(previous)
	} else {
		p.firstChild = newChild
	}
	newChild.setNext(reference)
	reference.setPrevious(newChild)

	p.adopt(newChild)

	return nil
}

func (p *parent) RemoveChild(child Node) error {
	assert.NotNil(child, "child must be non-nil to be removed")
	log.Debugln("removing", child, "from", "p")

	// if not a child of p
	if !p.isSame(child.Parent()) {
		return errors.New("can't remove child, the node is not a child of this node")
	}

	// Removing siblings link
	next := child.Next()
	prev := child.Previous()
	if next != nil {
		child.setNext(nil)
		next.setPrevious(prev)
	} else {
		p.lastChild = prev
	}

	if prev != nil {
		child.setPrevious(nil)
		prev.setNext(next)
	} else {
		p.firstChild = next
	}
	// Removing parent & root link
	child.setParent(nil)
	child.setRoot(nil)
	if p.isConnected() {
		p.root.unregister(child)
	}

	return nil
}
