package events

import (
	"log"

	"github.com/gdamore/tcell"
)

// OldScreenSize is the old screen size for comparing to recent event.
var oldScreenSize size = size{0, 0}

// Adapt adapth the given tcell event to a ginger
// event.
func Adapt(event tcell.Event) Event {
	switch rawEvent := event.(type) {

	case *tcell.EventKey:
		log.Printf("Key event: %+v %v\n", rawEvent.Name(), rawEvent.Key())
		direction := Direction(rawEvent.Key() - 257)

		se := NewScrollEvent(rawEvent.When(), direction)

		return se

	case *tcell.EventMouse:
		log.Printf("Mouse event: %+v\n", rawEvent)
		buttons := rawEvent.Buttons()
		scrollDir := Direction(buttons << 8)
		click := Button(0xFF & buttons)

		// scroll event
		if scrollDir > 0 {
			return NewScrollEvent(rawEvent.When(), scrollDir)
		}
		// click event
		X, Y := rawEvent.Position()
		return NewClickEvent(rawEvent.When(), click, Position{X, Y})

	case *tcell.EventResize:
		// log.Printf("Resize event: %+v\n", rawEvent)
		w, h := rawEvent.Size()
		newSize := size{
			width:  w,
			height: h,
		}
		re := NewResizeEvent(rawEvent.When(), newSize, oldScreenSize)

		// Update oldScreenSize
		oldScreenSize = newSize

		return re

	case *tcell.EventError:
		log.Printf("Error event: %+v\n", rawEvent)
		return nil

	case *tcell.EventInterrupt:
		log.Printf("Interrupt event: %+v\n", rawEvent)
		return nil

	case *tcell.EventTime:
		log.Printf("Time event: %+v\n", rawEvent)
		return nil

	default:
		log.Println("Unidentified event.")
		return nil
	}
}
