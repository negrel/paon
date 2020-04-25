package events

import (
	"errors"
	"log"
)

// Emit is used to emit event to listeners.
var Emit chan<- Event

// Emitter listen to the Emit channel and emit
// received events.
var Emitter *_emitter

// Tcell adpater.
func init() {
	channel := make(chan Event)
	Emit = channel

	Emitter = &_emitter{
		resizeListeners: []ResizeListener{},
		scrollListeners: []ScrollListener{},
		clickListeners:  []ClickListener{},
	}

	go func() {
		for {
			Emitter.Dispatch(<-channel)
		}
	}()
}

// Emitter is used to emit an event to is listeners.
type _emitter struct {
	Input <-chan Event

	resizeListeners []ResizeListener
	scrollListeners []ScrollListener
	clickListeners  []ClickListener
}

/*****************************************************
 ***************** Getters & Setters *****************
 *****************************************************/
// ANCHOR Getters & Setters

// AddResizeListener add a resize event listener.
func (e *_emitter) AddResizeListener(rl ResizeListener) {
	e.resizeListeners = append(e.resizeListeners, rl)
}

// AddScrollListener add a scroll event listener.
func (e *_emitter) AddScrollListener(sl ScrollListener) {
	e.scrollListeners = append(e.scrollListeners, sl)
}

// // AddScrollListener add a scroll event listener.
func (e *_emitter) AddClickListener(cl ClickListener) {
	e.clickListeners = append(e.clickListeners, cl)
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

	case *ClickEvent:
		e.DispatchClickEvent(event)

	default:
		return errors.New("the given event is undispatchable")
	}

	return nil
}

// DispatchResizeEvent method dispatch the given resize event
// to the subscribed listeners.
func (e *_emitter) DispatchResizeEvent(re *ResizeEvent) {
	for _, listener := range e.resizeListeners {
		listener.OnResize(re)
	}
}

// DispatchResizeEvent method dispatch the given resize event
// to the subscribed listeners.
func (e *_emitter) DispatchScrollEvent(se *ScrollEvent) {
	for _, listener := range e.scrollListeners {
		log.Println(listener)
		listener.OnScroll(se)
	}
}

// DispatchResizeEvent method dispatch the given click event
// to the destined widgets.
func (e *_emitter) DispatchClickEvent(ce *ClickEvent) {
	for _, listener := range e.clickListeners {
		listener.OnClick(ce)
	}
}
