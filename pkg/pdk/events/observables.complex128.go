package events

// Complex128ChangeListener is notified whenever the value of an ObservableComplex128 change.
type Complex128ChangeListener interface {
	Changed(old, new complex128)
}

type complex128ChangeListener func(old, new complex128)

func (cl complex128ChangeListener) Changed(old, new complex128) {
	cl(old, new)
}

// MakeComplex128ChangeListener wraps your ObservableComplex128 change handler
// and returns it.
func MakeComplex128ChangeListener(handler func(old, new complex128)) Complex128ChangeListener {
	return complex128ChangeListener(handler)
}

// ObservableComplex128 is an entity that wraps a complex128 value and allows to observe
// the value for changes.
type ObservableComplex128 interface {
	Observable

	AddChangeListener(Complex128ChangeListener)
	RemoveChangeListener(Complex128ChangeListener)

	Get() complex128
	Set(complex128)
}

type observableComplex128 struct {
	observable
	value complex128
	changeListeners map[Complex128ChangeListener]struct{}
}

func NewObservableComplex128(initialValue complex128) ObservableComplex128 {
	o := &observableComplex128{
		value:           initialValue,
		changeListeners: make(map[Complex128ChangeListener]struct{}),
	}
	o.observable = makeCompositeObservable(o)

	return o
}

func (o *observableComplex128) AddChangeListener(listener Complex128ChangeListener) {
	o.changeListeners[listener] = struct{}{}
}

func (o *observableComplex128) RemoveChangeListener(listener Complex128ChangeListener) {
	delete(o.changeListeners, listener)
}

func (o *observableComplex128) Get() complex128 {
	return o.value
}

func (o *observableComplex128) Set(newValue complex128) {
	if (newValue == o.Get()) {
		return
	}
	
	old := o.value
	o.value = newValue

	o.triggerChange(old, newValue)
}

func (o *observableComplex128) triggerChange(old, new complex128) {
	for changeListener := range o.changeListeners {
		changeListener.Changed(old, new)
	}
}