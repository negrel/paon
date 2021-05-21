package paon

import (
	"sync"
	"time"

	"github.com/negrel/paon/pkg/pdk/displays"
	"github.com/negrel/paon/pkg/pdk/displays/tcell"
	"github.com/negrel/paon/pkg/pdk/events"
	"github.com/negrel/paon/pkg/pdk/widgets"
)

// Application define a TUI application.
type Application struct {
	sync.Once

	screen displays.Screen
	done   chan struct{}
	clock  *time.Ticker
}

// MakeApplication returns a new Application object configured with the given object.
func MakeApplication(opts ...Option) (*Application, error) {
	app := &Application{}

	for _, opt := range opts {
		opt(app)
	}

	if app.screen == nil {
		var err error
		app.screen, err = tcell.MakeScreen()
		if err != nil {
			return nil, err
		}
	}

	if app.clock == nil {
		app.clock = time.NewTicker(time.Millisecond * 16)
	}

	return app, nil
}

// Screen returns the screen used by the application.
func (app *Application) Screen() displays.Screen {
	return app.screen
}

// Start starts this application.
func (app *Application) Start(root widgets.Widget) error {
	defer app.Recover()

	app.Once = sync.Once{}
	app.done = make(chan struct{})

	eventChannel := make(chan events.Event)
	err := app.screen.Start(eventChannel)
	if err != nil {
		return err
	}

	root = widgets.NewRoot(app.screen, root)

loop:
	for {
		select {
		case event := <-eventChannel:
			root.DispatchEvent(event)

		case <-app.clock.C:
			app.screen.Flush()

		case <-app.done:
			break loop
		}
	}

	return nil
}

func (app *Application) Recover() {
	if r := recover(); r != nil {
		app.Stop()
		panic(r)
	}
}

// Stop stops this application.
func (app *Application) Stop() {
	app.Once.Do(app.stop)
}

func (app *Application) stop() {
	app.screen.Stop()
	close(app.done)
}
