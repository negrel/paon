package events

// Uint16ChangeListener is notified whenever the value of an ObservableUint16 change.
type Uint16ChangeListener interface {
	Changed(old, new uint16)
}

type uint16ChangeListener func(old, new uint16)

func (cl uint16ChangeListener) Changed(old, new uint16) {
	cl(old, new)
}

// MakeUint16ChangeListener wraps your ObservableUint16 change handler
// and returns it.
func MakeUint16ChangeListener(handler func(old, new uint16)) Uint16ChangeListener {
	return uint16ChangeListener(handler)
}

// ObservableUint16 is an entity that wraps a uint16 value and allows to observe
// the value for changes.
type ObservableUint16 interface {
	Observable

	AddChangeListener(Uint16ChangeListener)
	RemoveChangeListener(Uint16ChangeListener)

	Get() uint16
	Set(uint16)
}

type observableUint16 struct {
	observable
	value uint16
	changeListeners map[Uint16ChangeListener]struct{}
}

func NewObservableUint16(initialValue uint16) ObservableUint16 {
	o := &observableUint16{
		value:           initialValue,
		changeListeners: make(map[Uint16ChangeListener]struct{}),
	}
	o.observable = makeCompositeObservable(o)

	return o
}

func (o *observableUint16) AddChangeListener(listener Uint16ChangeListener) {
	o.changeListeners[listener] = struct{}{}
}

func (o *observableUint16) RemoveChangeListener(listener Uint16ChangeListener) {
	delete(o.changeListeners, listener)
}

func (o *observableUint16) Get() uint16 {
	return o.value
}

func (o *observableUint16) Set(newValue uint16) {
	if (newValue == o.Get()) {
		return
	}
	
	old := o.value
	o.value = newValue

	o.triggerChange(old, newValue)
}

func (o *observableUint16) triggerChange(old, new uint16) {
	for changeListener := range o.changeListeners {
		changeListener.Changed(old, new)
	}
}