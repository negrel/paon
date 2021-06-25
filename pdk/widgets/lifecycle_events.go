package widgets

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/events"
)

var lifeCycleEventType = events.NewType("lifecycle")

// LifeCycleEventType returns the events.Type of node lifecycle events.
func LifeCycleEventType() events.Type {
	return lifeCycleEventType
}

// LifeCycleEventListener returns an events.Listener that will call the given handler
// when a LifeCycleEvent is dispatched.
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
	Widget Widget
	Stage  LifeCycleStage
}

// NewLifeCycleEvent returns a new LifeCycleEvent events.Event with the given stage.
func NewLifeCycleEvent(widget Widget, stage LifeCycleStage) LifeCycleEvent {
	return LifeCycleEvent{
		Event:  events.NewEvent(lifeCycleEventType),
		Widget: widget,
		Stage:  stage,
	}
}
