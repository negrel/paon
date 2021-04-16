package events

// Float32ChangeListener is notified whenever the value of an ObservableFloat32 change.
type Float32ChangeListener interface {
	Changed(old, new float32)
}

type float32ChangeListener func(old, new float32)

func (cl float32ChangeListener) Changed(old, new float32) {
	cl(old, new)
}

// MakeFloat32ChangeListener wraps your ObservableFloat32 change handler
// and returns it.
func MakeFloat32ChangeListener(handler func(old, new float32)) Float32ChangeListener {
	return float32ChangeListener(handler)
}

// ObservableFloat32 is an entity that wraps a float32 value and allows to observe
// the value for changes.
type ObservableFloat32 interface {
	Observable

	AddChangeListener(Float32ChangeListener)
	RemoveChangeListener(Float32ChangeListener)

	Get() float32
	Set(float32)
}

type observableFloat32 struct {
	observable
	value float32
	changeListeners map[Float32ChangeListener]struct{}
}

func NewObservableFloat32(initialValue float32) ObservableFloat32 {
	o := &observableFloat32{
		value:           initialValue,
		changeListeners: make(map[Float32ChangeListener]struct{}),
	}
	o.observable = makeCompositeObservable(o)

	return o
}

func (o *observableFloat32) AddChangeListener(listener Float32ChangeListener) {
	o.changeListeners[listener] = struct{}{}
}

func (o *observableFloat32) RemoveChangeListener(listener Float32ChangeListener) {
	delete(o.changeListeners, listener)
}

func (o *observableFloat32) Get() float32 {
	return o.value
}

func (o *observableFloat32) Set(newValue float32) {
	if (newValue == o.Get()) {
		return
	}
	
	old := o.value
	o.value = newValue

	o.triggerChange(old, newValue)
}

func (o *observableFloat32) triggerChange(old, new float32) {
	for changeListener := range o.changeListeners {
		changeListener.Changed(old, new)
	}
}