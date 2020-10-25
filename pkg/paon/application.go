package paon

import (
	"context"

	"github.com/negrel/debuggo/pkg/log"

	"github.com/negrel/paon/internal/render"
	"github.com/negrel/paon/internal/utils"
	"github.com/negrel/paon/internal/widgets"
)

var surface render.Surface

// App is the entry point of your TUI application.
type App struct {
	root *root
	stop func()
}

// New return a new application object.
func New(root widgets.Widget) *App {

	return &App{
		root: newRoot(root),
	}
}

// Start the application and block the goroutine.
func (a *App) Start() (err error) {
	log.Infoln("starting the app")
	defer a.recoverStart()

	// Initialising the surface
	if surface == nil {
		surface, err = render.NewTcellSurface()
		if err != nil {
			return
		}
	}

	// Set up application context
	ctx, cancel := context.WithCancel(context.Background())
	a.stop = func() {
		// Clean surface an cancel context
		surface.Fini()
		cancel()
	}

	// Create & start the rendering engine
	engine := render.NewEngine(surface, ctx)
	go engine.Start()

	// Force first render of the root widget
	w, h := surface.Size()
	engine.Render(a.root.Render(
		utils.Rect(0, 0, w, h),
	))

	// Wait until application stop
	<-ctx.Done()

	return nil
}

func (a *App) recoverStart() {
	if r := recover(); r != nil {
		a.Stop()
		log.Fatal(r)
	}
}

// Stop the application
func (a *App) Stop() {
	log.Infoln("stopping the app")

	a.stop()
}
