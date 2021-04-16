package events

// IntChangeListener is notified whenever the value of an ObservableInt change.
type IntChangeListener interface {
	Changed(old, new int)
}

type intChangeListener func(old, new int)

func (cl intChangeListener) Changed(old, new int) {
	cl(old, new)
}

// MakeIntChangeListener wraps your ObservableInt change handler
// and returns it.
func MakeIntChangeListener(handler func(old, new int)) IntChangeListener {
	return intChangeListener(handler)
}

// ObservableInt is an entity that wraps a int value and allows to observe
// the value for changes.
type ObservableInt interface {
	Observable

	AddChangeListener(IntChangeListener)
	RemoveChangeListener(IntChangeListener)

	Get() int
	Set(int)
}

type observableInt struct {
	observable
	value int
	changeListeners map[IntChangeListener]struct{}
}

func NewObservableInt(initialValue int) ObservableInt {
	o := &observableInt{
		value:           initialValue,
		changeListeners: make(map[IntChangeListener]struct{}),
	}
	o.observable = makeCompositeObservable(o)

	return o
}

func (o *observableInt) AddChangeListener(listener IntChangeListener) {
	o.changeListeners[listener] = struct{}{}
}

func (o *observableInt) RemoveChangeListener(listener IntChangeListener) {
	delete(o.changeListeners, listener)
}

func (o *observableInt) Get() int {
	return o.value
}

func (o *observableInt) Set(newValue int) {
	if (newValue == o.Get()) {
		return
	}
	
	old := o.value
	o.value = newValue

	o.triggerChange(old, newValue)
}

func (o *observableInt) triggerChange(old, new int) {
	for changeListener := range o.changeListeners {
		changeListener.Changed(old, new)
	}
}