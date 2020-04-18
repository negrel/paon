package widgets

import (
	"image"
	"log"

	"github.com/negrel/ginger/v2/render"
)

// SizeFactor are used by layout to have multiple
type SizeFactor struct {

	// If non-null, sets its width to the child's width multiplied by this factor.
	WidthFactor int

	// If non-null, sets its height to the child's height multiplied by this factor.
	HeightFactor int
}

var _ Widget = &_center{}
var _ Layout = &_center{}

// _center is a widget that center its child within
// itself.
type _center struct {
	*CoreLayout
	*SizeFactor
}

// Center return a layout that center its child within
// itself.
func Center(factor *SizeFactor, child Widget) Layout {
	cen := &_center{
		CoreLayout: NewCoreLayout([]Widget{child}),
		SizeFactor: factor,
	}

	cen.AdoptChild(child)

	cen.Draw = cen.draw

	// cen.AdoptChild(child)

	return cen
}

/*****************************************************
 **************** Getters & Setters ******************
 *****************************************************/
// ANCHOR Getters & Setters

func (c *_center) Child() Widget {
	return c.Children[0]
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// Widget interface

// Draw implements Widget interface.
func (c *_center) draw(co Constraint) *render.Frame {
	// Child bounds are relative
	childConstraint := Constraint{
		image.Rectangle{
			Min: image.Pt(0, 0),
			Max: co.Bounds.Max.Sub(co.Bounds.Min),
		},
	}

	width := co.Bounds.Dx()
	height := co.Bounds.Dy()

	// Drawing child
	childFrame := c.Child().Render(childConstraint)
	childWidth := childFrame.Patch.Width()
	childHeight := childFrame.Patch.Height()

	// Computing height & width factor
	if c.WidthFactor != 0 &&
		(childWidth*c.WidthFactor) < width {
		width = childWidth
	}

	if c.HeightFactor != 0 &&
		(childHeight*c.HeightFactor) < height {
		height = childHeight
	}

	// The final frame
	frame := render.NewFrame(co.Bounds.Min, width, height)

	// Centering position
	cPosition := image.Point{
		X: (width/2 - childWidth/2),
		Y: (height/2 - childHeight/2),
	}
	childFrame.Position = cPosition

	// Adding centered child frame
	err := frame.Add(childFrame)
	if err != nil {
		log.Print("CENTER: ", err)
		log.Printf("CENTER FRAME: %+v %+v %+v", frame.Position, frame.Patch.Width(), frame.Patch.Height())
		log.Fatalf("CHILD FRAME: %+v %+v %+v", childFrame.Position, childFrame.Patch.Width(), childFrame.Patch.Height())
	}

	return frame
}