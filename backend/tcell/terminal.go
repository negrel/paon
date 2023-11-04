package tcell

import (
	"context"
	"reflect"

	"github.com/gdamore/tcell/v2"
	"github.com/negrel/paon/backend"
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/events/keypress"
	"github.com/negrel/paon/events/mouse"
	"github.com/negrel/paon/events/resize"
	"github.com/negrel/paon/geometry"
)

var _ backend.Terminal = &Terminal{}

// Terminal is a wrapper around https://www.github.com/gdamore/tcell Screen
// that satisfy the backend.Terminal interface.
type Terminal struct {
	// the wrapped tcell.Screen.
	// It is initialized in NewTerminal and never reassigned.
	screen tcell.Screen

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

// Start implements the backend.Terminal interface.
func (c *Terminal) Start(evch chan<- events.Event) error {
	err := c.screen.Init()
	if err != nil {
		return err
	}
	c.screen.EnableMouse(tcell.MouseMotionEvents)
	c.screen.EnablePaste()

	go eventLoop(c.screen.PollEvent, evch)

	return nil
}

// Stop implements the backend.Terminal interface.
func (c *Terminal) Stop() {
	c.screen.Fini()
}

func eventLoop(pollFunc func() tcell.Event, evch chan<- events.Event) {
	oldSize := geometry.Size{}
	mousePress := mouse.Event{}

	for {
		event := pollFunc()
		if event == nil {
			evch <- nil
			return
		}

		switch ev := event.(type) {
		case *tcell.EventError:
			_ = ev

		case *tcell.EventResize:
			newSize := geometry.NewSize(ev.Size())
			resize := resize.New(oldSize, newSize)
			oldSize = newSize

			evch <- resize

		case *tcell.EventMouse:
			// A button was pressed in previous event.
			if mousePress.Event != nil && ev.Buttons() == tcell.ButtonNone {
				newX, newY := ev.Position()

				// Mouse up.
				evch <- mouse.NewUp(
					geometry.NewVec2D(ev.Position()),
					mouse.ButtonMask(ev.Buttons()),
					keypress.ModMask(ev.Modifiers()),
				)

				if mousePress.Buttons&mouse.ButtonPrimary != 0 {
					pos := geometry.NewVec2D(newX, newY)
					evch <- mouse.NewClick(
						pos,
						mouse.ButtonMask(ev.Buttons()),
						keypress.ModMask(ev.Modifiers()),
						mousePress,
					)
				}

				mousePress = mouse.Event{}
				continue
			}

			// Store until another event is triggered.
			if mousePress.Event == nil && ev.Buttons() != tcell.ButtonNone {
				mousePress = mouse.NewPress(
					geometry.NewVec2D(ev.Position()),
					mouse.ButtonMask(ev.Buttons()),
					keypress.ModMask(ev.Modifiers()),
				)
				evch <- mousePress
			}

		case *tcell.EventKey:
			evch <- keypress.New(
				keypress.ModMask(ev.Modifiers()),
				keypress.Key(ev.Key()),
				ev.Rune(),
			)

		default:
			println(reflect.TypeOf(ev).String())
		}
	}
}
