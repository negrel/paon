package events

var _ error = Error{}

// Error is an event representing some sort of error, and carries an error payload.
type Error struct {
	event
	msg string
}

// MakeError returns a new Error object.
func MakeError(msg string) Error {
	return Error{
		event: makeEvent(TypeError),
		msg:   msg,
	}
}
func (ee Error) Error() string {
	return ee.msg
}
