package paon

import (
	"context"
	"fmt"

	"github.com/negrel/debuggo/pkg/log"

	"github.com/negrel/paon/internal/draw"
	"github.com/negrel/paon/internal/events"
	"github.com/negrel/paon/pkg/widgets"
)

// Singletons
var screen draw.Screen
var engine *draw.Engine

// App is the entry point of your TUI application.
type App struct {
	root *root

	Stop func()
}

// New return a new application object.
func New() *App {
	return &App{
		Stop: func() {
			log.Errorln("can't stop an application that haven't start")
		},
	}
}

// Start the application and block the goroutine.
func (a *App) Start(root widgets.Widget) (err error) {
	log.Debugln("starting the app")
	defer a.recoverStart()

	// Set the root component of the widget tree
	a.root = newRoot(root)

	// Set up application context
	ctx, cancel := context.WithCancel(context.Background())
	a.Stop = func() {
		log.Debugln("stopping the app")
		// Clean screen an cancel context
		screen.Fini()
		cancel()
	}

	// Initialising the screen
	if screen == nil {
		screen, err = draw.NewTcellScreen()
		if err != nil {
			return
		}
	}

	// Start listening to events
	go a.listenToEvents(ctx)

	// Create & start the rendering engine
	engine = draw.NewEngine(screen, ctx)
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
	for {
		select {
		case <-ctx.Done():
			log.Debugln("stopping event loop")
			return

		default:
			a.dispatchEvent(screen.PollEvent())
		}
	}
}

func (a *App) dispatchEvent(event events.Event) {
	log.Debugfn(func() []interface{} {
		return []interface{}{
			fmt.Sprintf("%v: %+v\n", event.Type(), event),
		}
	})

	a.root.DispatchEvent(event)

	if event.Type() == events.InterruptEventType {
		a.Stop()
		return
	}
}
