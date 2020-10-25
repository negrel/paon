package render

import (
	"github.com/gdamore/tcell"
	"github.com/negrel/debuggo/pkg/log"
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
	log.Infoln("updating screen")

	t.Show()
}

// Apply the given patch to the screen.
func (t *tcellScreen) Apply(patch Patch) {
	log.Infoln("applying patch", patch)

	for i, row := range patch.Frame {
		y := patch.Origin.Y + i
		for j, cell := range row {
			x := patch.Origin.X + j

			t.SetContent(x, y, cell.Content, []rune{}, cell.Style)
		}
	}
}
