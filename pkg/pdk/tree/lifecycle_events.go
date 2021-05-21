package tree

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/pkg/pdk/events"
)

var lifeCycleEventType = events.MakeType("lifecycle")

// LifeCycleEventType returns the events.Type of node lifecycle events.
func LifeCycleEventType() events.Type {
	return lifeCycleEventType
}

func LifeCycleEventListener(handler func(LifeCycleEvent)) *events.Listener {
	l := events.Listener{
		Type: lifeCycleEventType,
		Handle: func(event events.Event) {
			assert.IsType(event, MakeLifeCycleEvent(nil, _maxLifeCycle))
			handler(event.(LifeCycleEvent))
		},
	}

	return &l
}

var _ events.Event = LifeCycleEvent{}

// LifeCycleEvent is triggered when the lifecycle step of an object change.
type LifeCycleEvent struct {
	events.Event
	Node  Node
	Stage LifeCycleStage
}

// MakeLifeCycleEvent returns a new LifeCycleEvent events.Event with the given stage.
func MakeLifeCycleEvent(node Node, stage LifeCycleStage) LifeCycleEvent {
	return LifeCycleEvent{
		Event: events.MakeEvent(lifeCycleEventType),
		Node:  node,
		Stage: stage,
	}
}
