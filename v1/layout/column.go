package layout

import (
	"image"
	"log"

	"github.com/negrel/ginger/v1/painting"
	"github.com/negrel/ginger/v1/widget"
)

var _ widget.Widget = &Column{}

// Column is a layout that arrange widget vertically.
type Column struct {
	*widget.Base

	Children []widget.Widget
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// Widget interface

func (c *Column) drawChilds(bounds image.Rectangle) ([]*painting.Frame, image.Point) {
	childCount := len(c.Children)
	cFrames := make([]*painting.Frame, childCount)
	size := image.Pt(0, 0)

	for i := 0; i < childCount; i++ {
		cFrame := c.Children[i].Draw(bounds)
		cFrames[i] = cFrame

		if cWidth := cFrame.Patch.Width(); cWidth > size.X {
			size.X = cWidth
		}

		// updating bounds for next children
		bounds.Min.Y += cFrame.Patch.Height()
		// updating total size.Y
		size.Y += cFrame.Patch.Height()
	}

	return cFrames, size
}

// Draw implements Widget interface.
func (c *Column) Draw(bounds image.Rectangle) *painting.Frame {
	// child bounds are relative
	cBounds := image.Rectangle{
		Min: image.Point{},
		Max: bounds.Max.Sub(bounds.Min),
	}

	cFrames, size := c.drawChilds(cBounds)

	frame := painting.NewFrame(bounds.Min, size.X, size.Y)

	for i := 0; i < len(cFrames); i++ {
		err := frame.Add(cFrames[i])

		if err != nil {
			log.Print("COLUMN: ", err)
			log.Printf("COLUMN FRAME: %+v %+v %+v", frame.Position, frame.Patch.Width(), frame.Patch.Height())
			log.Fatalf("CHILD n°%v FRAME: %+v %+v %+v", i, cFrames[i], cFrames[i].Patch.Width(), cFrames[i].Patch.Height())
		}
	}

	return frame
}
