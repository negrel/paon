package tcell

import (
	"context"

	"github.com/gdamore/tcell/v2"
	"github.com/negrel/debuggo/pkg/log"
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/pdk/backend"
	"github.com/negrel/paon/pdk/draw"
	pdkevents "github.com/negrel/paon/pdk/events"
	"github.com/negrel/paon/pdk/render"
)

var _ backend.Terminal = &Terminal{}

// Terminal is a wrapper around https://www.github.com/gdamore/tcell Screen
// that satisfy the backend.Terminal interface.
type Terminal struct {
	// the wrapped tcell.Screen.
	// It is initialized in NewTerminal and never reassigned.
	screen     tcell.Screen
	compositor *render.Compositor

	eventLoopCancel context.CancelFunc
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

	terminal.compositor = render.NewCompositor(terminal)

	return terminal, nil
}

// Size implements the geometry.Sized interface.
func (c *Terminal) Size() geometry.Size {
	w, h := c.screen.Size()
	return geometry.NewSize(w, h)
}

// Get implements the draw.Canvas interface.
func (c *Terminal) Get(pos geometry.Vec2D) draw.Cell {
	return fromTcell(c.screen.GetContent(pos.X(), pos.Y()))
}

// Set implements the draw.Canvas interface.
func (c *Terminal) Set(pos geometry.Vec2D, cell draw.Cell) {
	mainc, combc, style := toTcell(cell)
	c.screen.SetContent(pos.X(), pos.Y(), mainc, combc, style)
}

// Clear implements the backend.Terminal interface.
func (c *Terminal) Clear() {
	c.screen.Clear()
}

// Flush implements the backend.Terminal interface.
func (c *Terminal) Flush() {
	c.screen.Show()
}

// Compositor implements the backend.Terminal interface.
func (c *Terminal) Compositor() *render.Compositor {
	return c.compositor
}

// Start implements the backend.Terminal interface.
func (c *Terminal) Start(evch chan<- pdkevents.Event) error {
	err := c.screen.Init()
	if err != nil {
		return err
	}

	go eventLoop(c.screen.PollEvent, evch)

	return nil
}

// Stop implements the backend.Terminal interface.
func (c *Terminal) Stop() {
	c.compositor.Stop()
	c.screen.Fini()
}

var oldSize = geometry.Size{}

func adaptEvent(event tcell.Event) pdkevents.Event {
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

func eventLoop(pollFunc func() tcell.Event, evch chan<- pdkevents.Event) {
	pollCh := make(chan pdkevents.Event)
	go func(pollCh chan<- pdkevents.Event) {
		for {
			event := pollFunc()
			if event != nil {
				pollCh <- adaptEvent(event)
			} else {
				pollCh <- nil
			}
		}
	}(pollCh)

	for {
		ev := <-pollCh
		if ev == nil {
			log.Debug("EVENT LOOP DONE")
			break
		}
		evch <- ev
	}
}
