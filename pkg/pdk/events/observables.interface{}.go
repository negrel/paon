package events

// ObjectChangeListener is notified whenever the value of an ObservableObject change.
type ObjectChangeListener interface {
	Changed(old, new interface{})
}

type objectChangeListener func(old, new interface{})

func (cl objectChangeListener) Changed(old, new interface{}) {
	cl(old, new)
}

// MakeObjectChangeListener wraps your ObservableObject change handler
// and returns it.
func MakeObjectChangeListener(handler func(old, new interface{})) ObjectChangeListener {
	return objectChangeListener(handler)
}

// ObservableObject is an entity that wraps a interface{} value and allows to observe
// the value for changes.
type ObservableObject interface {
	Observable

	AddChangeListener(ObjectChangeListener)
	RemoveChangeListener(ObjectChangeListener)

	Get() interface{}
	Set(interface{})
}

