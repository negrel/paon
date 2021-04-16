package events

// Uint8ChangeListener is notified whenever the value of an ObservableUint8 change.
type Uint8ChangeListener interface {
	Changed(old, new uint8)
}

type uint8ChangeListener func(old, new uint8)

func (cl uint8ChangeListener) Changed(old, new uint8) {
	cl(old, new)
}

// MakeUint8ChangeListener wraps your ObservableUint8 change handler
// and returns it.
func MakeUint8ChangeListener(handler func(old, new uint8)) Uint8ChangeListener {
	return uint8ChangeListener(handler)
}

// ObservableUint8 is an entity that wraps a uint8 value and allows to observe
// the value for changes.
type ObservableUint8 interface {
	Observable

	AddChangeListener(Uint8ChangeListener)
	RemoveChangeListener(Uint8ChangeListener)

	Get() uint8
	Set(uint8)
}

type observableUint8 struct {
	observable
	value uint8
	changeListeners map[Uint8ChangeListener]struct{}
}

func NewObservableUint8(initialValue uint8) ObservableUint8 {
	o := &observableUint8{
		value:           initialValue,
		changeListeners: make(map[Uint8ChangeListener]struct{}),
	}
	o.observable = makeCompositeObservable(o)

	return o
}

func (o *observableUint8) AddChangeListener(listener Uint8ChangeListener) {
	o.changeListeners[listener] = struct{}{}
}

func (o *observableUint8) RemoveChangeListener(listener Uint8ChangeListener) {
	delete(o.changeListeners, listener)
}

func (o *observableUint8) Get() uint8 {
	return o.value
}

func (o *observableUint8) Set(newValue uint8) {
	if (newValue == o.Get()) {
		return
	}
	
	old := o.value
	o.value = newValue

	o.triggerChange(old, newValue)
}

func (o *observableUint8) triggerChange(old, new uint8) {
	for changeListener := range o.changeListeners {
		changeListener.Changed(old, new)
	}
}