package events

var _ Event = Keyboard{}

// Keyboard define a user interaction with the keyboard. Each event
// describes a single interaction between the user and a key
// (or combination of a key with modifier keys) on the keyboard.
type Keyboard struct {
	Event
	Key      int16
	Modifier int16
	Name     string
	Rune     rune
}

// MakeKeyboard returns a new Keyboard events.Event.
func MakeKeyboard(key, modifier int16, name string, r rune) Keyboard {
	return Keyboard{
		Event:    MakeEvent(TypeKeyboard),
		Key:      key,
		Modifier: modifier,
		Name:     name,
		Rune:     r,
	}
}
