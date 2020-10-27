package events

// ErrorEvent is an event representing some sort of error, and carries an error payload.
type ErrorEvent struct {
	event
	msg string
}

// MakeErrorEvent returns a new ErrorEvent object.
func MakeErrorEvent(msg string) ErrorEvent {
	return ErrorEvent{
		event: makeEvent(ErrorEventType),
		msg:   msg,
	}
}
func (ee ErrorEvent) Error() string {
	return ee.msg
}
