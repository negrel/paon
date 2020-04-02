package ginger

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/gdamore/tcell"
	"github.com/negrel/ginger/v1/paint"
	"github.com/negrel/ginger/v1/style"
)

var logF io.Writer

func init() {
	logf, err := os.OpenFile("debug.log", os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(logf)
}

// Screen represent the terminal screen.
type Screen = tcell.Screen

// Application contain the whole terminal application,
// it contain all your application activity.
type Application struct {
	screen Screen
}

// New return a new application
func New() (*Application, error) {
	log.Println("Creating the application...")
	screen, err := tcell.NewTerminfoScreen()

	if err != nil {
		return nil, err
	}

	return &Application{
		screen: screen,
	}, nil
}

/*****************************************************
 ********************** METHODS **********************
 *****************************************************/
// ANCHOR Methods

// AddActivities add activities
// func (a *Application) AddActivities(ac ...*Activity) {
// 	log.Println("Adding activity...")
// 	a.activity = append(a.activity, ac...)
// }

// Start the application.
func (a *Application) Start(ac *Activity) error {
	log.Println("Starting the application.")

	err := a.screen.Init()
	if err != nil {
		return err
	}

	// PAINTER
	paintChannel := make(chan style.Frame)

	painter := paint.Painter{
		Channel: paintChannel,
		Paint:   a.screen.SetContent,
		Refresh: a.screen.Show,
	}

	// Start the painter
	go painter.Start()

	// ACTIVITY
	ac.screen = &a.screen
	ac.paintChannel = paintChannel

	// Start the activity
	ac.Start()

	// Start the event listener
	go a.listen()

	return nil
}

// Stop the application
func (a *Application) Stop() error {
	log.Println("Stopping the application...")
	a.screen.Fini()
	return nil
}

func (a *Application) listen() {
	for {
		event := a.screen.PollEvent()
		switch event := event.(type) {

		case *tcell.EventKey:
			log.Printf("%+v\n", event)

		case *tcell.EventMouse:
			log.Printf("%+v\n", event)

		case *tcell.EventResize:
			log.Printf("%+v\n", event)

		case *tcell.EventTime:
			log.Printf("%+v\n", event)

		case *tcell.EventInterrupt:
			log.Printf("%+v\n", event)

		case *tcell.EventError:
			log.Printf("%+v\n", event)

		default:
			time.Sleep(time.Second)
		}
	}
}
