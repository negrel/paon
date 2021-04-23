package paon

import (
	"github.com/negrel/paon/pkg/pdk/displays"
	"github.com/negrel/paon/pkg/pdk/displays/tcell"
	"github.com/negrel/paon/pkg/pdk/render"
	"github.com/negrel/paon/pkg/pdk/widgets"
)

// Application define a TUI application.
type Application struct {
	engine render.Engine
	screen displays.Screen
	root   *widgets.Root
}

// MakeApplication returns a new Application object configured with the given object.
func MakeApplication(opts ...Option) (*Application, error) {
	app := new(Application)

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

	if app.engine == nil {
		app.engine = render.NewEngine(app.screen)
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

	app.root = widgets.NewRoot(app.screen, app.engine, root)

	err := app.screen.Start()
	if err != nil {
		return err
	}

	app.engine.Start()

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
	app.engine.Stop()
	app.screen.Stop()
}
