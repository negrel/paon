package layout

import "github.com/negrel/paon/geometry"

type boxOffset struct{ left, top, right, bottom int }

func (bo boxOffset) x() int {
	return bo.left + bo.right
}

func (bo boxOffset) y() int {
	return bo.top + bo.bottom
}

func (bo boxOffset) applyOn(rectangle geometry.Rectangle) geometry.Rectangle {
	min := rectangle.Min.Add(
		geometry.NewVec2D(bo.left, bo.top),
	)
	max := rectangle.Max.Sub(
		geometry.NewVec2D(bo.right, bo.bottom),
	)

	if min.X() > max.X() || min.Y() > max.Y() {
		return geometry.Rectangle{}
	}

	return geometry.Rectangle{
		Min: min,
		Max: max,
	}
}
