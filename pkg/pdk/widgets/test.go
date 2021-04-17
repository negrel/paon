package widgets

func newTestLayout() *layout {
	return newLayout("test-layout")
}

func newTestWidget() *widget {
	return newWidget("test-widget")
}

func newTestRoot() *Root {
	return &Root{
		layout: newLayout("test-root"),
	}
}
