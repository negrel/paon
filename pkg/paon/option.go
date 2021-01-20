package paon

import (
	"github.com/negrel/paon/internal/draw"
	"time"
)

type Option func(*App)

func Clock(ticker *time.Ticker) Option {
	return func(app *App) {
		app.clock = ticker
	}
}

func Screen(screen draw.Screen) Option {
	return func(app *App) {
		app.window.Screen = screen
	}
}
