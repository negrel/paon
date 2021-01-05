package events

import "sync/atomic"

//go:generate stringer -type=Type -trimprefix=Type
// Type is the type of an Event
type Type int32

// List of existing event type.
const (
	TypeError Type = iota - 2
	TypeUnsupported
	TypeInterrupt
	TypeClick
	TypeKeyboard
	TypeResize
	TypeWheel

	// Used for custom property created using the pdk
	unusedType
)

var eventTypeCounter int32 = int32(unusedType - 1)

func MakeType() Type {
	return Type(atomic.AddInt32(&eventTypeCounter, 1))
}
