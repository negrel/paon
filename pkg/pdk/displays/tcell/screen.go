package tcell

import (
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

var _ displays.Screen = &screen{}

// Screen implements the screens.screen interface.
type screen struct {
	sync.RWMutex

	events.Target
	tcell.Screen
	canvas draw.Canvas
	done   bool
}

// MakeScreen return a new displays.Screen using the tcell backend.
func MakeScreen() (displays.Screen, error) {
	s, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}

	return &screen{
		Screen: s,
		Target: events.MakeTarget(),
	}, nil
}

func (s *screen) Recover() {
	if r := recover(); r != nil {
		s.Stop()
		panic(r)
	}
}

// Start implements the displays.Screen interface.
func (s *screen) Start() error {
	defer s.Recover()
	s.Lock()
	defer s.Unlock()

	err := s.Screen.Init()
	if err != nil {
		return err
	}
	s.done = false

	go func() {
		defer s.Recover()
		s.eventLoop()
	}()

	width, height := s.Screen.Size()
	s.canvas = draw.NewCellGrid(geometry.Rect(0, 0, width, height))

	return nil
}

// Canvas implements the displays.Screen interface.
func (s *screen) Canvas() draw.Canvas {
	return s.canvas
}

func (s *screen) pollEvent(ch chan<- events.Event) {
	rawEvent := s.Screen.PollEvent()
	for ; rawEvent != nil; rawEvent = s.Screen.PollEvent() {
		event := s.adaptEvent(rawEvent)
		// TODO remove nil event safety
		if event == nil {
			continue
		}
		log.Infof("%+v\n", event)

		ch <- event
	}
}

func (s *screen) eventLoop() {
	ch := make(chan events.Event)
	go func() {
		defer s.Recover()
		s.pollEvent(ch)
	}()

	for {
		select {
		case event := <-ch:
			s.Target.DispatchEvent(event)

		default:
			s.RWMutex.RLock()
			done := s.done
			s.RWMutex.RUnlock()
			if done {
				close(ch)
				return
			}
		}
	}
}

func (s *screen) adaptEvent(event tcell.Event) events.Event {
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
func (s *screen) Apply(patch render.Patch) {
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
func (s *screen) Flush() {
	s.Screen.Show()
}

func (s *screen) Stop() {
	s.Screen.Fini()
	s.Lock()
	s.done = true
	s.Unlock()
}

// Bounds implements the displays.Screen interface.
func (s *screen) Bounds() geometry.Rectangle {
	w, h := s.Screen.Size()
	return geometry.Rect(0, 0, w, h)
}
