package events

// Target define an object that can receive events and may have listeners for them.
type Target interface {
	AddEventListener(Type, Listener)
	RemoveEventListener(Type, Listener)
	DispatchEvent(event Event)
}

var _ Target = &target{}

// target is an implementation of the Target interface.
type target struct {
	listeners [][]Listener
}

// NewTarget return a new event Target with no listeners.
func NewTarget() Target {
	return target{
		listeners: make([][]Listener, typeRegistry.Last()+1),
	}
}

func (t target) AddEventListener(tpe Type, listener Listener) {
	if t.listeners[tpe] == nil {
		t.listeners[tpe] = make([]Listener, 0, 8)
	}

	t.listeners[tpe] = append(t.listeners[tpe], listener)
}

// RemoveEventListener removes an event listener of a specific event type from the target.
func (t target) RemoveEventListener(tpe Type, listener Listener) {
	if t.listeners[tpe] == nil {
		return
	}

	for i, l := range t.listeners[tpe] {
		if l == listener {
			t.listeners[tpe] = append(t.listeners[tpe][:i], t.listeners[tpe][i+1:]...)
			return
		}
	}
}

// DispatchEvent dispatch the given event to listeners.
func (t target) DispatchEvent(event Event) {
	i := uint32(event.EventType())
	if t.listeners[i] == nil {
		return
	}

	for _, listener := range t.listeners[i] {
		listener.HandleEvent(event)
	}
}

type noOpTarget struct{}

// NewNoOpTarget returns a new Target that ignore events listener
// and events.
func NewNoOpTarget() Target {
	return noOpTarget{}
}

func (not noOpTarget) AddEventListener(tpe Type, listener Listener) {
}

// RemoveEventListener removes an event listener of a specific event type from the target.
func (not noOpTarget) RemoveEventListener(tpe Type, listener Listener) {
}

// DispatchEvent dispatch the given event to listeners.
func (not noOpTarget) DispatchEvent(event Event) {
}

