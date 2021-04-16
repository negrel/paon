package events

// Int8ChangeListener is notified whenever the value of an ObservableInt8 change.
type Int8ChangeListener interface {
	Changed(old, new int8)
}

type int8ChangeListener func(old, new int8)

func (cl int8ChangeListener) Changed(old, new int8) {
	cl(old, new)
}

// MakeInt8ChangeListener wraps your ObservableInt8 change handler
// and returns it.
func MakeInt8ChangeListener(handler func(old, new int8)) Int8ChangeListener {
	return int8ChangeListener(handler)
}

// ObservableInt8 is an entity that wraps a int8 value and allows to observe
// the value for changes.
type ObservableInt8 interface {
	Observable

	AddChangeListener(Int8ChangeListener)
	RemoveChangeListener(Int8ChangeListener)

	Get() int8
	Set(int8)
}

type observableInt8 struct {
	observable
	value int8
	changeListeners map[Int8ChangeListener]struct{}
}

func NewObservableInt8(initialValue int8) ObservableInt8 {
	o := &observableInt8{
		value:           initialValue,
		changeListeners: make(map[Int8ChangeListener]struct{}),
	}
	o.observable = makeCompositeObservable(o)

	return o
}

func (o *observableInt8) AddChangeListener(listener Int8ChangeListener) {
	o.changeListeners[listener] = struct{}{}
}

func (o *observableInt8) RemoveChangeListener(listener Int8ChangeListener) {
	delete(o.changeListeners, listener)
}

func (o *observableInt8) Get() int8 {
	return o.value
}

func (o *observableInt8) Set(newValue int8) {
	if (newValue == o.Get()) {
		return
	}
	
	old := o.value
	o.value = newValue

	o.triggerChange(old, newValue)
}

func (o *observableInt8) triggerChange(old, new int8) {
	for changeListener := range o.changeListeners {
		changeListener.Changed(old, new)
	}
}