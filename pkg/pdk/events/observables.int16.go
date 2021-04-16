package events

// Int16ChangeListener is notified whenever the value of an ObservableInt16 change.
type Int16ChangeListener interface {
	Changed(old, new int16)
}

type int16ChangeListener func(old, new int16)

func (cl int16ChangeListener) Changed(old, new int16) {
	cl(old, new)
}

// MakeInt16ChangeListener wraps your ObservableInt16 change handler
// and returns it.
func MakeInt16ChangeListener(handler func(old, new int16)) Int16ChangeListener {
	return int16ChangeListener(handler)
}

// ObservableInt16 is an entity that wraps a int16 value and allows to observe
// the value for changes.
type ObservableInt16 interface {
	Observable

	AddChangeListener(Int16ChangeListener)
	RemoveChangeListener(Int16ChangeListener)

	Get() int16
	Set(int16)
}

type observableInt16 struct {
	observable
	value int16
	changeListeners map[Int16ChangeListener]struct{}
}

func NewObservableInt16(initialValue int16) ObservableInt16 {
	o := &observableInt16{
		value:           initialValue,
		changeListeners: make(map[Int16ChangeListener]struct{}),
	}
	o.observable = makeCompositeObservable(o)

	return o
}

func (o *observableInt16) AddChangeListener(listener Int16ChangeListener) {
	o.changeListeners[listener] = struct{}{}
}

func (o *observableInt16) RemoveChangeListener(listener Int16ChangeListener) {
	delete(o.changeListeners, listener)
}

func (o *observableInt16) Get() int16 {
	return o.value
}

func (o *observableInt16) Set(newValue int16) {
	if (newValue == o.Get()) {
		return
	}
	
	old := o.value
	o.value = newValue

	o.triggerChange(old, newValue)
}

func (o *observableInt16) triggerChange(old, new int16) {
	for changeListener := range o.changeListeners {
		changeListener.Changed(old, new)
	}
}