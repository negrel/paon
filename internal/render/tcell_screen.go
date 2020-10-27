package render

import (
	"github.com/gdamore/tcell"
	"github.com/negrel/debuggo/pkg/log"

	"github.com/negrel/paon/internal/events"
	"github.com/negrel/paon/internal/utils"
)

var _ Surface = &tcellScreen{}

type tcellScreen struct {
	tcell.Screen
	size utils.Size
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
		Screen: scr,
		size:   utils.Size{},
	}, nil
}

// Size implements the Screen interface.
func (t *tcellScreen) Size() utils.Size {
	return t.size
}

// Width implements the Screen interface.
func (t *tcellScreen) Width() int {
	return t.size.Width()
}

// Height implements the Screen interface.
func (t *tcellScreen) Height() int {
	return t.size.Height()
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

// PollEvent implements the Screen interface.
func (t *tcellScreen) PollEvent() events.Event {
	ev := t.Screen.PollEvent()

	switch event := ev.(type) {
	case *tcell.EventResize:
		oldSize := t.size

		w, h := event.Size()
		newSize := utils.Size{
			X: w,
			Y: h,
		}
		t.size = newSize

		return events.MakeResizeEvent(newSize, oldSize)

	case *tcell.EventMouse:
		X, Y := event.Position()
		return events.MakeClickEvent(utils.Point{X: X, Y: Y})

	case *tcell.EventInterrupt:
		return events.MakeInterruptEvent(ev.When().UnixNano())

	default:
		return events.UnsupportedEvent{}
	}
}
