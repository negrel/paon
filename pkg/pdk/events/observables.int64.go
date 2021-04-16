package events

// Int64ChangeListener is notified whenever the value of an ObservableInt64 change.
type Int64ChangeListener interface {
	Changed(old, new int64)
}

type int64ChangeListener func(old, new int64)

func (cl int64ChangeListener) Changed(old, new int64) {
	cl(old, new)
}

// MakeInt64ChangeListener wraps your ObservableInt64 change handler
// and returns it.
func MakeInt64ChangeListener(handler func(old, new int64)) Int64ChangeListener {
	return int64ChangeListener(handler)
}

// ObservableInt64 is an entity that wraps a int64 value and allows to observe
// the value for changes.
type ObservableInt64 interface {
	Observable

	AddChangeListener(Int64ChangeListener)
	RemoveChangeListener(Int64ChangeListener)

	Get() int64
	Set(int64)
}

type observableInt64 struct {
	observable
	value int64
	changeListeners map[Int64ChangeListener]struct{}
}

func NewObservableInt64(initialValue int64) ObservableInt64 {
	o := &observableInt64{
		value:           initialValue,
		changeListeners: make(map[Int64ChangeListener]struct{}),
	}
	o.observable = makeCompositeObservable(o)

	return o
}

func (o *observableInt64) AddChangeListener(listener Int64ChangeListener) {
	o.changeListeners[listener] = struct{}{}
}

func (o *observableInt64) RemoveChangeListener(listener Int64ChangeListener) {
	delete(o.changeListeners, listener)
}

func (o *observableInt64) Get() int64 {
	return o.value
}

func (o *observableInt64) Set(newValue int64) {
	if (newValue == o.Get()) {
		return
	}
	
	old := o.value
	o.value = newValue

	o.triggerChange(old, newValue)
}

func (o *observableInt64) triggerChange(old, new int64) {
	for changeListener := range o.changeListeners {
		changeListener.Changed(old, new)
	}
}