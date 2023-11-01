package events

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/pdk/events"
)

// LifeCycleStage define a stage in the life of a Node.
type LifeCycleStage uint8

const (
	// LCSInitial stage correspond to the stage of a new Node that haven't
	// been mounted in a tree (no root).
	LCSInitial LifeCycleStage = iota
	// LCSBeforeMount stage is set just before inserting the Node in a tree with a tree.Root.
	LCSBeforeMount
	// LCSMounted stage is set just after the Node has been mounted, its root is non-nil.
	LCSMounted
	// LCSBeforeUnmount stage is set just before the Node is unmounted from a tree, its root is still non-nil at this stage.
	LCSBeforeUnmount
	// LCSUnmounted stage is set just after a Node or one of its parent has been removed from the tree.
	LCSUnmounted
	_maxLifeCycle
)

// String implements the fmt.Stringer interface.
func (lcs LifeCycleStage) String() string {
	switch lcs {
	case LCSInitial:
		return "initial"
	case LCSBeforeMount:
		return "before mount"
	case LCSMounted:
		return "mounted"
	case LCSBeforeUnmount:
		return "before unmount"
	case LCSUnmounted:
		return "unmounted"
	default:
		panic("invalid life cycle stage")
	}
}

var lifeCycleEventType = events.NewType("lifecycle")

// LifeCycleEventType returns the events.Type of node lifecycle events.
func LifeCycleEventType() events.Type {
	return lifeCycleEventType
}

// LifeCycleEventListener returns an events.Listener that will call the given handler
// when a LifeCycleEvent is dispatched.
func LifeCycleEventListener(handler func(LifeCycleEvent)) (events.Type, events.Handler) {
	return LifeCycleEventType(), events.HandlerFunc(func(event events.Event) {
		assert.IsType(event, NewLifeCycleEvent(nil, _maxLifeCycle))
		handler(event.(LifeCycleEvent))
	})
}

var _ events.Event = LifeCycleEvent{}

// LifeCycleEvent is triggered when the lifecycle step of an object change.
type LifeCycleEvent struct {
	events.Event
	Node  Node
	Stage LifeCycleStage
}

// NewLifeCycleEvent returns a new LifeCycleEvent events.Event with the given stage.
func NewLifeCycleEvent(node Node, stage LifeCycleStage) LifeCycleEvent {
	return LifeCycleEvent{
		Event: events.NewEvent(lifeCycleEventType),
		Node:  node,
		Stage: stage,
	}
}
