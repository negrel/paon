package events

import "time"

type Event2 interface {
	BaseEvent() *BaseEvent
}

type BaseEvent struct {
	Type   Type
	Target Target
	When   time.Time
}

// Event is a generic interface for all events
type Event interface {
	// Type returns event type.
	Type() Type

	// Target returns Target on which event was dispatched.
	Target() Target

	// WithTarget returns a copy of this event with a different target. This is
	// mainly used to forward events to another target.
	WithTarget(Target) Event

	// When returns time when event was created. WithXXX events methods shouldn't
	// change this property.
	When() time.Time
}

// NewEvent returns a new Event object of the given type. This function
// should be used as a base for real Event objects.
func NewEvent(eventType Type) Event {
	return event{
		eType: eventType,
		when:  time.Now(),
	}
}

// event define a basic events with no extra field.
type event struct {
	eType  Type
	when   time.Time
	target Target
}

// Type implements Event.
func (e event) Type() Type {
	return e.eType
}

// Target implements Event.
func (e event) Target() Target {
	return e.target
}

// WithTarget implements Event.
func (e event) WithTarget(t Target) Event {
	e.target = t
	return e
}

// When implements Event.
func (e event) When() time.Time {
	return e.when
}
