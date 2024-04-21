package tree

import "github.com/negrel/paon/events"

var BeforeMountEventType = events.NewType("BeforeMountNode")

// BeforeMountEventListener returns an events.Listener that will call the given
// handler on before mount node events.
func BeforeMountEventListener(handler func(BeforeMountEvent)) (events.Type, events.Handler) {
	return BeforeMountEventType, events.HandlerFunc(func(ev events.Event) {
		handler(ev.(BeforeMountEvent))
	})
}

// BeforeMountEvent is dispatched before a node is mounted into a parent node.
type BeforeMountEvent struct {
	events.Event
}

// WithTarget implements events.Event.
func (bme BeforeMountEvent) WithTarget(t events.Target) events.Event {
	bme.Event = bme.Event.WithTarget(t)
	return bme
}

var MountedEventType = events.NewType("MountedNode")

// MountedEventListener returns an events.Listener that will call the given
// handler on mounted node events.
func MountedEventListener(handler func(MountedEvent)) (events.Type, events.Handler) {
	return MountedEventType, events.HandlerFunc(func(ev events.Event) {
		handler(ev.(MountedEvent))
	})
}

// MountedEvent is dispatched after a node is mounted into a parent node.
type MountedEvent struct {
	events.Event
}

// WithTarget implements events.Event.
func (me MountedEvent) WithTarget(t events.Target) events.Event {
	me.Event = me.Event.WithTarget(t)
	return me
}

var BeforeUnmountEventType = events.NewType("BeforeUnmountNode")

// BeforeUnmountEventListener returns an events.Listener that will call the given
// handler on before unmount node events.
func BeforeUnmountEventListener(handler func(BeforeUnmountEvent)) (events.Type, events.Handler) {
	return BeforeUnmountEventType, events.HandlerFunc(func(ev events.Event) {
		handler(ev.(BeforeUnmountEvent))
	})
}

// BeforeUnmountEvent is dispatched before a node is unmounted from parent node.
type BeforeUnmountEvent struct {
	events.Event
}

// WithTarget implements events.Event.
func (bue BeforeUnmountEvent) WithTarget(t events.Target) events.Event {
	bue.Event = bue.Event.WithTarget(t)
	return bue
}

var UnmountedEventType = events.NewType("UnmountedNode")

// UnmountedEventListener returns an events.Listener that will call the given
// handler on unmounted node events.
func UnmountedEventListener(handler func(UnmountedEvent)) (events.Type, events.Handler) {
	return UnmountedEventType, events.HandlerFunc(func(ev events.Event) {
		handler(ev.(UnmountedEvent))
	})
}

// UnmountedEvent is dispatched before a node is unmounted from parent node.
type UnmountedEvent struct {
	events.Event
}

// WithTarget implements events.Event.
func (ue UnmountedEvent) WithTarget(t events.Target) events.Event {
	ue.Event = ue.Event.WithTarget(t)
	return ue
}
