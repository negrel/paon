package events

import (
	"github.com/negrel/paon/pdk/events"
	"github.com/negrel/paon/pdk/tree"
)

// Node define a tree.Node with events and LifeCycleStage.
// tree.Node method are rewritten to ensure type safety until generics are released.
type Node interface {
	events.Target

	// LifeCycleStage returns the LifeCycleStage of this Node.
	LifeCycleStage() LifeCycleStage

	Node() tree.Node

	// Copied from tree.Node
	IsSame(Node) bool
	Next() Node
	SetNext(Node)
	Previous() Node
	SetPrevious(Node)
	Parent() Node
	SetParent(Node)
	Root() Node
	FirstChild() Node
	LastChild() Node
	AppendChild(newChild Node) error
	InsertBefore(reference, newChild Node) error
	RemoveChild(child Node) error
	IsAncestorOf(child Node) bool
	IsDescendantOf(node Node) bool
}

type node struct {
	node tree.Node
	events.Target

	root  Node
	stage LifeCycleStage
}

// NewLeafNode returns a new Node object.
func NewLeafNode(options ...Option) Node {
	return newNode(options...)
}

// NewNode returns a new Node object.
func NewNode(options ...Option) Node {
	node := newNode(options...)

	// Dispcatch lifecycle events to children.
	node.AddEventListener(LifeCycleEventListener(func(event LifeCycleEvent) {
		for child := node.FirstChild(); child != nil; child.Next() {
			ev := event
			ev.Node = child
			child.DispatchEvent(ev)
		}
	}))

	return node
}

func newNode(options ...Option) *node {
	node := &node{}

	for _, option := range options {
		option(node)
	}

	if node.node == nil {
		node.node = tree.NewLeafNode(node)
	}

	if node.Target == nil {
		node.Target = events.NewTarget()
	}

	node.AddEventListener(LifeCycleEventListener(func(event LifeCycleEvent) {
		node.stage = event.Stage

		// Update root field on mount/unmount
		if event.Stage == LCSMounted {
			node.root = node.Parent().Root()
		} else if event.Stage == LCSUnmounted {
			node.root = nil
		}
	}))

	return node
}

func (n *node) Node() tree.Node {
	return n.node
}

func (n *node) LifeCycleStage() LifeCycleStage {
	return n.stage
}

func (n *node) IsSame(other Node) bool {
	if other == nil {
		return false
	}

	return n.node.IsSame(other.Node())
}

func (n *node) adaptTreeNodeGetter(fn func() tree.Node) Node {
	if value := fn(); value != nil {
		return value.Unwrap().(Node)
	}

	return nil
}

func (n *node) adaptTreeNodeSetter(fn func(tree.Node), arg Node) {
	if arg == nil {
		fn(nil)
	}

	fn(arg.Node())
}

func (n *node) Next() Node {
	return n.adaptTreeNodeGetter(n.node.Next)
}

func (n *node) SetNext(next Node) {
	n.adaptTreeNodeSetter(n.node.SetNext, next)
}

func (n *node) Previous() Node {
	return n.adaptTreeNodeGetter(n.node.Previous)
}

func (n *node) SetPrevious(previous Node) {
	n.adaptTreeNodeSetter(n.node.SetPrevious, previous)
}

func (n *node) Parent() Node {
	return n.adaptTreeNodeGetter(n.node.Parent)
}

func (n *node) SetParent(parent Node) {
	if parent != nil {
		n.DispatchEvent(MakeLifeCycleEvent(n, LCSBeforeMount))
	} else {
		n.DispatchEvent(MakeLifeCycleEvent(n, LCSBeforeUnmount))
	}

	n.adaptTreeNodeSetter(n.node.SetParent, parent)
}

func (n *node) Root() Node {
	return n.root
}

func (n *node) FirstChild() Node {
	return n.adaptTreeNodeGetter(n.node.FirstChild)
}

func (n *node) LastChild() Node {
	return n.adaptTreeNodeGetter(n.node.LastChild)
}

func (n *node) AppendChild(newChild Node) error {
	err := n.node.AppendChild(newChild.Node())
	if err != nil {
		n.DispatchEvent(MakeLifeCycleEvent(newChild, LCSMounted))
	}

	return err
}

func (n *node) InsertBefore(reference, newChild Node) error {
	err := n.node.InsertBefore(reference.Node(), newChild.Node())
	if err != nil {
		n.DispatchEvent(MakeLifeCycleEvent(newChild, LCSMounted))
	}

	return err
}

func (n *node) RemoveChild(newChild Node) error {
	err := n.node.RemoveChild(newChild.Node())
	if err != nil {
		n.DispatchEvent(MakeLifeCycleEvent(newChild, LCSUnmounted))
	}

	return err
}

func (n *node) IsAncestorOf(child Node) bool {
	if child == nil {
		return false
	}

	return n.node.IsAncestorOf(child.Node())
}

func (n *node) IsDescendantOf(parent Node) bool {
	if parent == nil {
		return false
	}

	return n.node.IsDescendantOf(parent.Node())
}
