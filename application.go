package ginger

import (
	"log"

	"github.com/gdamore/tcell"
	"github.com/negrel/ginger/tree"
	"github.com/negrel/ringo"
)

// Application ...
type Application struct {
	l          *tree.Layout
	screen     tcell.Screen
	paintStack *ringo.AtomicBuffer
}

// New ...
func New() *Application {
	screen, err := tcell.NewTerminfoScreen()

	if err != nil {
		log.Fatal(err)
	}

	return &Application{
		l:          tree.ROOT,
		screen:     screen,
		paintStack: ringo.NewAtomic(2048),
	}
}
