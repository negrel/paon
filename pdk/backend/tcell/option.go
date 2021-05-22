package tcell

import (
	"github.com/gdamore/tcell/v2"
	"github.com/negrel/paon/pdk/events"
)

// Option define an option for Console object.
type Option func(*Console) error

// EventChannel returns an Option that sets the Console eventChannel.
func EventChannel(ch chan<- events.Event) Option {
	return func(c *Console) error {
		c.eventChannel = ch

		return nil
	}
}

// Screen returns an Option that sets the underlying tcell.Screen used
// by the Console.
func Screen(screen tcell.Screen) Option {
	return func(c *Console) error {
		c.Screen = screen

		return nil
	}
}
