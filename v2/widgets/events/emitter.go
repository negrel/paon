package events

import (
	"errors"

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
		scrollListener: []ScrollListener{},
	}

	go func() {
		for {
			ev := <-channel
			Emitter.Dispatch(Adapt(ev))
		}
	}()
}

// Emitter is used to emit an event to is listeners.
type _emitter struct {
	Input <-chan Event

	resizeListener []ResizeListener
	scrollListener []ScrollListener
}

/*****************************************************
 ***************** Getters & Setters *****************
 *****************************************************/
// ANCHOR Getters & Setters

// AddResizeListener add a resize event listener.
func (e *_emitter) AddResizeListener(rl ResizeListener) {
	e.resizeListener = append(e.resizeListener, rl)
}

// AddScrollListener add a scroll event listener.
func (e *_emitter) AddScrollListener(sl ScrollListener) {
	e.scrollListener = append(e.scrollListener, sl)
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

func (e *_emitter) Dispatch(ev Event) error {
	switch event := ev.(type) {
	case *ResizeEvent:
		e.DispatchResizeEvent(event)

	case *ScrollEvent:
		e.DispatchScrollEvent(event)

	default:
		return errors.New("the given event is undispatchable")
	}

	return nil
}

// DispatchResizeEvent method dispatch the given resize event
// to the subscribed listeners.
func (e *_emitter) DispatchResizeEvent(re *ResizeEvent) {
	for _, listener := range e.resizeListener {
		listener.OnResize(re)
	}
}

// DispatchResizeEvent method dispatch the given resize event
// to the subscribed listeners.
func (e *_emitter) DispatchScrollEvent(se *ScrollEvent) {
	for _, listener := range e.scrollListener {
		listener.OnScroll(se)
	}
}
