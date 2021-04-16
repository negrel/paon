package events

// StringChangeListener is notified whenever the value of an ObservableString change.
type StringChangeListener interface {
	Changed(old, new string)
}

type stringChangeListener func(old, new string)

func (cl stringChangeListener) Changed(old, new string) {
	cl(old, new)
}

// MakeStringChangeListener wraps your ObservableString change handler
// and returns it.
func MakeStringChangeListener(handler func(old, new string)) StringChangeListener {
	return stringChangeListener(handler)
}

// ObservableString is an entity that wraps a string value and allows to observe
// the value for changes.
type ObservableString interface {
	Observable

	AddChangeListener(StringChangeListener)
	RemoveChangeListener(StringChangeListener)

	Get() string
	Set(string)
}

type observableString struct {
	observable
	value string
	changeListeners map[StringChangeListener]struct{}
}

func NewObservableString(initialValue string) ObservableString {
	o := &observableString{
		value:           initialValue,
		changeListeners: make(map[StringChangeListener]struct{}),
	}
	o.observable = makeCompositeObservable(o)

	return o
}

func (o *observableString) AddChangeListener(listener StringChangeListener) {
	o.changeListeners[listener] = struct{}{}
}

func (o *observableString) RemoveChangeListener(listener StringChangeListener) {
	delete(o.changeListeners, listener)
}

func (o *observableString) Get() string {
	return o.value
}

func (o *observableString) Set(newValue string) {
	if (newValue == o.Get()) {
		return
	}
	
	old := o.value
	o.value = newValue

	o.triggerChange(old, newValue)
}

func (o *observableString) triggerChange(old, new string) {
	for changeListener := range o.changeListeners {
		changeListener.Changed(old, new)
	}
}