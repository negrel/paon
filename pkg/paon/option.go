package paon

import "github.com/negrel/paon/pkg/pdk/displays"

type Option func(app *Application)

// Screen sets the underlying displays.Screen to use.
func Screen(screen displays.Screen) Option {
	return func(app *Application) {
		app.screen = screen
	}
}
