package draw

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/gdamore/tcell/v2"
	"github.com/negrel/debuggo/pkg/log"
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pkg/pdk/events"
)

var _ Screen = &tcellScreen{}

type tcellScreen struct {
	tcell.Screen
	size geometry.Size
}

// NewTcellScreen return a new Screen based on the github.com/gdamore/tcell Screen.
func NewTcellScreen() (Screen, error) {
	// Initialise tcell Screen
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

func (t *tcellScreen) bounds() geometry.Rectangle {
	return geometry.Rect(0, 0, t.Width(), t.Height())
}

func (t *tcellScreen) get(point geometry.Point) Cell {
	c, _, style, _ := t.GetContent(point.X(), point.Y())

	return makeCellFromTcell(c, style)
}

// Update implements the Screen interface.
func (t *tcellScreen) Update() {
	log.Debugln("updating Screen")

	t.Show()
}

// Apply the given patch to the Screen.
func (t *tcellScreen) Apply(patch Canvas) {
	log.Traceln("applying patch", patch)

	for i, row := range patch.grid {
		y := patch.Origin().Y() + i
		for j, cell := range row {
			x := patch.Origin().X() + j

			t.SetContent(x, y, cell.Content, []rune{}, tcell.Style(cell.Style))
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

		return events.MakeResize(newSize, oldSize)

	case *tcell.EventMouse:
		X, Y := event.Position()
		return events.MakeClick(geometry.Pt(X, Y))

	case *tcell.EventInterrupt:
		return events.MakeInterrupt(ev.When().UnixNano())

	default:
		return events.MakeUnsupported(spew.Sdump(event))
	}
}
