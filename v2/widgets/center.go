package widgets

import (
	"image"
	"log"

	"github.com/negrel/ginger/v2/render"
	"github.com/negrel/ginger/v2/widgets/events"
)

var _ Widget = &_center{}
var _ Layout = &_center{}

// _center is a widget that center its child within
// itself.
type _center struct {
	*CoreLayout
}

// Center return a layout that center its child within
// itself.
func Center(child Widget) Layout {
	cen := &_center{
		CoreLayout: NewCoreLayout("center", []Widget{child}),
	}
	cen.Rendering = cen.rendering

	return cen
}

// OnScroll handle scroll events
func (c *_center) OnScroll(se *events.ScrollEvent) {
	log.Println("Scroll to the", se.Direction().String())
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

// rendering implements Widget interface.
func (c *_center) rendering(bounds image.Rectangle) *render.Frame {
	width := bounds.Dx()
	height := bounds.Dy()

	// Drawing child
	childFrame := c.Child().Render(bounds)
	childWidth := childFrame.Patch.Width()
	childHeight := childFrame.Patch.Height()

	// The final frame
	frame := render.NewFrame(width, height)

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
