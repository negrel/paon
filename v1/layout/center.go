package layout

import (
	"image"
	"log"

	"github.com/negrel/ginger/v1/painting"
	"github.com/negrel/ginger/v1/widget"
)

var _ Layout = &Center{}
var _ widget.Widget = &Center{}

// Center is a widget that center its child within
// itself.
type Center struct {
	*BaseSingleChild

	// If non-null, sets its width to the child's width multiplied by this factor.
	WidthFactor int

	// If non-null, sets its height to the child's height multiplied by this factor.
	HeightFactor int
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// Widget interface

// Draw implements Widget interface.
func (c *Center) Draw(bounds image.Rectangle) *painting.Frame {
	// Child bounds are relative
	cBounds := image.Rectangle{
		Min: image.Point{},
		Max: bounds.Max.Sub(bounds.Min),
	}

	// Drawing child
	cFrame := c.Child.Draw(cBounds)
	cWidth := cFrame.Patch.Width()
	cHeight := cFrame.Patch.Height()

	width := bounds.Dx()
	height := bounds.Dy()

	// Computing height & width factor
	if c.WidthFactor != 0 &&
		(cWidth*c.WidthFactor) < width {
		width = cWidth
	}

	if c.HeightFactor != 0 &&
		(cHeight*c.HeightFactor) < height {
		height = cHeight
	}

	// The final frame
	frame := painting.NewFrame(bounds.Min, width, height)

	// Centering position
	cPosition := image.Point{
		X: (width/2 - cWidth/2),
		Y: (height/2 - cHeight/2),
	}
	cFrame.Position = cPosition

	// Adding centered child frame
	err := frame.Add(cFrame)
	if err != nil {
		log.Print("CENTER: ", err)
		log.Printf("CENTER FRAME: %+v %+v %+v", frame.Position, frame.Patch.Width(), frame.Patch.Height())
		log.Fatalf("CHILD FRAME: %+v %+v %+v", cFrame.Position, cFrame.Patch.Width(), cFrame.Patch.Height())
	}

	return frame
}
