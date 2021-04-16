package events

// UintChangeListener is notified whenever the value of an ObservableUint change.
type UintChangeListener interface {
	Changed(old, new uint)
}

type uintChangeListener func(old, new uint)

func (cl uintChangeListener) Changed(old, new uint) {
	cl(old, new)
}

// MakeUintChangeListener wraps your ObservableUint change handler
// and returns it.
func MakeUintChangeListener(handler func(old, new uint)) UintChangeListener {
	return uintChangeListener(handler)
}

// ObservableUint is an entity that wraps a uint value and allows to observe
// the value for changes.
type ObservableUint interface {
	Observable

	AddChangeListener(UintChangeListener)
	RemoveChangeListener(UintChangeListener)

	Get() uint
	Set(uint)
}

type observableUint struct {
	observable
	value uint
	changeListeners map[UintChangeListener]struct{}
}

func NewObservableUint(initialValue uint) ObservableUint {
	o := &observableUint{
		value:           initialValue,
		changeListeners: make(map[UintChangeListener]struct{}),
	}
	o.observable = makeCompositeObservable(o)

	return o
}

func (o *observableUint) AddChangeListener(listener UintChangeListener) {
	o.changeListeners[listener] = struct{}{}
}

func (o *observableUint) RemoveChangeListener(listener UintChangeListener) {
	delete(o.changeListeners, listener)
}

func (o *observableUint) Get() uint {
	return o.value
}

func (o *observableUint) Set(newValue uint) {
	if (newValue == o.Get()) {
		return
	}
	
	old := o.value
	o.value = newValue

	o.triggerChange(old, newValue)
}

func (o *observableUint) triggerChange(old, new uint) {
	for changeListener := range o.changeListeners {
		changeListener.Changed(old, new)
	}
}