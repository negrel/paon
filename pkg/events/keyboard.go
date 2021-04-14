package events

import "github.com/negrel/paon/pkg/pdk/events"

var _TypeKeyboard = events.MakeType("keyboard")

func TypeKeyboard() events.Type {
	return _TypeKeyboard
}

var _ events.Event = Keyboard{}

// Keyboard define a user interaction with the keyboard. Each event
// describes a single interaction between the user and a key
// (or combination of a key with modifier keys) on the keyboard.
type Keyboard struct {
	events.Event
	Key      int16
	Modifier int16
	Name     string
	Rune     rune
}

// MakeKeyboard returns a new Keyboard events.Event.
func MakeKeyboard(key, modifier int16, name string, r rune) Keyboard {
	return Keyboard{
		Event:    events.MakeEvent(_TypeKeyboard),
		Key:      key,
		Modifier: modifier,
		Name:     name,
		Rune:     r,
	}
}
