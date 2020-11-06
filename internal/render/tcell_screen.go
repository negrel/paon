package render

import (
	"github.com/gdamore/tcell"
	"github.com/negrel/debuggo/pkg/log"

	"github.com/negrel/paon/internal/events"
	"github.com/negrel/paon/internal/geometry"
)

var _ Screen = &tcellScreen{}

type tcellScreen struct {
	tcell.Screen
	size geometry.Size
}

// NewTcellScreen return a new Screen based on the github.com/gdamore/tcell screen.
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
	scr.EnableMouse()

	return &tcellScreen{
		Screen: scr,
		size:   geometry.Size{},
	}, nil
}

// Size implements the Screen interface.
func (t *tcellScreen) Size() geometry.Size {
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
	log.Debugln("updating screen")

	t.Show()
}

// Apply the given patch to the screen.
func (t *tcellScreen) Apply(patch Patch) {
	log.Debugln("applying patch", patch)

	for i, row := range patch.Frame {
		y := patch.Origin.Y() + i
		for j, cell := range row {
			x := patch.Origin.X() + j

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
		newSize := geometry.NewSize(w, h)
		t.size = newSize

		return events.MakeResizeEvent(newSize, oldSize)

	case *tcell.EventMouse:
		X, Y := event.Position()
		return events.MakeClickEvent(geometry.Pt(X, Y))

	case *tcell.EventInterrupt:
		return events.MakeInterruptEvent(ev.When().UnixNano())

	default:
		return events.UnsupportedEvent{}
	}
}
