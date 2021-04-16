package events

// BoolChangeListener is notified whenever the value of an ObservableBool change.
type BoolChangeListener interface {
	Changed(old, new bool)
}

type boolChangeListener func(old, new bool)

func (cl boolChangeListener) Changed(old, new bool) {
	cl(old, new)
}

// MakeBoolChangeListener wraps your ObservableBool change handler
// and returns it.
func MakeBoolChangeListener(handler func(old, new bool)) BoolChangeListener {
	return boolChangeListener(handler)
}

// ObservableBool is an entity that wraps a bool value and allows to observe
// the value for changes.
type ObservableBool interface {
	Observable

	AddChangeListener(BoolChangeListener)
	RemoveChangeListener(BoolChangeListener)

	Get() bool
	Set(bool)
}

type observableBool struct {
	observable
	value bool
	changeListeners map[BoolChangeListener]struct{}
}

func NewObservableBool(initialValue bool) ObservableBool {
	o := &observableBool{
		value:           initialValue,
		changeListeners: make(map[BoolChangeListener]struct{}),
	}
	o.observable = makeCompositeObservable(o)

	return o
}

func (o *observableBool) AddChangeListener(listener BoolChangeListener) {
	o.changeListeners[listener] = struct{}{}
}

func (o *observableBool) RemoveChangeListener(listener BoolChangeListener) {
	delete(o.changeListeners, listener)
}

func (o *observableBool) Get() bool {
	return o.value
}

func (o *observableBool) Set(newValue bool) {
	if (newValue == o.Get()) {
		return
	}
	
	old := o.value
	o.value = newValue

	o.triggerChange(old, newValue)
}

func (o *observableBool) triggerChange(old, new bool) {
	for changeListener := range o.changeListeners {
		changeListener.Changed(old, new)
	}
}