package events

import (
	"log"

	"github.com/gdamore/tcell"
)

// Emit is used to emit event to listeners.
var Emit chan<- tcell.Event

// Emitter listen to the Emit channel and emit
// received events.
var Emitter *_emitter

// Tcell adpater.
func init() {
	channel := make(chan tcell.Event)
	Emit = channel

	Emitter = &_emitter{
		resizeListener: []ResizeListener{},
	}

	var oldSize size = size{0, 0}

	go func() {
		for {
			event := <-channel

			switch rawEvent := event.(type) {
			case *tcell.EventKey:
				log.Printf("Key event: %+v %v\n", rawEvent.Name(), rawEvent.Key())

			case *tcell.EventMouse:
				log.Printf("Mouse event: %+v\n", rawEvent)

			case *tcell.EventResize:
				// log.Printf("Resize event: %+v\n", rawEvent)
				w, h := rawEvent.Size()
				newSize := size{
					width:  w,
					height: h,
				}

				re := NewResizeEvent(rawEvent.When(), newSize, oldSize)
				Emitter.DispatchResizeEvent(re)

				// Update oldSize
				oldSize = newSize

			case *tcell.EventError:
				log.Printf("Error event: %+v\n", rawEvent)

			default:
				log.Println("Unidentified event.")
				continue
			}
		}
	}()
}

// Emitter is used to emit an event to is listeners.
type _emitter struct {
	Input <-chan Event

	resizeListener []ResizeListener
}

/*****************************************************
 ***************** Getters & Setters *****************
 *****************************************************/
// ANCHOR Getters & Setters

// AddResizeHandler add an input handler.
func (e *_emitter) AddResizeListener(rl ResizeListener) {
	e.resizeListener = append(e.resizeListener, rl)
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// DispatchResizeEvent method dispatch the given resize event
// to the subscribed listeners.
func (e *_emitter) DispatchResizeEvent(re *ResizeEvent) {
	for _, listener := range e.resizeListener {
		listener.OnResize(re)
	}
}
