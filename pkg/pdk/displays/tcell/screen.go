package tcell

import (
	"context"
	"sync"

	"github.com/gdamore/tcell/v2"
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/debuggo/pkg/log"
	"github.com/negrel/paon/internal/geometry"
	paonEvents "github.com/negrel/paon/pkg/events"
	"github.com/negrel/paon/pkg/pdk/displays"
	"github.com/negrel/paon/pkg/pdk/draw"
	"github.com/negrel/paon/pkg/pdk/events"
	"github.com/negrel/paon/pkg/pdk/render"
)

var _ displays.Screen = &Screen{}

// Screen implements the screens.screen interface.
type Screen struct {
	sync.RWMutex

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
		Target: events.MakeTarget(),
	}, nil
}

func (s *Screen) Recover() {
	if r := recover(); r != nil {
		s.stop()
		panic(r)
	}
}

// Start implements the displays.Screen interface.
func (s *Screen) Start(ctx context.Context) error {
	defer s.Recover()

	err := s.Screen.Init()
	if err != nil {
		return err
	}

	go func() {
		defer s.Recover()
		s.eventLoop(ctx)
	}()

	width, height := s.Screen.Size()

	s.Lock()
	s.canvas = draw.NewCellGrid(geometry.Rect(0, 0, width, height))
	s.Unlock()

	return nil
}

// Canvas implements the displays.Screen interface.
func (s *Screen) Canvas() draw.Canvas {
	return s.canvas
}

func (s *Screen) pollEvent(ch chan<- events.Event) {
	for rawEvent := s.Screen.PollEvent(); rawEvent != nil; rawEvent = s.Screen.PollEvent() {
		event := s.adaptEvent(rawEvent)
		// TODO remove nil event safety
		if event == nil {
			continue
		}
		log.Infof("%+v\n", event)

		ch <- event
	}
}

func (s *Screen) eventLoop(ctx context.Context) {
	ch := make(chan events.Event)
	go func() {
		defer s.Recover()
		s.pollEvent(ch)
	}()

	for {
		select {
		case <-ctx.Done():
			close(ch)
			s.stop()
			return

		case event := <-ch:
			s.Target.DispatchEvent(event)
		}
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

		s.Lock()
		oldSize := s.canvas.Bounds().Size()
		s.canvas = draw.NewCellGrid(geometry.Rect(0, 0, width, height))
		s.Unlock()

		return paonEvents.MakeResize(size, oldSize)

	default:
		assert.Nil(event)
	}

	return nil
}

// Apply implements the render.Surface interface.
func (s *Screen) Apply(patch render.Patch) {
	log.Info("applying patch")
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

func (s *Screen) stop() {
	s.Screen.Fini()
}

// Bounds implements the displays.Screen interface.
func (s *Screen) Bounds() geometry.Rectangle {
	w, h := s.Screen.Size()
	return geometry.Rect(0, 0, w, h)
}
