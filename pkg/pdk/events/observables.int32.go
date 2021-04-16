package events

// Int32ChangeListener is notified whenever the value of an ObservableInt32 change.
type Int32ChangeListener interface {
	Changed(old, new int32)
}

type int32ChangeListener func(old, new int32)

func (cl int32ChangeListener) Changed(old, new int32) {
	cl(old, new)
}

// MakeInt32ChangeListener wraps your ObservableInt32 change handler
// and returns it.
func MakeInt32ChangeListener(handler func(old, new int32)) Int32ChangeListener {
	return int32ChangeListener(handler)
}

// ObservableInt32 is an entity that wraps a int32 value and allows to observe
// the value for changes.
type ObservableInt32 interface {
	Observable

	AddChangeListener(Int32ChangeListener)
	RemoveChangeListener(Int32ChangeListener)

	Get() int32
	Set(int32)
}

type observableInt32 struct {
	observable
	value int32
	changeListeners map[Int32ChangeListener]struct{}
}

func NewObservableInt32(initialValue int32) ObservableInt32 {
	o := &observableInt32{
		value:           initialValue,
		changeListeners: make(map[Int32ChangeListener]struct{}),
	}
	o.observable = makeCompositeObservable(o)

	return o
}

func (o *observableInt32) AddChangeListener(listener Int32ChangeListener) {
	o.changeListeners[listener] = struct{}{}
}

func (o *observableInt32) RemoveChangeListener(listener Int32ChangeListener) {
	delete(o.changeListeners, listener)
}

func (o *observableInt32) Get() int32 {
	return o.value
}

func (o *observableInt32) Set(newValue int32) {
	if (newValue == o.Get()) {
		return
	}
	
	old := o.value
	o.value = newValue

	o.triggerChange(old, newValue)
}

func (o *observableInt32) triggerChange(old, new int32) {
	for changeListener := range o.changeListeners {
		changeListener.Changed(old, new)
	}
}