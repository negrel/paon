package events

var _ Event = KeyboardEvent{}

// KeyboardEvent define a user interaction with the keyboard. Each event
// describes a single interaction between the user and a key
// (or combination of a key with modifier keys) on the keyboard.
type KeyboardEvent struct {
	event
	Key      int16
	Modifier int16
	Name     string
	Rune     rune
}

// MakeKeyboardEvent returns a new KeyboardEvent object.
func MakeKeyboardEvent(key, modifier int16, name string, r rune) KeyboardEvent {
	return KeyboardEvent{
		event:    makeEvent(KeyboardEventType),
		Key:      key,
		Modifier: modifier,
		Name:     name,
		Rune:     r,
	}
}
