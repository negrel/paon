package events

// Uint32ChangeListener is notified whenever the value of an ObservableUint32 change.
type Uint32ChangeListener interface {
	Changed(old, new uint32)
}

type uint32ChangeListener func(old, new uint32)

func (cl uint32ChangeListener) Changed(old, new uint32) {
	cl(old, new)
}

// MakeUint32ChangeListener wraps your ObservableUint32 change handler
// and returns it.
func MakeUint32ChangeListener(handler func(old, new uint32)) Uint32ChangeListener {
	return uint32ChangeListener(handler)
}

// ObservableUint32 is an entity that wraps a uint32 value and allows to observe
// the value for changes.
type ObservableUint32 interface {
	Observable

	AddChangeListener(Uint32ChangeListener)
	RemoveChangeListener(Uint32ChangeListener)

	Get() uint32
	Set(uint32)
}

type observableUint32 struct {
	observable
	value uint32
	changeListeners map[Uint32ChangeListener]struct{}
}

func NewObservableUint32(initialValue uint32) ObservableUint32 {
	o := &observableUint32{
		value:           initialValue,
		changeListeners: make(map[Uint32ChangeListener]struct{}),
	}
	o.observable = makeCompositeObservable(o)

	return o
}

func (o *observableUint32) AddChangeListener(listener Uint32ChangeListener) {
	o.changeListeners[listener] = struct{}{}
}

func (o *observableUint32) RemoveChangeListener(listener Uint32ChangeListener) {
	delete(o.changeListeners, listener)
}

func (o *observableUint32) Get() uint32 {
	return o.value
}

func (o *observableUint32) Set(newValue uint32) {
	if (newValue == o.Get()) {
		return
	}
	
	old := o.value
	o.value = newValue

	o.triggerChange(old, newValue)
}

func (o *observableUint32) triggerChange(old, new uint32) {
	for changeListener := range o.changeListeners {
		changeListener.Changed(old, new)
	}
}