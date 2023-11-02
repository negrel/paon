package tcell

import (
	"github.com/gdamore/tcell/v2"
)

// Option define an option for Terminal object.
type Option func(*Terminal) error

// Screen returns an Option that sets the underlying tcell.Screen used
// by the Terminal.
func Screen(screen tcell.Screen) Option {
	return func(c *Terminal) error {
		c.screen = screen

		return nil
	}
}
