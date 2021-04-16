package events

// RuneChangeListener is notified whenever the value of an ObservableRune change.
type RuneChangeListener interface {
	Changed(old, new rune)
}

type runeChangeListener func(old, new rune)

func (cl runeChangeListener) Changed(old, new rune) {
	cl(old, new)
}

// MakeRuneChangeListener wraps your ObservableRune change handler
// and returns it.
func MakeRuneChangeListener(handler func(old, new rune)) RuneChangeListener {
	return runeChangeListener(handler)
}

// ObservableRune is an entity that wraps a rune value and allows to observe
// the value for changes.
type ObservableRune interface {
	Observable

	AddChangeListener(RuneChangeListener)
	RemoveChangeListener(RuneChangeListener)

	Get() rune
	Set(rune)
}

type observableRune struct {
	observable
	value rune
	changeListeners map[RuneChangeListener]struct{}
}

func NewObservableRune(initialValue rune) ObservableRune {
	o := &observableRune{
		value:           initialValue,
		changeListeners: make(map[RuneChangeListener]struct{}),
	}
	o.observable = makeCompositeObservable(o)

	return o
}

func (o *observableRune) AddChangeListener(listener RuneChangeListener) {
	o.changeListeners[listener] = struct{}{}
}

func (o *observableRune) RemoveChangeListener(listener RuneChangeListener) {
	delete(o.changeListeners, listener)
}

func (o *observableRune) Get() rune {
	return o.value
}

func (o *observableRune) Set(newValue rune) {
	if (newValue == o.Get()) {
		return
	}
	
	old := o.value
	o.value = newValue

	o.triggerChange(old, newValue)
}

func (o *observableRune) triggerChange(old, new rune) {
	for changeListener := range o.changeListeners {
		changeListener.Changed(old, new)
	}
}