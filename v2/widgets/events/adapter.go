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
		direction := Direction(rawEvent.Buttons() << 8)

		return NewScrollEvent(rawEvent.When(), direction)

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

	default:
		log.Println("Unidentified event.")
		return rawEvent.(Event)
	}
}