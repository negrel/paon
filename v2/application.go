package ginger

import (
	"errors"
	"image"
	"log"

	"github.com/gdamore/tcell"
	"github.com/negrel/ginger/v2/events"
	"github.com/negrel/ginger/v2/render"
)

var renderer chan *render.Frame

func init() {
	renderer = make(chan *render.Frame)
}

// Application is the entry point of your terminal
// application
type Application struct {
	screen   tcell.Screen
	activity *Activity
}

// New return a new Application
func New() (*Application, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}

	app := &Application{
		screen: screen,
	}

	err = app.screen.Init()
	if err != nil {
		return nil, err
	}

	// init the renderer
	renderer := &render.Renderer{
		Input:   renderer,
		Paint:   app.paint,
		Refresh: app.screen.Show,
	}
	go renderer.Start()
	// start listening to events
	go app.listen()

	return app, nil
}

// Start the application
func (a *Application) Start(ac *Activity) {
	defer func() {
		if p := recover(); p != nil {
			if a.screen != nil {
				a.screen.Fini()
			}
			log.Panic(p)
		}
	}()

	// First activity started
	if a.activity != nil {
		a.activity.stop()
	}

	a.activity = ac
	a.activity.start(a.screen.Size)
}

// Stop the application.
func (a *Application) Stop() error {
	a.activity.stop()
	a.screen.Fini()

	return nil
}

// EnableMouse enable the mouse support.
func (a *Application) EnableMouse() error {
	if !a.screen.HasMouse() {
		return errors.New("mouse is not supported by your terminal")
	}

	a.screen.EnableMouse()
	return nil
}

// listen to the terminal events and send
// them to the events emitter.
func (a *Application) listen() {
	for {
		events.Emit <- adaptEvent(a.screen.PollEvent())
	}
}

// oldScreenSize is the old screen size for comparing to recent event.
var oldScreenSize events.Size = events.Size{
	Width:  0,
	Height: 0,
}

// Adapt the tcell event for the events emitter (ginger event).
func adaptEvent(event tcell.Event) events.Event {
	switch rawEvent := event.(type) {
	case *tcell.EventKey:
		log.Printf("Key event: %+v %v\n", rawEvent.Name(), rawEvent.Key())
		direction := events.Direction(rawEvent.Key() - 257)

		se := events.NewScrollEvent(rawEvent.When(), direction)

		return se

	case *tcell.EventMouse:
		log.Printf("Mouse event: %+v\n", rawEvent)
		buttons := rawEvent.Buttons()
		// scrollDir := events.Direction(buttons << 8)
		click := events.Button(0xFF & buttons)

		// scroll event
		// if scrollDir > 0 {
		// 	return events.NewScrollEvent(rawEvent.When(), scrollDir)
		// }
		// click event
		X, Y := rawEvent.Position()
		return events.NewClickEvent(rawEvent.When(), click, image.Pt(X, Y))

	case *tcell.EventResize:
		// log.Printf("Resize event: %+v\n", rawEvent)
		w, h := rawEvent.Size()
		newSize := events.Size{
			Width:  w,
			Height: h,
		}
		re := events.NewResizeEvent(rawEvent.When(), newSize, oldScreenSize)

		// Update oldScreenSize
		oldScreenSize = newSize

		return re

	case *tcell.EventError:
		log.Printf("Error event: %+v\n", rawEvent)
		return nil

	case *tcell.EventInterrupt:
		log.Printf("Interrupt event: %+v\n", rawEvent)
		return nil

	case *tcell.EventTime:
		log.Printf("Time event: %+v\n", rawEvent)
		return nil

	default:
		return nil
	}
}

// paint is passed to the renderer for painting raw cells.
func (a *Application) paint(rc *render.RawCell) {
	// log.Printf("%v %v %v %v", rc.X, rc.Y, rc.Mainc, rc.Style)
	a.screen.SetContent(rc.X, rc.Y, rc.Mainc, []rune{}, rc.Style)
}
