package render

import (
	"image"
)

// Rendable is implementeb by elements that have a Render
// method which return the rendered Frame.
type Rendable interface {
	Render(image.Rectangle) *Frame
}

// Func define the type of Rendable render function.
type Func = func(image.Rectangle) *Frame
