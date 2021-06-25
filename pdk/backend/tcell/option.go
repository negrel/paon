package tcell

import (
	"github.com/gdamore/tcell/v2"
)

// Option define an option for Console object.
type Option func(*Console) error

// Screen returns an Option that sets the underlying tcell.Screen used
// by the Console.
func Screen(screen tcell.Screen) Option {
	return func(c *Console) error {
		c.Screen = screen

		return nil
	}
}
