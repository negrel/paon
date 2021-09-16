package layout

import "github.com/negrel/paon/internal/geometry"

type boxOffset [4]int

func newBoxOffset(left, top, right, bottom int) boxOffset {
	return boxOffset{left, top, right, bottom}
}

func (bo boxOffset) left() int {
	return bo[0]
}

func (bo boxOffset) top() int {
	return bo[1]
}

func (bo boxOffset) right() int {
	return bo[2]
}

func (bo boxOffset) bottom() int {
	return bo[3]
}

func (bo boxOffset) x() int {
	return bo.left() + bo.right()
}

func (bo boxOffset) y() int {
	return bo.top() + bo.bottom()
}

func (bo boxOffset) applyOn(rectangle geometry.Rectangle) geometry.Rectangle {
	min := rectangle.Min.Add(
		geometry.NewVec2D(bo.left(), bo.top()),
	)
	max := rectangle.Max.Sub(
		geometry.NewVec2D(bo.right(), bo.bottom()),
	)

	if min.X() > max.X() || min.Y() > max.Y() {
		return geometry.Rectangle{}
	}

	return geometry.Rectangle{
		Min: min,
		Max: max,
	}
}
