package geometry

// Vec2D define a point in 2 dimensional Euclidian space.
type Vec2D struct {
	x, y int
}

// NewVec2D returns a new 2D vector containing the given values.
func NewVec2D(x, y int) Vec2D {
	return Vec2D{x, y}
}

// X returns the location of the vector on the X axis.
func (v2 Vec2D) X() int {
	return v2.x
}

// Y returns the location of the vector on the Y axis.
func (v2 Vec2D) Y() int {
	return v2.y
}

// Add returns a new Vec2D translated by adding the given vector.
func (v2 Vec2D) Add(other Vec2D) Vec2D {
	v2.x += other.x
	v2.y += other.y

	return v2
}

// Sub returns a new Vec2D translated by subtracting the given vector.
func (v2 Vec2D) Sub(other Vec2D) Vec2D {
	v2.x -= other.x
	v2.y -= other.y

	return v2
}

// Equals returns true if the given vector is equal to this Vec2D.
func (v2 Vec2D) Equals(other Vec2D) bool {
	return v2.X() == other.X() && v2.Y() == other.Y()
}
