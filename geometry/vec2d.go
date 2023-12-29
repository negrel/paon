package geometry

// Vec2D define a point in 2 dimensional Euclidian space.
type Vec2D struct {
	X, Y int
}

// Add returns a new Vec2D translated by adding the given vector.
func (v2 Vec2D) Add(other Vec2D) Vec2D {
	v2.X += other.X
	v2.Y += other.Y

	return v2
}

// Sub returns a new Vec2D translated by subtracting the given vector.
func (v2 Vec2D) Sub(other Vec2D) Vec2D {
	v2.X -= other.X
	v2.Y -= other.Y

	return v2
}
