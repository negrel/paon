package events

// Unsupported define any events that is not in the list of the supported events.
type Unsupported struct {
	event
	str string
}

// MakeUnsupported return a new Unsupported object.
func MakeUnsupported(content string) Unsupported {
	return Unsupported{
		event: makeEvent(TypeUnsupported()),
		str:   content,
	}
}
