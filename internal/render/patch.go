package render

import "image"

// Patch define a rectangle screen area to update.
type Patch struct {
	origin image.Point
	frame  [][]Cell
}
