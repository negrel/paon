package events

// Uint64ChangeListener is notified whenever the value of an ObservableUint64 change.
type Uint64ChangeListener interface {
	Changed(old, new uint64)
}

type uint64ChangeListener func(old, new uint64)

func (cl uint64ChangeListener) Changed(old, new uint64) {
	cl(old, new)
}

// MakeUint64ChangeListener wraps your ObservableUint64 change handler
// and returns it.
func MakeUint64ChangeListener(handler func(old, new uint64)) Uint64ChangeListener {
	return uint64ChangeListener(handler)
}

// ObservableUint64 is an entity that wraps a uint64 value and allows to observe
// the value for changes.
type ObservableUint64 interface {
	Observable

	AddChangeListener(Uint64ChangeListener)
	RemoveChangeListener(Uint64ChangeListener)

	Get() uint64
	Set(uint64)
}

type observableUint64 struct {
	observable
	value uint64
	changeListeners map[Uint64ChangeListener]struct{}
}

func NewObservableUint64(initialValue uint64) ObservableUint64 {
	o := &observableUint64{
		value:           initialValue,
		changeListeners: make(map[Uint64ChangeListener]struct{}),
	}
	o.observable = makeCompositeObservable(o)

	return o
}

func (o *observableUint64) AddChangeListener(listener Uint64ChangeListener) {
	o.changeListeners[listener] = struct{}{}
}

func (o *observableUint64) RemoveChangeListener(listener Uint64ChangeListener) {
	delete(o.changeListeners, listener)
}

func (o *observableUint64) Get() uint64 {
	return o.value
}

func (o *observableUint64) Set(newValue uint64) {
	if (newValue == o.Get()) {
		return
	}
	
	old := o.value
	o.value = newValue

	o.triggerChange(old, newValue)
}

func (o *observableUint64) triggerChange(old, new uint64) {
	for changeListener := range o.changeListeners {
		changeListener.Changed(old, new)
	}
}