package styles

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/pdk/events"
	"github.com/negrel/paon/styles/property"
)

var colorChangedEventType = events.NewType("color-changed")

// ColorChangedEventType returns the events.Type of ColorChangedEvent events.
func ColorChangedEventType() events.Type {
	return colorChangedEventType
}

// ColorChangedListener returns an events.Listener that will call the given handler
// on ColorChangedEvent events.
func ColorChangedListener(handler func(ColorChangedEvent)) (events.Type, events.Listener) {
	return ColorChangedEventType(), events.ListenerFunc(func(event events.Event) {
		assert.IsType(event, ColorChangedEvent{})
		handler(event.(ColorChangedEvent))
	})
}

// ColorChangedEvent event is triggered after the value of a
// property.Color property changed.
type ColorChangedEvent struct {
	events.Event
	ColorID  property.ColorID
	Old, New *property.Color
}

// NewColorChanged returns a new ColorChangedEvent event.
func NewColorChanged(id property.ColorID, old, new *property.Color) ColorChangedEvent {
	return ColorChangedEvent{
		Event:   events.NewEvent(colorChangedEventType),
		ColorID: id,
		Old:     old,
		New:     new,
	}
}

var ifaceChangedEventType = events.NewType("iface-changed")

// IfaceChangedEventType returns the events.Type of IfaceChangedEvent events.
func IfaceChangedEventType() events.Type {
	return ifaceChangedEventType
}

// IfaceChangedListener returns an events.Listener that will call the given handler
// on IfaceChangedEvent events.
func IfaceChangedListener(handler func(IfaceChangedEvent)) (events.Type, events.Listener) {
	return IfaceChangedEventType(), events.ListenerFunc(func(event events.Event) {
		assert.IsType(event, IfaceChangedEvent{})
		handler(event.(IfaceChangedEvent))
	})
}

// IfaceChangedEvent event is triggered after the value of a
// property.Iface property changed.
type IfaceChangedEvent struct {
	events.Event
	IfaceID  property.IfaceID
	Old, New interface{}
}

// NewIfaceChanged returns a new IfaceChangedEvent event.
func NewIfaceChanged(id property.IfaceID, old, new interface{}) IfaceChangedEvent {
	return IfaceChangedEvent{
		Event:   events.NewEvent(ifaceChangedEventType),
		IfaceID: id,
		Old:     old,
		New:     new,
	}
}

var intChangedEventType = events.NewType("int-changed")

// IntChangedEventType returns the events.Type of IntChangedEvent events.
func IntChangedEventType() events.Type {
	return intChangedEventType
}

// IntChangedListener returns an events.Listener that will call the given handler
// on IntChangedEvent events.
func IntChangedListener(handler func(IntChangedEvent)) (events.Type, events.Listener) {
	return IntChangedEventType(), events.ListenerFunc(func(event events.Event) {
		assert.IsType(event, IntChangedEvent{})
		handler(event.(IntChangedEvent))
	})
}

// IntChangedEvent event is triggered after the value of a
// property.Int property changed.
type IntChangedEvent struct {
	events.Event
	IntID    property.IntID
	Old, New *property.Int
}

// NewIntChanged returns a new IntChangedEvent event.
func NewIntChanged(id property.IntID, old, new *property.Int) IntChangedEvent {
	return IntChangedEvent{
		Event: events.NewEvent(intChangedEventType),
		IntID: id,
		Old:   old,
		New:   new,
	}
}

var intUnitChangedEventType = events.NewType("int-unit-changed")

// IntUnitChangedEventType returns the events.Type of IntUnitChangedEvent events.
func IntUnitChangedEventType() events.Type {
	return intUnitChangedEventType
}

// IntUnitChangedListener returns an events.Listener that will call the given handler
// on IntUnitChangedEvent events.
func IntUnitChangedListener(handler func(IntUnitChangedEvent)) (events.Type, events.Listener) {
	return IntUnitChangedEventType(), events.ListenerFunc(func(event events.Event) {
		assert.IsType(event, IntUnitChangedEvent{})
		handler(event.(IntUnitChangedEvent))
	})
}

// IntUnitChangedEvent event is triggered after the value of a
// property.IntUnit property changed.
type IntUnitChangedEvent struct {
	events.Event
	IntID    property.IntUnitID
	Old, New *property.IntUnit
}

// NewIntUnitChanged returns a new IntUnitChangedEvent event.
func NewIntUnitChanged(id property.IntUnitID, old, new *property.IntUnit) IntUnitChangedEvent {
	return IntUnitChangedEvent{
		Event: events.NewEvent(intUnitChangedEventType),
		IntID: id,
		Old:   old,
		New:   new,
	}
}
