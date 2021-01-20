package paon

import (
	"github.com/negrel/paon/internal/draw"
	"time"
)

type Option func(*App)

// Clock is an option to set the internal clock used in the App for
// rendering.
func Clock(ticker *time.Ticker) Option {
	return func(app *App) {
		app.clock = ticker
	}
}

// Screen is an option to use a different screen backend (default is tcell)
func Screen(screen draw.Screen) Option {
	return func(app *App) {
		app.window.Screen = screen
	}
}
