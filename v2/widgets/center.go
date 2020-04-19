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
	*events.ResizeListener
}

// Center return a layout that center its child within
// itself.
func Center(child Widget) Layout {
	cen := &_center{
		CoreLayout:     NewCoreLayout("center", []Widget{child}),
		ResizeListener: &events.ResizeListener{},
	}

	cen.OnResize = func(_ events.ResizeEvent) {
		cen.cache.Invalid()
	}

	cen.Rendering = cen.rendering

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

// rendering implements Widget interface.
func (c *_center) rendering(bounds image.Rectangle) *render.Frame {
	// Child bounds are relative
	childBounds := image.Rectangle{
		Min: image.Pt(0, 0),
		Max: bounds.Max.Sub(bounds.Min),
	}

	width := bounds.Dx()
	height := bounds.Dy()

	// Drawing child
	childFrame := c.Child().Render(childBounds)
	childWidth := childFrame.Patch.Width()
	childHeight := childFrame.Patch.Height()

	// The final frame
	frame := render.NewFrame(bounds.Min, width, height)

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
