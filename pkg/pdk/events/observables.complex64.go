package events

// Complex64ChangeListener is notified whenever the value of an ObservableComplex64 change.
type Complex64ChangeListener interface {
	Changed(old, new complex64)
}

type complex64ChangeListener func(old, new complex64)

func (cl complex64ChangeListener) Changed(old, new complex64) {
	cl(old, new)
}

// MakeComplex64ChangeListener wraps your ObservableComplex64 change handler
// and returns it.
func MakeComplex64ChangeListener(handler func(old, new complex64)) Complex64ChangeListener {
	return complex64ChangeListener(handler)
}

// ObservableComplex64 is an entity that wraps a complex64 value and allows to observe
// the value for changes.
type ObservableComplex64 interface {
	Observable

	AddChangeListener(Complex64ChangeListener)
	RemoveChangeListener(Complex64ChangeListener)

	Get() complex64
	Set(complex64)
}

type observableComplex64 struct {
	observable
	value complex64
	changeListeners map[Complex64ChangeListener]struct{}
}

func NewObservableComplex64(initialValue complex64) ObservableComplex64 {
	o := &observableComplex64{
		value:           initialValue,
		changeListeners: make(map[Complex64ChangeListener]struct{}),
	}
	o.observable = makeCompositeObservable(o)

	return o
}

func (o *observableComplex64) AddChangeListener(listener Complex64ChangeListener) {
	o.changeListeners[listener] = struct{}{}
}

func (o *observableComplex64) RemoveChangeListener(listener Complex64ChangeListener) {
	delete(o.changeListeners, listener)
}

func (o *observableComplex64) Get() complex64 {
	return o.value
}

func (o *observableComplex64) Set(newValue complex64) {
	if (newValue == o.Get()) {
		return
	}
	
	old := o.value
	o.value = newValue

	o.triggerChange(old, newValue)
}

func (o *observableComplex64) triggerChange(old, new complex64) {
	for changeListener := range o.changeListeners {
		changeListener.Changed(old, new)
	}
}