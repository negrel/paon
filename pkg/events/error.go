package events

import "github.com/negrel/paon/pkg/pdk/events"

var _TypeError = events.MakeType("error")

func TypeError() events.Type {
	return _TypeError
}

var _ error = Error{}

// Error is an event representing some sort of error, and carries an error payload.
type Error struct {
	events.Event
	msg string
}

// MakeError returns a new Error object.
func MakeError(msg string) Error {
	return Error{
		Event: events.MakeEvent(_TypeError),
		msg:   msg,
	}
}
func (e Error) Error() string {
	return e.msg
}
