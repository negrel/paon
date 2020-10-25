package screens

import (
	"github.com/gdamore/tcell"

	"github.com/negrel/paon/internal/render"
)

var _ Screen = &tcellScreen{}

type tcellScreen struct {
	screen tcell.Screen
	update func()
}

// NewTcellScreen return a new Screen based on the github.com/gdamore/tcell.
func NewTcellScreen() (Screen, error) {
	// Initialise tcell screen
	scr, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}
	err = scr.Init()
	if err != nil {
		return nil, err
	}

	return &tcellScreen{
		screen: scr,
		update: scr.Show,
	}, nil
}

// Update implements the Screen interface.
func (t *tcellScreen) Update() {
	t.update()
}

// Apply the given patch to the screen.
func (t *tcellScreen) Apply(patch render.Patch) {
	for i, row := range patch.Frame {
		x := patch.Origin.X + i
		for j, cell := range row {
			y := patch.Origin.Y + j

			t.screen.SetContent(x, y, cell.Content(), []rune{}, cell.Style())
		}
	}
}
