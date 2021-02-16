package layout

import "github.com/negrel/paon/internal/geometry"

type boxOffset [4]int

func makeBoxOffset(left, top, right, bottom int) boxOffset {
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
	return geometry.Rectangle{
		Min: rectangle.Min.Add(
			geometry.Pt(bo.left(), bo.top()),
		),
		Max: rectangle.Max.Sub(
			geometry.Pt(bo.right(), bo.bottom()),
		),
	}
}

func (bo boxOffset) reverse() boxOffset {
	return boxOffset{-bo.left(), -bo.top(), -bo.right(), -bo.bottom()}
}
