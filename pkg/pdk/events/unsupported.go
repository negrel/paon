package events

import (
	"fmt"
)

// Unsupported define any events that is not in the list of the supported events.
type Unsupported struct {
	event
	str string
}

// MakeUnsupported return a new Unsupported object.
func MakeUnsupported(content string) Unsupported {
	return Unsupported{
		event: makeEvent(TypeUnsupported),
		str:   content,
	}
}

// String implements the fmt.Stringer interface.
func (ue Unsupported) String() string {
	return fmt.Sprintf("%v: %v", ue.eType, ue.str)
}
