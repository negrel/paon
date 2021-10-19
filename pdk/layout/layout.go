package layout

import "github.com/negrel/paon/geometry"

// Layout define objects responsible for laying out sized object.
// The result is a rectangle with the same size as the shared object.
type Layout interface {
	Layout(geometry.Sized) geometry.Rectangle
}

// LayoutFunc define a function type that implements the Layout interface.
type LayoutFunc func(geometry.Sized) geometry.Rectangle

// Layout implements the Layout interface.
func (lf LayoutFunc) Layout(sized geometry.Sized) geometry.Rectangle {
	return lf(sized)
}

// BoxPacker is a generic interface for object
// that can produce BoxModel using the given layout.Constraint.
type BoxPacker interface {
	Pack(Constraint) BoxModel
}

// Sized define objects that can compute their size using Constraint.
type Sized interface {
	Size(Constraint) geometry.Size
}
