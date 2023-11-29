package paon

import (
	"context"
	"os"
	"time"

	"github.com/negrel/paon/backend"
	"github.com/negrel/paon/backend/tcell"
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
	"github.com/negrel/paon/widgets"
)

// Application define a TUI application object.
type Application struct {
	terminal backend.Terminal
	// Closure that throttle calls to render.
	throttledRender func()
	// Public do channel that execute function on main goroutine.
	// Every write to it trigger a throttled render.
	do chan func()
	// do channel that don't trigger a render.
	privateDo chan func()
	// Root of widgets tree.
	root widgets.Root
	// event channel written by terminal backend and read by event loop.
	evch chan events.Event
}

// NewApp returns a new Application object.
func NewApp() (*Application, error) {
	terminal, err := tcell.NewTerminal()
	if err != nil {
		return nil, err
	}

	app := &Application{
		terminal:  terminal,
		privateDo: make(chan func()),
		do:        make(chan func()),
		evch:      make(chan events.Event),
	}

	app.throttledRender = throttle(time.Second/60, func() {
		app.privateDo <- app.render
	})

	return app, nil
}

func (app *Application) recover() {
	if r := recover(); r != nil {
		t := time.AfterFunc(time.Second, func() {
			os.Exit(1)
		})

		app.stop()
		t.Stop()

		panic(r)
	}
}

// DoChannel returns a write-only channel that can be used to execute
// function on the main thread. Any write to the channel will enqueue
// a rendering of widgets tree. Note that rendering is throttled to avoid
// excessive rendering.
func (app *Application) DoChannel() chan<- func() {
	return app.do
}

// Start starts the application console, event loop.
func (app *Application) Start(ctx context.Context, root widgets.Root) error {
	defer app.recover()

	err := app.terminal.Start(app.evch)
	if err != nil {
		return err
	}

	app.root = root
	app.root.AddEventListener(widgets.NeedRenderEventListener(func(_ widgets.NeedRenderEvent) {
		app.throttledRender()
	}))

	app.render()
	app.eventLoop(ctx)

	return nil
}

func (app *Application) eventLoop(ctx context.Context) {
	for {
		select {
		case ev := <-app.evch:
			app.root.DispatchEvent(ev)

		case fn := <-app.do:
			fn()
			app.throttledRender()

		case fn := <-app.privateDo:
			fn()

		case <-ctx.Done():
			app.stop()
			return
		}
	}
}

func (app *Application) render() {
	app.terminal.Clear()
	renderable := app.root.Renderable()
	rootSize := renderable.Layout(layout.Constraint{
		MinSize:    geometry.Size{},
		MaxSize:    app.terminal.Size(),
		ParentSize: app.terminal.Size(),
		RootSize:   app.terminal.Size(),
	})
	renderable.Draw(draw.NewSubSurface(app.terminal, geometry.Rect(0, 0, rootSize.Width(), rootSize.Height())))
	app.terminal.Flush()
}

func (app *Application) stop() {
	app.terminal.Stop()
}
