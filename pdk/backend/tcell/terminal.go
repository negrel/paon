package tcell

import (
	"sync"

	"github.com/gdamore/tcell/v2"
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pdk/backend"
	"github.com/negrel/paon/pdk/draw"
	pdkevents "github.com/negrel/paon/pdk/events"
)

var _ backend.Terminal = &Terminal{}

// Terminal is a wrapper around https://www.github.com/gdamore/tcell Screen
// that satisfy the backend.Terminal interface.
type Terminal struct {
	mu sync.Mutex

	// the wrapped tcell.Screen.
	// It is initialized in NewTerminal and never reassigned.
	screen tcell.Screen

	done chan struct{}

	ctx *draw.Context
}

// NewTerminal returns a new Terminal object configured with the
// given options.
func NewTerminal(options ...Option) (*Terminal, error) {
	terminal := &Terminal{}

	var err error
	for _, option := range options {
		err = option(terminal)
		if err != nil {
			return nil, err
		}
	}

	if terminal.screen == nil {
		terminal.screen, err = tcell.NewScreen()
		if err != nil {
			return nil, err
		}
	}

	return terminal, nil
}

// Bounds implements the draw.Canvas interface.
func (c *Terminal) Bounds() geometry.Rectangle {
	w, h := c.screen.Size()
	return geometry.Rect(0, 0, w, h)
}

// Get implements the draw.Canvas interface.
func (c *Terminal) Get(pos geometry.Point) draw.Cell {
	return fromTcell(c.screen.GetContent(pos.X(), pos.Y()))
}

// Set implements the draw.Canvas interface.
func (c *Terminal) Set(pos geometry.Point, cell draw.Cell) {
	mainc, combc, style := toTcell(cell)
	c.screen.SetContent(pos.X(), pos.Y(), mainc, combc, style)
}

// Clear implements the backend.Terminal interface.
func (c *Terminal) Clear() {
	c.screen.Clear()
}

// Flush implements the backend.Terminal interface.
func (c *Terminal) Flush() {
	if c.done != nil {
		c.screen.Show()
	}
}

// Start implements the backend.Terminal interface.
func (c *Terminal) Start(evch chan<- pdkevents.Event) error {
	assert.NotNil(evch)
	c.mu.Lock()
	defer c.mu.Unlock()

	err := c.screen.Init()
	if err != nil {
		return err
	}

	c.done = make(chan struct{})
	go c.eventLoop(c.done, evch, c.screen.PollEvent)

	return nil
}

// Stop implements the backend.Terminal interface.
func (c *Terminal) Stop() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.done == nil {
		return
	}

	c.screen.Fini()

	c.done <- struct{}{}
	close(c.done)
	c.done = nil
}

var oldSize = geometry.Size{}

func (c *Terminal) adaptEvent(event tcell.Event) pdkevents.Event {
	switch ev := event.(type) {
	case *tcell.EventError:
		_ = ev
		return nil

	case *tcell.EventResize:
		newSize := geometry.NewSize(ev.Size())
		e := events.NewResize(oldSize, newSize)
		oldSize = newSize

		return e

	default:
		return nil
	}
}

func (c *Terminal) eventLoop(done <-chan struct{}, eventChannel chan<- pdkevents.Event, pollEvent func() tcell.Event) {
	ch := make(chan pdkevents.Event)

	go func(ch chan<- pdkevents.Event) {
		for {
			event := pollEvent()
			if event == nil {
				return
			}
			ch <- c.adaptEvent(event)
		}
	}(ch)

loop:
	for {
		select {
		case ev := <-ch:
			eventChannel <- ev

		case <-done:
			close(ch)
			break loop
		}
	}
}
