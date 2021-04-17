package events

// InvalidationListener is notified whenever an Observable becomes invalid.
type InvalidationListener interface {
	Invalidated(Observable)
}

type invalidationListener func(Observable)

func (il invalidationListener) Invalidated(observable Observable) {
	il(observable)
}

// MakeInvalidationListener wraps the given invalidation handler and returns
// an InvalidationListener.
func MakeInvalidationListener(handler func(Observable)) InvalidationListener {
	return invalidationListener(handler)
}

// Observable is a generic interface implemented by all Observable
// types.
type Observable interface {
	Invalidate()
	IsValid() bool

	AddInvalidationListener(InvalidationListener)
	RemoveInvalidationListener(InvalidationListener)
}

var _ Observable = observable{}

type observable struct {
	ptr Observable

	invalidationListeners map[InvalidationListener]struct{}
}

// MakeCompositeObservable ...
func MakeCompositeObservable(ptr Observable) Observable {
	return makeCompositeObservable(ptr)
}

func makeCompositeObservable(ptr Observable) observable {
	return observable{
		ptr: ptr,

		invalidationListeners: make(map[InvalidationListener]struct{}),
	}
}

func (o observable) Invalidate() {
	for invalidationListener := range o.invalidationListeners {
		invalidationListener.Invalidated(o.ptr)
	}
	o.invalidationListeners = nil
}

func (o observable) IsValid() bool {
	return o.invalidationListeners != nil
}

func (o observable) AddInvalidationListener(listener InvalidationListener) {
	o.invalidationListeners[listener] = struct{}{}
}

func (o observable) RemoveInvalidationListener(listener InvalidationListener) {
	o.invalidationListeners[listener] = struct{}{}
}
