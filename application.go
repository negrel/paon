package paon

import (
	"os"
	"runtime/debug"
	"time"

	"github.com/negrel/debuggo/pkg/log"
	"github.com/negrel/paon/internal/metrics"
	"github.com/negrel/paon/pdk/backend"
	"github.com/negrel/paon/pdk/backend/tcell"
	"github.com/negrel/paon/pdk/draw"
	pdkevents "github.com/negrel/paon/pdk/events"
	pdkwidgets "github.com/negrel/paon/pdk/widgets"
)

var renderingTimeout = time.Millisecond * 50

// Application define a TUI application object.
type Application struct {
	terminal backend.Terminal
	clock    *time.Ticker

	done chan struct{}
	do   chan func()

	root   *pdkwidgets.Root
	target pdkevents.Target
	evch   chan pdkevents.Event
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
		target:   pdkevents.NewTarget(),
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

		app.Stop()
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
func (app *Application) Start(widget pdkwidgets.Widget) error {
	defer app.recover()

	app.evch = make(chan pdkevents.Event, 1024)
	err := app.terminal.Start(app.evch)
	if err != nil {
		return err
	}

	app.done = make(chan struct{})

	var updateRoot func()
	app.root = pdkwidgets.NewRoot()
	app.root.SetChild(widget)

loop:
	for {
		select {
		case <-app.clock.C:
			updateRoot()
			app.terminal.Flush()

		case event := <-app.evch:
			app.handleEvent(event)

		case fn := <-app.do:
			fn()

		case <-app.done:
			metrics.Report(os.Stderr)
			break loop
		}
	}

	return nil
}

// Stop stops the application console, event loop and render loop.
func (app *Application) Stop() {
	app.terminal.Stop()
	if app.done == nil {
		return
	}

	app.done <- struct{}{}
	close(app.done)
	app.done = nil
}

func (app *Application) handleEvent(event pdkevents.Event) {
	if event == nil {
		app.Stop()
	}

	switch event.Type() {

	}
}

func (app *Application) Canvas() draw.Surface {
	return app.terminal
}
