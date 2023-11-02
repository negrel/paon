package paon

import (
	"context"
	"os"
	"runtime/debug"
	"time"

	"github.com/negrel/debuggo/pkg/log"
	"github.com/negrel/paon/backend"
	"github.com/negrel/paon/backend/tcell"
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
	"github.com/negrel/paon/widgets"
)

// Application define a TUI application object.
type Application struct {
	terminal backend.Terminal
	clock    *time.Ticker
	do       chan func()
	root     *widgets.Root
	target   events.Target
	evch     chan events.Event
}

// NewApp returns a new Application object.
func NewApp() (*Application, error) {
	terminal, err := tcell.NewTerminal()
	if err != nil {
		return nil, err
	}

	app := &Application{
		terminal: terminal,
		clock:    time.NewTicker(time.Millisecond * 16), // 60 fps
		do:       make(chan func()),
		target:   events.NewTarget(),
		evch:     make(chan events.Event),
	}

	return app, nil
}

func (app *Application) recover() {
	if r := recover(); r != nil {
		log.Error(r)

		stack := debug.Stack()
		t := time.AfterFunc(time.Second, func() {
			log.Error(string(stack))
			os.Exit(1)
		})

		app.stop()
		t.Stop()

		panic(r)
	}
}

// DoChannel returns a write-only channel that can be used to execute
// function on the main thread.
func (app *Application) DoChannel() chan<- func() {
	return app.do
}

// Start starts the application console, event loop and render loop.
func (app *Application) Start(ctx context.Context, widget widgets.Widget) error {
	defer app.recover()

	err := app.terminal.Start(app.evch)
	if err != nil {
		return err
	}

	if app.root == nil {
		app.root = widgets.NewRoot()
	}
	err = app.root.AppendChild(widget.Node())
	if err != nil {
		return err
	}

	app.eventLoop(ctx)

	return nil
}

func (app *Application) eventLoop(ctx context.Context) {
	for {
		select {
		case <-app.clock.C:
			app.terminal.Clear()
			_ = app.root.Layout(layout.Constraint{
				MinSize:    geometry.NewSize(0, 0),
				MaxSize:    app.terminal.Size(),
				ParentSize: app.terminal.Size(),
				RootSize:   app.terminal.Size(),
			})
			app.root.Draw(app.terminal)
			go app.terminal.Flush()

		case ev := <-app.evch:
			app.target.DispatchEvent(ev)
			app.root.DispatchEvent(ev)

		case fn := <-app.do:
			fn()

		case <-ctx.Done():
			app.stop()
			return
		}
	}
}

func (app *Application) stop() {
	app.terminal.Stop()
}
