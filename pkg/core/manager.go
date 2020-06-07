package core

// Manager handle the widget tree, it is responsible
// for updating widget that need to be laid out or rendered.
// before before the new frame.
type Manager interface {
	// The given widget will be rendered for the
	// next frame.
	ScheduleRenderFor(Widget)
}
