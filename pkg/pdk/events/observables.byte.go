package events

// ByteChangeListener is notified whenever the value of an ObservableByte change.
type ByteChangeListener interface {
	Changed(old, new byte)
}

type byteChangeListener func(old, new byte)

func (cl byteChangeListener) Changed(old, new byte) {
	cl(old, new)
}

// MakeByteChangeListener wraps your ObservableByte change handler
// and returns it.
func MakeByteChangeListener(handler func(old, new byte)) ByteChangeListener {
	return byteChangeListener(handler)
}

// ObservableByte is an entity that wraps a byte value and allows to observe
// the value for changes.
type ObservableByte interface {
	Observable

	AddChangeListener(ByteChangeListener)
	RemoveChangeListener(ByteChangeListener)

	Get() byte
	Set(byte)
}

type observableByte struct {
	observable
	value byte
	changeListeners map[ByteChangeListener]struct{}
}

func NewObservableByte(initialValue byte) ObservableByte {
	o := &observableByte{
		value:           initialValue,
		changeListeners: make(map[ByteChangeListener]struct{}),
	}
	o.observable = makeCompositeObservable(o)

	return o
}

func (o *observableByte) AddChangeListener(listener ByteChangeListener) {
	o.changeListeners[listener] = struct{}{}
}

func (o *observableByte) RemoveChangeListener(listener ByteChangeListener) {
	delete(o.changeListeners, listener)
}

func (o *observableByte) Get() byte {
	return o.value
}

func (o *observableByte) Set(newValue byte) {
	if (newValue == o.Get()) {
		return
	}
	
	old := o.value
	o.value = newValue

	o.triggerChange(old, newValue)
}

func (o *observableByte) triggerChange(old, new byte) {
	for changeListener := range o.changeListeners {
		changeListener.Changed(old, new)
	}
}