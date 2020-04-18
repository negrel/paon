package ginger

import (
	"fmt"
	"image"
	"log"

	"github.com/gdamore/tcell"
	"github.com/negrel/ginger/v2/render"
	"github.com/negrel/ginger/v2/widgets"
)

// Application is the entry point of your terminal
// application
type Application struct {
	screen   tcell.Screen
	renderer chan *render.Frame
	root     *widgets.Root
}

// New return a new Application
func New(child widgets.Widget) *Application {
	screen, err := tcell.NewTerminfoScreen()

	if err != nil {
		log.Fatal(err)
	}

	app := &Application{
		screen:   screen,
		renderer: make(chan *render.Frame),
		root:     widgets.ROOT(child),
	}

	return app
}

// Start the application
func (a *Application) Start() error {
	defer func() {
		if r := recover(); r != nil {
			a.screen.Fini()
			fmt.Println(r)
		}
	}()

	err := a.screen.Init()

	if err != nil {
		return err
	}

	// Listening to events
	go a.listen()

	// Renderer
	render := render.Renderer{
		Input:   a.renderer,
		Paint:   a.paint,
		Refresh: a.screen.Show,
	}

	go render.Start()

	// First paint
	a.Show()

	return nil
}

// Show force the total redraw of the application.
func (a *Application) Show() {
	width, height := a.screen.Size()

	a.renderer <- a.root.Render(widgets.Constraint{
		Bounds: image.Rect(0, 0, width, height),
		// Inherited: &style.DefaultTheme,
	})
}

// paint to the screen the raw cell of the render
func (a *Application) paint(rc *render.RawCell) {
	a.screen.SetContent(rc.X, rc.Y, rc.Mainc, rc.Combc, rc.Style)
}

// listen to the terminal events
func (a *Application) listen() {
	for {
		event := a.screen.PollEvent()
		switch event := event.(type) {

		case *tcell.EventKey:
			log.Printf("%+v %v\n", event.Name(), event.Key())

		case *tcell.EventMouse:
			log.Printf("%+v\n", event)

		case *tcell.EventResize:
			log.Printf("%+v\n", event)
			a.Show()

		case *tcell.EventTime:
			log.Printf("%+v\n", event)

		case *tcell.EventInterrupt:
			log.Printf("%+v\n", event)

		case *tcell.EventError:
			log.Printf("%+v\n", event)
		}
	}
}

// Stop the application
func (a *Application) Stop() error {
	a.screen.Fini()

	return nil
}
