package paon

import (
	"context"
	"github.com/negrel/debuggo/pkg/log"
	"time"

	"github.com/negrel/paon/internal/draw"
	"github.com/negrel/paon/pkg/pdk/events"
	"github.com/negrel/paon/pkg/pdk/widgets"
)

// App is the entry point of your TUI application.
type App struct {
	root   *root
	clock  *time.Ticker
	Stop   func()
	window *window
}

// New return a new application object.
func New(opts ...Option) *App {
	app := newApp()

	for _, opt := range opts {
		opt(app)
	}

	return app
}

func newApp() *App {
	return &App{
		clock: time.NewTicker(16 * time.Millisecond),
		Stop: func() {
			log.Errorln("can't stop an application that haven't start")
		},
		window: newWindow(),
	}
}

// Start the application and block the goroutine.
func (a *App) Start(root widgets.Widget) (err error) {
	log.Debugln("starting the app")
	defer a.recoverStart()

	// Set the root component of the widget tree
	a.root = newRoot(root)

	if a.window.Screen == nil {
		a.window.Screen, err = draw.NewTcellScreen()
		if err != nil {
			return
		}
	}

	// Set up application context
	ctx, cancel := context.WithCancel(context.Background())
	a.Stop = func() {
		log.Debugln("stopping the app")
		// Clean screen an cancel context
		a.window.Fini()
		cancel()
	}

	// Start listening to events
	go a.listenToEvents(ctx)

	// Create & start the rendering engine
	engine := draw.NewEngine(a.clock, ctx)
	go engine.Start()

	// Wait until application stop
	<-ctx.Done()

	return nil
}

func (a *App) recoverStart() {
	if r := recover(); r != nil {
		a.Stop()
		panic(r)
	}
}

func (a *App) listenToEvents(ctx context.Context) {
	log.Debugln("starting event loop")

	pollEvent := make(chan events.Event, 8)
	go a.window.Screen.PollEvent(pollEvent)

	for {
		select {
		case <-ctx.Done():
			log.Debugln("stopping event loop")
			return

		case event := <-pollEvent:
			a.dispatchEvent(event)
		}
	}
}

func (a *App) dispatchEvent(event events.Event) {
	log.Debugln("dispatching", event)

	a.window.DispatchEvent(event)
	a.root.DispatchEvent(event)
}

func (a *App) Window() Window {
	return a.window
}
