package ginger

import (
	"image"
	"log"

	"github.com/gdamore/tcell"
	"github.com/negrel/ginger/v2/render"
	"github.com/negrel/ginger/v2/widgets"
	"github.com/negrel/ginger/v2/widgets/events"
)

var _ events.ResizeListener = &Application{}

// Application is the entry point of your terminal
// application
type Application struct {
	devMode bool

	screen   tcell.Screen
	renderer chan *render.Frame
	root     *widgets.Root
}

// New return a new Application
func New(child widgets.Widget) *Application {
	screen, err := tcell.NewScreen()

	if err != nil {
		log.Fatal(err)
	}

	app := &Application{
		screen:   screen,
		renderer: make(chan *render.Frame),
		root:     widgets.ROOT(child),
	}

	events.Emitter.AddResizeListener(app)

	return app
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// Start the application
func (a *Application) Start() error {
	err := a.screen.Init()

	if err != nil {
		return err
	}

	// Renderer
	renderer := render.Renderer{
		Input:   a.renderer,
		Paint:   a.paint,
		Refresh: a.screen.Show,
	}

	go renderer.Start()

	// Listening to events
	go a.listen()

	return nil
}

// OnResize implements events.ResizeListener interface.
func (a *Application) OnResize(re *events.ResizeEvent) {
	a.ForceRender(re.Width(), re.Height())
}

// ForceRender force the total redraw of the application.
func (a *Application) ForceRender(width, height int) {
	frame := a.root.Rendering(image.Rect(0, 0, width, height))

	a.renderer <- frame
}

// paint to the screen the raw cell of the render
func (a *Application) paint(rc *render.RawCell) {
	// log.Printf("%v %v %v %v", rc.X, rc.Y, rc.Mainc, rc.Style)
	a.screen.SetContent(rc.X, rc.Y, rc.Mainc, []rune{}, rc.Style)
}

// listen to the terminal events
func (a *Application) listen() {
	for {
		event := a.screen.PollEvent()
		events.Emit <- event
	}
}

// Stop the application
func (a *Application) Stop() error {
	a.screen.Fini()

	return nil
}
