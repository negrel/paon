package draw

// Painter is responsible for painting, sizing and positioning an element
// on a Canvas.
type Painter interface {
	Layout(ctx Context)
	Paint(ctx Context)
}

// Drawable represents object that can draw themself.
type Drawable interface {
	Draw(Canvas)
}
