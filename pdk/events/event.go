package events

import (
	"time"
)

// Event is a generic interface for all events
type Event interface {
	When() int64
	Type() Type
}

type event struct {
	eType     Type
	timeStamp int64
}

// NewEvent returns a new Event object of the given type. This function
// should be used as a base for real Event objects.
func NewEvent(eventType Type) Event {
	return newEvent(eventType)
}

func newEvent(eventType Type) event {
	return event{
		eType:     eventType,
		timeStamp: time.Now().UnixNano(),
	}
}

// Type returns the type of the event.
func (e event) Type() Type {
	return e.eType
}

// When returns the timestamp of the event.
func (e event) When() int64 {
	return e.timeStamp
}

// Phase define the flow phase of the actual dispatched TreeEvent.
type Phase uint8

const (
	// PhaseNone define the initial state of a TreeEvent.
	// During this phase, the event is not processed.
	PhaseNone Phase = iota

	// PhaseCapturing define the phase where the target tree.Node
	// haven't yet been found. The event is propagated from the root
	// of the tree to the target parent.
	PhaseCapturing

	// PhaseAtTarget define the phase where the event finally reach the target element.
	PhaseAtTarget

	// PhaseBubbling define the phase where the event propagate throught all ancestors
	// of the target until it reaches the root of the tree.
	PhaseBubbling
)

// TreeEvent define a generic event interface for
// tree events.
type TreeEvent interface {
	Event

	// Phase returns the current propagation phase of the Event.
	Phase() Phase
}

var _ TreeEvent = BaseTreeEvent{}

// BaseTreeEvent define a basic TreeEvent that can be used as
// a base for custom TreeEvent.
type BaseTreeEvent struct {
	event
	CurrentPhase Phase
}

// NewBaseTreeEvent returns a new BaseTreeEvent with the given event type
// and a phase sets to PhaseNone.
func NewBaseTreeEvent(eventType Type) BaseTreeEvent {
	return BaseTreeEvent{
		event:        newEvent(eventType),
		CurrentPhase: PhaseNone,
	}
}

// Phase implements the BaseTreeEvent interface.
func (bte BaseTreeEvent) Phase() Phase {
	return bte.CurrentPhase
}
