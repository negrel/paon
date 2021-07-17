package tcell

import (
	"github.com/gdamore/tcell/v2"
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pdk/backend"
	"github.com/negrel/paon/pdk/draw"
	pdkevents "github.com/negrel/paon/pdk/events"
)

var _ backend.Terminal = &Terminal{}

// Terminal is a wrapper around https://www.github.com/gdamore/tcell Screen
// that satisfy the backend.Terminal interface.
type Terminal struct {
	tcell.Screen

	done chan struct{}
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

	if terminal.Screen == nil {
		terminal.Screen, err = tcell.NewScreen()
		if err != nil {
			return nil, err
		}
	}

	return terminal, nil
}

// Bounds implements the draw.Canvas interface.
func (c *Terminal) Bounds() geometry.Rectangle {
	w, h := c.Screen.Size()
	return geometry.Rect(0, 0, w, h)
}

// Get implements the draw.Canvas interface.
func (c *Terminal) Get(pos geometry.Point) draw.Cell {
	return fromTcell(c.Screen.GetContent(pos.X(), pos.Y()))
}

// Set implements the draw.Canvas interface.
func (c *Terminal) Set(pos geometry.Point, cell draw.Cell) {
	mainc, combc, style := toTcell(cell)
	c.Screen.SetContent(pos.X(), pos.Y(), mainc, combc, style)
}

// NewContext implements the draw.Canvas interface.
func (c *Terminal) NewContext(bounds geometry.Rectangle) draw.Context {
	return draw.NewContext(c, bounds)
}

// Clear implements the backend.Terminal interface.
func (c *Terminal) Clear() {
	c.Screen.Clear()
}

// Flush implements the backend.Terminal interface.
func (c *Terminal) Flush() {
	if c.done != nil {
		c.Screen.Show()
	}
}

// Start implements the backend.Terminal interface.
func (c *Terminal) Start(evch chan<- pdkevents.Event) error {
	assert.NotNil(evch)

	err := c.Screen.Init()
	if err != nil {
		return err
	}

	c.done = make(chan struct{})
	go eventLoop(c.done, evch, c.Screen.PollEvent)

	return nil
}

// Stop implements the backend.Terminal interface.
func (c *Terminal) Stop() {
	if c.done == nil {
		return
	}

	c.Screen.Fini()
	c.done <- struct{}{}
	close(c.done)
	c.done = nil
}

func adaptEvent(event tcell.Event) pdkevents.Event {
	switch ev := event.(type) {
	case *tcell.EventError:
		_ = ev
		return nil

	default:
		return nil
	}
}

func eventLoop(done <-chan struct{}, eventChannel chan<- pdkevents.Event, pollEvent func() tcell.Event) {
	ch := make(chan pdkevents.Event)

	go func(ch chan<- pdkevents.Event) {
	loop:
		for {
			event := pollEvent()
			if event == nil {
				break loop
			}
			ch <- adaptEvent(event)
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
