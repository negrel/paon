package widgets

import "github.com/negrel/paon/pkg/pdk/events"

// LifeCycleStep define the life cycle step of a Widget.
type LifeCycleStep uint8

const (
	beforeCreateLifeCycleStep LifeCycleStep = iota
	createdLifeCycleStep
	beforeMountLifeCycleStep
	mountedLifeCycleStep
	beforeUpdateLifeCycleStep
	updatedLifeCycleStep
	beforeUnmountLifeCycleStep
	unmountedLifeCycleStep
	_maxLifeCycle
)

var _lifeCycleEventType = events.MakeType("lifecycle")

func LifeCycleEventType() events.Type {
	return _lifeCycleEventType
}

// LifeCycleEvent is triggered when the lifecycle step of an object change.
type LifeCycleEvent struct {
	events.Event
	Step LifeCycleStep
}

// MakeLifeCycleEvent returns a new LifeCycleEvent events.Event.
func MakeLifeCycleEvent(step LifeCycleStep) LifeCycleEvent {
	return LifeCycleEvent{
		Event: events.MakeEvent(_lifeCycleEventType),
		Step:  step,
	}
}
