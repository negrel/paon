package events

import "github.com/negrel/paon/geometry"

type PointerEventData interface {
	RelativePosition() geometry.Vec2D
	WithPositionRelativeToOrigin(geometry.Vec2D) PointerEventData
}
