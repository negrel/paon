package events

import (
	"fmt"

	"github.com/negrel/paon/events"
	"github.com/negrel/paon/tree"
)

var _ events.Target = target{}

type target struct {
	events.Target

	node            tree.Node
	bubbleListeners map[events.Type][]events.Handler
}

// newTarget returns a new events.Target that wraps the given events.Target.
// Events dispatched using the new target will be propagated to child nodes
// if they implement the Event interface of this package.
func newTarget(node tree.Node, t events.Target) target {
	return target{
		node:   node,
		Target: t,
	}
}

// AddEventListener implements the events.Target interface.
func (t target) AddEventListener(tpe events.Type, listener events.Handler) {
	if l, ok := listener.(Listener); ok {
		if l.bubbles {
			t.bubbleListeners[tpe] = append(t.bubbleListeners[tpe], l)
		}
	} else {
		t.Target.AddEventListener(tpe, listener)
	}
}

// DispatchEvent implements the events.Target interface.
func (t target) DispatchEvent(event events.Event) {
	if ev, ok := event.(Event); ok {
		t.dispatchTreeEvent(ev)
	} else {
		// Dispatch regular events
		t.Target.DispatchEvent(event)
	}
}

func (t target) dispatchTreeEvent(event Event) {
	switch event.Phase {
	case PhaseNone:
		event.Phase = PhaseCapturing
		fallthrough

	case PhaseCapturing:
		t.dispatchCaptureEvent(event)

	case PhaseAtTarget:
		t.Target.DispatchEvent(event)
		if event.bubbles {
			event.Phase = PhaseBubbling
			t.dispatchBubbleEvent(event)
		}

	case PhaseBubbling:
		t.dispatchBubbleEvent(event)

	default:
		panic(fmt.Sprintf("invalid phase state: %v", event.Phase))
	}
}

func (t target) dispatchBubbleEvent(event Event) {
	for _, listener := range t.bubbleListeners[event.Type()] {
		listener.HandleEvent(event)
	}

	t.node.Parent().(Node).DispatchEvent(event)
}

func (t target) dispatchCaptureEvent(event Event) {
	eventTarget := event.Target()

	for child := t.node.FirstChild(); child != nil; child = child.Next() {
		if eventTarget.IsSame(child) {
			event.Phase = PhaseAtTarget
			child.(Node).DispatchEvent(event)
			return
		}

		if eventTarget.IsDescendantOf(child) {
			child.(Node).DispatchEvent(event)
		}
	}
}
