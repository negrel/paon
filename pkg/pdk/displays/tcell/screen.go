package tcell

import (
	"github.com/gdamore/tcell/v2"
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/debuggo/pkg/log"
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pkg/pdk/displays"
	"github.com/negrel/paon/pkg/pdk/draw"
	"github.com/negrel/paon/pkg/pdk/events"
	"github.com/negrel/paon/pkg/pdk/render"
)

var _ displays.Screen = &Screen{}

// Screen implements the screens.screen interface.
type Screen struct {
	events.Target
	tcell.Screen
	canvas draw.Canvas
}

// MakeScreen return a new displays.Screen using the tcell backend.
func MakeScreen() (displays.Screen, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}

	return &Screen{
		Screen: screen,
	}, nil
}

// Start implements the displays.Screen interface.
func (s *Screen) Start() error {
	defer func() {
		if r := recover(); r != nil {
			s.Stop()
			panic(r)
		}
	}()

	err := s.Screen.Init()
	if err != nil {
		return err
	}

	go s.eventLoop()

	width, height := s.Screen.Size()
	s.canvas = draw.MakeCellGrid(geometry.Rect(0, 0, width, height))

	return nil
}

// Stop implements the displays.Screen interface.
func (s *Screen) Stop() {
	s.Screen.Fini()
}

// Canvas implements the displays.Screen interface.
func (s *Screen) Canvas() draw.Canvas {
	return s.canvas
}

func (s *Screen) eventLoop() {
	for rawEvent := s.Screen.PollEvent(); rawEvent != nil; s.Screen.PollEvent() {
		event := s.adaptEvent(rawEvent)
		if event == nil {
			continue
		}
		log.Info(event)

		s.Target.DispatchEvent(event)
	}
}

func (s *Screen) adaptEvent(event tcell.Event) events.Event {
	switch ev := event.(type) {
	case *tcell.EventKey:

	case *tcell.EventMouse:

	case *tcell.EventPaste:

	case *tcell.EventInterrupt:

	case *tcell.EventResize:
		width, height := ev.Size()
		size := geometry.MakeSize(width, height)
		s.canvas = draw.MakeCellGrid(geometry.Rect(0, 0, width, height))

		return events.MakeResize(size, s.canvas.Bounds().Size())

	default:
		assert.Nil(event)
	}

	return nil
}

// Apply implements the render.Surface interface.
func (s *Screen) Apply(patch render.Patch) {
	patch.ForEachCell(func(pos geometry.Point, cell *render.Cell) {
		if cell == nil {
			return
		}

		s.Screen.SetContent(
			pos.X(), pos.Y(),
			cell.Content, []rune{},
			tcellStyle(cell.Style),
		)
	})
}

// Flush implements the render.Surface interface.
func (s *Screen) Flush() {
	s.Screen.Show()
}
