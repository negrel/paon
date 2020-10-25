package surfaces

import (
	"github.com/gdamore/tcell"

	"github.com/negrel/paon/internal/render"
)

var _ Surface = &tcellScreen{}

type tcellScreen struct {
	tcell.Screen
}

// NewTcellSurface return a new Surface based on the github.com/gdamore/tcell screen.
func NewTcellSurface() (Surface, error) {
	// Initialise tcell screen
	scr, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}
	err = scr.Init()
	if err != nil {
		return nil, err
	}
	scr.EnableMouse()

	return &tcellScreen{
		scr,
	}, nil
}

// Update implements the Screen interface.
func (t *tcellScreen) Update() {
	t.Show()
}

// Apply the given patch to the screen.
func (t *tcellScreen) Apply(patch render.Patch) {
	for i, row := range patch.Frame {
		x := patch.Origin.X + i
		for j, cell := range row {
			y := patch.Origin.Y + j

			t.SetContent(x, y, cell.Content, []rune{}, *cell.Style)
		}
	}
}
