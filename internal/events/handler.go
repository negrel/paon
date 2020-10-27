package events

// Handler define object that can handle an event with one of the
// following function.
type Handler interface {
	OnError(ErrorEvent)
	OnInterrupt(InterruptEvent)
	OnResize(ResizeEvent)
	OnClick(ClickEvent)
	OnScroll(WheelEvent)
	OnInput(KeyboardEvent)
}

var _ Handler = EmptyHandler{}

type EmptyHandler struct {
}

// OnError implements the Handler interface.
func (e EmptyHandler) OnError(_ ErrorEvent) {
}

// OnInterrupt implements the Handler interface.
func (e EmptyHandler) OnInterrupt(_ InterruptEvent) {
}

// OnResize implements the Handler interface.
func (e EmptyHandler) OnResize(_ ResizeEvent) {
}

// OnClick implements the Handler interface.
func (e EmptyHandler) OnClick(_ ClickEvent) {
}

// OnScroll implements the Handler interface.
func (e EmptyHandler) OnScroll(_ WheelEvent) {
}

// OnInput implements the Handler interface.
func (e EmptyHandler) OnInput(_ KeyboardEvent) {
}
