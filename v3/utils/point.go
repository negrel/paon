package utils

// Infinite represent an infinite value.
const Infinite int = -1

// Point define a point in Cartesian space a
// specified distance from a separately-maintained
// origin. No negative point should be used.
type Point struct {
	X, Y int
}

// Pt is a shortcut for Point.
func Pt(x, y int) Point {
	return Point{
		X: x,
		Y: y,
	}
}
