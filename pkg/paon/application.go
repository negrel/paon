package paon

import (
	"context"
	"time"

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
	cancel context.CancelFunc
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

	var ctx context.Context
	ctx, app.cancel = context.WithCancel(context.Background())

	app.root = widgets.NewRoot(app.screen, app.engine, root)

	err := app.screen.Start(ctx)
	if err != nil {
		return err
	}

	go func() {
		defer app.Recover()
		app.engine.Start()
	}()

	<-ctx.Done()

	return nil
}

func (app *Application) Recover() {
	if r := recover(); r != nil {
		app.Stop()
		// Sleep so other co routine can handle context cancelling
		time.Sleep(time.Millisecond)
		panic(r)
	}
}

// Stop stops this application.
func (app *Application) Stop() {
	app.cancel()
	app.screen.Stop()
	app.engine.Stop()
}
