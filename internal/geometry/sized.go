package geometry

// Sized define any sized object.
type Sized interface {
	Size() Size
	Width() int
	Height() int
}
