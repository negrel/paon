package geometry

// Infinite represent an infinite value.
const Infinite int = -1

// Point define a point in Cartesian space a
// specified distance from a separately-maintained
// origin. No negative point should be used.
type Point struct {
	x, y int
}

// Pt returns a new Point using the given X and Y value.
func Pt(x, y int) Point {
	return Point{
		x: x,
		y: y,
	}
}

// X returns the location of the point on the X axis.
func (p Point) X() int {
	return p.x
}

// Y returns the location of the point on the Y axis.
func (p Point) Y() int {
	return p.y
}

// Add returns a new Point translated by adding the given point.
func (p Point) Add(other Point) Point {
	p.x += other.x
	p.y += other.y

	return p
}

// Sub returns a new Point translated by subtracting the given point.
func (p Point) Sub(other Point) Point {
	p.x -= other.x
	p.y -= other.y

	return p
}

// Equals returns true if the given Point is equal to this Point.
func (p Point) Equals(other Point) bool {
	return p.X() == other.X() && p.Y() == other.Y()
}
