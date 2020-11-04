package geometry

import (
	"github.com/negrel/debuggo/pkg/assert"
)

// Infinite represent an infinite value.
const Infinite int = -1

// Point define a point in Cartesian space a
// specified distance from a separately-maintained
// origin. No negative point should be used.
type Point struct {
	x, y int
}

// Pt is a shortcut for Point.
func Pt(x, y int) Point {
	assert.GreaterOrEqual(0, x)
	assert.GreaterOrEqual(0, y)

	return Point{
		x: x,
		y: y,
	}
}

func (p Point) X() int {
	return p.x
}

func (p Point) Y() int {
	return p.y
}

// Add translate the point coordinate by adding the
// given point.
func (p Point) Add(other Point) Point {

	p.x += other.x
	p.y += other.y

	return p
}

// Sub translate the point coordinate by substracting
// the given point.
func (p Point) Sub(other Point) Point {
	p.x -= other.x
	p.y -= other.y

	return p
}
