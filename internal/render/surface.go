package render

// Surface define terminal window/surface to draw on.
type Surface interface {
	Update()
	Apply(Patch)
	Size() (w, h int)
	Clear()
	Fini()
}
