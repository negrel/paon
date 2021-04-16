package events

// Float64ChangeListener is notified whenever the value of an ObservableFloat64 change.
type Float64ChangeListener interface {
	Changed(old, new float64)
}

type float64ChangeListener func(old, new float64)

func (cl float64ChangeListener) Changed(old, new float64) {
	cl(old, new)
}

// MakeFloat64ChangeListener wraps your ObservableFloat64 change handler
// and returns it.
func MakeFloat64ChangeListener(handler func(old, new float64)) Float64ChangeListener {
	return float64ChangeListener(handler)
}

// ObservableFloat64 is an entity that wraps a float64 value and allows to observe
// the value for changes.
type ObservableFloat64 interface {
	Observable

	AddChangeListener(Float64ChangeListener)
	RemoveChangeListener(Float64ChangeListener)

	Get() float64
	Set(float64)
}

type observableFloat64 struct {
	observable
	value float64
	changeListeners map[Float64ChangeListener]struct{}
}

func NewObservableFloat64(initialValue float64) ObservableFloat64 {
	o := &observableFloat64{
		value:           initialValue,
		changeListeners: make(map[Float64ChangeListener]struct{}),
	}
	o.observable = makeCompositeObservable(o)

	return o
}

func (o *observableFloat64) AddChangeListener(listener Float64ChangeListener) {
	o.changeListeners[listener] = struct{}{}
}

func (o *observableFloat64) RemoveChangeListener(listener Float64ChangeListener) {
	delete(o.changeListeners, listener)
}

func (o *observableFloat64) Get() float64 {
	return o.value
}

func (o *observableFloat64) Set(newValue float64) {
	if (newValue == o.Get()) {
		return
	}
	
	old := o.value
	o.value = newValue

	o.triggerChange(old, newValue)
}

func (o *observableFloat64) triggerChange(old, new float64) {
	for changeListener := range o.changeListeners {
		changeListener.Changed(old, new)
	}
}