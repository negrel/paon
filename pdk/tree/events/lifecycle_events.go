package events

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/pdk/events"
)

var lifeCycleEventType = events.NewType("lifecycle")

// LifeCycleEventType returns the events.Type of node lifecycle events.
func LifeCycleEventType() events.Type {
	return lifeCycleEventType
}

func LifeCycleEventListener(handler func(LifeCycleEvent)) *events.Listener {
	l := events.Listener{
		Type: lifeCycleEventType,
		Handle: func(event events.Event) {
			assert.IsType(event, NewLifeCycleEvent(nil, _maxLifeCycle))
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

// NewLifeCycleEvent returns a new LifeCycleEvent events.Event with the given stage.
func NewLifeCycleEvent(node Node, stage LifeCycleStage) LifeCycleEvent {
	return LifeCycleEvent{
		Event: events.NewEvent(lifeCycleEventType),
		Node:  node,
		Stage: stage,
	}
}
