package events

import (
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/tree"
)

// Node define tree.Node with support for events.
type Node interface {
	tree.Node
	events.Target

	LifeCycleStage() LifeCycleStage
}

var _ Node = &BaseNode{}

// BaseNode define a basic Node implementation.
type BaseNode struct {
	tree.Node
	target

	root  *tree.Root
	stage LifeCycleStage
}

// NewBaseNode returns a new BaseNode configured with the given options.
func NewBaseNode(options ...NodeOption) *BaseNode {
	bn := newBaseNode(options...)
	bn.AddEventListener(LifeCycleEventListener(func(lce LifeCycleEvent) {
		bn.stage = lce.Stage

		switch lce.Stage {
		case LCSMounted:
			bn.root = bn.Parent().Root()

		case LCSUnmounted:
			bn.root = nil
		}

		for child := bn.FirstChild(); child != nil; child = child.Next() {
			lce.Node = child.(Node)
			lce.Node.DispatchEvent(lce)
		}
	}))

	return bn
}

func newBaseNode(options ...NodeOption) *BaseNode {
	bn := &BaseNode{}
	nodeConf := baseNodeOption{
		BaseNode:        bn,
		nodeConstructor: tree.NewLeafNode,
		data:            bn,
	}

	for _, option := range options {
		option(&nodeConf)
	}

	bn.Node = nodeConf.nodeConstructor(nodeConf.data)
	bn.target.node = bn.Node
	if bn.target.Target == nil {
		bn.target.Target = events.NewTarget()
	}

	if bn.Node.Root() != nil {
		bn.stage = LCSMounted
	}

	return bn
}

// LifeCycleStage implements the Node interface.
func (bn *BaseNode) LifeCycleStage() LifeCycleStage {
	return bn.stage
}

// SetParent implements the Node interface.
func (bn *BaseNode) SetParent(parent tree.Node) {
	if parent == nil && bn.stage == LCSMounted {
		bn.DispatchEvent(NewLifeCycleEvent(bn, LCSBeforeUnmount))
	} else if parent != nil && parent.Root() != nil {
		bn.DispatchEvent(NewLifeCycleEvent(bn, LCSBeforeMount))
	}

	bn.Node.SetParent(parent)
}

func (bn *BaseNode) setNewChildLCStage(newChild Node) {
	if bn.LifeCycleStage() == LCSMounted {
		newChild.DispatchEvent(NewLifeCycleEvent(newChild, LCSMounted))
	}
}

func (bn *BaseNode) setRemovedChildLCStage(child Node) {
	if bn.LifeCycleStage() == LCSMounted {
		child.DispatchEvent(NewLifeCycleEvent(child, LCSUnmounted))
	}
}

// AppendChild implements the tree.Node interface.
func (bn *BaseNode) AppendChild(newChild tree.Node) error {
	err := bn.Node.AppendChild(newChild)
	if err == nil {
		bn.setNewChildLCStage(newChild.(Node))
	}

	return err
}

// InsertBefore implements the tree.Node interface.
func (bn *BaseNode) InsertBefore(reference, newChild tree.Node) error {
	err := bn.Node.InsertBefore(reference, newChild)
	if err == nil {
		bn.setNewChildLCStage(newChild.(Node))
	}

	return err
}

// RemoveChild implements the tree.Node interface.
func (bn *BaseNode) RemoveChild(child tree.Node) error {
	err := bn.Node.RemoveChild(child)
	if err == nil {
		bn.setRemovedChildLCStage(child.(Node))
	}

	return err
}
