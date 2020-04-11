package widget

import (
	"image"
	"log"

	"github.com/negrel/ginger/v1/painting"
)

// SizeFactor are used by layout to have multiple
type SizeFactor struct {

	// If non-null, sets its width to the child's width multiplied by this factor.
	WidthFactor int

	// If non-null, sets its height to the child's height multiplied by this factor.
	HeightFactor int
}

var _ Layout = &_center{}
var _ Widget = &_center{}

// _center is a widget that center its child within
// itself.
type _center struct {
	*LayoutSingleChild
	*SizeFactor
}

// Center return a layout that center its child within
// itself.
func Center(factor *SizeFactor, child Widget) Layout {
	cen := &_center{
		LayoutSingleChild: &LayoutSingleChild{
			Child: child,
		},
		SizeFactor: factor,
	}

	child.AdoptedBy(cen)

	return cen
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// Widget interface

// Draw implements Widget interface.
func (c *_center) Draw(bounds image.Rectangle) *painting.Frame {
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
