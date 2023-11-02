package events

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/tree"
)

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

	// PhaseBubbling define the phase where the event propagate through all ancestors
	// of the target until it reaches the root of the tree.
	PhaseBubbling
)

// Event define a wrapper around events.Event interface for
// tree events. Tree events are propagated to the target node
// if it is present in the tree.
type Event struct {
	events.Event
	Phase  Phase
	target tree.Node

	bubbles bool
}

// NewEvent returns a new Event with the given event type
// and a phase sets to PhaseNone. The phase will be sets accordingly
// by the dispatcher if it supports it.
func NewEvent(ev events.Event, target tree.Node, bubbles bool) Event {
	assert.NotNil(target)

	return Event{
		Event:   ev,
		Phase:   PhaseNone,
		target:  target,
		bubbles: bubbles,
	}
}

// Target returns the Event target tree.Node.
func (e Event) Target() tree.Node {
	return e.target
}

// Bubbles returns true if the event will propagating back up through the
// target's ancestors in reverse order.
func (e Event) Bubbles() bool {
	return e.bubbles
}
